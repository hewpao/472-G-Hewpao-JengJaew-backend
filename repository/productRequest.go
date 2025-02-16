package repository

import (
	"github.com/hewpao/hewpao-backend/domain"
	"github.com/minio/minio-go/v7"
)

type ProductRequestRepository interface {
	Create(productRequest *domain.ProductRequest, uploadInfo minio.UploadInfo) error
}
