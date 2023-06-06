package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	ID                 int       `json:"id"`
	ParentID           int       `json:"parent_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Views              int       `json:"views"`
	SortOrder          int       `json:"sort_order"`
	PageTitle          string    `json:"page_title"`
	SearchKeywords     string    `json:"search_keywords"`
	MetaKeywords       []string  `json:"meta_keywords"`
	MetaDescription    string    `json:"meta_description"`
	LayoutFile         string    `json:"layout_file"`
	IsVisible          bool      `json:"is_visible"`
	DefaultProductSort string    `json:"default_product_sort"`
	ImageURL           string    `json:"image_url"`
	CustomURL          CustomURL `json:"custom_url"`
}

func (client *Client) GetCategory(id int) (Category, error) {
	type ResponseObject struct {
		Data Category `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	var categoryURL string = client.BaseURL.JoinPath("/catalog/categories", fmt.Sprint(id)).String()

	resp, err := client.Request("GET", categoryURL)
	if err != nil {
		return response.Data, err
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

func (client *Client) GetCategories() ([]Category, MetaData, error) {
	type ResponseObject struct {
		Data []Category `json:"data"`
		Meta MetaData   `json:"meta"`
	}
	var response ResponseObject

	var categoriesURL string = client.BaseURL.JoinPath("/catalog/categories").String()

	resp, err := client.Request("GET", categoriesURL)
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
