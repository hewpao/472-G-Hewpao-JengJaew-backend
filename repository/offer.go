package repository

import (
	"github.com/hewpao/hewpao-backend/domain"
)

type OfferRepository interface {
	Create(req *domain.Offer) error
}
