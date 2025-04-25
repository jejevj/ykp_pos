package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	CustomerController interface {
		AddCustomer(ctx *fiber.Ctx) error
		GetCustomerById(ctx *fiber.Ctx) error
		GetAllCustomerWithPagination(ctx *fiber.Ctx) error
		UpdateCustomer(ctx *fiber.Ctx) error
		DeleteCustomer(ctx *fiber.Ctx) error
	}

	customerController struct {
		customerService service.CustomerService
	}
)

func NewCustomerController(us service.CustomerService) CustomerController {
	return &customerController{
		customerService: us,
	}
}

func (c *customerController) AddCustomer(ctx *fiber.Ctx) error {
	var customer dto.CustomerCreateRequest

	if err := ctx.BodyParser(&customer); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.customerService.AddCustomer(ctx.Context(), customer)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *customerController) GetCustomerById(ctx *fiber.Ctx) error {
	var req dto.GetCustomerByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.customerService.GetCustomerById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *customerController) GetAllCustomerWithPagination(ctx *fiber.Ctx) error {
	// var req dto.PaginationRequest

	result, err := c.customerService.GetAllCustomerWithPagination(ctx.Context())
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

func (c *customerController) UpdateCustomer(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request
	var req dto.CustomerUpdateRequest
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
	existingCustomer, err := c.customerService.GetCustomerById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "Customer not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Use the existing entity and update the fields
	existingCustomer.NamaToko = req.NamaToko
	existingCustomer.NamaPemilik = req.NamaPemilik
	existingCustomer.Alamat = req.Alamat
	existingCustomer.HP = req.HP

	// Call the service to update the Customer
	result, err := c.customerService.UpdateCustomer(ctx.Context(), req, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated Customer
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *customerController) DeleteCustomer(ctx *fiber.Ctx) error {
	var req dto.GetCustomerByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.customerService.DeleteCustomer(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
