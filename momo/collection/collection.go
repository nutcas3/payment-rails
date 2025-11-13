package collection

import (
	"context"

	"github.com/nutcas3/payment-rails/momo/common"
	"github.com/nutcas3/payment-rails/momo/common/types"

	"github.com/google/uuid"
)

const (
	subHeader      = "Ocp-Apim-Subscription-Key"
	envHeader      = "X-Target-Environment"
	callbackHeader = "X-Callback-Url"
	authHeader     = "Authorization"
	refHeader      = "X-Reference-Id"
	contentHeader  = "application/json"
)

type Collection struct {
	subscriptionKey, apiKey, apiSecret, environment string
	backend                                         common.Backend
	cache                                           common.CacheStore
}

type Service interface {
	GetAccountBalance(ctx context.Context) (*types.Balance, error)
	GetAccountBalanceInSpecificCurrency(ctx context.Context, currency types.Currency) (*types.Balance, error)
	CreateInvoice(ctx context.Context, refID uuid.UUID, callbackURL string, body types.CreateInvoiceInput) error
	CancelInvoice(ctx context.Context, invoiceID uuid.UUID, transactionID uuid.UUID, callbackURL string) error
	GetInvoiceStatus(ctx context.Context, refID uuid.UUID) (*types.InvoiceStatus, error)
	CreatePayments(ctx context.Context, refID uuid.UUID, callbackURL string, body types.PaymentInput) error
	GetPaymentStatus(ctx context.Context, refID uuid.UUID) (*types.PaymentStatus, error)
	CancelPreApproval(ctx context.Context, id uuid.UUID) error
	GetPreApprovalStatus(ctx context.Context, refID uuid.UUID) (*types.PreApprovalStatus, error)
	PreApproval(ctx context.Context, refID uuid.UUID, callbackURL string, body types.PreApprovalInput) error
	GetApprovedPreApprovals(ctx context.Context, accHolderIDType string, accHolderID string) ([]*types.PreApprovalDetails, error)
	RequestToPay(ctx context.Context, refID uuid.UUID, callbackURL string, handleStatusPolling bool, body types.RequestToPayInput) (*types.RequestToPayStatus, error)
	RequestToPayTransactionStatus(ctx context.Context, refID uuid.UUID) (*types.RequestToPayStatus, error)
	RequestToPayDeliveryNotification(ctx context.Context, refID uuid.UUID, message string, language string) (*types.DeliveryNotification, error)
	RequestToWithdrawTransactionStatus(ctx context.Context, refID uuid.UUID) (*types.RequestToPayStatus, error)
	RequestToWithdrawV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RequestToPayInput) error
	RequestToWithdrawV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RequestToPayInput) error
	CreateAccessToken(ctx context.Context) (string, error)
	CreateOauth2Token(ctx context.Context) (string, error)
	BcAuthorize(ctx context.Context, callbackURL string) (string, error)
	ValidateAccountHolderStatus(ctx context.Context, accHolderID, accHolderIDType string) (bool, error)
	GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error)
	GetBasicUserInfo(ctx context.Context, accHolderIDType, accHolderID string) (*types.BasicUserInfo, error)
}

func NewCollection(
	subKey, apiKey, apiSecret, env string, backend common.Backend, cache common.CacheStore,
) Service {
	return &Collection{
		subscriptionKey: subKey,
		apiKey:          apiKey,
		apiSecret:       apiSecret,
		environment:     env,
		backend:         backend,
		cache:           cache,
	}
}
