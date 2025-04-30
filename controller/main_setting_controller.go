package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/service"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	MainSettingController interface {
		AddMainSetting(ctx *fiber.Ctx) error
		GetMainSettingById(ctx *fiber.Ctx) error
		GetAllMainSettingWithPagination(ctx *fiber.Ctx) error
		UpdateMainSetting(ctx *fiber.Ctx) error
		DeleteMainSetting(ctx *fiber.Ctx) error
	}

	mainSettingController struct {
		mainSettingService service.MainSettingService
	}
)

func NewMainSettingController(us service.MainSettingService) MainSettingController {
	return &mainSettingController{
		mainSettingService: us,
	}
}

func (c *mainSettingController) AddMainSetting(ctx *fiber.Ctx) error {
	var mainSetting dto.MainSettingCreateRequest

	// Manually handle multipart file
	logo, err := ctx.FormFile("logo") // Use FormFile instead of BodyParser for files
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Parse other form fields into the DTO
	if err := ctx.BodyParser(&mainSetting); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Set the logo file to the DTO for the service layer
	mainSetting.Logo = logo

	// Call the service layer to handle business logic
	result, err := c.mainSettingService.AddMainSetting(ctx.Context(), mainSetting)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return success response
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *mainSettingController) GetMainSettingById(ctx *fiber.Ctx) error {
	var req dto.GetMainSettingByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.mainSettingService.GetMainSettingById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
func (c *mainSettingController) GetAllMainSettingWithPagination(ctx *fiber.Ctx) error {
	// var req dto.PaginationRequest

	result, err := c.mainSettingService.GetAllMainSettingWithPagination(ctx.Context())
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

func (c *mainSettingController) UpdateMainSetting(ctx *fiber.Ctx) error {
	// Parse the request body to get the update request (excluding the file)
	var req dto.MainSettingUpdateRequest
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
	existingMainSetting, err := c.mainSettingService.GetMainSettingById(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed("failed update data", "MainSetting not found: "+err.Error(), nil)
		return ctx.Status(http.StatusNotFound).JSON(res)
	}

	// Log the parsed request to debug the fields
	fmt.Printf("Parsed request: %+v\n", existingMainSetting)

	// Handle logo upload if it's present
	var logoFile *multipart.FileHeader
	var logoPath string
	if logoFile, err = ctx.FormFile("logo"); err == nil && logoFile != nil {
		// If logo is uploaded, handle the file upload
		imageId := uuid.New()
		ext := utils.GetExtensions(logoFile.Filename)
		logoPath = fmt.Sprintf("main-setting/%s.%s", imageId, ext)

		// Upload the file
		if err := utils.UploadFile(logoFile, logoPath); err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
			return ctx.Status(http.StatusBadRequest).JSON(res)
		}

		// Update LogoUrl to the new path
		existingMainSetting.LogoUrl = logoPath
	} else if err != nil {

		logoPath = existingMainSetting.LogoUrl
		// If error occurs while handling the file, return error
		// res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		// return ctx.Status(http.StatusBadRequest).JSON(res)
	} else {
		// If no logo is uploaded, retain the existing LogoUrl
		logoPath = existingMainSetting.LogoUrl
	}
	// Update the other fields
	existingMainSetting.NamaUsaha = req.NamaUsaha
	existingMainSetting.JenisUsaha = req.JenisUsaha
	existingMainSetting.Alamat = req.Alamat
	existingMainSetting.Hp = req.Hp

	// Prepare dto.MainSettingUpdateRequest for the service call
	updateReq := dto.MainSettingUpdateRequest{
		ID:         req.ID,
		NamaUsaha:  req.NamaUsaha,
		JenisUsaha: req.JenisUsaha,
		Alamat:     req.Alamat,
		Hp:         req.Hp,
		Logo:       logoFile, // This is the logo file, if uploaded
	}

	// Log the updated existingMainSetting for debugging purposes
	fmt.Printf("Updated MainSetting entity: %+v\n", existingMainSetting)

	// Call the service to update the MainSetting with the updated entity
	result, err := c.mainSettingService.UpdateMainSetting(ctx.Context(), updateReq, req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Return the success response with the updated MainSetting
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *mainSettingController) DeleteMainSetting(ctx *fiber.Ctx) error {
	var req dto.GetMainSettingByIdRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		res := utils.BuildResponseFailed("failed delete data", "ID is missing or empty", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.mainSettingService.DeleteMainSetting(ctx.Context(), req.ID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
