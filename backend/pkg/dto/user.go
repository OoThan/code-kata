package dto

import "time"

type UserListReq struct {
	RequestLimit
}

type UserListResp struct {
	Id                uint64 `json:"id"`
	Username          string `json:"username"`
	UserNRC           string `json:"user_nrc"`
	UserPhoneNumber   string `json:"user_phone_number"`
	ReferenceUserName string `json:"reference_user_name"`
	Street            string `json:"street"`
	City              string `json:"city"`
	Region            string `json:"region"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAddReq struct {
	Username          string `json:"username" binding:"required"`
	UserNRC           string `json:"user_nrc" binding:"required"`
	UserPhoneNumber   string `json:"user_phone_number" binding:"required"`
	ReferenceUserName string `json:"reference_user_name" binding:"required"`
	Street            string `json:"street" binding:"required"`
	City              string `json:"city" binding:"required"`
	Region            string `json:"region" binging:"required"`
}

type UserEditReq struct {
	Id                uint64 `json:"id" binding:"required"`
	UserNRC           string `json:"user_nrc" binding:"required"`
	UserPhoneNumber   string `json:"user_phone_number" binding:"required"`
	ReferenceUserName string `json:"reference_user_name" binding:"required"`
	Street            string `json:"street" binding:"required"`
	City              string `json:"city" binding:"required"`
	Region            string `json:"region" binging:"required"`
}

type UserDeleteReq struct {
	Ids []uint64 `json:"ids" binding:"required,gre=1"`
}
