package bigcommerce

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
	Weight                    int             `json:"weight"`
	CalculatedWeight          int             `json:"calculated_weight"`
	Width                     int             `json:"width"`
	Height                    int             `json:"height"`
	Depth                     int             `json:"depth"`
	IsFreeShipping            bool            `json:"is_free_shipping"`
	FixedCostShippingPrice    int             `json:"fixed_cost_shipping_price"`
	PurchasingDisabled        bool            `json:"purchasing_disabled"`
	PurchasingDisabledMessage string          `json:"purchasing_disabled_message"`
	ImageURL                  string          `json:"image_url"`
	CostPrice                 int             `json:"cost_price"`
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
