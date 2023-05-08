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
	Activated2fa bool      `json:"activated2fa"`
	SecretKey2fa string    `json:"secret_key2fa"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AdminAddReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Status   int    `json:"status"`
	Password string `json:"password" binding:"required"`
}

type AdminEditReq struct {
	Id       uint64 `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Status   int    `json:"status"`
	Password string `json:"password"`
}

type AdminDeleteReq struct {
	Ids []uint64 `json:"ids" binding:"required,gte=1"`
}
