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
	SatuanRepository interface {
		AddSatuan(ctx context.Context, satuan entity.Satuan) (entity.Satuan, error)
		GetAllSatuanWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllSatuanRepositoryResponse, error)
		GetSatuanById(ctx context.Context, satuanId string) (entity.Satuan, error)
		UpdateSatuan(ctx context.Context, satuan entity.Satuan) (entity.Satuan, error)
		DeleteSatuan(ctx context.Context, satuanId string) error
	}
	satuanRepository struct {
		db *gorm.DB
	}
)

func NewSatuanRepository(db *gorm.DB) SatuanRepository {
	return &satuanRepository{
		db: db,
	}
}

func (r *satuanRepository) AddSatuan(ctx context.Context, satuan entity.Satuan) (entity.Satuan, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&satuan).Error; err != nil {
		return entity.Satuan{}, err
	}
	return satuan, nil
}

func (r *satuanRepository) GetAllSatuanWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllSatuanRepositoryResponse, error) {
	tx := r.db

	var satuans []entity.Satuan
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Satuan{}).Count(&count).Error; err != nil {
		return dto.GetAllSatuanRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&satuans).Error; err != nil {
		return dto.GetAllSatuanRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllSatuanRepositoryResponse{
		Satuans: satuans,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *satuanRepository) GetSatuanById(ctx context.Context, satuanId string) (entity.Satuan, error) {
	tx := r.db

	var satuan entity.Satuan
	if err := tx.WithContext(ctx).Where("id = ?", satuanId).Take(&satuan).Error; err != nil {
		return entity.Satuan{}, err
	}

	return satuan, nil
}
func (r *satuanRepository) UpdateSatuan(ctx context.Context, satuan entity.Satuan) (entity.Satuan, error) {
	tx := r.db

	// First, check if the record exists
	var existingSatuan entity.Satuan
	if err := tx.WithContext(ctx).Where("id = ?", satuan.ID).Take(&existingSatuan).Error; err != nil {
		// If the record doesn't exist, return a specific error
		if err == gorm.ErrRecordNotFound {
			return entity.Satuan{}, fmt.Errorf("Satuan with ID %s not found", satuan.ID)
		}
		return entity.Satuan{}, err
	}

	// Proceed with updating the record
	if err := tx.WithContext(ctx).Model(&existingSatuan).Updates(satuan).Error; err != nil {
		return entity.Satuan{}, err
	}

	// Return the updated entity
	return existingSatuan, nil
}
func (r *satuanRepository) DeleteSatuan(ctx context.Context, satuanId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Satuan{}, "id = ?", satuanId).Error; err != nil {
		return err
	}

	return nil
}
