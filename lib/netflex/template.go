package netflex

type Template struct {
	ID         Int64  `json:"id,omitempty"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	Type       string `json:"type"`
	TemplateID Int64  `json:"template_id,omitempty"`
	AreaType   string `json:"area_type,omitempty"`
}
