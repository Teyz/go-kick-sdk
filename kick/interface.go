package kick

import (
	"context"

	"github.com/teyz/go-kick-sdk/entities"
)

type Kick interface {
	//Token
	ExchangeCode(ctx context.Context, code string, codeVerifier string) error
	RefreshAccessToken(ctx context.Context) error
	IsTokenExpired() bool
	EnsureAccessTokenValid(ctx context.Context) error

	//User
	GetCurrentUser(ctx context.Context, accessToken string) (*entities.GetUsersResponse, error)
	GetUsersByIDs(ctx context.Context, userIDs []int, accessToken string) (*entities.GetUsersResponse, error)
}
