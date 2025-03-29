package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	TransaksiController interface {
		AddTransaksi(ctx *fiber.Ctx) error
		GetTransaksiById(ctx *fiber.Ctx) error
		GetAllTransaksiWithPagination(ctx *fiber.Ctx) error
		UpdateTransaksi(ctx *fiber.Ctx) error
		DeleteTransaksi(ctx *fiber.Ctx) error
	}

	transaksiController struct {
		transaksiService service.TransaksiService
	}
)

func NewTransaksiController(us service.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: us,
	}
}

func (c *transaksiController) AddTransaksi(ctx *fiber.Ctx) error {
	var transaksi dto.TransaksiCreateRequest

	if err := ctx.BodyParser(&transaksi); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.transaksiService.AddTransaksi(ctx.Context(), transaksi)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *transaksiController) GetTransaksiById(ctx *fiber.Ctx) error {
	var req dto.GetTransaksiByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.transaksiService.GetTransaksiById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *transaksiController) GetAllTransaksiWithPagination(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.transaksiService.GetAllTransaksiWithPagination(ctx.Context(), req)
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

func (c *transaksiController) UpdateTransaksi(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request
	var req dto.TransaksiUpdateRequest
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
	existingTransaksi, err := c.transaksiService.GetTransaksiById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "Transaksi not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Use the existing entity and update the fields
	existingTransaksi.IdLoading = req.IdLoading
	existingTransaksi.IdBarang = req.IdBarang
	existingTransaksi.Jumlah = req.Jumlah

	// Call the service to update the Transaksi
	result, err := c.transaksiService.UpdateTransaksi(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated Transaksi
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *transaksiController) DeleteTransaksi(ctx *fiber.Ctx) error {
	var req dto.GetTransaksiByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.transaksiService.DeleteTransaksi(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
