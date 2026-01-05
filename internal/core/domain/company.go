package domain

type Company struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
