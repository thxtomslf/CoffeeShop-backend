package entities

type ShopInfo struct {
	Title        string   `json:"title"`
	OpeningHours string   `json:"opening_hours"`
	Barista      []string `json:"barista"`
	IsOpen       bool     `json:"is_open"`
	Number       string   `json:"number"`
}
