package collection_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/nutcas3/payment-rails/momo/collection"
	"github.com/nutcas3/payment-rails/momo/common/mocks"
	"github.com/nutcas3/payment-rails/momo/common/types"
	"github.com/stretchr/testify/mock"
)

func TestCreatePayments(t *testing.T) {
	input := types.PaymentInput{
		ExternalTransactionID: gofakeit.UUID(),
		Money: types.Money{
			Amount:   "100",
			Currency: types.PAB,
		},
		CustomerReference:       gofakeit.UUID(),
		ServiceProviderUsername: gofakeit.Username(),
		CouponID:                gofakeit.UUID(),
		ProductID:               gofakeit.UUID(),
		ProductOfferingID:       gofakeit.UUID(),
		ReceiverMessage:         gofakeit.BeerName(),
		SenderNote:              gofakeit.BeerName(),
		MaxNumberOfRetries:      int(gofakeit.Int64()),
		IncludeSenderCharges:    gofakeit.Bool(),
	}

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		callback string
		body     types.PaymentInput
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create payments",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.PaymentInput"), nil).Return(nil).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					callback: gofakeit.URL(),
					body:     input,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to create payment",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.PaymentInput"), nil).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					callback: gofakeit.URL(),
					body:     input,
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: fail to get headers",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					callback: gofakeit.URL(),
					body:     input,
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil uuid passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:      ctx,
					id:       uuid.Nil,
					callback: gofakeit.URL(),
					body:     input,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBackend := mocks.NewMockBackend(t)
			mockCache := mocks.NewMockCacheStore(t)

			mh := &mockHandler{
				Backend: mockBackend,
				Cache:   mockCache,
			}

			args := tt.setup(mh)

			c := collection.NewCollection(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			err := c.CreatePayments(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCreatePayments() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPaymentStatus(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get payment status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.PaymentStatus")).Return(nil).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get payment status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.PaymentStatus")).Return(errors.New("an error")).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: fail to get headers",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil uuid passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx: ctx,
					id:  uuid.Nil,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBackend := mocks.NewMockBackend(t)
			mockCache := mocks.NewMockCacheStore(t)

			mh := &mockHandler{
				Backend: mockBackend,
				Cache:   mockCache,
			}

			args := tt.setup(mh)

			c := collection.NewCollection(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := c.GetPaymentStatus(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetPaymentStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
