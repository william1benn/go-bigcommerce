package bigcommerce

type CustomURL struct {
	URL          string `json:"url"`
	IsCustomized bool   `json:"is_customized"`
}

type MetaData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Total       int   `json:"total"`
	Count       int   `json:"count"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	Links       Links `json:"links"`
}

type Links struct {
	Current string `json:"current"`
}
