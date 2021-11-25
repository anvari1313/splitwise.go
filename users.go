package splitwise

import (
	"context"
	"encoding/json"
	"net/http"
)

type Users interface {
	// CurrentUser returns information about the current user
	CurrentUser(ctx context.Context) (*CurrentUser, error)
	UserByID(ctx context.Context) (interface{}, error)
	UpdateUser(ctx context.Context) (interface{}, error)
}

type currentUserResponse struct {
	User CurrentUser `json:"user"`
}

type CurrentUser struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"picture"`
	CustomPicture      bool        `json:"custom_picture"`
	Email              string      `json:"email"`
	RegistrationStatus string      `json:"registration_status"`
	ForceRefreshAt     interface{} `json:"force_refresh_at"` // TODO: Data type is not known.
	Locale             string      `json:"locale"`
	CountryCode        string      `json:"country_code"`
	DateFormat         string      `json:"date_format"`
	DefaultCurrency    string      `json:"default_currency"`
	DefaultGroupID     int64       `json:"default_group_id"`
	NotificationsRead  string      `json:"notifications_read"` // TODO: Convert it to time.Date type
	NotificationsCount uint        `json:"notifications_count"`
	Notifications      struct {
		AddedAsFriend  bool `json:"added_as_friend"`
		AddedToGroup   bool `json:"added_to_group"`
		ExpenseAdded   bool `json:"expense_added"`
		ExpenseUpdated bool `json:"expense_updated"`
		Bills          bool `json:"bills"`
		Payments       bool `json:"payments"`
		MonthlySummary bool `json:"monthly_summary"`
		Announcements  bool `json:"announcements"`
	} `json:"notifications"`
}

// CurrentUser returns information about the current user
func (c client) CurrentUser(ctx context.Context) (*CurrentUser, error) {
	url := c.baseURL + "/api/v3.0/get_current_user"
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
	defer res.Body.Close()

	var response currentUserResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.User, nil
}

func (c client) UserByID(ctx context.Context) (interface{}, error) {
	panic("implement me")
}

func (c client) UpdateUser(ctx context.Context) (interface{}, error) {
	panic("implement me")
}
