package bigcommerce

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Product struct {
	ID                      int       `json:"id"`
	Name                    string    `json:"name"`
	Type                    string    `json:"type"`
	SKU                     string    `json:"sku"`
	Description             string    `json:"description"`
	Weight                  float64   `json:"weight"`
	Width                   float64   `json:"width"`
	Depth                   float64   `json:"depth"`
	Height                  float64   `json:"height"`
	Price                   float64   `json:"price"`
	CostPrice               float64   `json:"cost_price"`
	RetailPrice             float64   `json:"retail_price"`
	SalePrice               float64   `json:"sale_price"`
	MapPrice                float64   `json:"map_price"`
	TaxClassID              float64   `json:"tax_class_id"`
	ProductTaxCode          string    `json:"product_tax_code"`
	CalculatedPrice         float64   `json:"calculated_price"`
	Categories              []int     `json:"categories"`
	BrandID                 int       `json:"brand_id"`
	OptionSetID             int       `json:"option_set_id"`
	OptionSetDisplay        string    `json:"option_set_display"`
	InventoryLevel          int       `json:"inventory_level"`
	InventoryWarningLevel   int       `json:"inventory_warning_level"`
	InventoryTracking       string    `json:"inventory_tracking"`
	ReviewsRatingSum        float64   `json:"reviews_rating_sum"`
	ReviewsCount            int       `json:"reviews_count"`
	TotalSold               int       `json:"total_sold"`
	FixedCostShippingPrice  float64   `json:"fixed_cost_shipping_price"`
	IsFreeShipping          bool      `json:"is_free_shipping"`
	IsVisible               bool      `json:"is_visible"`
	IsFeatured              bool      `json:"is_featured"`
	RelatedProducts         []int     `json:"related_products"`
	Warranty                string    `json:"warranty"`
	BinPickingNumber        string    `json:"bin_picking_number"`
	LayoutFile              string    `json:"layout_file"`
	UPC                     string    `json:"upc"`
	MPN                     string    `json:"mpn"`
	GTIN                    string    `json:"gtin"`
	SearchKeywords          string    `json:"search_keywords"`
	Availability            string    `json:"availability"`
	AvailabilityDescription string    `json:"availability_description"`
	GiftWrappingOptionsType string    `json:"gift_wrapping_options_type"`
	GiftWrappingOptionsList []string  `json:"gift_wrapping_options_list"`
	SortOrder               int       `json:"sort_order"`
	Condition               string    `json:"condition"`
	IsConditionShown        bool      `json:"is_condition_shown"`
	OrderQuantityMinimum    int       `json:"order_quantity_minimum"`
	OrderQuantityMaximum    int       `json:"order_quantity_maximum"`
	PageTitle               string    `json:"page_title"`
	MetaKeywords            []string  `json:"meta_keywords"`
	MetaDescription         string    `json:"meta_description"`
	DateCreated             string    `json:"date_created"`
	DateModified            string    `json:"date_modified"`
	ViewCount               int       `json:"view_count"`
	PreorderReleaseDate     string    `json:"preorder_release_date"`
	PreorderMessage         string    `json:"preorder_message"`
	IsPreorderOnly          bool      `json:"is_preorder_only"`
	IsPriceHidden           bool      `json:"is_price_hidden"`
	PriceHiddenLabel        string    `json:"price_hidden_label"`
	CustomURL               CustomURL `json:"custom_url"`
	BaseVariantID           int       `json:"base_variant_id"`
	OpenGraphType           string    `json:"open_graph_type"`
	OpenGraphTitle          string    `json:"open_graph_title"`
	OpenGraphDescription    string    `json:"open_graph_description"`
	OpenGraphUseMetaDesc    bool      `json:"open_graph_use_meta_description"`
	OpenGraphUseProductName bool      `json:"open_graph_use_product_name"`
	OpenGraphUseImage       bool      `json:"open_graph_use_image"`
}

func (client *Client) GetProduct(id int) (Product, error) {

	type ResponseObject struct {
		Data Product  `json:"data"`
		Meta MetaData `json:"meta"`
	}
	var response ResponseObject

	getProductUrl := client.BaseURL.JoinPath("/catalog/products/", fmt.Sprint(id)).String()

	// Send the request
	resp, err := client.Request("GET", getProductUrl)
	if err != nil {
		return response.Data, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return response.Data, errors.New(
			"API responded with a non 200 status code",
		)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	// Read the response body
	if err != nil {
		return response.Data, err
	}

	return response.Data, nil
}

func (client *Client) GetAllProducts() ([]Product, MetaData, error) {
	type ResponseObject struct {
		Data []Product `json:"data"`
		Meta MetaData  `json:"meta"`
	}
	var response ResponseObject

	getProductsUrl := client.BaseURL.JoinPath("/catalog/products").String()

	resp, err := client.Request("GET", getProductsUrl)
	if err != nil {
		return response.Data, response.Meta, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return response.Data, response.Meta, errors.New("API responded with a non 200 status code")
	}

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return response.Data, response.Meta, err
	}

	return response.Data, response.Meta, nil
}

type ProductQueryParams struct {
	ID                    int      `url:"id,omitempty"`
	IDIn                  []int    `url:"id:in,omitempty,comma"`
	IDNotIn               []int    `url:"id:not_in,omitempty,comma"`
	IDMin                 []int    `url:"id:min,omitempty,comma"`
	IDMax                 []int    `url:"id:max,omitempty,comma"`
	IDGreater             []int    `url:"id:greater,omitempty,comma"`
	IDLess                []int    `url:"id:less,omitempty,comma"`
	Name                  string   `url:"name,omitempty"`
	UPC                   string   `url:"upc,omitempty"`
	Price                 float64  `url:"price,omitempty"`
	Weight                float64  `url:"weight,omitempty"`
	Condition             string   `url:"condition,omitempty"`
	BrandID               int      `url:"brand_id,omitempty"`
	DateModified          string   `url:"date_modified,omitempty"`
	DateModifiedMax       string   `url:"date_modified:max,omitempty"`
	DateModifiedMin       string   `url:"date_modified:min,omitempty"`
	DateLastImported      string   `url:"date_last_imported,omitempty"`
	DateLastImportedMax   string   `url:"date_last_imported:max,omitempty"`
	DateLastImportedMin   string   `url:"date_last_imported:min,omitempty"`
	IsVisible             bool     `url:"is_visible,omitempty"`
	IsFeatured            int      `url:"is_featured,omitempty"`
	IsFreeShipping        int      `url:"is_free_shipping,omitempty"`
	InventoryLevel        int      `url:"inventory_level,omitempty"`
	InventoryLevelIn      []int    `url:"inventory_level:in,omitempty,comma"`
	InventoryLevelNotIn   []int    `url:"inventory_level:not_in,omitempty,comma"`
	InventoryLevelMin     []int    `url:"inventory_level:min,omitempty,comma"`
	InventoryLevelMax     []int    `url:"inventory_level:max,omitempty,comma"`
	InventoryLevelGreater []int    `url:"inventory_level:greater,omitempty,comma"`
	InventoryLevelLess    []int    `url:"inventory_level:less,omitempty,comma"`
	InventoryLow          int      `url:"inventory_low,omitempty"`
	OutOfStock            int      `url:"out_of_stock,omitempty"`
	TotalSold             int      `url:"total_sold,omitempty"`
	Type                  string   `url:"type,omitempty"`
	Categories            int      `url:"categories,omitempty"`
	Keyword               string   `url:"keyword,omitempty"`
	KeywordContext        string   `url:"keyword_context,omitempty"`
	Status                int      `url:"status,omitempty"`
	Include               string   `url:"include,omitempty"`
	IncludeFields         string   `url:"include_fields,omitempty"`
	ExcludeFields         string   `url:"exclude_fields,omitempty"`
	Availability          string   `url:"availability,omitempty"`
	Page                  int      `url:"page,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	Direction             string   `url:"direction,omitempty"`
	Sort                  string   `url:"sort,omitempty"`
	CategoriesIn          []int    `url:"categories:in,omitempty,comma"`
	SKU                   string   `url:"sku,omitempty"`
	SKUIn                 []string `url:"sku:in,omitempty,comma"`
}
