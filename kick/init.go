package kick

import (
	"context"
	"net/http"
	"time"
)

type KickClient struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	HTTPClient   *http.Client
}

func NewKickClient(ctx context.Context, clientID string, clientSecret string, redirectURI string) (*KickClient, error) {
	return &KickClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		HTTPClient:   http.DefaultClient,
		AccessToken:  "ZWQ3OTE4ZGUTZJLLZS0ZOTC0LTLKZJATMMQ4ZJE3YWUXNTQ4",
	}, nil
}
