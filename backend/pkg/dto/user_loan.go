package dto

import "time"

type UserLoanReq struct {
	RequestLimit
}

type UserLoanListResp struct {
	Id                   uint64    `json:"id"`
	LoanPackageId        uint64    `json:"loan_package_id"`
	LoanPackageNo        string    `json:"loan_package_no"`
	LoanPackagePercent   float64   `json:"loan_package_percent"`
	LoanUserId           uint64    `json:"loan_user_id"`
	LoanUsername         string    `json:"loan_username"`
	InitializeLoanAmount float64   `json:"initialize_loan_amount"`
	CurrentLoanAmount    float64   `json:"current_loan_amount"`
	PaidLoanAmount       float64   `json:"paid_loan_amount"`
	PaidCount            uint32    `json:"paid_count"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type UserLoanAddReq struct {
	LoanPackageId        uint64  `json:"loan_package_id" binding:"required"`
	LoanPackageNo        string  `json:"loan_package_no" binding:"required"`
	LoanPackagePercent   string  `json:"loan_package_percent" binding:"required"`
	LoanUserId           uint64  `json:"loan_user_id" binding:"required"`
	LoanUsername         uint64  `json:"loan_username" binding:"required"`
	InitializeLoanAmount float64 `json:"initialize_loan_amount" binding:"required"`
}
