package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id                uint64 `gorm:"column:id;primaryKey" json:"id"`
	Username          string `gorm:"column:username;type:varchar(100);unique;not null" json:"username"`
	UserNRC           string `gorm:"column:user_nrc;unique;not null" json:"user_nrc"`
	UserPhoneNumber   string `gorm:"column:user_phone_number;not null" json:"user_phone_number"`
	ReferenceUserName string `gorm:"column:reference_user_name" json:"reference_user_name"`
	Street            string `gorm:"column:street;not null"json:"street"`
	City              string `gorm:"column:city;not null" json:"city"`
	Region            string `gorm:"column:region;not null" json:"region"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *User) TableName() string {
	return "users"
}
