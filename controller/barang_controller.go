package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	BarangController interface {
		AddBarang(ctx *fiber.Ctx) error
		GetBarangById(ctx *fiber.Ctx) error
		GetAllBarangWithPagination(ctx *fiber.Ctx) error
		UpdateBarang(ctx *fiber.Ctx) error
		DeleteBarang(ctx *fiber.Ctx) error
	}

	barangController struct {
		barangService service.BarangService
	}
)

func NewBarangController(us service.BarangService) BarangController {
	return &barangController{
		barangService: us,
	}
}

func (c *barangController) AddBarang(ctx *fiber.Ctx) error {
	var barang dto.BarangCreateRequest

	if err := ctx.BodyParser(&barang); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.barangService.AddBarang(ctx.Context(), barang)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *barangController) GetBarangById(ctx *fiber.Ctx) error {
	var req dto.GetBarangByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.barangService.GetBarangById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *barangController) GetAllBarangWithPagination(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.barangService.GetAllBarangWithPagination(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *barangController) UpdateBarang(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request
	var req dto.BarangUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Check if ID is provided in the request body
	if req.ID == "" {
		res := utils.BuildResponseFailed("failed update data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Get the existing data by ID
	existingBarang, err := c.barangService.GetBarangById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "Barang not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Use the existing entity and update the fields
	existingBarang.NamaBarang = req.NamaBarang
	existingBarang.KodeBarang = req.KodeBarang
	existingBarang.HargaBeli = req.HargaBeli
	existingBarang.HargaJual = req.HargaJual
	existingBarang.IdSatuan = req.IdSatuan
	existingBarang.Stok = req.Stok

	// Call the service to update the Barang
	result, err := c.barangService.UpdateBarang(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated Barang
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *barangController) DeleteBarang(ctx *fiber.Ctx) error {
	var req dto.GetBarangByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.barangService.DeleteBarang(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
