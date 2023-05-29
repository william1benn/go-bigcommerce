package bigcommerce

type Brand struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MetaKeywords    []string  `json:"meta_keywords"`
	MetaDescription string    `json:"meta_description"`
	ImageURL        string    `json:"image_url"`
	SearchKeywords  string    `json:"search_keywords"`
	CustomURL       CustomURL `json:"custom_url"`
}
