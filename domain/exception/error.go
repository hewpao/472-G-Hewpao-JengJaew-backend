package exception

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exist")

	ErrUserNoPassword    = errors.New("account has no password, probably using social login")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrFileSizeLimit     = errors.New("file size limit exceeded")
	ErrTypeConversion    = errors.New("could not convert variable type")
	ErrFileIsNull        = errors.New("file could not be null")
	ErrFileCountLimit    = errors.New("file count limit exceeded")
)
