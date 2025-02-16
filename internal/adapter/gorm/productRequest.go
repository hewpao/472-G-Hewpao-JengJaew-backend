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

func (pr *ProductRequestGormRepo) Create(productRequest *domain.ProductRequest, uploadInfos []minio.UploadInfo) error {
	uris := []string{}

	for _, uploadInfo := range uploadInfos {
		uri := uploadInfo.Bucket + "/" + uploadInfo.Key
		uris = append(uris, uri)
	}
	productRequest.Images = uris
	result := pr.db.Create(&productRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
