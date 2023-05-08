package model

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id           uint64         `gorm:"column:id;primaryKey" json:"id"`
	Username     string         `gorm:"column:username;unique;not null" json:"username"`
	Email        string         `gorm:"column:email;unique;not null" json:"email"`
	Status       int            `gorm:"column:status;default:1" json:"status"`
	Password     string         `gorm:"column:password" json:"-"`
	IP           string         `gorm:"column:ip;type:varchar(20)" json:"ip"`
	Location     string         `gorm:"column:location;type:varchar(255)" json:"location"`
	Activated2fa bool           `gorm:"column:activated2fa;default:false;not null" json:"activated2fa"`
	SecretKey2fa string         `gorm:"column:secret_key2fa;size:191" json:"secret_key2fa"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (m *Admin) TableName() string {
	return "admins"
}

type AdminLog struct {
	Id        uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AdminId   uint64         `gorm:"column:admin_id"`
	Username  string         `gorm:"column:username" json:"username"`
	LoginUrl  string         `gorm:"column:login_url" json:"login_url"`
	IP        string         `gorm:"column:ip" json:"ip"`
	Location  string         `gorm:"column:location;type:varchar(255)" json:"location"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *AdminLog) TableName() string {
	return "admin_logs"
}
