package splitwise

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// Users resources to access and modify user information.
type Users interface {
	// CurrentUser returns information about the current user
	CurrentUser(ctx context.Context) (*CurrentUser, error)

	// UserByID returns a user information by their id
	UserByID(ctx context.Context, id uint64) (*User, error)

	// UpdateUser updates a user's information by their ID and returns the result
	UpdateUser(ctx context.Context, id uint64, fields ...UserUpdatableField) (*CurrentUser, error)
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
	defer func() {
		_ = res.Body.Close()
	}()

	err = c.checkError(res)
	if err != nil {
		return nil, err
	}

	var response currentUserResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.User, nil
}

type userResponse struct {
	User User `json:"user"`
}

type User struct {
	ID        uint64 `json:"id"`
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
}

// UserByID returns a user information by their id.
func (c client) UserByID(ctx context.Context, id uint64) (*User, error) {
	url := c.baseURL + "/api/v3.0/get_user/" + strconv.FormatUint(id, 10)
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

	var response userResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.User, nil
}

type updateUserResponse struct {
	User CurrentUser `json:"user"`
}

func (c client) UpdateUser(ctx context.Context, id uint64, fields ...UserUpdatableField) (*CurrentUser, error) {
	url := c.baseURL + "/api/v3.0/update_user/" + strconv.FormatUint(id, 10)

	body := map[string]interface{}{}
	for _, field := range fields {
		body[field.Key()] = field.Value()
	}

	rawBody, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(rawBody))
	if err != nil {
		return nil, err
	}

	token, err := c.AuthProvider.Auth()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
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

	var response updateUserResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.User, nil
}

type UserUpdatableField interface {
	Key() string
	Value() interface{}
}

type userUpdatableField struct {
	key   string
	value interface{}
}

func (u userUpdatableField) Key() string {
	return u.key
}

func (u userUpdatableField) Value() interface{} {
	return u.value
}

func UserLastNameField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "last_name",
		value: value,
	}
}

func UserFirstNameField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "first_name",
		value: value,
	}
}

func UserEmailField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "email",
		value: value,
	}
}

func UserPasswordField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "password",
		value: value,
	}
}

func UserLocaleField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "locale",
		value: value,
	}
}

func UserDefaultCurrencyField(value string) UserUpdatableField {
	return &userUpdatableField{
		key:   "default_currency",
		value: value,
	}
}
