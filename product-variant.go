package bigcommerce

import (
	"encoding/json"
	"fmt"
)

func (client *Client) GetProductVariants(productID int, params ProductVariantQueryParams) ([]ProductVariant, MetaData, error) {
	type ResponseObject struct {
		Data []ProductVariant `json:"data"`
		Meta MetaData         `json:"meta"`
	}
	var response ResponseObject

	queryParams, err := paramString(params)
	if err != nil {
		return response.Data, response.Meta, err
	}

	getProductVariantsURL := client.BaseURL.JoinPath("/catalog/products", fmt.Sprint(productID), "/variants").String() + queryParams

	resp, err := client.Get(getProductVariantsURL)
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

type ProductVariant struct {
	ID                        int             `json:"id"`
	ProductID                 int             `json:"product_id"`
	SKU                       string          `json:"sku"`
	SKUID                     int             `json:"sku_id"`
	Price                     float64         `json:"price"`
	CalculatedPrice           float64         `json:"calculated_price"`
	SalePrice                 float64         `json:"sale_price"`
	RetailPrice               float64         `json:"retail_price"`
	MapPrice                  interface{}     `json:"map_price"`
	Weight                    float64         `json:"weight"`
	CalculatedWeight          float64         `json:"calculated_weight"`
	Width                     float64         `json:"width"`
	Height                    float64         `json:"height"`
	Depth                     float64         `json:"depth"`
	IsFreeShipping            bool            `json:"is_free_shipping"`
	FixedCostShippingPrice    float64         `json:"fixed_cost_shipping_price"`
	PurchasingDisabled        bool            `json:"purchasing_disabled"`
	PurchasingDisabledMessage string          `json:"purchasing_disabled_message"`
	ImageURL                  string          `json:"image_url"`
	CostPrice                 float64         `json:"cost_price"`
	UPC                       string          `json:"upc"`
	MPN                       string          `json:"mpn"`
	GTIN                      string          `json:"gtin"`
	InventoryLevel            int             `json:"inventory_level"`
	InventoryWarningLevel     int             `json:"inventory_warning_level"`
	BinPickingNumber          string          `json:"bin_picking_number"`
	OptionValues              []VariantOption `json:"option_values"`
}

type VariantOption struct {
	ID                int    `json:"id"`
	Label             string `json:"label"`
	OptionID          int    `json:"option_id"`
	OptionDisplayName string `json:"option_display_name"`
}

type ProductVariantQueryParams struct {
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
}
