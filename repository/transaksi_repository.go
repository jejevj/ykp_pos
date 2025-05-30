package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/entity"
	"gorm.io/gorm"
)

type (
	TransaksiRepository interface {
		AddTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
		GetAllTransaksiWithPagination(ctx context.Context) (dto.GetAllTransaksiRepositoryResponse, error)
		GetTransaksiById(ctx context.Context, transaksiId string) (entity.Transaksi, error)
		UpdateTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
		DeleteTransaksi(ctx context.Context, transaksiId string) error
	}
	transaksiRepository struct {
		db *gorm.DB
	}
)

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		db: db,
	}
}

func (r *transaksiRepository) AddTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error) {
	tx := r.db

	// Create the transaksi
	if err := tx.WithContext(ctx).Create(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	// Preload related data (Loading.User and Barang.Satuan)
	if err := tx.WithContext(ctx).
		Preload("Barang.Satuan").
		Preload("Loading.User").
		Where("id = ?", transaksi.ID).
		Take(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}

func (r *transaksiRepository) GetAllTransaksiWithPagination(ctx context.Context) (dto.GetAllTransaksiRepositoryResponse, error) {
	tx := r.db

	var transaksis []entity.Transaksi
	var err error
	var count int64

	// Count the total number of transaksis
	if err := tx.WithContext(ctx).Model(&entity.Transaksi{}).Count(&count).Error; err != nil {
		return dto.GetAllTransaksiRepositoryResponse{}, err
	}

	// Fetch the transaksis with preloaded relationships (Loading.User and Barang.Satuan)
	// Fetch transaksis with preloaded relationships (Barang and Satuan)
	if err := tx.WithContext(ctx).
		Preload("Loading.User").
		Preload("Barang.Satuan").
		Scopes(Paginate(1, 10)).
		Find(&transaksis).Error; err != nil {
		return dto.GetAllTransaksiRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllTransaksiRepositoryResponse{
		Transaksis: transaksis,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *transaksiRepository) GetTransaksiById(ctx context.Context, transaksiId string) (entity.Transaksi, error) {
	tx := r.db

	// Convert transaksiId string to uuid.UUID
	id, err := uuid.Parse(transaksiId)
	if err != nil {
		return entity.Transaksi{}, fmt.Errorf("invalid transaksi ID format")
	}

	// Fetch the transaksi with preloaded relationships
	var transaksi entity.Transaksi
	if err := tx.WithContext(ctx).Preload("Loading.User").Preload("Barang.Satuan").Where("id = ?", id).First(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}
func (r *transaksiRepository) UpdateTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error) {
	tx := r.db

	// First, check if the record exists
	var existingTransaksi entity.Transaksi
	if err := tx.WithContext(ctx).Where("id = ?", transaksi.ID).Take(&existingTransaksi).Error; err != nil {
		// If the record doesn't exist, return a specific error
		if err == gorm.ErrRecordNotFound {
			return entity.Transaksi{}, fmt.Errorf("Transaksi with ID %s not found", transaksi.ID)
		}
		return entity.Transaksi{}, err
	}

	// Proceed with updating the record
	if err := tx.WithContext(ctx).Model(&existingTransaksi).Updates(transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}

	// Return the updated entity
	return existingTransaksi, nil
}
func (r *transaksiRepository) DeleteTransaksi(ctx context.Context, transaksiId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Transaksi{}, "id = ?", transaksiId).Error; err != nil {
		return err
	}

	return nil
}
