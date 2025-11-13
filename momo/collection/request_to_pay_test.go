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

func TestRequestToPay(t *testing.T) {
	input := types.RequestToPayInput{
		Amount:       "100",
		ExternalID:   gofakeit.UUID(),
		PayerMessage: gofakeit.BeerName(),
		PayeeNote:    gofakeit.BeerName(),
		Currency:     types.YER,
		Payer: types.Party{
			PartyIDType: types.EMAIL,
			PartyID:     gofakeit.Email(),
		},
	}

	type args struct {
		ctx                 context.Context
		id                  uuid.UUID
		body                types.RequestToPayInput
		callback            string
		handleStatusPolling bool
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully request payment",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(nil).Once()

				return args{
					ctx:                 ctx,
					id:                  uuid.New(),
					body:                input,
					callback:            gofakeit.URL(),
					handleStatusPolling: false,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to request payment",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(errors.New("an error")).Once()

				return args{
					ctx:                 ctx,
					id:                  uuid.New(),
					body:                input,
					callback:            gofakeit.URL(),
					handleStatusPolling: gofakeit.Bool(),
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
					ctx:                 ctx,
					id:                  uuid.New(),
					body:                input,
					callback:            gofakeit.URL(),
					handleStatusPolling: gofakeit.Bool(),
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

			_, err := c.RequestToPay(args.ctx, args.id, args.callback, args.handleStatusPolling, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToPay() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestToPayTransactionStatus(t *testing.T) {
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
			name: "happy case: successfully get transaction status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.RequestToPayStatus")).Return(nil).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get transaction status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.RequestToPayStatus")).Return(
					errors.New("an error occurred")).Once()

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

			_, err := c.RequestToPayTransactionStatus(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToPayTransactionStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestToPayDeliveryNotification(t *testing.T) {
	type args struct {
		ctx               context.Context
		id                uuid.UUID
		message, language string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get transaction status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.DeliveryNotification")).Return(nil).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					message:  gofakeit.BeerName(),
					language: gofakeit.BeerName(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get transaction status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.DeliveryNotification")).Return(
					errors.New("an error occurred")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					message:  gofakeit.BeerName(),
					language: gofakeit.BeerName(),
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
					message:  gofakeit.BeerName(),
					language: gofakeit.BeerName(),
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
		{
			name: "sad case: empty message passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					language: gofakeit.BeerName(),
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

			_, err := c.RequestToPayDeliveryNotification(args.ctx, args.id, args.message, args.language)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToPayDeliveryNotification() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
