package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
)

// Currencies methods related to supported currencies
type Currencies interface {
	// Currencies returns a list of all currencies allowed by the system. These are mostly ISO 4217 codes, but we do
	//sometimes use pending codes or unofficial, colloquial codes (like BTC instead of XBT for Bitcoin)
	Currencies(ctx context.Context) ([]Currency, error)
}

type Currency struct {
	CurrencyCode string `json:"currency_code"`
	Unit         string `json:"unit"`
}

type currenciesResponse struct {
	Currencies []Currency `json:"currencies"`
}

// Currencies returns a list of all currencies allowed by the system. These are mostly ISO 4217 codes, but we do
// sometimes use pending codes or unofficial, colloquial codes (like BTC instead of XBT for Bitcoin)
func (c client) Currencies(ctx context.Context) ([]Currency, error) {
	url := c.baseURL + "/api/v3.0/get_currencies"
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

	var response currenciesResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Currencies, nil
}
