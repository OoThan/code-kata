package model

import (
	"gorm.io/gorm"
	"time"
)

type UserLoan struct {
	Id                   uint64  `gorm:"column:id;primaryKey" json:"id"`
	LoanPackageId        uint64  `gorm:"column:loan_package_id" json:"loan_package_id"`
	LoanPackageNo        string  `gorm:"column:loan_package_no" json:"loan_package_no"`
	LoanPackagePercent   float64 `gorm:"column:loan_package_percent" json:"loan_package_percent"`
	LoanUserId           uint64  `gorm:"column:loan_user_id" json:"loan_user_id"`
	LoanUsername         string  `gorm:"column:loan_username" json:"loan_username"`
	InitializeLoanAmount float64 `gorm:"column:initialize_loan_amount" json:"initialize_loan_amount"`
	CurrentLoanAmount    float64 `gorm:"column:current_loan_amount" json:"current_loan_amount"`
	PaidLoanAmount       float64 `gorm:"column:paid_loan_amount" json:"paid_loan_amount"`
	PaidCount            float64 `gorm:"column:paid_count" json:"paid_count"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *UserLoan) TableName() string {
	return "user_loans"
}

type UserLoanLog struct {
	Id                   uint64  `gorm:"column:id;primaryKey" json:"id"`
	UserLoanId           uint64  `gorm:"column:user_loan_id" json:"user_loan_id"`
	LoanPackageId        uint64  `gorm:"column:loan_package_id" json:"loan_package_id"`
	LoanPackageNo        string  `gorm:"column:loan_package_no" json:"loan_package_no"`
	LoanPackagePercent   float64 `gorm:"column:loan_package_percent" json:"loan_package_percent"`
	LoanUserId           uint64  `gorm:"column:loan_user_id" json:"loan_user_id"`
	LoanUsername         string  `gorm:"column:loan_username" json:"loan_username"`
	InitializeLoanAmount float64 `gorm:"column:initialize_loan_amount" json:"initialize_loan_amount"`
	CurrentLoanAmount    float64 `gorm:"column:current_loan_amount" json:"current_loan_amount"`
	PaidLoanAmount       float64 `gorm:"column:paid_loan_amount" json:"paid_loan_amount"`
	PaidCount            uint32  `gorm:"column:paid_count" json:"paid_count"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (m *UserLoanLog) TableName() string {
	return "user_loan_logs"
}
