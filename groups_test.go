package splitwise

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Groups(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_groups" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
				"groups": [
					{
						"id": 0,
						"name": "Non-group expenses",
						"created_at": "2019-09-08T10:06:05Z",
						"updated_at": "2021-11-25T18:31:14Z",
						"members": [
							{
								"id": 1313,
								"first_name": "John",
								"last_name": "Petrucci",
								"picture": {
									"small": "https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/small_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
									"medium":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/medium_a89sdf15-ce23-58c6-1d42-992785702938.jpeg",
									"large":"https://splitwise.s3.amazonaws.com/uploads/user/avatar/1313/large_a89sdf15-ce23-58c6-1d42-992785702938.jpeg"
								},
								"custom_picture": true,
								"email": "john@gmail.com",
								"registration_status": "confirmed",
								"balance": [
									{
										"amount": "-10465000.0",
										"currency_code": "IRR"
									}
								]
							}
						],
						"simplify_by_default": false,
						"original_debts": [
							{
								"currency_code": "IRR",
								"from": 1313,
								"to": 1314,
								"amount": "9940000.0"
							},
							{
								"currency_code": "IRR",
								"from": 1313,
								"to": 1314,
								"amount": "525000.0"
							}
						],
						"simplified_debts": [
							{
								"currency_code": "IRR",
								"from": 1313,
								"to": 1314,
								"amount": "9940000.0"
							},
							{
								"currency_code": "IRR",
								"from": 1314,
								"to": 1313,
								"amount": "525000.0"
							}
						],
						"avatar": {
							"small": "https://s3.amazonaws.com/splitwise/uploads/group/default_avatars/v2021/avatar-nongroup-50px.png",
							"medium": "https://s3.amazonaws.com/splitwise/uploads/group/default_avatars/v2021/avatar-nongroup-100px.png",
							"large": "https://s3.amazonaws.com/splitwise/uploads/group/default_avatars/v2021/avatar-nongroup-200px.png",
							"xlarge": "https://s3.amazonaws.com/splitwise/uploads/group/default_avatars/v2021/avatar-nongroup-500px.png",
							"xxlarge": "https://s3.amazonaws.com/splitwise/uploads/group/default_avatars/v2021/avatar-nongroup-1000px.png",
							"original": null
						},
						"tall_avatar": {
							"xlarge": "https://s3.amazonaws.com/splitwise/uploads/group/default_tall_avatars/avatar-nongroup-288px.png",
							"large": "https://s3.amazonaws.com/splitwise/uploads/group/default_tall_avatars/avatar-nongroup-192px.png"
						},
						"custom_avatar": false,
						"cover_photo": {
							"xxlarge": "https://s3.amazonaws.com/splitwise/uploads/group/default_cover_photos/coverphoto-nongroup-1000px.png",
							"xlarge": "https://s3.amazonaws.com/splitwise/uploads/group/default_cover_photos/coverphoto-nongroup-500px.png"
						}
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

		_, err := c.Groups(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	})
}
