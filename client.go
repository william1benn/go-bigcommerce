package bigcommerce

import (
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	AuthToken  string
	httpClient *http.Client
}

func NewClient(version string, storeHash string, authToken string) *Client {
	var client Client
	url, err := url.Parse("https://api.bigcommerce.com/stores/" + storeHash + "/v" + version)
	if err != nil {
		log.Fatal(err)
	}
	client.BaseURL = url
	client.AuthToken = authToken
	client.httpClient = http.DefaultClient
	return &client
}
