package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/entity"
	"gorm.io/gorm"
)

type (
	BarangRepository interface {
		AddBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error)
		GetAllBarangWithPagination(ctx context.Context) (dto.GetAllBarangRepositoryResponse, error)
		GetBarangById(ctx context.Context, barangId string) (entity.Barang, error)
		UpdateBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error)
		UpdateStokBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error)
		DeleteBarang(ctx context.Context, barangId string) error
	}
	barangRepository struct {
		db *gorm.DB
	}
)

func NewBarangRepository(db *gorm.DB) BarangRepository {
	return &barangRepository{
		db: db,
	}
}

func (r *barangRepository) AddBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error) {
	tx := r.db

	// Fetch Satuan entity based on IdSatuan (assuming you have a related Satuan table/entity)
	var satuan entity.Satuan
	if err := tx.WithContext(ctx).Where("id = ?", barang.IdSatuan).First(&satuan).Error; err != nil {
		// Return an error if the Satuan entity can't be found
		return entity.Barang{}, err
	}

	// Assign the retrieved Satuan to the Barang entity
	barang.Satuan = satuan

	// Calculate the Stok based on JumlahKrat, JumlahSatuan, and Satuan.Value
	barang.Stok = barang.JumlahKrat*satuan.Value + barang.JumlahSatuan

	// I want to make barang.Stok as
	// barang.Stok = JumlahKrat * barang.Satuan.Value + JumlahSatuan
	if err := tx.WithContext(ctx).Create(&barang).Error; err != nil {
		return entity.Barang{}, err
	}
	return barang, nil
}

func (r *barangRepository) GetAllBarangWithPagination(ctx context.Context) (dto.GetAllBarangRepositoryResponse, error) {
	tx := r.db

	var barangs []entity.Barang
	var err error
	var count int64

	if err := tx.WithContext(ctx).Model(&entity.Barang{}).Count(&count).Error; err != nil {
		return dto.GetAllBarangRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).
		Preload("Satuan").
		Scopes(Paginate(1, 10)).
		Find(&barangs).Error; err != nil {
		return dto.GetAllBarangRepositoryResponse{}, err
	}
	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllBarangRepositoryResponse{
		Barangs: barangs,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *barangRepository) GetBarangById(ctx context.Context, barangId string) (entity.Barang, error) {
	tx := r.db

	var barang entity.Barang
	// Preload the Satuan data based on the foreign key IdSatuan
	if err := tx.WithContext(ctx).Preload("Satuan").Where("id = ?", barangId).Take(&barang).Error; err != nil {
		return entity.Barang{}, err
	}

	return barang, nil
}
func (r *barangRepository) UpdateBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error) {
	tx := r.db

	var existingBarang entity.Barang
	if err := tx.WithContext(ctx).Where("id = ?", barang.ID).Take(&existingBarang).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Barang{}, fmt.Errorf("Barang with ID %s not found", barang.ID)
		}
		return entity.Barang{}, err
	}

	if err := tx.WithContext(ctx).Model(&existingBarang).Updates(barang).Error; err != nil {
		return entity.Barang{}, err
	}

	return existingBarang, nil
}

func (r *barangRepository) UpdateStokBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error) {
	tx := r.db

	// Fetch the existing Barang from the database based on ID
	var existingBarang entity.Barang
	if err := tx.WithContext(ctx).Where("id = ?", barang.ID).Take(&existingBarang).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Barang{}, fmt.Errorf("Barang with ID %s not found", barang.ID)
		}
		return entity.Barang{}, err
	}

	// Fetch the associated Satuan based on the IdSatuan of the existing Barang
	var satuan entity.Satuan
	if err := tx.WithContext(ctx).Where("id = ?", existingBarang.IdSatuan).First(&satuan).Error; err != nil {
		return entity.Barang{}, err
	}

	// Update Satuan in existingBarang to make sure it has the most up-to-date value
	existingBarang.Satuan = satuan

	// Calculate the updated Stok

	barang.Stok = barang.JumlahKrat*satuan.Value + barang.JumlahSatuan
	barang.Stok += existingBarang.Stok
	barang.JumlahKrat += existingBarang.JumlahKrat
	barang.JumlahSatuan += existingBarang.JumlahSatuan

	// If you want to update other fields of existingBarang, make sure to set those values in the `barang` parameter
	// You can now update the existing Barang, passing only the necessary fields
	if err := tx.WithContext(ctx).Model(&existingBarang).Updates(barang).Error; err != nil {
		return entity.Barang{}, err
	}

	// After updating other fields, explicitly set the Stok again if it's not being passed in `barang`
	if err := tx.WithContext(ctx).Model(&existingBarang).Update("stok", existingBarang.Stok).Update("jumlah_krat", existingBarang.JumlahKrat).Update("jumlah_satuan", existingBarang.JumlahSatuan).Error; err != nil {
		return entity.Barang{}, err
	}

	// Return the updated Barang
	return existingBarang, nil
}

func (r *barangRepository) DeleteBarang(ctx context.Context, barangId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Barang{}, "id = ?", barangId).Error; err != nil {
		return err
	}

	return nil
}
