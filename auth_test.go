package splitwise

import (
	"testing"
)

func TestApiKeyAuthProvider_Auth(t *testing.T) {
	ap := &apiKeyAuthProvider{apiKey: "api-key"}

	token, err := ap.Auth()
	if token != "api-key" {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}
