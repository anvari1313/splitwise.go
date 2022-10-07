package splitwise

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Currencies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_currencies" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
	"currencies": [
		{
			"currency_code": "EGP",
			"unit": "E£"
		},
		{
			"currency_code": "ETB",
			"unit": "Br"
		},
		{
			"currency_code": "EUR",
			"unit": "€"
		}
	]
}`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.Currencies(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	})
}
