package bigcommerce

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return response.Data, errors.New("resource not found")
		}
		return response.Data, errors.New("something went wrong")
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

	if resp.StatusCode != 200 {
		return response.Data, response.Meta, errors.New("API responded with a non 200 status code")
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil

}
