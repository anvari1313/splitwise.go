package splitwise

// AuthProvider knows how to provide authentication for calling the service APIs
type AuthProvider interface {
	// Auth Provides the auth token header used in calling the APIs
	Auth() (string, error)
}

// NewAPIKeyAuth returns a new AuthProvider that is working with API key
func NewAPIKeyAuth(apiKey string) AuthProvider {
	return &apiKeyAuthProvider{apiKey: apiKey}
}

type apiKeyAuthProvider struct {
	apiKey string
}

// Auth Provides the auth token header used in calling the APIs
func (a apiKeyAuthProvider) Auth() (string, error) {
	return a.apiKey, nil
}
