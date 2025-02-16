package rest

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/dto"
	"github.com/hewpao/hewpao-backend/usecase"
	"github.com/hewpao/hewpao-backend/util"
)

type ProductRequestHandler interface {
	CreateProductRequest(c *fiber.Ctx) error
}

type productRequestHandler struct {
	service usecase.ProductRequestUsecase
}

func NewProductRequestHandler(service usecase.ProductRequestUsecase) ProductRequestHandler {
	return &productRequestHandler{
		service: service,
	}
}

func (pr *productRequestHandler) CreateProductRequest(c *fiber.Ctx) error {
	fileHeaders, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	files := fileHeaders.File["images"] // "images" should match the form field name for multiple file uploads
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No files uploaded")
	}

	var fileReaders []io.Reader
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		content, err := io.ReadAll(file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		reader := bytes.NewReader(content)
		fileReaders = append(fileReaders, reader)

		defer file.Close()
	}

	req := dto.CreateProductRequestDTO{}
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	validationErr := util.ValidateStruct(req)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(validationErr.Error)
	}

	productRequest := domain.ProductRequest{
		Name:     req.Name,
		Desc:     req.Desc,
		Budget:   req.Budget,
		Quantity: req.Quantity,
		Category: req.Category,
		Offers:   []domain.Offer{},
		Images:   []string{},
	}

	err = pr.service.CreateProductRequest(&productRequest, files, fileReaders)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message":         "Product request created sucessfully",
		"product-request": productRequest,
	})
}
