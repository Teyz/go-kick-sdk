package kick

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/teyz/go-kick-sdk/constants"
)

func (kc *KickClient) ExchangeCode(ctx context.Context, code string, codeVerifier string) error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", kc.ClientID)
	data.Set("client_secret", kc.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", kc.RedirectURI)
	data.Set("code_verifier", codeVerifier)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, constants.LinkToken.ToString(), strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := kc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("token exchange failed: %s", string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	kc.AccessToken = result.AccessToken
	kc.ExpiresAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	return nil
}

func (kc *KickClient) RefreshAccessToken(ctx context.Context) error {
	if kc.RefreshToken == "" {
		return errors.New("missing refresh token")
	}

	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", kc.RefreshToken)
	form.Set("client_id", kc.ClientID)
	form.Set("client_secret", kc.ClientSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, constants.LinkToken.ToString(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := kc.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return errors.New("refresh failed: " + string(body))
	}

	var result struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	kc.AccessToken = result.AccessToken
	if result.RefreshToken != "" {
		kc.RefreshToken = result.RefreshToken
	}
	kc.ExpiresAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	return nil
}

func (kc *KickClient) IsTokenExpired() bool {
	return time.Now().After(kc.ExpiresAt)
}

func (kc *KickClient) EnsureAccessTokenValid(ctx context.Context) error {
	if kc.IsTokenExpired() {
		return kc.RefreshAccessToken(ctx)
	}
	return nil
}
