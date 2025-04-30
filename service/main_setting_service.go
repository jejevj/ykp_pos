package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/entity"
	"github.com/jejevj/ykp_pos/repository"
	"github.com/jejevj/ykp_pos/utils"
)

type (
	MainSettingService interface {
		AddMainSetting(ctx context.Context, req dto.MainSettingCreateRequest) (dto.MainSettingResponse, error)
		GetAllMainSettingWithPagination(ctx context.Context) (dto.MainSettingPaginationResponse, error)
		GetMainSettingById(ctx context.Context, mainSettingId string) (dto.MainSettingResponse, error)
		UpdateMainSetting(ctx context.Context, req dto.MainSettingUpdateRequest, mainSettingId string) (dto.MainSettingUpdateResponse, error)
		DeleteMainSetting(ctx context.Context, mainSettingId string) error
	}
	mainSettingService struct {
		mainSettingRepo repository.MainSettingRepository
		jwtService      JWTService
	}
)

func NewMainSettingService(mainSettingRepo repository.MainSettingRepository, jwtService JWTService) MainSettingService {
	return &mainSettingService{
		mainSettingRepo: mainSettingRepo,
		jwtService:      jwtService,
	}
}
func (s *mainSettingService) AddMainSetting(ctx context.Context, req dto.MainSettingCreateRequest) (dto.MainSettingResponse, error) {
	var filename string

	fmt.Printf("AddMainSetting called with request: %+v\n", req)

	if req.Logo != nil {
		imageId := uuid.New()
		ext := utils.GetExtensions(req.Logo.Filename)

		fmt.Printf("Uploaded logo file: %s, extension: %s, generated image ID: %s\n", req.Logo.Filename, ext, imageId)

		dir := "main-setting"
		fmt.Printf("Ensuring directory exists: %s\n", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return dto.MainSettingResponse{}, fmt.Errorf("failed to create directory: %w", err)
		}

		filename = fmt.Sprintf("%s/%s.%s", dir, imageId, ext)
		fmt.Printf("Saving file with filename: %s\n", filename)

		if err := utils.UploadFile(req.Logo, filename); err != nil {
			return dto.MainSettingResponse{}, err
		}
	} else {
		fmt.Println("No logo uploaded, skipping file save.")
	}

	mainSetting := entity.MainSetting{
		NamaUsaha:  req.NamaUsaha,
		JenisUsaha: req.JenisUsaha,
		Alamat:     req.Alamat,
		LogoUrl:    filename, // Save the generated logo URL in the entity
		Hp:         req.Hp,
	}

	fmt.Printf("MainSetting entity to be saved: %+v\n", mainSetting)

	mainSettingAdd, err := s.mainSettingRepo.AddMainSetting(ctx, mainSetting)
	if err != nil {
		fmt.Printf("Error saving main setting to repository: %v\n", err)
		return dto.MainSettingResponse{}, dto.ErrCreateMainSetting
	}

	fmt.Printf("MainSetting successfully saved: %+v\n", mainSettingAdd)

	return dto.MainSettingResponse{
		ID:         mainSettingAdd.ID.String(),
		NamaUsaha:  mainSettingAdd.NamaUsaha,
		JenisUsaha: mainSettingAdd.JenisUsaha,
		Alamat:     mainSettingAdd.Alamat,
		LogoUrl:    mainSettingAdd.LogoUrl,
		Hp:         mainSettingAdd.Hp,
	}, nil
}

func (s *mainSettingService) GetAllMainSettingWithPagination(ctx context.Context) (dto.MainSettingPaginationResponse, error) {
	dataWithPaginate, err := s.mainSettingRepo.GetAllMainSettingWithPagination(ctx)
	if err != nil {
		return dto.MainSettingPaginationResponse{}, err
	}

	var datas []dto.MainSettingResponse
	for _, mainSetting := range dataWithPaginate.MainSettings {
		data := dto.MainSettingResponse{
			ID:         mainSetting.ID.String(),
			NamaUsaha:  mainSetting.NamaUsaha,
			JenisUsaha: mainSetting.JenisUsaha,
			Alamat:     mainSetting.Alamat,
			LogoUrl:    mainSetting.LogoUrl,
			Hp:         mainSetting.Hp,
		}

		datas = append(datas, data)
	}

	return dto.MainSettingPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *mainSettingService) GetMainSettingById(ctx context.Context, mainSettingId string) (dto.MainSettingResponse, error) {
	mainSetting, err := s.mainSettingRepo.GetMainSettingById(ctx, mainSettingId)
	if err != nil {
		return dto.MainSettingResponse{}, dto.ErrGetMainSettingById
	}

	return dto.MainSettingResponse{
		ID:         mainSetting.ID.String(),
		NamaUsaha:  mainSetting.NamaUsaha,
		JenisUsaha: mainSetting.JenisUsaha,
		Alamat:     mainSetting.Alamat,
		LogoUrl:    mainSetting.LogoUrl,
		Hp:         mainSetting.Hp,
	}, nil
}

func (s *mainSettingService) UpdateMainSetting(ctx context.Context, req dto.MainSettingUpdateRequest, mainSettingId string) (dto.MainSettingUpdateResponse, error) {
	// Convert string ID to uuid.UUID
	id, err := uuid.Parse(mainSettingId)
	if err != nil {
		return dto.MainSettingUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	var filename string
	var existingLogoUrl string

	// Fetch the existing main setting to retain the current logo URL (if no new logo is uploaded)
	existingMainSetting, err := s.mainSettingRepo.GetMainSettingById(ctx, mainSettingId)
	if err != nil {
		return dto.MainSettingUpdateResponse{}, fmt.Errorf("failed to retrieve existing MainSetting: %v", err)
	}

	// Retain the existing LogoUrl if no new logo is uploaded
	existingLogoUrl = existingMainSetting.LogoUrl

	// If logo is uploaded, handle the file upload
	if req.Logo != nil {
		// Generate a new image ID and get file extension
		imageId := uuid.New()
		ext := utils.GetExtensions(req.Logo.Filename)

		// Set the file path for the new logo
		filename = fmt.Sprintf("main-setting/%s.%s", imageId, ext)

		// Upload the file
		if err := utils.UploadFile(req.Logo, filename); err != nil {
			return dto.MainSettingUpdateResponse{}, fmt.Errorf("failed to upload file: %v", err)
		}
	} else {
		// If no new logo is uploaded, use the existing logo URL
		filename = existingLogoUrl
	}

	// Prepare the entity to be updated
	data := entity.MainSetting{
		ID:         id,
		NamaUsaha:  req.NamaUsaha,
		JenisUsaha: req.JenisUsaha,
		Alamat:     req.Alamat,
		LogoUrl:    filename, // Set the updated logo URL (or existing logo if not updated)
		Hp:         req.Hp,
	}

	// Call the repository to update the main setting
	mainSettingUpdate, err := s.mainSettingRepo.UpdateMainSetting(ctx, data)
	if err != nil {
		return dto.MainSettingUpdateResponse{}, fmt.Errorf("failed to update MainSetting: %v", err)
	}

	// Return the updated response
	return dto.MainSettingUpdateResponse{
		ID:         mainSettingUpdate.ID.String(),
		NamaUsaha:  mainSettingUpdate.NamaUsaha,
		JenisUsaha: mainSettingUpdate.JenisUsaha,
		Alamat:     mainSettingUpdate.Alamat,
		Logo:       mainSettingUpdate.LogoUrl,
		Hp:         mainSettingUpdate.Hp,
	}, nil
}

func (s *mainSettingService) DeleteMainSetting(ctx context.Context, mainSettingId string) error {
	mainSetting, err := s.mainSettingRepo.GetMainSettingById(ctx, mainSettingId)
	if err != nil {
		return dto.ErrMainSettingNotFound
	}

	err = s.mainSettingRepo.DeleteMainSetting(ctx, mainSetting.ID.String())
	if err != nil {
		return dto.ErrDeleteMainSetting
	}

	return nil
}
