package service

import (
	"errors"
	"tiktok_electric_business/user/repository"
)

var (
	ErrUserDuplicate         = repository.ErrUserDuplicate
	ErrInvalidUserOrPassword = errors.New("the username or password is incorrect")
)
