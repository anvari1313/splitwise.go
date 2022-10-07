package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Friends contains method to work with friend resource
type Friends interface {
	// Friends returns current user's friends
	Friends(ctx context.Context) ([]Friend, error)

	// DeleteFriend Given a friend ID, break off the friendship between the current user and the specified user.
	DeleteFriend(ctx context.Context, id uint64) (bool, error)
}

type Friend struct {
	ID                 int    `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	RegistrationStatus string `json:"registration_status"`
	Picture            struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"picture"`
	Groups []struct {
		GroupId int `json:"group_id"`
		Balance []struct {
			CurrencyCode string `json:"currency_code"`
			Amount       string `json:"amount"`
		} `json:"balance"`
	} `json:"groups"`
	Balance []struct {
		CurrencyCode string `json:"currency_code"`
		Amount       string `json:"amount"`
	} `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

type friendsResponse struct {
	Friends []Friend `json:"friends"`
}

func (c client) Friends(ctx context.Context) ([]Friend, error) {
	url := c.baseURL + "/api/v3.0/get_friends"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	err = c.checkError(res)
	if err != nil {
		return nil, err
	}

	var response friendsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Friends, nil
}

type deleteFriendResponse struct {
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
}

func (c client) DeleteFriend(ctx context.Context, id uint64) (bool, error) {
	url := c.baseURL + "/api/v3.0/delete_friend/" + strconv.FormatUint(id, 10)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return false, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	res, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	err = c.checkError(res)
	if err != nil {
		return false, err
	}

	var response deleteFriendResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}
