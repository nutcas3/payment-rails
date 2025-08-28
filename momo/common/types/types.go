package types

import (
	"errors"
)

var ErrRefIDRequired = errors.New("refID is required")

// Currency represents ISO4217 currency codes
//
// https://en.wikipedia.org/wiki/ISO_4217
type Currency string

const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BOV Currency = "BOV"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHE Currency = "CHE"
	CHF Currency = "CHF"
	CHW Currency = "CHW"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	COU Currency = "COU"
	CRC Currency = "CRC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRU Currency = "MRU"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MXV Currency = "MXV"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLE Currency = "SLE"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	SSP Currency = "SSP"
	STN Currency = "STN"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	USN Currency = "USN"
	UYI Currency = "UYI"
	UYU Currency = "UYU"
	UYW Currency = "UYW"
	UZS Currency = "UZS"
	VED Currency = "VED"
	VES Currency = "VES"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAD Currency = "XAD"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XBA Currency = "XBA"
	XBB Currency = "XBB"
	XBC Currency = "XBC"
	XBD Currency = "XBD"
	XCD Currency = "XCD"
	XCG Currency = "XCG"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPD Currency = "XPD"
	XPF Currency = "XPF"
	XPT Currency = "XPT"
	XSU Currency = "XSU"
	XTS Currency = "XTS"
	XUA Currency = "XUA"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMW Currency = "ZMW"
	ZWG Currency = "ZWG"
)

type Balance struct {
	AvailableBalance string   `json:"availableBalance"`
	Currency         Currency `json:"currency"`
}

// bc-authorize token response
//
//	{
//	    "auth_req_id": "string",
//	    "interval": 0,
//	    "expires_in": 0
//	}
type BcAuthResp struct {
	AuthRequestID string `json:"auth_req_id"`
	Interval      int    `json:"interval"`
	ExpiresIn     int    `json:"expires_in"`
}

// Oauth2 token response
//
//	{
//	    "access_token": "string",
//	    "token_type": "string",
//	    "expires_in": 0,
//	    "scope": "string",
//	    "refresh_token": "string",
//	    "refresh_token_expired_in": 0
//	}
type Oauth2Resp struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	Scope                 string `json:"scope"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiredIn int    `json:"refresh_token_expired_in"`
}

// Access token response
//
//	{
//	    "access_token": "string",
//	    "token_type": "string",
//	    "expires_in": 0
//	}
type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// PartyIDType represents the types for party IDs
type PartyIDType string

const (
	MSISDN    PartyIDType = "MSISDN"     // Valid mobile number according to ITU-T E.164
	EMAIL     PartyIDType = "EMAIL"      // Valid email address
	PARTYCODE PartyIDType = "PARTY_CODE" // UUID
)

type Party struct {
	PartyIDType PartyIDType `json:"partyIdType"`
	PartyID     string      `json:"partyId"`
}

type ErrorReason struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

//	{
//	    "sub": "string",
//	    "name": "string",
//	    "given_name": "string",
//	    "family_name": "string",
//	    "middle_name": "string",
//	    "email": "string",
//	    "email_verified": true,
//	    "gender": "string",
//	    "locale": "string",
//	    "phone_number": "string",
//	    "phone_number_verified": true,
//	    "address": "string",
//	    "updated_at": 0,
//	    "status": "string",
//	    "birthdate": "string",
//	    "credit_score": "string",
//	    "active": true,
//	    "country_of_birth": "string",
//	    "region_of_birth": "string",
//	    "city_of_birth": "string",
//	    "occupation": "string",
//	    "employer_name": "string",
//	    "identification_type": "string",
//	    "identification_value": "string"
//	}
type UserConsentInfo struct {
	Subject             string  `json:"sub"`
	Name                string  `json:"name"`
	GivenName           string  `json:"given_name"`
	FamilyName          string  `json:"family_name"`
	MiddleName          string  `json:"middle_name"`
	Email               string  `json:"email"`
	EmailVerified       bool    `json:"email_verified"`
	Gender              string  `json:"gender"`
	Locale              string  `json:"locale"`
	PhoneNumber         string  `json:"phone_number"`
	PhoneNumberVerified bool    `json:"phone_number_verified"`
	Address             string  `json:"address"`
	UpdatedAt           float64 `json:"updated_at"`
	Status              string  `json:"status"`
	BirthDate           string  `json:"birthdate"`
	CreditScore         string  `json:"credit_score"`
	Active              bool    `json:"active"`
	CountryOfBirth      string  `json:"country_of_birth"`
	RegionOfBirth       string  `json:"region_of_birth"`
	CityOfBirth         string  `json:"city_of_birth"`
	Occupation          string  `json:"occupation"`
	EmployeeName        string  `json:"employee_name"`
	IdentificationType  string  `json:"identification_type"`
	IdentificationValue string  `json:"identification_value"`
}

//	{
//	    "given_name": "string",
//	    "family_name": "string",
//	    "birthdate": "string",
//	    "locale": "string",
//	    "gender": "string",
//	    "status": "string"
//	}
type BasicUserInfo struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	BirthDate  string `json:"birthdate"`
	Locale     string `json:"locale"`
	Gender     string `json:"gender"`
	Status     string `json:"status"`
}

//	{
//	    "amount": "string",
//	    "currency": "string",
//	    "payee": {
//	        "partyId": "string",
//	        "partyIdType": "MSISDN/EMAIL/PARTY_CODE"
//	    },
//	    "externalId": "string",
//	    "orginatingCountry": "string",
//	    "originalAmount": "string",
//	    "originalCurrency": "string",
//	    "payerMessage": "string",
//	    "payeeNote": "string",
//	    "payerIdentificationType": "PersonIdentificationType1Code",
//	    "payerIdentificationNumber": "string",
//	    "payerIdentity": "string",
//	    "payerFirstName": "string",
//	    "payerSurName": "string",
//	    "payerLanguageCode": "string",
//	    "payerEmail": "string (Email)",
//	    "payerMsisdn": "string (Msisdn)",
//	    "payerGender": "string"
//	}
type CashTransferInput struct {
	Amount                    string   `json:"amount"`
	Currency                  Currency `json:"currency"`
	Payee                     Party    `json:"payee"`
	ExternalID                string   `json:"externalId"`
	OriginatingCountry        string   `json:"originatingCountry"`
	OriginalAmount            string   `json:"originalAmount"`
	PayerMessage              string   `json:"payerMessage"`
	PayeeNote                 string   `json:"payeeNote"`
	PayerIdentificationType   string   `json:"payerIdentificationType"`
	PayerIdentificationNumber string   `json:"payerIdentificationNumber"`
	PayerIdentity             string   `json:"payerIdentity"`
	PayerFirstName            string   `json:"payerFirstName"`
	PayerSurname              string   `json:"payerSurName"`
	PayerLanguageCode         string   `json:"payerLanguageCode"`
	PayerEmail                string   `json:"payerEmail"`
	PayerMSISDN               string   `json:"payerMsisdn"`
	PayerGender               string   `json:"payerGender"`
}

//	{
//		"financialTransactionId" "string",
//		"status": "string",
//		"reason": "string"
//	    "amount": "string",
//	    "currency": "string",
//	    "payee": {
//	        "partyId": "string",
//	        "partyIdType": "MSISDN/EMAIL/PARTY_CODE"
//	    },
//	    "externalId": "string",
//	    "orginatingCountry": "string",
//	    "originalAmount": "string",
//	    "originalCurrency": "string",
//	    "payerMessage": "string",
//	    "payeeNote": "string",
//	    "payerIdentificationType": "PersonIdentificationType1Code",
//	    "payerIdentificationNumber": "string",
//	    "payerIdentity": "string",
//	    "payerFirstName": "string",
//	    "payerSurName": "string",
//	    "payerLanguageCode": "string",
//	    "payerEmail": "string (Email)",
//	    "payerMsisdn": "string (Msisdn)",
//	    "payerGender": "string"
//	}
type CashTransferStatus struct {
	FinancialTransactionId    string   `json:"financialTransactionId"`
	Status                    string   `json:"status"`
	Reason                    string   `json:"reason"`
	Amount                    string   `json:"amount"`
	Currency                  Currency `json:"currency"`
	Payee                     Party    `json:"payee"`
	ExternalID                string   `json:"externalId"`
	OriginatingCountry        string   `json:"originatingCountry"`
	OriginalAmount            string   `json:"originalAmount"`
	PayerMessage              string   `json:"payerMessage"`
	PayeeNote                 string   `json:"payeeNote"`
	PayerIdentificationType   string   `json:"payerIdentificationType"`
	PayerIdentificationNumber string   `json:"payerIdentificationNumber"`
	PayerIdentity             string   `json:"payerIdentity"`
	PayerFirstName            string   `json:"payerFirstName"`
	PayerSurname              string   `json:"payerSurName"`
	PayerLanguageCode         string   `json:"payerLanguageCode"`
	PayerEmail                string   `json:"payerEmail"`
	PayerMSISDN               string   `json:"payerMsisdn"`
	PayerGender               string   `json:"payerGender"`
}

//	{
//	    "amount": "string",
//	    "currency": "string",
//	    "externalId": "string",
//	    "payee": {
//	        "partyIdType": "MSISDN",
//	        "partyId": "string"
//	    },
//	    "payerMessage": "string",
//	    "payeeNote": "string"
//	}
type TransferInput struct {
	Amount       string   `json:"amount"`
	Currency     Currency `json:"currency"`
	ExternalID   string   `json:"externalId"`
	Payee        Party    `json:"payee"`
	PayerMessage string   `json:"payerMessage"`
	PayeeNote    string   `json:"payeeNote"`
}

type DisbursementTransactionStatus struct {
	FinancialTransactionID string      `json:"financialTransactionId"`
	ExternalID             string      `json:"externalId"`
	Amount                 string      `json:"amount"`
	PayerMessage           string      `json:"payerMessage"`
	PayeeNote              string      `json:"payeeNote"`
	Status                 string      `json:"status"`
	Currency               Currency    `json:"currency"`
	Payee                  Party       `json:"payee"`
	Reason                 ErrorReason `json:"reason"`
}

//	{
//	    "amount": "string",
//	    "currency": "string",
//	    "externalId": "string",
//	    "payerMessage": "string",
//	    "payeeNote": "string",
//	    "referenceIdToRefund": "UUID-REQUEST-TO-PAY"
//	}
type RefundInput struct {
	Amount              string   `json:"amount"`
	Currency            Currency `json:"currency"`
	ExternalID          string   `json:"externalId"`
	PayerMessage        string   `json:"payerMessage"`
	PayeeNote           string   `json:"payeeNote"`
	ReferenceIDToRefund string   `json:"referenceIdToRefund"`
}

type TransferStatus struct {
	FinancialTransactionID string      `json:"financialTransactionId"`
	ExternalID             string      `json:"externalId"`
	Amount                 string      `json:"amount"`
	PayerMessage           string      `json:"payerMessage"`
	PayeeNote              string      `json:"payeeNote"`
	Status                 string      `json:"status"`
	Currency               Currency    `json:"currency"`
	Payee                  Party       `json:"payee"`
	Reason                 ErrorReason `json:"reason"`
}

// Response from the callback request or status polling
type RequestToPayStatus struct {
	FinancialTransactionID string      `json:"financialTransactionId"`
	ExternalID             string      `json:"externalId"`
	Amount                 string      `json:"amount"`
	PayerMessage           string      `json:"payerMessage"`
	PayeeNote              string      `json:"payeeNote"`
	Status                 string      `json:"status"`
	Currency               Currency    `json:"currency"`
	Payer                  Party       `json:"payer"`
	Reason                 ErrorReason `json:"reason"`
}

type RequestToPayInput struct {
	Amount       string   `json:"amount"`
	ExternalID   string   `json:"externalId"`
	PayerMessage string   `json:"partyMessage"`
	PayeeNote    string   `json:"payeeNote"`
	Currency     Currency `json:"currency"`
	Payer        Party    `json:"payer"`
}

type PreApprovalInput struct {
	Payer         Party    `json:"payer,omitempty"`
	PayerCurrency Currency `json:"payerCurrency,omitempty"`
	PayerMessage  string   `json:"payerMessage,omitempty"`
	ValidityTime  int      `json:"validityTime,omitempty"`
}

type PreApprovalStatus struct {
	Payer              Party       `json:"payer"`
	PayerCurrency      Currency    `json:"payerCurrency"`
	PayerMessage       string      `json:"payerMessage"`
	Status             string      `json:"status"`
	ExiprationDateTime string      `json:"expirationDateTime"`
	Reason             ErrorReason `json:"reason"`
}

type PaymentStatus struct {
	RefID                  string      `json:"referenceId"`
	Status                 string      `json:"status"`
	FinancialTransactionID string      `json:"financialTransactionId"`
	Reason                 ErrorReason `json:"reason"`
}

type Money struct {
	Amount   string   `json:"amount"`
	Currency Currency `json:"currency"`
}

type PaymentInput struct {
	ExternalTransactionID   string `json:"externalTransactionId"`
	Money                   Money  `json:"money"`
	CustomerReference       string `json:"customerReference"`
	ServiceProviderUsername string `json:"serviceProviderUserName"`
	CouponID                string `json:"couponId"`
	ProductID               string `json:"productId"`
	ProductOfferingID       string `json:"productOfferingId"`
	ReceiverMessage         string `json:"receiverMessage"`
	SenderNote              string `json:"senderNote"`
	MaxNumberOfRetries      int    `json:"maxNumberOfRetries"`
	IncludeSenderCharges    bool   `json:"includeSenderCharges"`
}

type InvoiceStatus struct {
	ReferenceID    string      `json:"referenceId,omitempty"`
	ExternalID     string      `json:"externalId,omitempty"`
	Amount         string      `json:"amount,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Status         string      `json:"status,omitempty"`
	PaymentRef     string      `json:"paymentReference,omitempty"`
	InvoiceID      string      `json:"invoiceID,omitempty"`
	ExpiryDateTime string      `json:"expiryDateTime,omitempty"`
	PayeeFirstName string      `json:"payeeFirstName,omitempty"`
	PayeeLastName  string      `json:"payeeLastName,omitempty"`
	Description    string      `json:"description,omitempty"`
	IntendedPayer  Party       `json:"intendedPayer"`
	ErrorReason    ErrorReason `json:"errorReason"`
}

type CreateInvoiceInput struct {
	ExternalID       string   `json:"externalID,omitempty"`
	Amount           string   `json:"amount,omitempty"`
	Currency         Currency `json:"currency,omitempty"`
	ValidityDuration string   `json:"validityDuration,omitempty"`
	IntendedPayer    Party    `json:"intendedPayer,omitempty"`
	Payee            Party    `json:"payee,omitempty"`
	Description      string   `json:"description,omitempty"`
}

type PreApprovalDetails struct {
	ID             string `json:"preApprovalId"`
	ToFri          string `json:"toFri"`
	FromFri        string `json:"fromFri"`
	FromCurrency   string `json:"fromCurrency"`
	CreatedTime    string `json:"createdTime"`
	ApprovedTime   string `json:"approvedTime"`
	ExpiryTime     string `json:"expiryTime"`
	Status         string `json:"status"`
	Message        string `json:"message"`
	Frequency      string `json:"frequency"`
	StartDate      string `json:"startDate"`
	LastUsedDate   string `json:"lastUsedDate"`
	Offer          string `json:"offer"`
	ExternalID     string `json:"externalID"`
	MaxDebitAmount string `json:"maxDebitAmount"`
}

type DeliveryNotification struct {
	NotificationMessage string `json:"notificationMessage"`
}
