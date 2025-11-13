package remittance

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

type Remittance struct {
	subscriptionKey, apiKey, apiSecret, environment string
	backend                                         common.Backend
	cache                                           common.CacheStore
}

type Service interface {
	GetAccountBalance(ctx context.Context) (*types.Balance, error)
	GetAccountBalanceInSpecificCurrency(ctx context.Context, currency types.Currency) (*types.Balance, error)
	CashTransfer(ctx context.Context, refID uuid.UUID, callbackURL string, body types.CashTransferInput) error
	GetCashTransferStatus(ctx context.Context, refID uuid.UUID) (*types.CashTransferStatus, error)
	Transfer(ctx context.Context, refID uuid.UUID, callbackURL string, body types.TransferInput) error
	GetTransferStatus(ctx context.Context, refID uuid.UUID) (*types.TransferStatus, error)
	CreateAccessToken(ctx context.Context) (string, error)
	CreateOauth2Token(ctx context.Context) (string, error)
	BcAuthorize(ctx context.Context, callbackURL string) (string, error)
	GetBasicUserInfo(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error)
	GetBasicUserInfov3(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error)
	GetBasicUserInfoClone(ctx context.Context, accHolderMsisdn string) (*types.BasicUserInfo, error)
	GetUserInfoWithConsent(ctx context.Context) (*types.UserConsentInfo, error)
	ValidateAccountHolderStatus(ctx context.Context, accHolderId, accHolderIdType string) (bool, error)
}

func NewRemittance(
	subKey, apiKey, apiSecret, env string, backend common.Backend, cache common.CacheStore,
) Service {
	return &Remittance{
		subscriptionKey: subKey,
		apiKey:          apiKey,
		apiSecret:       apiSecret,
		environment:     env,
		backend:         backend,
		cache:           cache,
	}
}
