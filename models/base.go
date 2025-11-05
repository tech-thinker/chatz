package models

// Option holds optional parameters for messaging operations.
type Option struct {
	Priority *int    `json:"priority"`
	Title    *string `json:"title"`
	Subject  *string `json:"subject"`
}
