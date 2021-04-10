package account

import (
	"net/http"
	"time"
)

type errResponse struct {
	Message string `json:"error_message"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		baseURL: "http://localhost:8080/v1",
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
