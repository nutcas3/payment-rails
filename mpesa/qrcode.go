package mpesa

import (
	"payment-rails/mpesa/pkg/daraja"
)

type QRCodeRequest struct {
	MerchantName string
	RefNo        string
	Amount       int
	TrxCode      string
	CPI          string
	Size         string
}

type QRCodeResponse struct {
	ResponseCode        string
	ResponseDescription string
	QRCode              string
}

const (
	TrxCodeBuyGoods      = "BG" // Buy Goods and Services
	TrxCodeWithdrawCash  = "WA" // Withdraw Cash at Agent Till
	TrxCodePaybill       = "PB" // Paybill or Business Number
	TrxCodeSendMoney     = "SM" // Send Money (Mobile Number)
	TrxCodeSendBusiness  = "SB" // Sent to Business. Business Buy Goods
)

func (c *Client) GenerateQRCode(req QRCodeRequest) (*QRCodeResponse, error) {
	internalReq := daraja.QRCodeRequest{
		MerchantName: req.MerchantName,
		RefNo:        req.RefNo,
		Amount:       req.Amount,
		TrxCode:      req.TrxCode,
		CPI:          req.CPI,
		Size:         req.Size,
	}

	resp, err := c.Service.GenerateQRCode(internalReq)
	if err != nil {
		return nil, err
	}

	return &QRCodeResponse{
		ResponseCode:        resp.ResponseCode,
		ResponseDescription: resp.ResponseDescription,
		QRCode:              resp.QRCode,
	}, nil
}
