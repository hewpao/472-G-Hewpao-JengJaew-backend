package dto

import (
	"mime/multipart"

	"github.com/hewpao/hewpao-backend/types"
)

type CreateProductRequestDTO struct {
	Name     string                `json:"name" validate:"required"`
	Desc     string                `json:"desc" validate:"required"`
	Image    *multipart.FileHeader `form:"image"`
	Budget   float64               `json:"budget" validate:"required,gt=0"`
	Quantity uint                  `json:"quantity" validate:"required,gt=0"`
	Category types.Category        `json:"category" validate:"required,category"`
}
