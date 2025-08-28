package disbursement_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/nutcas3/payment-rails/momo/common/mocks"
	"github.com/nutcas3/payment-rails/momo/common/types"
	"github.com/nutcas3/payment-rails/momo/disbursement"
	"github.com/stretchr/testify/mock"
)

func TestDepositV1(t *testing.T) {
	body := types.TransferInput{
		Amount:     "100",
		Currency:   types.AED,
		ExternalID: gofakeit.UUID(),
		Payee: types.Party{
			PartyIDType: types.EMAIL,
			PartyID:     gofakeit.Email(),
		},
		PayerMessage: gofakeit.BeerName(),
		PayeeNote:    gofakeit.BeerName(),
	}

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		body     types.TransferInput
		callback string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully deposit",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), mock.AnythingOfType("types.TransferInput"), nil).Return(nil)

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to deposit",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), mock.AnythingOfType("types.TransferInput"), nil).Return(errors.New("an error"))

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
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
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error occurred"))

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil UUID",
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

			d := disbursement.NewDisbursement(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			err := d.DepositV1(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRefundV1() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDepositV2(t *testing.T) {
	body := types.TransferInput{
		Amount:     "100",
		Currency:   types.AED,
		ExternalID: gofakeit.UUID(),
		Payee: types.Party{
			PartyIDType: types.EMAIL,
			PartyID:     gofakeit.Email(),
		},
		PayerMessage: gofakeit.BeerName(),
		PayeeNote:    gofakeit.BeerName(),
	}

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		body     types.TransferInput
		callback string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully deposit",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), mock.AnythingOfType("types.TransferInput"), nil).Return(nil)

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to deposit",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), mock.AnythingOfType("types.TransferInput"), nil).Return(errors.New("an error"))

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
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
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error occurred"))

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     body,
					callback: gofakeit.URL(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil UUID",
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

			d := disbursement.NewDisbursement(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			err := d.DepositV2(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestDepositV2() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDepositStatus(t *testing.T) {
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
			name: "happy case: successfully get deposit status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.DisbursementTransactionStatus")).Return(nil)

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get deposit status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"),
					mock.AnythingOfType("*common.Params"), nil, mock.AnythingOfType("*types.DisbursementTransactionStatus")).Return(
					errors.New("an error"))

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
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error occurred"))

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: nil UUID",
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

			d := disbursement.NewDisbursement(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := d.GetDepositStatus(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetDepositStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
