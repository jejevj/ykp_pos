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
		GetAllBarangWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllBarangRepositoryResponse, error)
		GetBarangById(ctx context.Context, barangId string) (entity.Barang, error)
		UpdateBarang(ctx context.Context, barang entity.Barang) (entity.Barang, error)
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

	if err := tx.WithContext(ctx).Create(&barang).Error; err != nil {
		return entity.Barang{}, err
	}
	return barang, nil
}

func (r *barangRepository) GetAllBarangWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.GetAllBarangRepositoryResponse, error) {
	tx := r.db

	var barangs []entity.Barang
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if err := tx.WithContext(ctx).Model(&entity.Barang{}).Count(&count).Error; err != nil {
		return dto.GetAllBarangRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&barangs).Error; err != nil {
		return dto.GetAllBarangRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllBarangRepositoryResponse{
		Barangs: barangs,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *barangRepository) GetBarangById(ctx context.Context, barangId string) (entity.Barang, error) {
	tx := r.db

	var barang entity.Barang
	if err := tx.WithContext(ctx).Where("id = ?", barangId).Take(&barang).Error; err != nil {
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
func (r *barangRepository) DeleteBarang(ctx context.Context, barangId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Barang{}, "id = ?", barangId).Error; err != nil {
		return err
	}

	return nil
}
