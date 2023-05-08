package repository

import (
	"context"
	"fmt"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/logger"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/utils"

	"gorm.io/gorm"
)

type loanPackageRepository struct {
	DB *gorm.DB
}

func newLoanPackageRepository(c *RepoConfig) *loanPackageRepository {
	return &loanPackageRepository{
		DB: c.DS.DB,
	}
}

func (r *loanPackageRepository) FindByField(ctx context.Context, field, value any) (*model.LoanPackage, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	loanPkg := &model.LoanPackage{}
	err := db.First(&loanPkg, fmt.Sprintf("%s = ?", field), value).Error
	return loanPkg, err
}

func (r *loanPackageRepository) List(ctx context.Context, req *dto.LoanPackageListReq) ([]*dto.LoanPackageListResp, int64, error) {
	list := make([]*dto.LoanPackageListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	db.Select("loan_packages.*, admins.username as creator_name")
	db.Joins("inner join admins on admins.id = loan_packages.creator")
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *loanPackageRepository) PackageNoFilterList(ctx context.Context, req *dto.PackageNameFilterListReq) ([]*dto.PackageNameFilterListResp, error) {
	list := make([]*dto.PackageNameFilterListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	db.Where("package_no LIKE ?", "%"+req.PackageNo+"%")
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *loanPackageRepository) Create(ctx context.Context, loanPkg *model.LoanPackage) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	return db.Create(&loanPkg).Error
}

func (r *loanPackageRepository) CreateLoanPkgLog(ctx context.Context, log *model.LoanPackageLog) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackageLog{})
	return db.Create(log).Error
}

func (r *loanPackageRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	tx := r.DB.Begin()
	db := tx.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	loanPkg := &model.LoanPackage{}
	db.Where(updateFields.Field, updateFields.Value).First(&loanPkg)
	if err := db.Updates(&updateFields.Data).Error; err != nil {
		tx.Rollback()
		return err
	}

	log := &model.LoanPackageLog{
		PackageNo:     loanPkg.PackageNo,
		Creator:       loanPkg.Creator,
		BeforeAmount:  loanPkg.Amount,
		AfterAmount:   updateFields.Data["amount"].(float64),
		BeforePercent: loanPkg.Percent,
		AfterPercent:  updateFields.Data["percent"].(float64),
	}

	logger.Sugar.Debugf("%+v", log)
	if err := tx.WithContext(ctx).Debug().Model(&model.LoanPackageLog{}).Create(&log).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *loanPackageRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.LoanPackage{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.LoanPackage{}).Error
}
