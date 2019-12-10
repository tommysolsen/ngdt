package netflex

type Page struct {
	ID          Int64   `json:"id"`
	Published   Boolean `json:"published"`
	URL         string  `json:"url"`
	Template    Int64   `json:"template"`
	Language    string  `json:"lang"`
	Description string  `json:"description"`
	Public      Boolean `json:"public"`
}
