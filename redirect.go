package bigcommerce

import "encoding/json"

type Redirect struct {
	ID       int      `json:"id"`
	SiteID   int      `json:"site_id"`
	FromPath string   `json:"from_path"`
	To       ToObject `json:"to"`
	ToURL    string   `json:"to_url"`
}

type ToObject struct {
	Type     string `json:"type"`
	EntityID int    `json:"entity_id"`
	URL      string `json:"url"`
}

func (client *Client) GetRedirects(params RedirectQueryParams) ([]Redirect, error) {
	type ResponseObject struct {
		Data []Redirect `json:"data"`
		Meta MetaData   `json:"meta"`
	}
	var response ResponseObject

	queryParams, err := paramString(params)
	if err != nil {
		return response.Data, err
	}

	getRedirectsURL := client.BaseURL.JoinPath("/storefront/redirects").String() + queryParams

	resp, err := client.Get(getRedirectsURL)
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

type RedirectQueryParams struct {
	SiteID    int    `url:"ite_id,omitempty"`
	IDs       []int  `url:"id,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Page      int    `url:"page,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	Include   string `url:"include,omitempty"`
	Keyword   string `url:"keyword,omitempty"`
}
