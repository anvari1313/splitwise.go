package splitwise

import (
	"net/http"
)

type Client interface {
	Users
	Groups
	Friends
	Expenses
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
