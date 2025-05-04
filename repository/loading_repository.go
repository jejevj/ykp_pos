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
	LoadingRepository interface {
		AddLoading(ctx context.Context, loading entity.Loading) (entity.Loading, error)
		GetAllLoadingWithPagination(ctx context.Context) (dto.GetAllLoadingRepositoryResponse, error)
		GetLoadingById(ctx context.Context, loadingId string) (entity.Loading, error)
		UpdateLoading(ctx context.Context, loading entity.Loading) (entity.Loading, error)
		DeleteLoading(ctx context.Context, loadingId string) error
	}
	loadingRepository struct {
		db *gorm.DB
	}
)

func NewLoadingRepository(db *gorm.DB) LoadingRepository {
	return &loadingRepository{
		db: db,
	}
}

func (r *loadingRepository) AddLoading(ctx context.Context, loading entity.Loading) (entity.Loading, error) {
	tx := r.db

	// Start a transaction
	err := tx.WithContext(ctx).Create(&loading).Error
	if err != nil {
		return entity.Loading{}, err
	}

	// Preload User based on the created Loading entity
	err = tx.WithContext(ctx).
		Preload("User").             // Preload associated User data
		Where("id = ?", loading.ID). // Make sure to use the correct field for loading's ID
		Take(&loading).Error
	if err != nil {
		return entity.Loading{}, err
	}

	return loading, nil
}

func (r *loadingRepository) GetAllLoadingWithPagination(ctx context.Context) (dto.GetAllLoadingRepositoryResponse, error) {
	tx := r.db

	var loadings []entity.Loading
	var err error
	var count int64

	if err := tx.WithContext(ctx).Model(&entity.Loading{}).Count(&count).Error; err != nil {
		return dto.GetAllLoadingRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Scopes(Paginate(1, 10)).Find(&loadings).Error; err != nil {
		return dto.GetAllLoadingRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllLoadingRepositoryResponse{
		Loadings: loadings,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *loadingRepository) GetLoadingById(ctx context.Context, loadingId string) (entity.Loading, error) {
	tx := r.db

	var loading entity.Loading
	if err := tx.WithContext(ctx).Preload("User").Where("id = ?", loadingId).Take(&loading).Error; err != nil {
		return entity.Loading{}, err
	}

	return loading, nil
}
func (r *loadingRepository) UpdateLoading(ctx context.Context, loading entity.Loading) (entity.Loading, error) {
	tx := r.db

	// First, check if the record exists
	var existingLoading entity.Loading
	if err := tx.WithContext(ctx).Where("id = ?", loading.ID).Take(&existingLoading).Error; err != nil {
		// If the record doesn't exist, return a specific error
		if err == gorm.ErrRecordNotFound {
			return entity.Loading{}, fmt.Errorf("Loading with ID %s not found", loading.ID)
		}
		return entity.Loading{}, err
	}

	// Proceed with updating the record
	if err := tx.WithContext(ctx).Model(&existingLoading).Updates(loading).Error; err != nil {
		return entity.Loading{}, err
	}

	// Return the updated entity
	return existingLoading, nil
}
func (r *loadingRepository) DeleteLoading(ctx context.Context, loadingId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Loading{}, "id = ?", loadingId).Error; err != nil {
		return err
	}

	return nil
}
