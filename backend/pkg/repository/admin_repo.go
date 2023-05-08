package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/utils"
	"time"
)

type adminRepository struct {
	DB *gorm.DB
}

func newAdminRepository(c *RepoConfig) *adminRepository {
	return &adminRepository{
		DB: c.DS.DB,
	}
}

func (r *adminRepository) FindByField(ctx context.Context, field, value any) (*model.Admin, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	admin := &model.Admin{}
	err := db.First(&admin, fmt.Sprintf("%s = ?", field), value).Error
	return admin, err
}

func (r *adminRepository) FindOrByField(ctx context.Context, field1, field2, value any) (*model.Admin, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	admin := &model.Admin{}
	err := db.First(&admin, fmt.Sprintf("%s = ? OR %s =?", field1, field2), value, value).Error
	return admin, err
}

func (r *adminRepository) List(ctx context.Context, req *dto.AdminListReq) ([]*dto.AdminListResp, int64, error) {
	list := make([]*dto.AdminListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	//db.Select("admins.*")
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *adminRepository) Create(ctx context.Context, admin *model.Admin) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	var (
		count      int64
		checkAdmin model.Admin
	)
	db.Unscoped().Where("username = ? AND deleted_at IS NOT NULL", admin.Username).Count(&count)
	if count > 0 {
		db.First(&checkAdmin)
		admin.Id = checkAdmin.Id
		admin.CreatedAt = time.Now()
		admin.UpdatedAt = time.Now()
		admin.DeletedAt = gorm.DeletedAt{
			Time:  time.Time{},
			Valid: false,
		}
		return db.Save(&admin).Error
	}

	return db.Create(&admin).Error
}

func (r *adminRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *adminRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Admin{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.Admin{}).Error
}
