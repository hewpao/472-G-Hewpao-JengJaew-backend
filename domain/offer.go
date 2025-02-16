package domain

import (
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	ProductRequestID *uint
	ProductRequest   *ProductRequest
	AccountID        *string
	Account          *Account
}
