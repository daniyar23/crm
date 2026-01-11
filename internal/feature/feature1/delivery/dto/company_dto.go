package dto

type CreateCompanyRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type CompanyResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
