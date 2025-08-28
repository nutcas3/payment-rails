package remittance_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/nutcas3/payment-rails/momo/common/mocks"
	"github.com/nutcas3/payment-rails/momo/remittance"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccessToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create access token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(nil).Once()

				mh.Cache.EXPECT().Set(mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration"))

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "happy case: successfully get cached access token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to create access token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.AccessTokenResp")).Return(errors.New("an error occurred")).Once()

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

			r := remittance.NewRemittance(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := r.CreateAccessToken(args.ctx)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCreateAccessToken() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOauth2Token(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create oauth2 token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Oauth2Resp")).Return(nil).Once()

				mh.Cache.EXPECT().Set(mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration"))

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "happy case: successfully get cached oauth2 token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				return args{
					ctx: ctx,
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to create oauth2 token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Oauth2Resp")).Return(errors.New("an error occurred")).Once()

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

			r := remittance.NewRemittance(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := r.CreateOauth2Token(args.ctx)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestCreateOauth2Token() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBcAuthorize(t *testing.T) {
	type args struct {
		ctx         context.Context
		callbackURL string
	}

	tests := []struct {
		name    string
		setup   func(mh *mockHandler) args
		wantErr bool
	}{
		{
			name: "happy case: successfully create bcauthorize token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false).Once()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.BcAuthResp")).Return(nil).Once()

				mh.Cache.EXPECT().Set(mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration"))

				return args{
					ctx:         ctx,
					callbackURL: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "happy case: successfully get cached oauth2 token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true)

				return args{
					ctx:         ctx,
					callbackURL: gofakeit.URL(),
				}
			},
			wantErr: false,
		},
		{
			name: "sad case: fail to create bcauthorize token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false).Once()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, true).Once()

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.BcAuthResp")).Return(errors.New("an error occurred")).Once()

				return args{
					ctx:         ctx,
					callbackURL: gofakeit.URL(),
				}
			},
			wantErr: true,
		},
		{
			name: "sad case: fail to get oauth2 token",
			setup: func(mh *mockHandler) args {
				ctx := context.Background()

				mh.Cache.EXPECT().Get(mock.Anything).Return(mock.Anything, false).Times(2)

				mh.Backend.EXPECT().Call(ctx, http.MethodPost, mock.Anything, mock.AnythingOfType("http.Header"), mock.Anything, nil,
					mock.AnythingOfType("*types.Oauth2Resp")).Return(errors.New("an error occurred")).Once()

				return args{
					ctx:         ctx,
					callbackURL: gofakeit.URL(),
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

			r := remittance.NewRemittance(
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				gofakeit.BeerName(),
				mockBackend,
				mockCache,
			)

			_, err := r.BcAuthorize(args.ctx, args.callbackURL)

			mockBackend.AssertExpectations(t)
			mockCache.AssertExpectations(t)

			if (err != nil) != tt.wantErr {
				t.Logf("TestBcAuthorize() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
