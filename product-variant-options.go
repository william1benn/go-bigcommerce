package bigcommerce

type ProductVariantOption struct {
	ID           int           `json:"id"`
	ProductID    int           `json:"product_id"`
	DisplayName  string        `json:"display_name"`
	Type         string        `json:"type"`
	Config       OptionConfig  `json:"config"`
	SortOrder    int           `json:"sort_order"`
	OptionValues []OptionValue `json:"option_values"`
	Name         string        `json:"name"`
}

type OptionConfig struct {
	DefaultValue                string   `json:"default_value"`
	CheckedByDefault            bool     `json:"checked_by_default"`
	CheckboxLabel               string   `json:"checkbox_label"`
	DateLimited                 bool     `json:"date_limited"`
	DateLimitMode               string   `json:"date_limit_mode"`
	DateEarliestValue           string   `json:"date_earliest_value"`
	DateLatestValue             string   `json:"date_latest_value"`
	FileTypesMode               string   `json:"file_types_mode"`
	FileTypesSupported          []string `json:"file_types_supported"`
	FileTypesOther              []string `json:"file_types_other"`
	FileMaxSize                 int      `json:"file_max_size"`
	TextCharactersLimited       bool     `json:"text_characters_limited"`
	TextMinLength               int      `json:"text_min_length"`
	TextMaxLength               int      `json:"text_max_length"`
	TextLinesLimited            bool     `json:"text_lines_limited"`
	TextMaxLines                int      `json:"text_max_lines"`
	NumberLimited               bool     `json:"number_limited"`
	NumberLimitMode             string   `json:"number_limit_mode"`
	NumberLowestValue           int      `json:"number_lowest_value"`
	NumberHighestValue          int      `json:"number_highest_value"`
	NumberIntegersOnly          bool     `json:"number_integers_only"`
	ProductListAdjustsInventory bool     `json:"product_list_adjusts_inventory"`
	ProductListAdjustsPricing   bool     `json:"product_list_adjusts_pricing"`
	ProductListShippingCalc     string   `json:"product_list_shipping_calc"`
}

type OptionValue struct {
	IsDefault bool        `json:"is_default"`
	Label     string      `json:"label"`
	SortOrder int         `json:"sort_order"`
	ValueData interface{} `json:"value_data"`
	ID        int         `json:"id"`
}
