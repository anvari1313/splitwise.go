package splitwise

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Categories(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/api/v3.0/get_categories" {
				t.Error("invalid URL request")
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{
  "categories": [
    {
      "id": 1,
      "name": "Utilities",
      "icon": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square/utilities/other.png",
      "icon_types": {
        "slim": {
          "small": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/utilities/other.png",
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/utilities/other@2x.png"
        },
        "square": {
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/utilities/other@2x.png",
          "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/utilities/other@3x.png"
        },
        "transparent": {
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/utilities/other@2x.png",
          "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/utilities/other@3x.png"
        }
      },
      "subcategories": [
        {
          "id": 48,
          "name": "Cleaning",
          "icon": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square/utilities/cleaning.png",
          "icon_types": {
            "slim": {
              "small": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/utilities/cleaning.png",
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/utilities/cleaning@2x.png"
            },
            "square": {
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/utilities/cleaning@2x.png",
              "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/utilities/cleaning@3x.png"
            },
            "transparent": {
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/utilities/cleaning@2x.png",
              "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/utilities/cleaning@3x.png"
            }
          }
        }
      ]
    },
    {
      "id": 2,
      "name": "Uncategorized",
      "icon": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square/uncategorized/general.png",
      "icon_types": {
        "slim": {
          "small": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/uncategorized/general.png",
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/uncategorized/general@2x.png"
        },
        "square": {
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/uncategorized/general@2x.png",
          "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/uncategorized/general@3x.png"
        },
        "transparent": {
          "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/uncategorized/general@2x.png",
          "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/uncategorized/general@3x.png"
        }
      },
      "subcategories": [
        {
          "id": 18,
          "name": "General",
          "icon": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square/uncategorized/general.png",
          "icon_types": {
            "slim": {
              "small": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/uncategorized/general.png",
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/slim/uncategorized/general@2x.png"
            },
            "square": {
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/uncategorized/general@2x.png",
              "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/square_v2/uncategorized/general@3x.png"
            },
            "transparent": {
              "large": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/uncategorized/general@2x.png",
              "xlarge": "https://s3.amazonaws.com/splitwise/uploads/category/icon/transparent/uncategorized/general@3x.png"
            }
          }
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

		_, err := c.Categories(context.Background())
		if err != nil {
			t.Fatal(err)
		}
	})
}
