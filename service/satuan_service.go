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
	SatuanService interface {
		AddSatuan(ctx context.Context, req dto.SatuanCreateRequest) (dto.SatuanResponse, error)
		GetAllSatuanWithPagination(ctx context.Context) (dto.SatuanPaginationResponse, error)
		GetSatuanById(ctx context.Context, satuanId string) (dto.SatuanResponse, error)
		UpdateSatuan(ctx context.Context, req dto.SatuanUpdateRequest, satuanId string) (dto.SatuanUpdateResponse, error)
		DeleteSatuan(ctx context.Context, satuanId string) error
	}
	satuanService struct {
		satuanRepo repository.SatuanRepository
		jwtService JWTService
	}
)

func NewSatuanService(satuanRepo repository.SatuanRepository, jwtService JWTService) SatuanService {
	return &satuanService{
		satuanRepo: satuanRepo,
		jwtService: jwtService,
	}
}
func (s *satuanService) AddSatuan(ctx context.Context, req dto.SatuanCreateRequest) (dto.SatuanResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	satuan := entity.Satuan{
		NamaSatuan: req.NamaSatuan,
		Value:      req.Value,
	}

	satuanAdd, err := s.satuanRepo.AddSatuan(ctx, satuan)
	if err != nil {
		return dto.SatuanResponse{}, dto.ErrCreateSatuan
	}

	return dto.SatuanResponse{
		ID:         satuanAdd.ID.String(),
		NamaSatuan: satuanAdd.NamaSatuan,
		Value:      satuanAdd.Value,
	}, nil
}
func (s *satuanService) GetAllSatuanWithPagination(ctx context.Context) (dto.SatuanPaginationResponse, error) {
	dataWithPaginate, err := s.satuanRepo.GetAllSatuanWithPagination(ctx)
	if err != nil {
		return dto.SatuanPaginationResponse{}, err
	}

	var datas []dto.SatuanResponse
	for _, satuan := range dataWithPaginate.Satuans {
		data := dto.SatuanResponse{
			ID:         satuan.ID.String(),
			NamaSatuan: satuan.NamaSatuan,
			Value:      satuan.Value,
		}

		datas = append(datas, data)
	}

	return dto.SatuanPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *satuanService) GetSatuanById(ctx context.Context, satuanId string) (dto.SatuanResponse, error) {
	satuan, err := s.satuanRepo.GetSatuanById(ctx, satuanId)
	if err != nil {
		return dto.SatuanResponse{}, dto.ErrGetSatuanById
	}

	return dto.SatuanResponse{
		ID:         satuan.ID.String(),
		NamaSatuan: satuan.NamaSatuan,
		Value:      satuan.Value,
	}, nil
}
func (s *satuanService) UpdateSatuan(ctx context.Context, req dto.SatuanUpdateRequest, satuanId string) (dto.SatuanUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(satuanId)
	if err != nil {
		return dto.SatuanUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Satuan{
		ID:         id,
		NamaSatuan: req.NamaSatuan,
		Value:      req.Value,
	}

	// Call the repository to update
	satuanUpdate, err := s.satuanRepo.UpdateSatuan(ctx, data)
	if err != nil {
		return dto.SatuanUpdateResponse{}, fmt.Errorf("failed to update Satuan: %v", err)
	}

	return dto.SatuanUpdateResponse{
		ID:         satuanUpdate.ID.String(),
		NamaSatuan: satuanUpdate.NamaSatuan,
		Value:      satuanUpdate.Value,
	}, nil
}

func (s *satuanService) DeleteSatuan(ctx context.Context, satuanId string) error {
	satuan, err := s.satuanRepo.GetSatuanById(ctx, satuanId)
	if err != nil {
		return dto.ErrSatuanNotFound
	}

	err = s.satuanRepo.DeleteSatuan(ctx, satuan.ID.String())
	if err != nil {
		return dto.ErrDeleteSatuan
	}

	return nil
}
