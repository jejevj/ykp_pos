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
	BarangService interface {
		AddBarang(ctx context.Context, req dto.BarangCreateRequest) (dto.BarangResponse, error)
		GetAllBarangWithPagination(ctx context.Context) (dto.BarangPaginationResponse, error)
		GetBarangById(ctx context.Context, barangId string) (dto.BarangResponse, error)
		UpdateBarang(ctx context.Context, req dto.BarangUpdateRequest, barangId string) (dto.BarangUpdateResponse, error)
		UpdateStokBarang(ctx context.Context, req dto.BarangUpdateStokRequest, barangId string) (dto.BarangUpdateResponse, error)
		DeleteBarang(ctx context.Context, barangId string) error
	}
	barangService struct {
		barangRepo repository.BarangRepository
		jwtService JWTService
	}
)

func NewBarangService(barangRepo repository.BarangRepository, jwtService JWTService) BarangService {
	return &barangService{
		barangRepo: barangRepo,
		jwtService: jwtService,
	}
}
func (s *barangService) AddBarang(ctx context.Context, req dto.BarangCreateRequest) (dto.BarangResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	barang := entity.Barang{
		NamaBarang:   req.NamaBarang,
		KodeBarang:   req.KodeBarang,
		HargaBeli:    req.HargaBeli,
		HargaJual:    req.HargaJual,
		IdSatuan:     req.IdSatuan,
		JumlahKrat:   req.JumlahKrat,
		JumlahSatuan: req.JumlahSatuan,
		// Stok:         req.Stok,
	}

	barangAdd, err := s.barangRepo.AddBarang(ctx, barang)
	if err != nil {
		return dto.BarangResponse{}, dto.ErrCreateBarang
	}

	// Convert Satuan entity to SatuanResponse DTO
	satuanResponse := dto.SatuanResponse{
		ID:         barangAdd.Satuan.ID.String(),
		NamaSatuan: barangAdd.Satuan.NamaSatuan,
		Value:      barangAdd.Satuan.Value,
		// Add other fields from Satuan entity if necessary
	}

	// req.JumlahKrat

	return dto.BarangResponse{
		ID:           barangAdd.ID.String(),
		NamaBarang:   barangAdd.NamaBarang,
		KodeBarang:   barangAdd.KodeBarang,
		HargaBeli:    barangAdd.HargaBeli,
		HargaJual:    barangAdd.HargaJual,
		IdSatuan:     barangAdd.IdSatuan,
		Satuan:       satuanResponse,
		JumlahKrat:   barangAdd.JumlahKrat,
		JumlahSatuan: barangAdd.JumlahSatuan,
		Stok:         barangAdd.Stok,
	}, nil
}
func (s *barangService) GetAllBarangWithPagination(ctx context.Context) (dto.BarangPaginationResponse, error) {
	dataWithPaginate, err := s.barangRepo.GetAllBarangWithPagination(ctx)
	if err != nil {
		return dto.BarangPaginationResponse{}, err
	}

	var datas []dto.BarangResponse
	for _, barang := range dataWithPaginate.Barangs {
		data := dto.BarangResponse{
			ID:           barang.ID.String(),
			NamaBarang:   barang.NamaBarang,
			KodeBarang:   barang.KodeBarang,
			HargaBeli:    barang.HargaBeli,
			HargaJual:    barang.HargaJual,
			IdSatuan:     barang.IdSatuan,
			JumlahKrat:   barang.JumlahKrat,
			JumlahSatuan: barang.JumlahSatuan,
			Satuan: dto.SatuanResponse{
				ID:         barang.Satuan.ID.String(),
				NamaSatuan: barang.Satuan.NamaSatuan,
				Value:      barang.Satuan.Value,
			},
			Stok: barang.Stok,
		}
		datas = append(datas, data)
	}

	// Return the response in a format compatible with DataTable
	return dto.BarangPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *barangService) GetBarangById(ctx context.Context, barangId string) (dto.BarangResponse, error) {
	barang, err := s.barangRepo.GetBarangById(ctx, barangId)
	if err != nil {
		return dto.BarangResponse{}, dto.ErrGetBarangById
	}

	satuanResponse := dto.SatuanResponse{
		ID:         barang.Satuan.ID.String(),
		NamaSatuan: barang.Satuan.NamaSatuan,
		Value:      barang.Satuan.Value,
		// Add other fields from Satuan entity if necessary
	}

	return dto.BarangResponse{
		ID:           barang.ID.String(),
		NamaBarang:   barang.NamaBarang,
		KodeBarang:   barang.KodeBarang,
		HargaBeli:    barang.HargaBeli,
		HargaJual:    barang.HargaJual,
		IdSatuan:     barang.IdSatuan,
		JumlahKrat:   barang.JumlahKrat,
		JumlahSatuan: barang.JumlahSatuan,
		Satuan:       satuanResponse,
		Stok:         barang.Stok,
	}, nil
}
func (s *barangService) UpdateBarang(ctx context.Context, req dto.BarangUpdateRequest, barangId string) (dto.BarangUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(barangId)
	if err != nil {
		return dto.BarangUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Barang{
		ID:         id,
		NamaBarang: req.NamaBarang,
		KodeBarang: req.KodeBarang,
		HargaBeli:  req.HargaBeli,
		HargaJual:  req.HargaJual,
		IdSatuan:   req.IdSatuan,
		Stok:       req.Stok,
	}

	// Call the repository to update
	barangUpdate, err := s.barangRepo.UpdateBarang(ctx, data)
	if err != nil {
		return dto.BarangUpdateResponse{}, fmt.Errorf("failed to update Barang: %v", err)
	} // Convert Satuan entity to SatuanResponse DTO
	satuanResponse := dto.SatuanResponse{
		ID:         barangUpdate.Satuan.ID.String(),
		NamaSatuan: barangUpdate.Satuan.NamaSatuan,
		// Add other fields from Satuan entity if necessary
	}

	return dto.BarangUpdateResponse{
		ID:         barangUpdate.ID.String(),
		NamaBarang: barangUpdate.NamaBarang,
		KodeBarang: barangUpdate.KodeBarang,
		HargaBeli:  barangUpdate.HargaBeli,
		HargaJual:  barangUpdate.HargaJual,
		IdSatuan:   barangUpdate.IdSatuan,
		Satuan:     satuanResponse,
		Stok:       barangUpdate.Stok,
	}, nil
}

func (s *barangService) UpdateStokBarang(ctx context.Context, req dto.BarangUpdateStokRequest, barangId string) (dto.BarangUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(barangId)
	if err != nil {
		return dto.BarangUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Barang{
		ID: id,

		JumlahKrat:   req.JumlahKrat,
		JumlahSatuan: req.JumlahSatuan,
	}

	// Call the repository to update
	barangUpdate, err := s.barangRepo.UpdateStokBarang(ctx, data)
	if err != nil {
		return dto.BarangUpdateResponse{}, fmt.Errorf("failed to update Barang: %v", err)
	} // Convert Satuan entity to SatuanResponse DTO
	satuanResponse := dto.SatuanResponse{
		ID:         barangUpdate.Satuan.ID.String(),
		NamaSatuan: barangUpdate.Satuan.NamaSatuan,
		Value:      barangUpdate.Satuan.Value,
		// Add other fields from Satuan entity if necessary
	}

	return dto.BarangUpdateResponse{
		ID:           barangUpdate.ID.String(),
		NamaBarang:   barangUpdate.NamaBarang,
		KodeBarang:   barangUpdate.KodeBarang,
		HargaBeli:    barangUpdate.HargaBeli,
		HargaJual:    barangUpdate.HargaJual,
		IdSatuan:     barangUpdate.IdSatuan,
		JumlahKrat:   barangUpdate.JumlahKrat,
		JumlahSatuan: barangUpdate.JumlahSatuan,
		Satuan:       satuanResponse,
		Stok:         barangUpdate.Stok,
	}, nil
}

func (s *barangService) DeleteBarang(ctx context.Context, barangId string) error {
	barang, err := s.barangRepo.GetBarangById(ctx, barangId)
	if err != nil {
		return dto.ErrBarangNotFound
	}

	err = s.barangRepo.DeleteBarang(ctx, barang.ID.String())
	if err != nil {
		return dto.ErrDeleteBarang
	}

	return nil
}
