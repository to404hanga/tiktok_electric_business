package cache

import "errors"

var (
	ErrCodeSendTooMany   = errors.New("the captcha is sent too frequently")
	ErrUnknowForCode     = errors.New("an unknown error occurred while sending the captcha")
	ErrCodeVerifyTooMany = errors.New("excessive verification")
)
