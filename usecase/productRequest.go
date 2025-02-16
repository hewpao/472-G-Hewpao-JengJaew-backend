package usecase

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/repository"
	"github.com/minio/minio-go/v7"
)

type ProductRequestUsecase interface {
	CreateProductRequest(productRequest *domain.ProductRequest, files []*multipart.FileHeader, readers []io.Reader) error
}

type productRequestService struct {
	repo      repository.ProductRequestRepository
	minioRepo repository.S3Repository
	ctx       context.Context
}

func NewProductRequestService(repo repository.ProductRequestRepository, minioRepo repository.S3Repository, ctx context.Context) ProductRequestUsecase {
	return &productRequestService{
		repo:      repo,
		minioRepo: minioRepo,
		ctx:       ctx,
	}
}

func (pr *productRequestService) CreateProductRequest(productRequest *domain.ProductRequest, files []*multipart.FileHeader, readers []io.Reader) error {
	uploadInfos := []minio.UploadInfo{}
	for i, file := range files {
		reader := readers[i] // Get the corresponding reader for this file

		uploadInfo, err := pr.minioRepo.UploadFile(pr.ctx, file.Filename, reader, file.Size, file.Header.Get("Content-Type"))
		if err != nil {
			return err
		}

		uploadInfos = append(uploadInfos, uploadInfo)
	}

	err := pr.repo.Create(productRequest, uploadInfos)
	if err != nil {
		return err
	}
	return nil
}
