package dto

import "time"

type AdminListReq struct {
	RequestLimit
}

type AdminListResp struct {
	Id           uint64    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Status       int       `json:"status"`
	IP           string    `json:"ip"`
	Activated2fa bool      `json:"activated_2fa"`
	SecretKey2fa string    `json:"secret_key_2fa"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
