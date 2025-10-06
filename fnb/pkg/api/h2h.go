package api

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

type H2HClient struct {
	config *H2HConfig
}

func NewH2HClient(config *H2HConfig) *H2HClient {
	return &H2HClient{
		config: config,
	}
}

type H2HPaymentFile struct {
	Header   H2HFileHeader
	Payments []H2HPaymentRecord
	Trailer  H2HFileTrailer
}

type H2HFileHeader struct {
	RecordType      string
	FileType        string
	FileReference   string
	CreationDate    time.Time
	OriginatorCode  string
	OriginatorName  string
	TestIndicator   string
}

type H2HPaymentRecord struct {
	RecordType          string
	SequenceNumber      int
	BeneficiaryAccount  string
	BeneficiaryName     string
	BankCode            string
	BranchCode          string
	Amount              float64
	PaymentReference    string
	BeneficiaryReference string
	ActionDate          string
}

type H2HFileTrailer struct {
	RecordType   string
	RecordCount  int
	TotalAmount  float64
	HashTotal    string
}

type H2HCollectionFile struct {
	Header      H2HFileHeader
	Collections []H2HCollectionRecord
	Trailer     H2HFileTrailer
}

type H2HCollectionRecord struct {
	RecordType          string
	SequenceNumber      int
	DebtorAccount       string
	DebtorName          string
	BankCode            string
	BranchCode          string
	Amount              float64
	CollectionReference string
	ContractReference   string
	ActionDate          string
	MandateReference    string
}

type H2HResponseFile struct {
	Header   H2HFileHeader
	Records  []H2HResponseRecord
	Trailer  H2HFileTrailer
}

type H2HResponseRecord struct {
	RecordType       string
	SequenceNumber   int
	TransactionID    string
	Status           string
	StatusCode       string
	StatusDescription string
	Reference        string
	Amount           float64
}

func (h *H2HClient) GeneratePaymentFile(file H2HPaymentFile) (string, error) {
	var sb strings.Builder

	headerLine := fmt.Sprintf("H|%s|%s|%s|%s|%s|%s\n",
		file.Header.FileType,
		file.Header.FileReference,
		file.Header.CreationDate.Format("20060102150405"),
		file.Header.OriginatorCode,
		file.Header.OriginatorName,
		file.Header.TestIndicator,
	)
	sb.WriteString(headerLine)

	for _, payment := range file.Payments {
		paymentLine := fmt.Sprintf("P|%d|%s|%s|%s|%s|%.2f|%s|%s|%s\n",
			payment.SequenceNumber,
			payment.BeneficiaryAccount,
			payment.BeneficiaryName,
			payment.BankCode,
			payment.BranchCode,
			payment.Amount,
			payment.PaymentReference,
			payment.BeneficiaryReference,
			payment.ActionDate,
		)
		sb.WriteString(paymentLine)
	}

	trailerLine := fmt.Sprintf("T|%d|%.2f|%s\n",
		file.Trailer.RecordCount,
		file.Trailer.TotalAmount,
		file.Trailer.HashTotal,
	)
	sb.WriteString(trailerLine)

	return sb.String(), nil
}

func (h *H2HClient) GenerateCollectionFile(file H2HCollectionFile) (string, error) {
	var sb strings.Builder

	headerLine := fmt.Sprintf("H|%s|%s|%s|%s|%s|%s\n",
		file.Header.FileType,
		file.Header.FileReference,
		file.Header.CreationDate.Format("20060102150405"),
		file.Header.OriginatorCode,
		file.Header.OriginatorName,
		file.Header.TestIndicator,
	)
	sb.WriteString(headerLine)

	for _, collection := range file.Collections {
		collectionLine := fmt.Sprintf("C|%d|%s|%s|%s|%s|%.2f|%s|%s|%s|%s\n",
			collection.SequenceNumber,
			collection.DebtorAccount,
			collection.DebtorName,
			collection.BankCode,
			collection.BranchCode,
			collection.Amount,
			collection.CollectionReference,
			collection.ContractReference,
			collection.ActionDate,
			collection.MandateReference,
		)
		sb.WriteString(collectionLine)
	}

	trailerLine := fmt.Sprintf("T|%d|%.2f|%s\n",
		file.Trailer.RecordCount,
		file.Trailer.TotalAmount,
		file.Trailer.HashTotal,
	)
	sb.WriteString(trailerLine)

	return sb.String(), nil
}

func (h *H2HClient) UploadFile(ctx context.Context, filename string, content string) error {
	conn, err := ftp.Dial(h.config.FTPHost, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		return fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	if err := conn.Login(h.config.FTPUser, h.config.FTPPassword); err != nil {
		return fmt.Errorf("failed to login to FTP server: %w", err)
	}
	if h.config.FTPDirectory != "" {
		if err := conn.ChangeDir(h.config.FTPDirectory); err != nil {
			return fmt.Errorf("failed to change directory: %w", err)
		}
	}

	reader := strings.NewReader(content)
	if err := conn.Stor(filename, reader); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}

func (h *H2HClient) DownloadResponseFile(ctx context.Context, filename string) (string, error) {
	conn, err := ftp.Dial(h.config.FTPHost, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		return "", fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	if err := conn.Login(h.config.FTPUser, h.config.FTPPassword); err != nil {
		return "", fmt.Errorf("failed to login to FTP server: %w", err)
	}
	if h.config.FTPDirectory != "" {
		if err := conn.ChangeDir(h.config.FTPDirectory); err != nil {
			return "", fmt.Errorf("failed to change directory: %w", err)
		}
	}

	resp, err := conn.Retr(filename)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Close()

	content, err := io.ReadAll(resp)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %w", err)
	}

	return string(content), nil
}

func (h *H2HClient) ParseResponseFile(content string) (*H2HResponseFile, error) {
	file := &H2HResponseFile{
		Records: make([]H2HResponseRecord, 0),
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		if len(fields) == 0 {
			continue
		}

		recordType := fields[0]
		switch recordType {
		case "H":
			if len(fields) >= 7 {
				creationDate, _ := time.Parse("20060102150405", fields[3])
				file.Header = H2HFileHeader{
					RecordType:     recordType,
					FileType:       fields[1],
					FileReference:  fields[2],
					CreationDate:   creationDate,
					OriginatorCode: fields[4],
					OriginatorName: fields[5],
					TestIndicator:  fields[6],
				}
			}
		case "R":
			if len(fields) >= 8 {
				var amount float64
				fmt.Sscanf(fields[7], "%f", &amount)

				record := H2HResponseRecord{
					RecordType:        recordType,
					TransactionID:     fields[2],
					Status:            fields[3],
					StatusCode:        fields[4],
					StatusDescription: fields[5],
					Reference:         fields[6],
					Amount:            amount,
				}
				file.Records = append(file.Records, record)
			}
		case "T":
			if len(fields) >= 3 {
				var recordCount int
				var totalAmount float64
				fmt.Sscanf(fields[1], "%d", &recordCount)
				fmt.Sscanf(fields[2], "%f", &totalAmount)

				file.Trailer = H2HFileTrailer{
					RecordType:  recordType,
					RecordCount: recordCount,
					TotalAmount: totalAmount,
					HashTotal:   fields[3],
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return file, nil
}

func SaveToFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func ReadFromFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}
