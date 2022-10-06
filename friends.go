package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Friends contains method to work with friend resource
type Friends interface {
	// Friends returns current user's friends
	Friends(ctx context.Context) ([]Friend, error)
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
