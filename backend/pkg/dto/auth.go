package dto

type LoginReq struct {
	EmailUsername string `json:"email_username" binding:"required"`
	Password      string `json:"password" binding:"required"`
}
