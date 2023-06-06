package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type Brand struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MetaKeywords    []string  `json:"meta_keywords"`
	MetaDescription string    `json:"meta_description"`
	ImageURL        string    `json:"image_url"`
	SearchKeywords  string    `json:"search_keywords"`
	CustomURL       CustomURL `json:"custom_url"`
}

func (client *Client) GetBrand(id int) (Brand, error) {
	type ResponseObject struct {
		Data Brand    `json:"data"`
		Meta MetaData `json:"meta"`
	}

	var response ResponseObject

	brandURL := client.BaseURL.JoinPath("/catalog/brands", fmt.Sprint(id)).String()

	resp, err := client.Request("GET", brandURL)
	if err != nil {
		return response.Data, nil
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *Client) GetBrands() ([]Brand, MetaData, error) {
	type ResponseObject struct {
		Data []Brand  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	brandsURL := client.BaseURL.JoinPath("/catalog/brands").String()

	resp, err := client.Request("GET", brandsURL)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	if err = expectStatusCode(200, resp); err != nil {
		return response.Data, response.Meta, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil

}
