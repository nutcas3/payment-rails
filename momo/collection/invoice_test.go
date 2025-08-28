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

func TestCreateInvoice(t *testing.T) {
	input := types.CreateInvoiceInput{
		ExternalID:       gofakeit.UUID(),
		Amount:           "100",
		Currency:         types.EUR,
		ValidityDuration: "1",
		IntendedPayer: types.Party{
			PartyIDType: types.EMAIL,
			PartyID:     gofakeit.Email(),
		},
		Payee: types.Party{
			PartyIDType: types.MSISDN,
			PartyID:     gofakeit.PhoneFormatted(),
		},
		Description: gofakeit.BeerName(),
	}

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		callback string
		body     types.CreateInvoiceInput
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create invoice",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.CreateInvoiceInput"), nil).Return(nil).Once()

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
			name: "sad case: fail to create invoice",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.CreateInvoiceInput"), nil).Return(errors.New("an error")).Once()

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

			err := c.CreateInvoice(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCreateInvoice() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCancelInvoice(t *testing.T) {
	type args struct {
		ctx           context.Context
		callback      string
		invoiceID     uuid.UUID
		transactionID uuid.UUID
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully cancel invoice",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodDelete, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, nil).Return(nil).Once()

				return args{
					ctx:           ctx,
					callback:      gofakeit.URL(),
					invoiceID:     uuid.New(),
					transactionID: uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to cancel invoice",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodDelete, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, nil).Return(errors.New("an error")).Once()

				return args{
					ctx:           ctx,
					callback:      gofakeit.URL(),
					invoiceID:     uuid.New(),
					transactionID: uuid.New(),
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
					ctx:           ctx,
					callback:      gofakeit.URL(),
					invoiceID:     uuid.New(),
					transactionID: uuid.New(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil uuid passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:           ctx,
					callback:      gofakeit.URL(),
					invoiceID:     uuid.Nil,
					transactionID: uuid.Nil,
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

			err := c.CancelInvoice(args.ctx, args.invoiceID, args.transactionID, args.callback)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCancelInvoice() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetInvoiceStatus(t *testing.T) {
	type args struct {
		ctx   context.Context
		refID uuid.UUID
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get invoice status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.InvoiceStatus")).Return(nil).Once()

				return args{
					ctx:   ctx,
					refID: uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get invoice status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.InvoiceStatus")).Return(
					errors.New("an error")).Once()

				return args{
					ctx:   ctx,
					refID: uuid.New(),
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
					ctx:   ctx,
					refID: uuid.New(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil uuid passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:   ctx,
					refID: uuid.Nil,
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

			_, err := c.GetInvoiceStatus(args.ctx, args.refID)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCancelInvoice() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
