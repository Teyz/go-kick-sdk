package kick

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/teyz/go-kick-sdk/constants"
	"github.com/teyz/go-kick-sdk/entities"
)

func (kc *KickClient) GetCurrentUser(ctx context.Context) (*entities.GetUsersResponse, error) {
	if err := kc.EnsureAccessTokenValid(ctx); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, constants.LinkUsers.ToString(), nil)
	if err != nil {
		return nil, errors.New("error creating request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", kc.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error sending request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	getUsersResponse := &entities.GetUsersResponse{}
	err = json.Unmarshal(body, &getUsersResponse)
	if err != nil {
		return nil, errors.New("error unmarshaling JSON")
	}

	return getUsersResponse, nil
}

func (kc *KickClient) GetUsers(ctx context.Context, userIDs []int) (*entities.GetUsersResponse, error) {
	if err := kc.EnsureAccessTokenValid(ctx); err != nil {
		return nil, err
	}

	strIDs := make([]string, 0, len(userIDs))
	for _, id := range userIDs {
		strIDs = append(strIDs, strconv.Itoa(id))
	}

	queryUserIDs := "id=" + strings.Join(strIDs, ",")

	link := fmt.Sprintf("%s?%s", constants.LinkUsers, queryUserIDs)
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, errors.New("error creating request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", kc.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error sending request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	getUsersResponse := &entities.GetUsersResponse{}
	err = json.Unmarshal(body, &getUsersResponse)
	if err != nil {
		return nil, errors.New("error unmarshaling JSON")
	}

	return getUsersResponse, nil
}
