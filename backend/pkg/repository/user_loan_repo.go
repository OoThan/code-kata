package repository

import (
	"context"
	"fmt"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/utils"

	"gorm.io/gorm"
)

type userLoanRepository struct {
	DB *gorm.DB
}

func newUserLoanRepository(c *RepoConfig) *userLoanRepository {
	return &userLoanRepository{
		DB: c.DS.DB,
	}
}

func (r *userLoanRepository) FindByField(ctx context.Context, field, value any) (*model.UserLoan, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.UserLoan{})
	userLoan := &model.UserLoan{}
	err := db.First(&userLoan, fmt.Sprintf("%s = ?", field), value).Error
	return userLoan, err
}

func (r *userLoanRepository) List(ctx context.Context, req *dto.UserListReq) ([]*dto.UserLoanListResp, int64, error) {
	list := make([]*dto.UserLoanListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.UserLoan{})
	//db.Select("user_loans.*")
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).First(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *userLoanRepository) Create(ctx context.Context, userLoan *model.UserLoan) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.UserLoan{})
	return db.Create(&userLoan).Error
}

func (r *userLoanRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.UserLoan{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *userLoanRepository) Deletes(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.UserLoan{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.UserLoan{}).Error
}
