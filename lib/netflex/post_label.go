package netflex

// PostLabelPayload is the payload expected by the postLabel endpoint
type PostLabelPayload struct {
	Label string `json:"label"`
}
