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

func TestRequestToWithdrawV1(t *testing.T) {
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
		ctx      context.Context
		id       uuid.UUID
		body     types.RequestToPayInput
		callback string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully request withdrawal",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(nil).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
					callback: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to request to withdraw",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
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
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
					callback: gofakeit.URL(),
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

			err := c.RequestToWithdrawV1(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToWithdrawV1() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestToWithdrawV2(t *testing.T) {
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
		ctx      context.Context
		id       uuid.UUID
		body     types.RequestToPayInput
		callback string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully request withdrawal",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(nil).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
					callback: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to request to withdraw",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.RequestToPayInput"), nil).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
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
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error")).Once()

				return args{
					ctx:      ctx,
					id:       uuid.New(),
					body:     input,
					callback: gofakeit.URL(),
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

			err := c.RequestToWithdrawV2(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToWithdrawV2() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRequestToWithdrawTransactionStatus(t *testing.T) {
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

			_, err := c.RequestToWithdrawTransactionStatus(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestRequestToWithdrawTransactionStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
