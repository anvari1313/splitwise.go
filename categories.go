package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
)

// Categories methods related to supported categories
type Categories interface {
	// Categories returns a list of all categories Splitwise allows for expenses. There are parent categories that
	//represent groups of categories with subcategories for more specific categorization. When creating expenses, you
	//must use a subcategory, not a parent category. If you intend for an expense to be represented by the parent
	//category and nothing more specific, please use the "Other" subcategory.
	Categories(ctx context.Context) ([]Category, error)
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconTypes struct {
		Slim struct {
			Small string `json:"small"`
			Large string `json:"large"`
		} `json:"slim"`
		Square struct {
			Large  string `json:"large"`
			XLarge string `json:"xlarge"`
		} `json:"square"`
		Transparent struct {
			Large  string `json:"large"`
			XLarge string `json:"xlarge"`
		} `json:"transparent"`
	} `json:"icon_types"`
	Subcategories []Category `json:"subcategories"`
}

type categoriesResponse struct {
	Categories []Category `json:"categories"`
}

// Categories returns a list of all categories Splitwise allows for expenses. There are parent categories that represent
// groups of categories with subcategories for more specific categorization. When creating expenses, you must use
// a subcategory, not a parent category. If you intend for an expense to be represented by the parent category and
// nothing more specific, please use the "Other" subcategory.
func (c client) Categories(ctx context.Context) ([]Category, error) {
	url := c.baseURL + "/api/v3.0/get_categories"
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

	var response categoriesResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Categories, nil
}
