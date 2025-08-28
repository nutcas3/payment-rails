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

func TestCancelPreApproval(t *testing.T) {
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
			name: "happy case: successfully cancel preapproval",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodDelete, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.Anything, nil).Return(nil).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to cancel preapproval",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodDelete, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.Anything, nil).Return(errors.New("an error")).Once()

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

			err := c.CancelPreApproval(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCancelPreApproval() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPreApprovalStatus(t *testing.T) {
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
			name: "happy case: successfully get preapproval status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.Anything, mock.AnythingOfType("*types.PreApprovalStatus")).Return(nil).Once()

				return args{
					ctx: ctx,
					id:  uuid.New(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get preapproval status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.Anything, mock.AnythingOfType("*types.PreApprovalStatus")).Return(errors.New("an error")).Once()

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

			_, err := c.GetPreApprovalStatus(args.ctx, args.id)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetPreApprovalStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPreApproval(t *testing.T) {
	input := types.PreApprovalInput{
		Payer: types.Party{
			PartyIDType: types.EMAIL,
			PartyID:     gofakeit.Email(),
		},
		PayerCurrency: types.AED,
		PayerMessage:  gofakeit.BeerAlcohol(),
		ValidityTime:  2,
	}

	type args struct {
		ctx      context.Context
		id       uuid.UUID
		body     types.PreApprovalInput
		callback string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create preapproval",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.PreApprovalInput"), nil).Return(nil).Once()

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
			name: "sad case: fail to create preapproval",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything,
					mock.AnythingOfType("types.PreApprovalInput"), nil).Return(errors.New("an error")).Once()

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

			err := c.PreApproval(args.ctx, args.id, args.callback, args.body)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestPreApproval() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetApprovedPreApprovals(t *testing.T) {
	type args struct {
		ctx        context.Context
		idType, id string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get approved preapprovals",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, mock.Anything,
					mock.AnythingOfType("[]*types.PreApprovalDetails")).Return(nil).Once()

				return args{
					ctx:    ctx,
					id:     gofakeit.UUID(),
					idType: gofakeit.BeerAlcohol(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get approved preapprovals",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, mock.Anything,
					mock.AnythingOfType("[]*types.PreApprovalDetails")).Return(errors.New("an error")).Once()

				return args{
					ctx:    ctx,
					id:     gofakeit.UUID(),
					idType: gofakeit.BeerAlcohol(),
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
					ctx:    ctx,
					id:     gofakeit.UUID(),
					idType: gofakeit.BeerAlcohol(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: empty id passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:    ctx,
					id:     "",
					idType: gofakeit.BeerAlcohol(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: empty idType passed",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:    ctx,
					id:     gofakeit.UUID(),
					idType: "",
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

			_, err := c.GetApprovedPreApprovals(args.ctx, args.id, args.idType)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetApprovedPreApprovals() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
