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
	MainSettingRepository interface {
		AddMainSetting(ctx context.Context, msetting entity.MainSetting) (entity.MainSetting, error)
		GetAllMainSettingWithPagination(ctx context.Context) (dto.GetAllMainSettingRepositoryResponse, error)
		GetMainSettingById(ctx context.Context, msettingId string) (entity.MainSetting, error)
		UpdateMainSetting(ctx context.Context, msetting entity.MainSetting) (entity.MainSetting, error)
		DeleteMainSetting(ctx context.Context, msettingId string) error
	}
	mainSettingRepository struct {
		db *gorm.DB
	}
)

func NewMainSettingRepository(db *gorm.DB) MainSettingRepository {
	return &mainSettingRepository{
		db: db,
	}
}

func (r *mainSettingRepository) AddMainSetting(ctx context.Context, msetting entity.MainSetting) (entity.MainSetting, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&msetting).Error; err != nil {
		return entity.MainSetting{}, err
	}
	return msetting, nil
}

func (r *mainSettingRepository) GetAllMainSettingWithPagination(ctx context.Context) (dto.GetAllMainSettingRepositoryResponse, error) {
	tx := r.db

	var msettings []entity.MainSetting
	var err error
	var count int64

	if err := tx.WithContext(ctx).Model(&entity.MainSetting{}).Count(&count).Error; err != nil {
		return dto.GetAllMainSettingRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(1, 10)).Find(&msettings).Error; err != nil {
		return dto.GetAllMainSettingRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllMainSettingRepositoryResponse{
		MainSettings: msettings,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *mainSettingRepository) GetMainSettingById(ctx context.Context, msettingId string) (entity.MainSetting, error) {
	tx := r.db

	var msetting entity.MainSetting
	if err := tx.WithContext(ctx).Where("id = ?", msettingId).Take(&msetting).Error; err != nil {
		return entity.MainSetting{}, err
	}

	return msetting, nil
}
func (r *mainSettingRepository) UpdateMainSetting(ctx context.Context, msetting entity.MainSetting) (entity.MainSetting, error) {
	tx := r.db

	var existingMainSetting entity.MainSetting
	if err := tx.WithContext(ctx).Where("id = ?", msetting.ID).Take(&existingMainSetting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.MainSetting{}, fmt.Errorf("MainSetting with ID %s not found", msetting.ID)
		}
		return entity.MainSetting{}, err
	}

	if err := tx.WithContext(ctx).Model(&existingMainSetting).Updates(msetting).Error; err != nil {
		return entity.MainSetting{}, err
	}

	return existingMainSetting, nil
}
func (r *mainSettingRepository) DeleteMainSetting(ctx context.Context, msettingId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.MainSetting{}, "id = ?", msettingId).Error; err != nil {
		return err
	}

	return nil
}
