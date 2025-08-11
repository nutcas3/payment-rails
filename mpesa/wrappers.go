package mpesa

import (
	"payment-rails/mpesa/pkg/daraja"
)

func (c *Client) InitiateStkPush(businessShortCode, transactionType, amount, partyA, partyB, phoneNumber, callBackURL, accountReference, transactionDesc string) (*daraja.STKPushResponse, error) {
	body := daraja.STKPushBody{
		BusinessShortCode: businessShortCode,
		TransactionType:   transactionType,
		Amount:            amount,
		PartyA:            partyA,
		PartyB:            partyB,
		PhoneNumber:       phoneNumber,
		CallBackURL:       callBackURL,
		AccountReference:  accountReference,
		TransactionDesc:   transactionDesc,
	}
	
	return c.Service.InitiateStkPush(body)
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

func (c *Client) B2CPayment(initiatorName, securityCredential, commandID string, amount, partyA, partyB int, remarks, queueTimeOutURL, resultURL, occasion string) (*daraja.B2CResponse, error) {
	body := daraja.B2CRequestBody{
		InitiatorName:      initiatorName,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		Amount:             amount,
		PartyA:             partyA,
		PartyB:             partyB,
		Remarks:            remarks,
		QueueTimeOutURL:    queueTimeOutURL,
		ResultURL:          resultURL,
		Occassion:          occasion,
	}
	
	return c.Service.B2CPayment(body)
}

func (c *Client) B2BPayment(initiator, securityCredential, commandID, senderIdentifierType, receiverIdentifierType, amount, partyA, partyB, accountReference, requester, remarks, queueTimeOutURL, resultURL string) (*daraja.BusinessToBusinessResponse, error) {
	body := daraja.BusinessToBusinessRequestBody{
		Initiator:             initiator,
		SecurityCredential:    securityCredential,
		CommandID:             commandID,
		SenderIdentifierType:  senderIdentifierType,
		RecieverIdentifierType: receiverIdentifierType,
		Amount:                amount,
		PartyA:                partyA,
		PartyB:                partyB,
		AccountReference:      accountReference,
		Requester:             requester,
		Remarks:               remarks,
		QueueTimeOutURL:       queueTimeOutURL,
		ResultURL:             resultURL,
	}
	
	return c.Service.BusinessToBusinessPayment(body)
}

func (c *Client) TransactionStatus(initiator, securityCredential, commandID, transactionID string, partyA, identifierType int, resultURL, queueTimeOutURL, remarks, occasion string) (*daraja.TransactionStatusResponse, error) {
	body := daraja.TransactionStatusRequestBody{
		Initiator:          initiator,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		TransactionID:      transactionID,
		PartyA:             partyA,
		IdentifierType:     identifierType,
		ResultURL:          resultURL,
		QueueTimeOutURL:    queueTimeOutURL,
		Remarks:            remarks,
		Occassion:          occasion,
	}
	
	return c.Service.TransactionStatus(body)
}

func (c *Client) AccountBalance(initiator, securityCredential, commandID string, partyA, identifierType int, remarks, queueTimeOutURL, resultURL string) (*daraja.AccountBalanceResponse, error) {
	body := daraja.AccountBalanceRequestBody{
		Initiator:          initiator,
		SecurityCredential: securityCredential,
		CommandID:          commandID,
		PartyA:             partyA,
		IdentifierType:     identifierType,
		Remarks:            remarks,
		QueueTimeOutURL:    queueTimeOutURL,
		ResultURL:          resultURL,
	}
	
	return c.Service.AccountBalance(body)
}

func (c *Client) Reversal(initiator, securityCredential, commandID, transactionID string, amount int, receiverParty, receiverIdentifierType int, resultURL, queueTimeOutURL, remarks, occasion string) (*daraja.ReversalResponse, error) {
	body := daraja.ReversalRequestBody{
		Initiator:              initiator,
		SecurityCredential:     securityCredential,
		CommandID:              commandID,
		TransactionID:          transactionID,
		Amount:                 amount,
		ReceiverParty:          receiverParty,
		RecieverIdentifierType: receiverIdentifierType,
		ResultURL:              resultURL,
		QueueTimeOutURL:        queueTimeOutURL,
		Remarks:                remarks,
		Occasion:               occasion,
	}
	
	return c.Service.Reversal(body)
}
