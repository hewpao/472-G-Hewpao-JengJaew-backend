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
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	reader := bytes.NewReader(content)
	defer file.Close()

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
	}

	err = pr.service.CreateProductRequest(&productRequest, fileHeader, reader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message":         "Product request created sucessfully",
		"product-request": productRequest,
	})
}
