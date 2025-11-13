package momo

import "payment-rails/momo/pkg/api"

func (c *Client) RequestToPay(req api.RequestToPayRequest) (*api.RequestToPayResponse, error) {
	return c.API.RequestToPay(req)
}
func (c *Client) GetRequestToPayStatus(referenceID string) (*api.TransactionStatus, error) {
	return c.API.GetRequestToPayStatus(referenceID)
}
func (c *Client) GetAccountBalance() (*api.Balance, error) {
	return c.API.GetAccountBalance()
}
func (c *Client) ValidateAccountHolderStatus(msisdn string) (*api.AccountHolder, error) {
	return c.API.ValidateAccountHolderStatus(msisdn)
}
func (c *Client) GetBasicUserInfo(msisdn string) (*api.BasicUserInfo, error) {
	return c.API.GetBasicUserInfo(msisdn)
}
func (c *Client) Transfer(req api.TransferRequest) (*api.TransferResponse, error) {
	return c.API.Transfer(req)
}
func (c *Client) GetTransferStatus(referenceID string) (*api.TransactionStatus, error) {
	return c.API.GetTransferStatus(referenceID)
}
func (c *Client) GetDisbursementBalance() (*api.Balance, error) {
	return c.API.GetDisbursementBalance()
}
func (c *Client) ValidateDisbursementAccountHolder(msisdn string) (*api.AccountHolder, error) {
	return c.API.ValidateDisbursementAccountHolder(msisdn)
}
func (c *Client) GetDisbursementUserInfo(msisdn string) (*api.BasicUserInfo, error) {
	return c.API.GetDisbursementUserInfo(msisdn)
}

func (c *Client) Remit(req api.RemittanceRequest) (*api.RemittanceResponse, error) {
	return c.API.Remit(req)
}
func (c *Client) GetRemittanceStatus(referenceID string) (*api.TransactionStatus, error) {
	return c.API.GetRemittanceStatus(referenceID)
}
func (c *Client) GetRemittanceBalance() (*api.Balance, error) {
	return c.API.GetRemittanceBalance()
}
func (c *Client) ValidateRemittanceAccountHolder(msisdn string) (*api.AccountHolder, error) {
	return c.API.ValidateRemittanceAccountHolder(msisdn)
}

func (c *Client) GetRemittanceUserInfo(msisdn string) (*api.BasicUserInfo, error) {
	return c.API.GetRemittanceUserInfo(msisdn)
}

func (c *Client) Refund(req api.RefundRequest) (*api.RefundResponse, error) {
	return c.API.Refund(req)
}

func (c *Client) RefundV2(req api.RefundRequest) (*api.RefundResponse, error) {
	return c.API.RefundV2(req)
}

func (c *Client) GetRefundStatus(referenceID string) (*api.RefundStatus, error) {
	return c.API.GetRefundStatus(referenceID)
}
