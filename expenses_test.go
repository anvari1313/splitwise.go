package splitwise

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_CreateExpenseSplitEqually(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedReqBody := []ExpenseSplitEqually{
			{
				Expense: Expense{
					Cost:           "25",
					Description:    "Grocery run",
					Details:        "string",
					Date:           "2012-05-02T13:00:00Z",
					RepeatInterval: "never",
					CurrencyCode:   "USD",
					CategoryId:     15,
					GroupId:        0,
				},
				SplitEqually: true,
			},
		}
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/create_expense" {
				t.Error("invalid URL request")
			}

			reqBody, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fail()
			}

			a, err := json.Marshal(expectedReqBody[0])
			if err != nil {
				t.Fail()
			}

			if string(reqBody) != string(a) {
				t.FailNow()
			}
			t.Log("ok equality")

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"expenses": [
				  {
					"cost": "25.0",
					"description": "Brunch",
					"details": "string",
					"date": "2012-05-02T13:00:00Z",
					"repeat_interval": "never",
					"currency_code": "USD",
					"category_id": 15,
					"id": 51023,
					"group_id": 391,
					"friendship_id": 4818,
					"expense_bundle_id": 491030,
					"repeats": true,
					"email_reminder": true,
					"email_reminder_in_advance": null,
					"next_repeat": "string",
					"comments_count": 0,
					"payment": true,
					"transaction_confirmed": true,
					"repayments": [
					  {
						"from": 6788709,
						"to": 270896089,
						"amount": "25.0"
					  }
					],
					"created_at": "2012-07-27T06:17:09Z",
					"created_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"updated_at": "2012-12-23T05:47:02Z",
					"updated_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"deleted_at": "2012-12-23T05:47:02Z",
					"deleted_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"category": {
					  "id": 5,
					  "name": "Electricity"
					},
					"receipt": {
					  "large": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/large_95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png",
					  "original": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png"
					},
					"users": [
					  {
						"user": {
						  "id": 491923,
						  "first_name": "Jane",
						  "last_name": "Doe",
						  "picture": {
							"medium": "image_url"
						  }
						},
						"user_id": 491923,
						"paid_share": "8.99",
						"owed_share": "4.5",
						"net_balance": "4.49"
					  }
					],
					"comments": [
					  {
						"id": 79800950,
						"content": "John D. updated this transaction: - The cost changed from $6.99 to $8.99",
						"comment_type": "System",
						"relation_type": "ExpenseComment",
						"relation_id": 855870953,
						"created_at": "2019-08-24T14:15:22Z",
						"deleted_at": "2019-08-24T14:15:22Z",
						"user": null
					  }
					]
				  }
				],
				"errors": {}
			  }`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.CreateExpenseSplitEqually(context.Background(), expectedReqBody[0])

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestClient_CreateExpenseByShare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		type ExpenseByShare struct {
			Expense
			UserID0    uint64 `json:"users__0__user_id"`
			PaidShare0 string `json:"users__0__paid_share"`
			OwedShare0 string `json:"users__0__owed_share"`
			UserID1    uint64 `json:"users__1__user_id"`
			PaidShare1 string `json:"users__1__paid_share"`
			OwedShare1 string `json:"users__1__owed_share"`
		}

		expense := Expense{
			Cost:           "25",
			Description:    "Grocery run",
			Details:        "string",
			Date:           "2012-05-02T13:00:00Z",
			RepeatInterval: "never",
			CurrencyCode:   "USD",
			CategoryId:     15,
			GroupId:        0,
		}

		user1 := UserShare{
			UserID:    54123,
			PaidShare: "25",
			OwedShare: "15",
		}

		user2 := UserShare{
			UserID:    34262,
			PaidShare: "0",
			OwedShare: "10",
		}

		userShares := []UserShare{
			user1,
			user2,
		}

		expectedReqBody := []ExpenseByShare{
			{
				Expense:    expense,
				UserID0:    user1.UserID,
				PaidShare0: user1.PaidShare,
				OwedShare0: user1.OwedShare,
				UserID1:    user2.UserID,
				PaidShare1: user2.PaidShare,
				OwedShare1: user2.OwedShare,
			},
		}
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/create_expense" {
				t.Error("invalid URL request")
			}

			reqBody, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fail()
			}
			t.Log(string(reqBody))
			a, err := json.Marshal(expectedReqBody[0])
			if err != nil {
				t.Fail()
			}

			if string(reqBody) != string(a) {
				t.FailNow()
			}
			t.Log("ok equality")

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"expenses": [
				  {
					"cost": "25.0",
					"description": "Brunch",
					"details": "string",
					"date": "2012-05-02T13:00:00Z",
					"repeat_interval": "never",
					"currency_code": "USD",
					"category_id": 15,
					"id": 51023,
					"group_id": 391,
					"friendship_id": 4818,
					"expense_bundle_id": 491030,
					"repeats": true,
					"email_reminder": true,
					"email_reminder_in_advance": null,
					"next_repeat": "string",
					"comments_count": 0,
					"payment": true,
					"transaction_confirmed": true,
					"repayments": [
					  {
						"from": 6788709,
						"to": 270896089,
						"amount": "25.0"
					  }
					],
					"created_at": "2012-07-27T06:17:09Z",
					"created_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"updated_at": "2012-12-23T05:47:02Z",
					"updated_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"deleted_at": "2012-12-23T05:47:02Z",
					"deleted_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"category": {
					  "id": 5,
					  "name": "Electricity"
					},
					"receipt": {
					  "large": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/large_95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png",
					  "original": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png"
					},
					"users": [
					  {
						"user": {
						  "id": 491923,
						  "first_name": "Jane",
						  "last_name": "Doe",
						  "picture": {
							"medium": "image_url"
						  }
						},
						"user_id": 491923,
						"paid_share": "8.99",
						"owed_share": "4.5",
						"net_balance": "4.49"
					  }
					],
					"comments": [
					  {
						"id": 79800950,
						"content": "John D. updated this transaction: - The cost changed from $6.99 to $8.99",
						"comment_type": "System",
						"relation_type": "ExpenseComment",
						"relation_id": 855870953,
						"created_at": "2019-08-24T14:15:22Z",
						"deleted_at": "2019-08-24T14:15:22Z",
						"user": null
					  }
					]
				  }
				],
				"errors": {}
			  }`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.CreateExpenseByShare(context.Background(), expense, userShares)

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestClient_GetExpenseCurrentUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_expenses" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"expenses": [
				  {
					"cost": "25.0",
					"description": "Brunch",
					"details": "string",
					"date": "2012-05-02T13:00:00Z",
					"repeat_interval": "never",
					"currency_code": "USD",
					"category_id": 15,
					"id": 51023,
					"group_id": 391,
					"friendship_id": 4818,
					"expense_bundle_id": 491030,
					"repeats": true,
					"email_reminder": true,
					"email_reminder_in_advance": null,
					"next_repeat": "string",
					"comments_count": 0,
					"payment": true,
					"transaction_confirmed": true,
					"repayments": [
					  {
						"from": 6788709,
						"to": 270896089,
						"amount": "25.0"
					  }
					],
					"created_at": "2012-07-27T06:17:09Z",
					"created_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"updated_at": "2012-12-23T05:47:02Z",
					"updated_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"deleted_at": "2012-12-23T05:47:02Z",
					"deleted_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"category": {
					  "id": 5,
					  "name": "Electricity"
					},
					"receipt": {
					  "large": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/large_95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png",
					  "original": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png"
					},
					"users": [
					  {
						"user": {
						  "id": 491923,
						  "first_name": "Jane",
						  "last_name": "Doe",
						  "picture": {
							"medium": "image_url"
						  }
						},
						"user_id": 491923,
						"paid_share": "8.99",
						"owed_share": "4.5",
						"net_balance": "4.49"
					  }
					],
					"comments": [
					  {
						"id": 79800950,
						"content": "John D. updated this transaction: - The cost changed from $6.99 to $8.99",
						"comment_type": "System",
						"relation_type": "ExpenseComment",
						"relation_id": 855870953,
						"created_at": "2019-08-24T14:15:22Z",
						"deleted_at": "2019-08-24T14:15:22Z",
						"user": null
					  }
					]
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

		_, err := c.Expenses(context.Background())

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestClient_GetExpenseByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_expense/10" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"expense": {
					"cost": "25.0",
					"description": "Brunch",
					"details": "string",
					"date": "2012-05-02T13:00:00Z",
					"repeat_interval": "never",
					"currency_code": "USD",
					"category_id": 15,
					"id": 51023,
					"group_id": 391,
					"friendship_id": 4818,
					"expense_bundle_id": 491030,
					"repeats": true,
					"email_reminder": true,
					"email_reminder_in_advance": null,
					"next_repeat": "string",
					"comments_count": 0,
					"payment": true,
					"transaction_confirmed": true,
					"repayments": [
					  {
						"from": 6788709,
						"to": 270896089,
						"amount": "25.0"
					  }
					],
					"created_at": "2012-07-27T06:17:09Z",
					"created_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"updated_at": "2012-12-23T05:47:02Z",
					"updated_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"deleted_at": "2012-12-23T05:47:02Z",
					"deleted_by": {
					  "id": 0,
					  "first_name": "Ada",
					  "last_name": "Lovelace",
					  "email": "ada@example.com",
					  "registration_status": "confirmed",
					  "picture": {
						"small": "string",
						"medium": "string",
						"large": "string"
					  }
					},
					"category": {
					  "id": 5,
					  "name": "Electricity"
					},
					"receipt": {
					  "large": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/large_95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png",
					  "original": "https://splitwise.s3.amazonaws.com/uploads/expense/receipt/3678899/95f8ecd1-536b-44ce-ad9b-0a9498bb7cf0.png"
					},
					"users": [
					  {
						"user": {
						  "id": 491923,
						  "first_name": "Jane",
						  "last_name": "Doe",
						  "picture": {
							"medium": "image_url"
						  }
						},
						"user_id": 491923,
						"paid_share": "8.99",
						"owed_share": "4.5",
						"net_balance": "4.49"
					  }
					],
					"comments": [
					  {
						"id": 79800950,
						"content": "John D. updated this transaction: - The cost changed from $6.99 to $8.99",
						"comment_type": "System",
						"relation_type": "ExpenseComment",
						"relation_id": 855870953,
						"created_at": "2019-08-24T14:15:22Z",
						"deleted_at": "2019-08-24T14:15:22Z",
						"user": null
					  }
					]
				  }
			  }`))
		}))
		defer server.Close()

		c := &client{
			AuthProvider: NewAPIKeyAuth("api-key"),
			baseURL:      server.URL,
			client:       http.DefaultClient,
		}

		_, err := c.ExpenseByID(context.Background(), 10)

		if err != nil {
			t.Fatal(err)
		}
	})
}
