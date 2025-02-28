package dao

import (
	"context"
	"tiktok_electric_business/errs"
	"time"

	"gorm.io/gorm"
)

type GormUserDAO struct {
	db *gorm.DB
}

var _ UserDAO = (*GormUserDAO)(nil)

func NewGormUserDAO(db *gorm.DB) UserDAO {
	return &GormUserDAO{db: db}
}

// Insert 往用户表中新增一条数据
func (u *GormUserDAO) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.CreatedAt = now
	user.UpdatedAt = now
	err := u.db.WithContext(ctx).Create(&user).Error
	if errs.IsDuplicateError(err) {
		return ErrUserDuplicate
	}
	return err
}

// FindByEmail 根据邮箱查询用户
func (u *GormUserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

// UpdateNonZeroFields 使用非零值字段更新用户表对应 id 的数据
func (u *GormUserDAO) UpdateNonZeroFields(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.UpdatedAt = now
	return u.db.WithContext(ctx).Model(&user).Updates(&user).Error
}

// FindById 根据 id 查询用户
func (u *GormUserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

// FindByPhone 根据电话查询用户
func (u *GormUserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	return user, err
}

// FindByWechat 根据微信 open_id 查询用户
func (u *GormUserDAO) FindByWechat(ctx context.Context, openId string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).Where("wechat_open_id =?", openId).First(&user).Error
	return user, err
}

// DeleteById 根据 id 删除用户
func (u *GormUserDAO) DeleteById(ctx context.Context, id int64) error {
	return u.db.WithContext(ctx).Delete(&User{}, id).Error
}
