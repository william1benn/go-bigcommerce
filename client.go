package bigcommerce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (c *Client) getProductById(id int) (Product, error) {
	var product Product

	type ResponseObject struct {
		Data Product  `json:"data"`
		Meta MetaData `json:"meta"`
	}

	getProductUrl := c.BaseURL.JoinPath("/catalog/products/", fmt.Sprint(id)).String()

	// Create a GET request
	req, err := http.NewRequest("GET", getProductUrl, nil)
	if err != nil {
		return product, err
	}

	req.Header.Set("x-auth-token", c.AuthToken)

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return product, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return product, err
	}

	var response ResponseObject

	// Print the response
	err = json.Unmarshal(body, &response)

	if err != nil {
		return product, err
	}

	product = response.Data

	return product, nil
}
