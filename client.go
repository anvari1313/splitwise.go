package splitwise

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client interface {
	Users
	Groups
	Friends
	Expenses
	Currencies
	Categories
}

const (
	ServerAddress = "https://secure.splitwise.com"
)

// NewClient returns a new Client with the given AuthProvider
func NewClient(authProvider AuthProvider) Client {
	return &client{
		AuthProvider: authProvider,
		baseURL:      ServerAddress,
		client:       http.DefaultClient,
	}
}

type client struct {
	AuthProvider
	baseURL string
	client  *http.Client
}

func (c client) checkError(res *http.Response) error {
	switch {
	case res.StatusCode/100 == 1:
		return nil
	case res.StatusCode/100 == 2:
		return nil
	case res.StatusCode/100 == 3:
		return nil
	case res.StatusCode == 401:
		return ErrInvalidToken
	case res.StatusCode == 403:
		return ErrPermissionDenied
	case res.StatusCode == 404:
		return ErrRecordNotFound
	case res.StatusCode == 500:
		return ErrSplitwiseServer
	default:
		// This case normally should not be happened
		var response interface{}
		err := json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return err
		}

		return fmt.Errorf("unknown API status code: %d - payload: %+v", res.StatusCode, response)
	}
}
