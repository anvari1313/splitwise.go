package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Groups A Group represents a collection of users who share expenses together. For example, some users use a Group to
// aggregate expenses related to an apartment. Others use it to represent a trip. Expenses assigned to a group are split
// among the users of that group. Importantly, two users in a Group can also have expenses with one another outside of
// the Group.
type Groups interface {
	// Groups returns the current user groups
	Groups(ctx context.Context) ([]Group, error)

	// GroupByID returns information about a group by its ID
	GroupByID(ctx context.Context, id uint64) (*Group, error)
}

type Group struct {
	ID                uint64        `json:"id"`
	Name              string        `json:"name"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Members           []GroupMember `json:"members"`
	SimplifyByDefault bool          `json:"simplify_by_default"`
	OriginalDebts     []Debt        `json:"original_debts"`
	SimplifiedDebts   []Debt        `json:"simplified_debts"`
	Avatar            struct {
		Small    string      `json:"small"`
		Medium   string      `json:"medium"`
		Large    string      `json:"large"`
		Xlarge   string      `json:"xlarge"`
		Xxlarge  string      `json:"xxlarge"`
		Original interface{} `json:"original"`
	} `json:"avatar"`
	TallAvatar struct {
		Xlarge string `json:"xlarge"`
		Large  string `json:"large"`
	} `json:"tall_avatar"`
	CustomAvatar bool `json:"custom_avatar"`
	CoverPhoto   struct {
		Xxlarge string `json:"xxlarge"`
		Xlarge  string `json:"xlarge"`
	} `json:"cover_photo"`
}

type GroupMember struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"picture"`
	CustomPicture      bool   `json:"custom_picture"`
	Email              string `json:"email"`
	RegistrationStatus string `json:"registration_status"`
	Balance            []struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
	} `json:"balance"`
}

type Debt struct {
	CurrencyCode string `json:"currency_code"`
	From         int    `json:"from"`
	To           int    `json:"to"`
	Amount       string `json:"amount"`
}

type groupsResponse struct {
	Groups []Group `json:"groups"`
}

func (c client) Groups(ctx context.Context) ([]Group, error) {
	url := c.baseURL + "/api/v3.0/get_groups"
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

	var response groupsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Groups, nil
}

type groupByIDResponse struct {
	Group Group `json:"group"`
}

func (c client) GroupByID(ctx context.Context, id uint64) (*Group, error) {
	url := c.baseURL + "/api/v3.0/get_group/" + strconv.FormatUint(id, 10)
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

	var response groupByIDResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.Group, nil
}
