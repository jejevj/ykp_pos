package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	SatuanController interface {
		AddSatuan(ctx *fiber.Ctx) error
		GetSatuanById(ctx *fiber.Ctx) error
		GetAllSatuanWithPagination(ctx *fiber.Ctx) error
		UpdateSatuan(ctx *fiber.Ctx) error
		DeleteSatuan(ctx *fiber.Ctx) error
	}

	satuanController struct {
		satuanService service.SatuanService
	}
)

func NewSatuanController(us service.SatuanService) SatuanController {
	return &satuanController{
		satuanService: us,
	}
}

func (c *satuanController) AddSatuan(ctx *fiber.Ctx) error {
	var satuan dto.SatuanCreateRequest

	if err := ctx.BodyParser(&satuan); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.satuanService.AddSatuan(ctx.Context(), satuan)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *satuanController) GetSatuanById(ctx *fiber.Ctx) error {
	var req dto.GetSatuanByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.satuanService.GetSatuanById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *satuanController) GetAllSatuanWithPagination(ctx *fiber.Ctx) error {

	result, err := c.satuanService.GetAllSatuanWithPagination(ctx.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *satuanController) UpdateSatuan(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request
	var req dto.SatuanUpdateRequest
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
	existingSatuan, err := c.satuanService.GetSatuanById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "Satuan not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Use the existing entity and update the fields
	existingSatuan.NamaSatuan = req.NamaSatuan
	existingSatuan.Value = req.Value

	// Call the service to update the Satuan
	result, err := c.satuanService.UpdateSatuan(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated Satuan
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *satuanController) DeleteSatuan(ctx *fiber.Ctx) error {
	var req dto.GetSatuanByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.satuanService.DeleteSatuan(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
