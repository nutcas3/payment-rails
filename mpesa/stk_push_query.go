package mpesa

type STKPushQueryRequest struct {
	BusinessShortCode string // The organization's shortcode (Paybill or Buygoods)
	CheckoutRequestID string // The global unique identifier of the processed checkout transaction request
}

type STKPushQueryResponse struct {
	ResponseCode        string // Status code of the request submission (0 means success)
	ResponseDescription string // Status message of the request submission
	MerchantRequestID   string // Global unique identifier for the payment request
	CheckoutRequestID   string // Global unique identifier of the processed checkout transaction
	ResultCode          string // Status code of the transaction processing (0 means success)
	ResultDesc          string // Status message of the transaction processing
}

func (c *Client) QueryStkPushStatus(req STKPushQueryRequest) (*STKPushQueryResponse, error) {
	resp, err := c.Service.QueryStkPush(req.BusinessShortCode, req.CheckoutRequestID)
	if err != nil {
		return nil, err
	}

	return &STKPushQueryResponse{
		ResponseCode:        resp.ResponseCode,
		ResponseDescription: resp.ResponseDescription,
		MerchantRequestID:   resp.MerchantRequestID,
		CheckoutRequestID:   resp.CheckoutRequestID,
		ResultCode:          resp.ResultCode,
		ResultDesc:          resp.ResultDesc,
	}, nil
}
