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

	// Prepare the transaksi entity
	transaksi := entity.Transaksi{
		IdLoading: req.IdLoading,
		IdBarang:  req.IdBarang,
		Jumlah:    req.Jumlah,
	}

	// Add the transaksi via the repository
	transaksiAdd, err := s.transaksiRepo.AddTransaksi(ctx, transaksi)
	if err != nil {
		return dto.TransaksiResponse{}, dto.ErrCreateTransaksi
	}

	// Map LoadingResponse with UserResponse
	loadingResponse := dto.LoadingResponse{
		ID:     transaksiAdd.Loading.ID.String(),
		IdUser: transaksiAdd.Loading.IdUser,
		User: dto.UserResponse{
			ID:         transaksiAdd.Loading.User.ID.String(),
			Name:       transaksiAdd.Loading.User.Name,
			Email:      transaksiAdd.Loading.User.Email,
			TelpNumber: transaksiAdd.Loading.User.TelpNumber,
			Role:       transaksiAdd.Loading.User.Role,
			ImageUrl:   transaksiAdd.Loading.User.ImageUrl,
		},
	}

	// Map BarangResponse with SatuanResponse
	barangResponse := dto.BarangResponse{
		ID:         transaksiAdd.Barang.ID.String(),
		NamaBarang: transaksiAdd.Barang.NamaBarang,
		KodeBarang: transaksiAdd.Barang.KodeBarang,
		HargaBeli:  transaksiAdd.Barang.HargaBeli,
		HargaJual:  transaksiAdd.Barang.HargaJual,
		IdSatuan:   transaksiAdd.Barang.IdSatuan,
		Stok:       transaksiAdd.Barang.Stok,
		Satuan: dto.SatuanResponse{
			ID:         transaksiAdd.Barang.Satuan.ID.String(),
			NamaSatuan: transaksiAdd.Barang.Satuan.NamaSatuan,
			Value:      transaksiAdd.Barang.Satuan.Value,
		},
	}

	// Return the mapped TransaksiResponse
	return dto.TransaksiResponse{
		ID:        transaksiAdd.ID.String(),
		IdLoading: transaksiAdd.IdLoading,
		Loading:   loadingResponse,
		IdBarang:  transaksiAdd.IdBarang,
		Barang:    barangResponse,
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
		// Map LoadingResponse with UserResponse
		loadingResponse := dto.LoadingResponse{
			ID:     transaksi.Loading.ID.String(),
			IdUser: transaksi.Loading.IdUser,
			User: dto.UserResponse{
				ID:         transaksi.Loading.User.ID.String(),
				Name:       transaksi.Loading.User.Name,
				Email:      transaksi.Loading.User.Email,
				TelpNumber: transaksi.Loading.User.TelpNumber,
				Role:       transaksi.Loading.User.Role,
				ImageUrl:   transaksi.Loading.User.ImageUrl,
			},
		}

		// Map BarangResponse with SatuanResponse
		barangResponse := dto.BarangResponse{
			ID:         transaksi.Barang.ID.String(),
			NamaBarang: transaksi.Barang.NamaBarang,
			KodeBarang: transaksi.Barang.KodeBarang,
			HargaBeli:  transaksi.Barang.HargaBeli,
			HargaJual:  transaksi.Barang.HargaJual,
			IdSatuan:   transaksi.Barang.IdSatuan,
			Stok:       transaksi.Barang.Stok,
			Satuan: dto.SatuanResponse{
				ID:         transaksi.Barang.Satuan.ID.String(),
				NamaSatuan: transaksi.Barang.Satuan.NamaSatuan,
				Value:      transaksi.Barang.Satuan.Value,
			},
		}

		// Map TransaksiResponse
		data := dto.TransaksiResponse{
			ID:        transaksi.ID.String(),
			IdLoading: transaksi.IdLoading,
			Loading:   loadingResponse,
			IdBarang:  transaksi.IdBarang,
			Barang:    barangResponse,
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
	// Fetch the transaksi with preloaded relationships
	transaksi, err := s.transaksiRepo.GetTransaksiById(ctx, transaksiId)
	if err != nil {
		return dto.TransaksiResponse{}, dto.ErrGetTransaksiById
	}

	// Map LoadingResponse with UserResponse
	loadingResponse := dto.LoadingResponse{
		ID:     transaksi.Loading.ID.String(),
		IdUser: transaksi.Loading.IdUser,
		User: dto.UserResponse{
			ID:         transaksi.Loading.User.ID.String(),
			Name:       transaksi.Loading.User.Name,
			Email:      transaksi.Loading.User.Email,
			TelpNumber: transaksi.Loading.User.TelpNumber,
			Role:       transaksi.Loading.User.Role,
			ImageUrl:   transaksi.Loading.User.ImageUrl,
		},
	}

	// Map BarangResponse with SatuanResponse
	barangResponse := dto.BarangResponse{
		ID:         transaksi.Barang.ID.String(),
		NamaBarang: transaksi.Barang.NamaBarang,
		KodeBarang: transaksi.Barang.KodeBarang,
		HargaBeli:  transaksi.Barang.HargaBeli,
		HargaJual:  transaksi.Barang.HargaJual,
		IdSatuan:   transaksi.Barang.IdSatuan,
		Stok:       transaksi.Barang.Stok,
		// Satuan: Uncomment if needed
		Satuan: dto.SatuanResponse{
			ID:         transaksi.Barang.Satuan.ID.String(),
			NamaSatuan: transaksi.Barang.Satuan.NamaSatuan,
			Value:      transaksi.Barang.Satuan.Value,
		},
	}

	// Return the mapped TransaksiResponse
	return dto.TransaksiResponse{
		ID:        transaksi.ID.String(),
		IdLoading: transaksi.IdLoading,
		Loading:   loadingResponse,
		IdBarang:  transaksi.IdBarang,
		Barang:    barangResponse,
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
