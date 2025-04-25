package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/entity"
	"github.com/jejevj/ykp_pos/repository"
)

type (
	LoadingService interface {
		AddLoading(ctx context.Context, req dto.LoadingCreateRequest) (dto.LoadingResponse, error)
		GetAllLoadingWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.LoadingPaginationResponse, error)
		GetLoadingById(ctx context.Context, loadingId string) (dto.LoadingResponse, error)
		UpdateLoading(ctx context.Context, req dto.LoadingUpdateRequest, loadingId string) (dto.LoadingUpdateResponse, error)
		DeleteLoading(ctx context.Context, loadingId string) error
	}
	loadingService struct {
		loadingRepo repository.LoadingRepository
		jwtService  JWTService
	}
)

func NewLoadingService(loadingRepo repository.LoadingRepository, jwtService JWTService) LoadingService {
	return &loadingService{
		loadingRepo: loadingRepo,
		jwtService:  jwtService,
	}
}
func (s *loadingService) AddLoading(ctx context.Context, req dto.LoadingCreateRequest) (dto.LoadingResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	loading := entity.Loading{
		IdUser: req.IdUser,
	}

	loadingAdd, err := s.loadingRepo.AddLoading(ctx, loading)
	if err != nil {
		return dto.LoadingResponse{}, dto.ErrCreateLoading
	}

	return dto.LoadingResponse{
		ID:         loadingAdd.ID.String(),
		IdUser:     loadingAdd.IdUser,
		IsApproved: loadingAdd.IsApproved,
	}, nil
}
func (s *loadingService) GetAllLoadingWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.LoadingPaginationResponse, error) {
	dataWithPaginate, err := s.loadingRepo.GetAllLoadingWithPagination(ctx, req)
	if err != nil {
		return dto.LoadingPaginationResponse{}, err
	}

	var datas []dto.LoadingResponse
	for _, loading := range dataWithPaginate.Loadings {
		data := dto.LoadingResponse{
			ID:         loading.ID.String(),
			IdUser:     loading.IdUser,
			IsApproved: loading.IsApproved,
		}

		datas = append(datas, data)
	}

	return dto.LoadingPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *loadingService) GetLoadingById(ctx context.Context, loadingId string) (dto.LoadingResponse, error) {
	loading, err := s.loadingRepo.GetLoadingById(ctx, loadingId)
	if err != nil {
		return dto.LoadingResponse{}, dto.ErrGetLoadingById
	}
	userResponse := dto.UserResponse{
		ID:         loading.User.ID.String(),
		Name:       loading.User.Name,
		Email:      loading.User.Email,
		TelpNumber: loading.User.TelpNumber,
		ImageUrl:   loading.User.ImageUrl,
		// Add other fields from Satuan entity if necessary
	}

	return dto.LoadingResponse{
		ID:         loading.ID.String(),
		IdUser:     loading.IdUser,
		User:       userResponse,
		IsApproved: loading.IsApproved,
	}, nil
}
func (s *loadingService) UpdateLoading(ctx context.Context, req dto.LoadingUpdateRequest, loadingId string) (dto.LoadingUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(loadingId)
	if err != nil {
		return dto.LoadingUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Loading{
		ID:         id,
		IsApproved: req.IsApproved,
	}

	// Call the repository to update
	loadingUpdate, err := s.loadingRepo.UpdateLoading(ctx, data)
	if err != nil {
		return dto.LoadingUpdateResponse{}, fmt.Errorf("failed to update Loading: %v", err)
	}

	return dto.LoadingUpdateResponse{
		ID:         loadingUpdate.ID.String(),
		IdUser:     loadingUpdate.IdUser,
		IsApproved: loadingUpdate.IsApproved,
	}, nil
}

func (s *loadingService) DeleteLoading(ctx context.Context, loadingId string) error {
	loading, err := s.loadingRepo.GetLoadingById(ctx, loadingId)
	if err != nil {
		return dto.ErrLoadingNotFound
	}

	err = s.loadingRepo.DeleteLoading(ctx, loading.ID.String())
	if err != nil {
		return dto.ErrDeleteLoading
	}

	return nil
}
