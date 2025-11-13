package collection_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/nutcas3/payment-rails/momo/collection"
	"github.com/nutcas3/payment-rails/momo/common/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGetBasicUserInfo(t *testing.T) {
	type args struct {
		ctx                          context.Context
		accHolderID, accHolderIDType string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get basic user info",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.BasicUserInfo")).Return(nil)

				return args{
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get basic user info",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.BasicUserInfo")).Return(errors.New("an error occurred"))

				return args{
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: missing accHolderIDType",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:         ctx,
					accHolderID: gofakeit.UUID(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: missing accHolderID",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:             ctx,
					accHolderIDType: gofakeit.Email(),
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
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
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

			_, err := c.GetBasicUserInfo(args.ctx, args.accHolderIDType, args.accHolderID)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetBasicUserInfo() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserInfoWithConsent(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully get user info with consent",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.UserConsentInfo")).Return(nil)

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to get user info with consent",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.UserConsentInfo")).Return(errors.New("an error occurred"))

				return args{
					ctx: ctx,
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
					mock.AnythingOfType("*types.Oauth2Resp")).Return(errors.New("an error occurred"))

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

			_, err := c.GetUserInfoWithConsent(args.ctx)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestGetUserInforWithConsent() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAccountHolderStatus(t *testing.T) {
	type args struct {
		ctx                          context.Context
		accHolderID, accHolderIDType string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully validate account holder status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*bool")).Return(nil)

				return args{
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to validate account holder status",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodGet, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*bool")).Return(errors.New("an error occurred"))

				return args{
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: missing accHolderIDType",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:         ctx,
					accHolderID: gofakeit.UUID(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: missing accHolderID",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				return args{
					ctx:             ctx,
					accHolderIDType: gofakeit.Email(),
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
					ctx:             ctx,
					accHolderID:     gofakeit.UUID(),
					accHolderIDType: gofakeit.Email(),
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

			_, err := c.ValidateAccountHolderStatus(args.ctx, args.accHolderIDType, args.accHolderID)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestValidateAccountHolderStatus() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
