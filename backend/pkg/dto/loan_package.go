package dto

import "time"

type LoanPackageListReq struct {
	RequestLimit
}

type LoanPackageListResp struct {
	Id          uint64  `json:"id"`
	PackageNo   string  `json:"package_no"`
	Creator     uint64  `json:"creator"`
	CreatorName string  `json:"creator_name"`
	Amount      float64 `json:"amount"`
	Percent     float64 `json:"percent"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PackageNameFilterListReq struct {
	PackageNo string `json:"package_no"`
}

type PackageNameFilterListResp struct {
	Id        uint64  `json:"id"`
	PackageNo string  `json:"package_no"`
	Amount    float64 `json:"amount"`
	Percent   float64 `json:"percent"`
}

type LoanPackageAddReq struct {
	PackageNo string `json:"package_no" binding:"required"`
	//Creator   uint64  `json:"creator" binding:"required"`
	Amount  float64 `json:"amount"`
	Percent float64 `json:"percent"`
}

type LoanPackageEditReq struct {
	Id        uint64 `json:"id" binding:"required"`
	PackageNo string `json:"package_no" binding:"required"`
	//Creator   uint64  `json:"creator" binding:"required"`
	Amount  float64 `json:"amount"`
	Percent float64 `json:"percent"`
}

type LoanPackageDeleteReq struct {
	Ids []uint64 `json:"ids" binding:"required,gte=1"`
}
