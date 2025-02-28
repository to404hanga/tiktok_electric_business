package dao

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("the user's email or phone conflicts")
	ErrDataNotFound  = gorm.ErrRecordNotFound
)
