package mpesa

import (
	"github.com/nutcas3/payment-rails/mpesa/pkg/daraja"
)

// InitiateStkPush initiates an STK push request using a parameter struct
func (c *Client) InitiateStkPush(params StkPushParams) (*daraja.STKPushResponse, error) {
	body := daraja.STKPushBody{
		BusinessShortCode: params.BusinessShortCode,
		TransactionType:   params.TransactionType,
		Amount:            params.Amount,
		PartyA:            params.PartyA,
		PartyB:            params.PartyB,
		PhoneNumber:       params.PhoneNumber,
		CallBackURL:       params.CallBackURL,
		AccountReference:  params.AccountReference,
		TransactionDesc:   params.TransactionDesc,
	}

	return c.Service.InitiateStkPush(body)
}

// LegacyInitiateStkPush is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyInitiateStkPush(businessShortCode, transactionType, amount, partyA, partyB, phoneNumber, callBackURL, accountReference, transactionDesc string) (*daraja.STKPushResponse, error) {
	return c.InitiateStkPush(StkPushParams{
		BusinessShortCode: businessShortCode,
		TransactionType:   transactionType,
		Amount:            amount,
		PartyA:            partyA,
		PartyB:            partyB,
		PhoneNumber:       phoneNumber,
		CallBackURL:       callBackURL,
		AccountReference:  accountReference,
		TransactionDesc:   transactionDesc,
	})
}

func (c *Client) QueryStkPush(businessShortCode, checkoutRequestID string) (*daraja.STKPushQueryResponse, error) {
	return c.Service.QueryStkPush(businessShortCode, checkoutRequestID)
}

func (c *Client) C2BRegisterURL(shortCode, responseType, confirmationURL, validationURL string) (*daraja.RegisterC2BURLResponse, error) {
	body := daraja.RegisterC2BURLBody{
		ShortCode:       shortCode,
		ResponseType:    responseType,
		ConfirmationURL: confirmationURL,
		ValidationURL:   validationURL,
	}

	return c.Service.C2BRegisterURL(body)
}

func (c *Client) C2BSimulate(shortCode int, commandID string, amount int, msisdn int, billRefNumber string) (*daraja.C2BSimulateResponse, error) {
	body := daraja.C2BSimulateRequestBody{
		ShortCode:     shortCode,
		CommandID:     commandID,
		Amount:        amount,
		Msisdn:        msisdn,
		BillRefNumber: billRefNumber,
	}

	return c.Service.C2BSimulate(body)
}

// B2CPayment performs a business to customer payment using a parameter struct
func (c *Client) B2CPayment(params B2CPaymentParams) (*daraja.B2CResponse, error) {
	body := daraja.B2CRequestBody{
		InitiatorName:      params.InitiatorName,
		SecurityCredential: params.SecurityCredential,
		CommandID:          params.CommandID,
		Amount:             params.Amount,
		PartyA:             params.PartyA,
		PartyB:             params.PartyB,
		Remarks:            params.Remarks,
		QueueTimeOutURL:    params.QueueTimeOutURL,
		ResultURL:          params.ResultURL,
		Occassion:          params.Occasion,
	}

	return c.Service.B2CPayment(body)
}

// LegacyB2CPayment is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyB2CPayment(initiatorName, securityCredential, commandID string, amount, partyA, partyB int, remarks, queueTimeOutURL, resultURL, occasion string) (*daraja.B2CResponse, error) {
	return c.B2CPayment(B2CPaymentParams{
		InitiatorName:      initiatorName,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		Amount:             amount,
		PartyA:             partyA,
		PartyB:             partyB,
		Remarks:            remarks,
		QueueTimeOutURL:    queueTimeOutURL,
		ResultURL:          resultURL,
		Occasion:           occasion,
	})
}

// B2BPayment performs a business to business payment using a parameter struct
func (c *Client) B2BPayment(params B2BPaymentParams) (*daraja.BusinessToBusinessResponse, error) {
	body := daraja.BusinessToBusinessRequestBody{
		Initiator:             params.Initiator,
		SecurityCredential:    params.SecurityCredential,
		CommandID:             params.CommandID,
		SenderIdentifierType:  params.SenderIdentifierType,
		RecieverIdentifierType: params.ReceiverIdentifierType,
		Amount:                params.Amount,
		PartyA:                params.PartyA,
		PartyB:                params.PartyB,
		AccountReference:      params.AccountReference,
		Requester:             params.Requester,
		Remarks:               params.Remarks,
		QueueTimeOutURL:       params.QueueTimeOutURL,
		ResultURL:             params.ResultURL,
	}

	return c.Service.BusinessToBusinessPayment(body)
}

// LegacyB2BPayment is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyB2BPayment(initiator, securityCredential, commandID, senderIdentifierType, receiverIdentifierType, amount, partyA, partyB, accountReference, requester, remarks, queueTimeOutURL, resultURL string) (*daraja.BusinessToBusinessResponse, error) {
	return c.B2BPayment(B2BPaymentParams{
		Initiator:              initiator,
		SecurityCredential:     securityCredential,
		CommandID:              commandID,
		SenderIdentifierType:   senderIdentifierType,
		ReceiverIdentifierType: receiverIdentifierType,
		Amount:                 amount,
		PartyA:                 partyA,
		PartyB:                 partyB,
		AccountReference:       accountReference,
		Requester:              requester,
		Remarks:                remarks,
		QueueTimeOutURL:        queueTimeOutURL,
		ResultURL:              resultURL,
	})
}

// TransactionStatus checks the status of a transaction using a parameter struct
func (c *Client) TransactionStatus(params TransactionStatusParams) (*daraja.TransactionStatusResponse, error) {
	body := daraja.TransactionStatusRequestBody{
		Initiator:          params.Initiator,
		SecurityCredential: params.SecurityCredential,
		CommandID:          params.CommandID,
		TransactionID:      params.TransactionID,
		PartyA:             params.PartyA,
		IdentifierType:     params.IdentifierType,
		ResultURL:          params.ResultURL,
		QueueTimeOutURL:    params.QueueTimeOutURL,
		Remarks:            params.Remarks,
		Occassion:          params.Occasion,
	}

	return c.Service.TransactionStatus(body)
}

// LegacyTransactionStatus is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyTransactionStatus(initiator, securityCredential, commandID, transactionID string, partyA, identifierType int, resultURL, queueTimeOutURL, remarks, occasion string) (*daraja.TransactionStatusResponse, error) {
	return c.TransactionStatus(TransactionStatusParams{
		Initiator:          initiator,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		TransactionID:      transactionID,
		PartyA:             partyA,
		IdentifierType:     identifierType,
		ResultURL:          resultURL,
		QueueTimeOutURL:    queueTimeOutURL,
		Remarks:            remarks,
		Occasion:           occasion,
	})
}

// AccountBalance checks the account balance using a parameter struct
func (c *Client) AccountBalance(params AccountBalanceParams) (*daraja.AccountBalanceResponse, error) {
	body := daraja.AccountBalanceRequestBody{
		Initiator:          params.Initiator,
		SecurityCredential: params.SecurityCredential,
		CommandID:          params.CommandID,
		PartyA:             params.PartyA,
		IdentifierType:     params.IdentifierType,
		Remarks:            params.Remarks,
		QueueTimeOutURL:    params.QueueTimeOutURL,
		ResultURL:          params.ResultURL,
	}

	return c.Service.AccountBalance(body)
}

// LegacyAccountBalance is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyAccountBalance(initiator, securityCredential, commandID string, partyA, identifierType int, remarks, queueTimeOutURL, resultURL string) (*daraja.AccountBalanceResponse, error) {
	return c.AccountBalance(AccountBalanceParams{
		Initiator:          initiator,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		PartyA:             partyA,
		IdentifierType:     identifierType,
		Remarks:            remarks,
		QueueTimeOutURL:    queueTimeOutURL,
		ResultURL:          resultURL,
	})
}

// Reversal reverses a transaction using a parameter struct
func (c *Client) Reversal(params ReversalParams) (*daraja.ReversalResponse, error) {
	body := daraja.ReversalRequestBody{
		Initiator:              params.Initiator,
		SecurityCredential:     params.SecurityCredential,
		CommandID:              params.CommandID,
		TransactionID:          params.TransactionID,
		Amount:                 params.Amount,
		ReceiverParty:          params.ReceiverParty,
		RecieverIdentifierType: params.ReceiverIdentifierType,
		ResultURL:              params.ResultURL,
		QueueTimeOutURL:        params.QueueTimeOutURL,
		Remarks:                params.Remarks,
		Occasion:               params.Occasion,
	}

	return c.Service.Reversal(body)
}

// LegacyReversal is the original method with multiple parameters (kept for backward compatibility)
func (c *Client) LegacyReversal(initiator, securityCredential, commandID, transactionID string, amount int, receiverParty, receiverIdentifierType int, resultURL, queueTimeOutURL, remarks, occasion string) (*daraja.ReversalResponse, error) {
	return c.Reversal(ReversalParams{
		Initiator:              initiator,
		SecurityCredential:     securityCredential,
		CommandID:              commandID,
		TransactionID:          transactionID,
		Amount:                 amount,
		ReceiverParty:          receiverParty,
		ReceiverIdentifierType: receiverIdentifierType,
		ResultURL:              resultURL,
		QueueTimeOutURL:        queueTimeOutURL,
		Remarks:                remarks,
		Occasion:               occasion,
	})
}
