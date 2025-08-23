package models

type Option struct {
	Priority *int    `json:"priority"`
	Title    *string `json:"title"`
	Subject  *string `json:"subject"`
}
