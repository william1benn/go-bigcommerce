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

func NewClient(version string, storeHash string, authToken string) Client {
	var client Client
	url, err := url.Parse("https://api.bigcommerce.com/stores/" + storeHash + "/v" + version)
	if err != nil {
		log.Fatal(err)
	}
	client.BaseURL = url
	client.AuthToken = authToken
	client.httpClient = http.DefaultClient
	return client
}

func (c *Client) configureRequest(httpMethod string, relativeUrl string) (*http.Request, error) {
	// Create a GET request
	req, err := http.NewRequest(httpMethod, relativeUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-auth-token", c.AuthToken)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Request(httpMethod string, relativeUrl string) (*http.Response, error) {
	req, err := c.configureRequest(httpMethod, relativeUrl)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
