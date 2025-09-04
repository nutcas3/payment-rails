package mpesa

type StkPushParams struct {
	BusinessShortCode string
	TransactionType   string
	Amount            string
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}

type C2BRegisterURLParams struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

type C2BSimulateParams struct {
	ShortCode     int
	CommandID     string
	Amount        int
	Msisdn        int
	BillRefNumber string
}

type B2CPaymentParams struct {
	InitiatorName      string
	SecurityCredential string
	CommandID          string
	Amount             int
	PartyA             int
	PartyB             int
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
	Occasion           string
}

type B2BPaymentParams struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	ReceiverIdentifierType string
	Amount                 string
	PartyA                 string
	PartyB                 string
	AccountReference       string
	Requester              string
	Remarks                string
	QueueTimeOutURL        string
	ResultURL              string
}

type TransactionStatusParams struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	TransactionID      string
	PartyA             int
	IdentifierType     int
	ResultURL          string
	QueueTimeOutURL    string
	Remarks            string
	Occasion           string
}

type AccountBalanceParams struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	PartyA             int
	IdentifierType     int
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
}

type ReversalParams struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	TransactionID          string
	Amount                 int
	ReceiverParty          int
	ReceiverIdentifierType int
	ResultURL              string
	QueueTimeOutURL        string
	Remarks                string
	Occasion               string
}
