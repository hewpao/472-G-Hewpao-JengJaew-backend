package domain

import (
	"github.com/hewpao/hewpao-backend/types"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProductRequest struct {
	gorm.Model
	Name     string
	Desc     string
	Images   pq.StringArray `gorm:"type:text[]"`
	Budget   float64
	Quantity uint
	Category types.Category `gorm:"type:varchar(20);default:'Other'"`

	AccountID *string
	Account   *Account
	Offers    []Offer `gorm:"foreignKey:ProductRequestID"`
}
