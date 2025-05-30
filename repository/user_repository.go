package repository

import (
	"context"
	"math"

	"github.com/jejevj/ykp_pos/dto"
	"github.com/jejevj/ykp_pos/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
		GetAllUserWithPagination(ctx context.Context) (dto.GetAllUserRepositoryResponse, error)
		GetUserById(ctx context.Context, userId string) (entity.User, error)
		GetUserByEmail(ctx context.Context, email string) (entity.User, error)
		CheckEmail(ctx context.Context, email string) (entity.User, bool, error)
		UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
		DeleteUser(ctx context.Context, userId string) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetAllUserWithPagination(ctx context.Context) (dto.GetAllUserRepositoryResponse, error) {
	tx := r.db

	var users []entity.User
	var err error
	var count int64

	if err := tx.WithContext(ctx).Model(&entity.User{}).Count(&count).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(1, 10)).Find(&users).Error; err != nil {
		return dto.GetAllUserRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllUserRepositoryResponse{
		Users: users,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *userRepository) GetUserById(ctx context.Context, userId string) (entity.User, error) {
	tx := r.db

	var user entity.User
	if err := tx.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	tx := r.db

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, email string) (entity.User, bool, error) {
	tx := r.db

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.User{}, "id = ?", userId).Error; err != nil {
		return err
	}

	return nil
}
