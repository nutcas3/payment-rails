package disbursement

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
	contentHeader  = "Content-Type"
)

type Disbursement struct {
	subscriptionKey, apiKey, apiSecret, environment string
	backend                                         common.Backend
	cache                                           common.CacheStore
}

type Service interface {
	GetAccountBalance(ctx context.Context) (*types.Balance, error)
	GetAccountBalanceInSpecificCurrency(ctx context.Context, currency types.Currency) (*types.Balance, error)
	DepositV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error
	DepositV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error
	GetDepositStatus(ctx context.Context, refID uuid.UUID) (*types.DisbursementTransactionStatus, error)
	RefundV1(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RefundInput) error
	RefundV2(ctx context.Context, refID uuid.UUID, callbackURL string, body types.RefundInput) error
	GetRefundStatus(ctx context.Context, refID uuid.UUID) (*types.DisbursementTransactionStatus, error)
	CreateAccessToken(ctx context.Context) (string, error)
	CreateOauth2Token(ctx context.Context) (string, error)
	BcAuthorize(ctx context.Context, callbackURL string) (string, error)
	Transfer(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error
	GetTransferStatus(ctx context.Context, refID uuid.UUID) (*types.DisbursementTransactionStatus, error)
	GetBasicUserInfo(ctx context.Context, accHolderIdType, accHolderId string) (*types.BasicUserInfo, error)
	GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error)
	ValidateAccountHolderStatus(ctx context.Context, accHolderId, accHolderIdType string) (bool, error)
}

func NewDisbursement(
	subKey, apiKey, apiSecret, env string, backend common.Backend, cache common.CacheStore,
) Service {
	return &Disbursement{
		subscriptionKey: subKey,
		apiKey:          apiKey,
		apiSecret:       apiSecret,
		environment:     env,
		backend:         backend,
		cache:           cache,
	}
}
