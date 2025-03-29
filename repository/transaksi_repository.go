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
	TransaksiRepository interface {
		AddTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
		GetAllTransaksiWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransaksiRepositoryResponse, error)
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

	if err := tx.WithContext(ctx).Create(&transaksi).Error; err != nil {
		return entity.Transaksi{}, err
	}
	return transaksi, nil
}

func (r *transaksiRepository) GetAllTransaksiWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransaksiRepositoryResponse, error) {
	tx := r.db

	var transaksis []entity.Transaksi
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Transaksi{}).Count(&count).Error; err != nil {
		return dto.GetAllTransaksiRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&transaksis).Error; err != nil {
		return dto.GetAllTransaksiRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllTransaksiRepositoryResponse{
		Transaksis: transaksis,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *transaksiRepository) GetTransaksiById(ctx context.Context, transaksiId string) (entity.Transaksi, error) {
	tx := r.db

	var transaksi entity.Transaksi
	if err := tx.WithContext(ctx).Where("id = ?", transaksiId).Take(&transaksi).Error; err != nil {
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
