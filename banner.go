package bigcommerce

import (
	"encoding/json"
	"fmt"
)

type Banner struct {
	ID          int    `json:"id"`
	DateCreated string `json:"date_created"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	Page        string `json:"page"`
	Location    string `json:"location"`
	DateType    string `json:"date_type"`
	DateFrom    string `json:"date_from,omitempty"`
	DateTo      string `json:"date_to,omitempty"`
	Visible     string `json:"visible"`
	ItemID      string `json:"item_id,omitempty"`
}
type GetBannersParams struct {
	MinID int `url:"min_id,omitempty"`
	MaxID int `url:"max_id,omitempty"`
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

func (client *Client) GetBanners(params GetBannersParams) ([]Banner, MetaData, error) {
	type ResponseObject struct {
		Data []Banner `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	if client.Version != "2" {
		return response.Data, response.Meta, fmt.Errorf("need to be using version 2 api for this function")
	}

	queryParams, err := paramString(params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	path := client.BaseURL.JoinPath("banners").String() + queryParams

	resp, err := client.Get(path)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	err = expectStatusCode(200, resp)
	if err != nil {
		return response.Data, response.Meta, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}
