package repository

import (
	"context"
	"fmt"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func newUserRepository(c *RepoConfig) *userRepository {
	return &userRepository{
		DB: c.DS.DB,
	}
}

func (r *userRepository) FindByField(ctx context.Context, field, value any) (*model.User, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	user := &model.User{}
	err := db.First(&user, fmt.Sprintf("%s = ?", field), value).Error
	return user, err
}

func (r *userRepository) List(ctx context.Context, req *dto.UserListReq) ([]*dto.UserListResp, int64, error) {
	list := make([]*dto.UserListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	//db.Select("users.*")
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *userRepository) UsernameFilterList(ctx context.Context, req *dto.UsernameFilterListReq) ([]*dto.UsernameFilterListResp, error) {
	list := make([]*dto.UsernameFilterListResp, 0)
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	db.Where("username LIKE ?", "%"+req.Username+"%")
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	return db.Create(&user).Error
}

func (r *userRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *userRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.User{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.User{}).Error
}
