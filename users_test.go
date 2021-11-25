package splitwise

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_CurrentUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_current_user" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"user": {
					"id": 1313,
					"first_name": "John",
					"last_name":"Petrucci",
					"picture":{
						"small": "https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/small_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"medium":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/medium_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"large":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/large_a89sdf15-ce23-58c6-1d42-992785702938.jpeg"
					},
					"custom_picture": true,
					"email": "john@gmail.com",
					"registration_status": "confirmed",
					"force_refresh_at": null,
					"locale": "en",
					"country_code": "FR",
					"date_format": "MM/DD/YYYY",
					"default_currency": "USD",
					"default_group_id": -1,
					"notifications_read": "2021-10-23T07:15:37Z",
					"notifications_count": 5,
					"notifications": {
						"added_as_friend": true,
						"added_to_group": true,
						"expense_added": false,
						"expense_updated": false,
						"bills": true,
						"payments": true,
						"monthly_summary": true,
						"announcements": true
					}
				}
			}`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.CurrentUser(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestClient_UserByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_user/1313" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"user": {
					"id": 1313,
					"first_name": "John",
					"last_name":"Petrucci",
					"picture":{
						"small": "https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/small_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"medium":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/medium_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"large":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/large_a89sdf15-ce23-58c6-1d42-992785702938.jpeg"
					},
					"custom_picture": true,
					"email": "john@gmail.com",
					"registration_status": "confirmed"
				}
			}`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.UserByID(context.Background(), 1313)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestClient_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/update_user/1313" {
				t.Error("invalid URL request")
			}

			reqBody, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fail()
			}

			expectedReqBody := `{"default_currency":"USD","email":"john@gmail.com","first_name":"Petrucci","last_name":"John","locale":"en","password":"password"}`
			if string(reqBody) != expectedReqBody {
				t.FailNow()
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"user": {
					"id": 1313,
					"first_name": "John",
					"last_name":"Petrucci",
					"picture":{
						"small": "https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/small_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"medium":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/medium_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
						"large":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/large_a89sdf15-ce23-58c6-1d42-992785702938.jpeg"
					},
					"custom_picture": true,
					"email": "john@gmail.com",
					"registration_status": "confirmed",
					"force_refresh_at": null,
					"locale": "en",
					"country_code": "FR",
					"date_format": "MM/DD/YYYY",
					"default_currency": "USD",
					"default_group_id": -1,
					"notifications_read": "2021-10-23T07:15:37Z",
					"notifications_count": 5,
					"notifications": {
						"added_as_friend": true,
						"added_to_group": true,
						"expense_added": false,
						"expense_updated": false,
						"bills": true,
						"payments": true,
						"monthly_summary": true,
						"announcements": true
					}
				}
			}`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.UpdateUser(context.Background(), 1313,
			UserLastNameField("John"),
			UserFirstNameField("Petrucci"),
			UserEmailField("john@gmail.com"),
			UserPasswordField("password"),
			UserLocaleField("en"),
			UserDefaultCurrencyField("USD"))
		if err != nil {
			t.Fatal(err)
		}
	})
}
