package domain

import (
	"github.com/hewpao/hewpao-backend/types"
	"gorm.io/gorm"
)

type ProductRequest struct {
	gorm.Model
	Name     string
	Desc     string
	Image    string
	Budget   float64
	Quantity uint
	Category types.Category `gorm:"type:varchar(20);default:'Other'"`

	UserID *uint
	User   User
	Offers []Offer
}
