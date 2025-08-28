package absa

import (
	"fmt"
	"net/http"
	"github.com/nutcas3/payment-rails/absa/pkg/api"
)

type Client struct {
	apiClient *api.Client
	webhookHandler *api.WebhookHandler
}

func NewClient(clientID, clientSecret, apiKey, environment string) (*Client, error) {
	apiClient, err := api.NewClient(clientID, clientSecret, apiKey, environment)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiClient: apiClient,
		webhookHandler: nil,
	}, nil
}

func (c *Client) SetWebhookSecret(webhookSecret string) {
	c.webhookHandler = api.NewWebhookHandler(webhookSecret)
}

func (c *Client) HandleWebhook(w http.ResponseWriter, r *http.Request, handlers api.WebhookHandlers) error {
	if c.webhookHandler == nil {
		return fmt.Errorf("webhook handler not initialized, call SetWebhookSecret first")
	}
	c.webhookHandler.HandleWebhook(w, r, handlers)
	return nil
}

func (c *Client) GetAccountBalance(req api.AccountBalanceRequest) (*api.AccountBalanceResponse, error) {
	return c.apiClient.GetAccountBalance(req)
}

func (c *Client) GetMiniStatement(req api.MiniStatementRequest) (*api.MiniStatementResponse, error) {
	return c.apiClient.GetMiniStatement(req)
}

func (c *Client) GetFullStatement(req api.FullStatementRequest) (*api.FullStatementResponse, error) {
	return c.apiClient.GetFullStatement(req)
}

func (c *Client) ValidateAccount(req api.AccountValidateRequest) (*api.AccountValidateResponse, error) {
	return c.apiClient.ValidateAccount(req)
}

func (c *Client) SendMoney(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendMoney(req)
}

func (c *Client) SendToMobileWallet(req api.MobileWalletRequest) (*api.MobileWalletResponse, error) {
	return c.apiClient.SendToMobileWallet(req)
}

func (c *Client) SendInternalBankTransfer(req api.SendMoneyRequest) (*api.SendMoneyResponse, error) {
	return c.apiClient.SendInternalBankTransfer(req)
}

func (c *Client) PayBill(req api.BillPaymentRequest) (*api.BillPaymentResponse, error) {
	return c.apiClient.PayBill(req)
}

func (c *Client) ReceiveMoney(req api.ReceiveMoneyRequest) (*api.ReceiveMoneyResponse, error) {
	return c.apiClient.ReceiveMoney(req)
}

func (c *Client) QueryTransaction(req api.TransactionQueryRequest) (*api.TransactionQueryResponse, error) {
	return c.apiClient.QueryTransaction(req)
}

func (c *Client) PurchaseAirtime(req api.AirtimePurchaseRequest) (*api.AirtimePurchaseResponse, error) {
	return c.apiClient.PurchaseAirtime(req)
}

func GenerateReference() string {
	return api.GenerateReference()
}


func (c *Client) ProcessBulkPayment(req api.BulkPaymentRequest) (*api.BulkPaymentResponse, error) {
	return c.apiClient.ProcessBulkPayment(req)
}

func (c *Client) GetBulkPaymentStatus(req api.BulkPaymentStatusRequest) (*api.BulkPaymentStatusResponse, error) {
	return c.apiClient.GetBulkPaymentStatus(req)
}


func (c *Client) CreateStandingOrder(req api.StandingOrderRequest) (*api.StandingOrderResponse, error) {
	return c.apiClient.CreateStandingOrder(req)
}

func (c *Client) GetStandingOrderStatus(req api.StandingOrderStatusRequest) (*api.StandingOrderStatusResponse, error) {
	return c.apiClient.GetStandingOrderStatus(req)
}

func (c *Client) CancelStandingOrder(req api.StandingOrderCancelRequest) (*api.StandingOrderCancelResponse, error) {
	return c.apiClient.CancelStandingOrder(req)
}

func (c *Client) ListStandingOrders(req api.StandingOrderListRequest) (*api.StandingOrderListResponse, error) {
	return c.apiClient.ListStandingOrders(req)
}


func (c *Client) CreateBeneficiary(req api.BeneficiaryCreateRequest) (*api.BeneficiaryCreateResponse, error) {
	return c.apiClient.CreateBeneficiary(req)
}

func (c *Client) ListBeneficiaries(req api.BeneficiaryListRequest) (*api.BeneficiaryListResponse, error) {
	return c.apiClient.ListBeneficiaries(req)
}

func (c *Client) GetBeneficiary(req api.BeneficiaryGetRequest) (*api.BeneficiaryGetResponse, error) {
	return c.apiClient.GetBeneficiary(req)
}

func (c *Client) UpdateBeneficiary(req api.BeneficiaryUpdateRequest) (*api.BeneficiaryUpdateResponse, error) {
	return c.apiClient.UpdateBeneficiary(req)
}

func (c *Client) DeleteBeneficiary(req api.BeneficiaryDeleteRequest) (*api.BeneficiaryDeleteResponse, error) {
	return c.apiClient.DeleteBeneficiary(req)
}


func (c *Client) GetForexRate(req api.ForexRateRequest) (*api.ForexRateResponse, error) {
	return c.apiClient.GetForexRate(req)
}

func (c *Client) ProcessForexTransfer(req api.ForexTransferRequest) (*api.ForexTransferResponse, error) {
	return c.apiClient.ProcessForexTransfer(req)
}

func (c *Client) GetForexTransferStatus(req api.ForexTransferStatusRequest) (*api.ForexTransferStatusResponse, error) {
	return c.apiClient.GetForexTransferStatus(req)
}


func (c *Client) RequestOTP(req api.OTPRequest) (*api.OTPResponse, error) {
	return c.apiClient.RequestOTP(req)
}

func (c *Client) VerifyOTP(req api.OTPVerifyRequest) (*api.OTPVerifyResponse, error) {
	return c.apiClient.VerifyOTP(req)
}

func (c *Client) AuthenticateTransaction(req api.TransactionAuthRequest) (*api.TransactionAuthResponse, error) {
	return c.apiClient.AuthenticateTransaction(req)
}

func (c *Client) RegisterDevice(req api.DeviceRegistrationRequest) (*api.DeviceRegistrationResponse, error) {
	return c.apiClient.RegisterDevice(req)
}
