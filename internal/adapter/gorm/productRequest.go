package gorm

import (
	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/repository"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ProductRequestGormRepo struct {
	db *gorm.DB
}

func NewProductRequestGormRepo(db *gorm.DB) repository.ProductRequestRepository {
	return &ProductRequestGormRepo{db: db}
}

func (pr *ProductRequestGormRepo) Create(productRequest *domain.ProductRequest, uploadInfo minio.UploadInfo) error {
	uri := uploadInfo.Bucket + "/" + uploadInfo.Key

	productRequest.Image = uri
	result := pr.db.Create(&productRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
