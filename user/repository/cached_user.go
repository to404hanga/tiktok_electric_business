package repository

import (
	"context"
	"tiktok_electric_business/user/domain"
	"tiktok_electric_business/user/repository/cache"
	"tiktok_electric_business/user/repository/dao"
	"time"

	"github.com/to404hanga/pkg404/logger"
)

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
	l     logger.Logger
}

var _ UserRepository = (*CachedUserRepository)(nil)

func NewCachedUserRepository(dao dao.UserDAO, cache cache.UserCache, l logger.Logger) UserRepository {
	return &CachedUserRepository{
		dao:   dao,
		cache: cache,
		l:     l,
	}
}

// Create 往表中新增一套记录
func (c *CachedUserRepository) Create(ctx context.Context, user domain.User) error {
	return c.dao.Insert(ctx, convertToEntity(user))
}

// FindByPhone 根据电话查询用户
func (c *CachedUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := c.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return convertToDomain(user), err
}

// FindByEmail 根据邮箱查询用户
func (c *CachedUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := c.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return convertToDomain(user), err
}

// UpdateNonZeroFields 使用非零值字段更新用户表对应 id 的数据
//
// 采用延迟 500ms 的双删策略更新缓存
func (c *CachedUserRepository) UpdateNonZeroFields(ctx context.Context, user domain.User) error {
	if err := c.dao.UpdateNonZeroFields(ctx, convertToEntity(user)); err != nil {
		return err
	}
	time.AfterFunc(time.Millisecond*500, func() {
		c.cache.Del(ctx, user.Id)
	})
	return c.cache.Del(ctx, user.Id)
}

// FindById 根据 id 查询用户
//
// 采用异步写缓存方案
func (c *CachedUserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	user, err := c.cache.Get(ctx, id)
	if err == nil {
		return user, nil
	}

	entity, err := c.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	user = convertToDomain(entity)

	go func() {
		if err = c.cache.Set(ctx, user); err != nil {
			c.l.Error("写入缓存失败", logger.Error(err))
		}
	}()

	return user, nil
}

// FindByWechat 根据微信 open_id 查询用户
func (c *CachedUserRepository) FindByWechat(ctx context.Context, openId string) (domain.User, error) {
	user, err := c.dao.FindByWechat(ctx, openId)
	if err != nil {
		return domain.User{}, err
	}
	return convertToDomain(user), err
}

// DeleteById 根据 id 删除用户
func (c *CachedUserRepository) DeleteById(ctx context.Context, id int64) error {
	return c.dao.DeleteById(ctx, id)
}
