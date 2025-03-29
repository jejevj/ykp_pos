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
	TransaksiService interface {
		AddTransaksi(ctx context.Context, req dto.TransaksiCreateRequest) (dto.TransaksiResponse, error)
		GetAllTransaksiWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TransaksiPaginationResponse, error)
		GetTransaksiById(ctx context.Context, transaksiId string) (dto.TransaksiResponse, error)
		UpdateTransaksi(ctx context.Context, req dto.TransaksiUpdateRequest, transaksiId string) (dto.TransaksiUpdateResponse, error)
		DeleteTransaksi(ctx context.Context, transaksiId string) error
	}
	transaksiService struct {
		transaksiRepo repository.TransaksiRepository
		jwtService    JWTService
	}
)

func NewTransaksiService(transaksiRepo repository.TransaksiRepository, jwtService JWTService) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
		jwtService:    jwtService,
	}
}
func (s *transaksiService) AddTransaksi(ctx context.Context, req dto.TransaksiCreateRequest) (dto.TransaksiResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	transaksi := entity.Transaksi{
		IdLoading: req.IdLoading,
		IdBarang:  req.IdBarang,
		Jumlah:    req.Jumlah,
	}

	transaksiAdd, err := s.transaksiRepo.AddTransaksi(ctx, transaksi)
	if err != nil {
		return dto.TransaksiResponse{}, dto.ErrCreateTransaksi
	}

	return dto.TransaksiResponse{
		ID:        transaksiAdd.ID.String(),
		IdLoading: transaksiAdd.IdLoading,
		IdBarang:  transaksiAdd.IdBarang,
		Jumlah:    transaksiAdd.Jumlah,
	}, nil
}
func (s *transaksiService) GetAllTransaksiWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.TransaksiPaginationResponse, error) {
	dataWithPaginate, err := s.transaksiRepo.GetAllTransaksiWithPagination(ctx, req)
	if err != nil {
		return dto.TransaksiPaginationResponse{}, err
	}

	var datas []dto.TransaksiResponse
	for _, transaksi := range dataWithPaginate.Transaksis {
		data := dto.TransaksiResponse{
			ID:        transaksi.ID.String(),
			IdLoading: transaksi.IdLoading,
			IdBarang:  transaksi.IdBarang,
			Jumlah:    transaksi.Jumlah,
		}

		datas = append(datas, data)
	}

	return dto.TransaksiPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *transaksiService) GetTransaksiById(ctx context.Context, transaksiId string) (dto.TransaksiResponse, error) {
	transaksi, err := s.transaksiRepo.GetTransaksiById(ctx, transaksiId)
	if err != nil {
		return dto.TransaksiResponse{}, dto.ErrGetTransaksiById
	}

	return dto.TransaksiResponse{
		ID:        transaksi.ID.String(),
		IdLoading: transaksi.IdLoading,
		IdBarang:  transaksi.IdBarang,
		Jumlah:    transaksi.Jumlah,
	}, nil
}
func (s *transaksiService) UpdateTransaksi(ctx context.Context, req dto.TransaksiUpdateRequest, transaksiId string) (dto.TransaksiUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(transaksiId)
	if err != nil {
		return dto.TransaksiUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Transaksi{
		ID:        id,
		IdLoading: req.IdLoading,
		IdBarang:  req.IdBarang,
		Jumlah:    req.Jumlah,
	}

	// Call the repository to update
	transaksiUpdate, err := s.transaksiRepo.UpdateTransaksi(ctx, data)
	if err != nil {
		return dto.TransaksiUpdateResponse{}, fmt.Errorf("failed to update Transaksi: %v", err)
	}

	return dto.TransaksiUpdateResponse{
		ID:        transaksiUpdate.ID.String(),
		IdLoading: transaksiUpdate.IdLoading,
		IdBarang:  transaksiUpdate.IdBarang,
		Jumlah:    transaksiUpdate.Jumlah,
	}, nil
}

func (s *transaksiService) DeleteTransaksi(ctx context.Context, transaksiId string) error {
	transaksi, err := s.transaksiRepo.GetTransaksiById(ctx, transaksiId)
	if err != nil {
		return dto.ErrTransaksiNotFound
	}

	err = s.transaksiRepo.DeleteTransaksi(ctx, transaksi.ID.String())
	if err != nil {
		return dto.ErrDeleteTransaksi
	}

	return nil
}
