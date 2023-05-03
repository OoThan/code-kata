package model

import (
	"gorm.io/gorm"
	"time"
)

type LoanPackage struct {
	Id        uint64         `gorm:"column:id;primaryKey" json:"id"`
	PackageNo string         `gorm:"column:package_no" json:"package_no"`
	Creator   uint64         `gorm:"column:creator;not null" json:"creator"`
	Amount    float64        `gorm:"column:amount" json:"amount"`
	Percent   float64        `gorm:"column:percent" json:"percent"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *LoanPackage) TableName() string {
	return "loan_packages"
}

type LoanPackageLog struct {
	Id        uint64 `gorm:"column:id;primaryKey" json:"id"`
	PackageNo string `gorm:"column:package_no" json:"package_no"`
	Creator   uint64 `gorm:"column:creator;not null" json:"creator"`

	BeforeAmount float64 `gorm:"column:before_amount" json:"before_amount"`
	AfterAmount  float64 `gorm:"column:after_amount" json:"after_amount"`

	BeforePercent float64 `gorm:"column:before_percent" json:"before_percent"`
	AfterPercent  float64 `gorm:"column:after_percent" json:"after_percent"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *LoanPackageLog) TableName() string {
	return "loan_package_logs"
}
