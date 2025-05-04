package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	LoadingController interface {
		AddLoading(ctx *fiber.Ctx) error
		GetLoadingById(ctx *fiber.Ctx) error
		GetAllLoadingWithPagination(ctx *fiber.Ctx) error
		UpdateLoading(ctx *fiber.Ctx) error
		DeleteLoading(ctx *fiber.Ctx) error
	}

	loadingController struct {
		loadingService service.LoadingService
	}
)

func NewLoadingController(us service.LoadingService) LoadingController {
	return &loadingController{
		loadingService: us,
	}
}

func (c *loadingController) AddLoading(ctx *fiber.Ctx) error {
	var loading dto.LoadingCreateRequest

	if err := ctx.BodyParser(&loading); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.loadingService.AddLoading(ctx.Context(), loading)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *loadingController) GetLoadingById(ctx *fiber.Ctx) error {
	var req dto.GetLoadingByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.loadingService.GetLoadingById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *loadingController) GetAllLoadingWithPagination(ctx *fiber.Ctx) error {

	result, err := c.loadingService.GetAllLoadingWithPagination(ctx.Context())
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

func (c *loadingController) UpdateLoading(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request
	var req dto.LoadingUpdateRequest
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
	existingLoading, err := c.loadingService.GetLoadingById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "Loading not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Use the existing entity and update the fields
	existingLoading.IsApproved = req.IsApproved

	// Call the service to update the Loading
	result, err := c.loadingService.UpdateLoading(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated Loading
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *loadingController) DeleteLoading(ctx *fiber.Ctx) error {
	var req dto.GetLoadingByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.loadingService.DeleteLoading(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
