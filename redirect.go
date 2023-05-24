package bigcommerce

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
