package daraja

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const qrCodeURL = "/mpesa/qrcode/v1/generate"

type QRCodeRequest struct {
	MerchantName string `json:"MerchantName"`
	RefNo        string `json:"RefNo"`
	Amount       int    `json:"Amount"`
	TrxCode      string `json:"TrxCode"`
	CPI          string `json:"CPI"`
	Size         string `json:"Size"`
}

type QRCodeResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	QRCode              string `json:"QRCode"`
}

func (s *Service) GenerateQRCode(req QRCodeRequest) (*QRCodeResponse, error) {
	if req.MerchantName == "" {
		return nil, fmt.Errorf("merchant name is required")
	}
	if req.RefNo == "" {
		return nil, fmt.Errorf("reference number is required")
	}
	if req.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}
	if req.TrxCode == "" {
		return nil, fmt.Errorf("transaction code is required")
	}
	if req.CPI == "" {
		return nil, fmt.Errorf("CPI (Customer Phone/Till/Paybill) is required")
	}

	if req.Size == "" {
		req.Size = "300"
	}

	respBody, err := s.makeRequest(http.MethodPost, qrCodeURL, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	var response QRCodeResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to parse QR code response: %w", err)
	}
	return &response, nil
}
