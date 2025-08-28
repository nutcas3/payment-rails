package collection_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/mock"

	"github.com/nutcas3/payment-rails/momo/collection"
	"github.com/nutcas3/payment-rails/momo/common/mocks"
	"github.com/nutcas3/payment-rails/momo/common/types"
)

type mockHandler struct {
	Backend *mocks.MockBackend
	Cache   *mocks.MockCacheStore
}

func TestGetAccountBalance(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get account balance",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Balance")).Return(nil).Once()

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: failed getting account balance",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Balance")).Return(errors.New("an error")).Once()

				return args{
					ctx: ctx,
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: failed getting headers",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx: ctx,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBackend := mocks.NewMockBackend(t)
			mockCache := mocks.NewMockCacheStore(t)

			mh := mockHandler{
				Backend: mockBackend,
				Cache:   mockCache,
			}

			args := tt.setup(&mh)

			c := collection.NewCollection(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := c.GetAccountBalance(args.ctx)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetAccountBalance() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAccountBalanceInSpecificCurrency(t *testing.T) {
	type args struct {
		ctx      context.Context
		currency types.Currency
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get account balance",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Balance")).Return(nil).Once()

				return args{
					ctx:      ctx,
					currency: types.KES,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: failed getting account balance",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Balance")).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					currency: types.USD,
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: failed getting headers",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					currency: types.AWG,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBackend := mocks.NewMockBackend(t)
			mockCache := mocks.NewMockCacheStore(t)

			mh := mockHandler{
				Backend: mockBackend,
				Cache:   mockCache,
			}

			args := tt.setup(&mh)

			c := collection.NewCollection(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := c.GetAccountBalanceInSpecificCurrency(args.ctx, args.currency)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("GetAccountBalanceInSpecificCurrency() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
