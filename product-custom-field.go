package bigcommerce

type ProductCustomField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetCustomFields(productID int, params ProductCustomFieldsRequestParams) {
	// /catalog/products/{product_id}/custom-fields
}
func CreateCustomFields(productID int, params CreateCustomFieldParams) {}
func GetCustomField(productID int, customFieldID int) {
	// /catalog/products/{product_id}/custom-fields/{custom_field_id}
}
func UpdateCustomField(productID int, customFieldID int, params UpdateCustomFieldParams) {}
func DeleteCustomField(productID int, customFieldID int)                                 {}

type ProductCustomFieldsRequestParams struct {
	IncludeFields string `url:"include_fields,omitempty"`
	ExcludeFields string `url:"exclude_fields,omitempty"`
	Page          int    `url:"page,omitempty"`
	Limit         int    `url:"limit,omitempty"`
}

type CreateCustomFieldParams struct {
	Name  string `json:"name" validate:"required,min=1,max=250"`
	Value string `json:"value" validate:"required,min=1,max=250"`
}

type UpdateCustomFieldParams struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=1,max=250"`
	Value string `json:"value" validate:"required,min=1,max=250"`
}
