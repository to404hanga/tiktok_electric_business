package service

import (
	"context"
	"tiktok_electric_business/user/domain"
	"tiktok_electric_business/user/repository"

	"github.com/to404hanga/pkg404/logger"
	"golang.org/x/crypto/bcrypt"
)

type BaseUserService struct {
	repo repository.UserRepository
	l    logger.Logger
}

var _ UserService = (*BaseUserService)(nil)

func NewBaseUserService(repo repository.UserRepository, l logger.Logger) UserService {
	return &BaseUserService{
		repo: repo,
		l:    l,
	}
}

// FindOrCreateByPhone 若存在，返回该用户，若不存在，则根据电话号码创建一个新用户并返回
func (b *BaseUserService) FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := b.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		return user, err
	}

	b.l.Info("New user registered by phone", logger.SafePhoneZH(phone))

	err = b.repo.Create(ctx, domain.User{
		Phone: phone,
	})
	// err != repository.ErrUserDuplicate 防止其他 goroutine 提前创建用户意外返回冲突
	if err != nil && err != repository.ErrUserDuplicate {
		return domain.User{}, err
	}
	return b.repo.FindByPhone(ctx, phone)
}

// FindOrCreateByWechat 若存在，返回该用户，若不存在，则根据微信 openID 创建一个新用户并返回
func (b *BaseUserService) FindOrCreateByWechat(ctx context.Context, info domain.WechatInfo) (domain.User, error) {
	user, err := b.repo.FindByWechat(ctx, info.OpenId)
	if err != repository.ErrUserNotFound {
		return user, err
	}

	b.l.Info("New user registered by wechat", logger.Any("wechat_info", info))

	err = b.repo.Create(ctx, domain.User{
		WechatInfo: info,
	})
	// err != repository.ErrUserDuplicate 防止其他 goroutine 提前创建用户意外返回冲突
	if err != nil && err != repository.ErrUserDuplicate {
		return domain.User{}, err
	}

	return b.repo.FindByWechat(ctx, info.OpenId)
}

// SignUp 新用户注册（邮箱）
func (b *BaseUserService) SignUp(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		b.l.Error("bcrypt generate password error", logger.Error(err))
		return err
	}
	user.Password = string(hash)
	return b.repo.Create(ctx, user)
}

// Login 用户登陆（邮箱）
func (b *BaseUserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	user, err := b.repo.FindByEmail(ctx, email)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return domain.User{}, ErrInvalidUserOrPassword
		}
		return domain.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return user, nil
}

// UpdateNonSensitiveInfo 更新非敏感数据（昵称、用户类型）
func (b *BaseUserService) UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error {
	// 仅传入可修改字段以及主键
	return b.repo.UpdateNonZeroFields(ctx, domain.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		UserType: user.UserType,
	})
}

// Profile 获取用户信息（包括所有敏感数据）
func (b *BaseUserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	return b.repo.FindById(ctx, id)
}
