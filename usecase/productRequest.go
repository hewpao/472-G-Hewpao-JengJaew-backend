package usecase

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/repository"
)

type ProductRequestUsecase interface {
	CreateProductRequest(productRequest *domain.ProductRequest, file *multipart.FileHeader, reader io.Reader) error
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

func (pr *productRequestService) CreateProductRequest(productRequest *domain.ProductRequest, file *multipart.FileHeader, reader io.Reader) error {
	uploadInfo, err := pr.minioRepo.UploadFile(pr.ctx, file.Filename, reader, file.Size, file.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	err = pr.repo.Create(productRequest, uploadInfo)
	if err != nil {
		return err
	}
	return nil
}
