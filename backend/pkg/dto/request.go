package dto

type RequestLimit struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
