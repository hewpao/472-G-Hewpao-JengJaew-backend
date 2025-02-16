package dto

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hewpao/hewpao-backend/types"
)

type CreateProductRequestDTO struct {
	Name     string            `json:"name" validate:"required"`
	Desc     string            `json:"desc" validate:"required"`
	Images   []*fiber.FormFile `form:"images"`
	Budget   float64           `json:"budget" validate:"required,gt=0"`
	Quantity uint              `json:"quantity" validate:"required,gt=0"`
	Category types.Category    `json:"category" validate:"required,category"`
}
