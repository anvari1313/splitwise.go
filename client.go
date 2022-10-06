package splitwise

import (
	"net/http"
)

type Client interface {
	Users
	Groups
	Friends
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
	case res.StatusCode == 500:
		return ErrSplitwiseServer
	default:
		return nil
	}
}
