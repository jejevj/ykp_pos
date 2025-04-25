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
	CustomerRepository interface {
		AddCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error)
		GetAllCustomerWithPagination(ctx context.Context) (dto.GetAllCustomerRepositoryResponse, error)
		GetCustomerById(ctx context.Context, customerId string) (entity.Customer, error)
		UpdateCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error)
		DeleteCustomer(ctx context.Context, customerId string) error
	}
	customerRepository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) AddCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	tx := r.db

	if err := tx.WithContext(ctx).Create(&customer).Error; err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func (r *customerRepository) GetAllCustomerWithPagination(ctx context.Context) (dto.GetAllCustomerRepositoryResponse, error) {
	tx := r.db

	var customers []entity.Customer
	var err error
	var count int64

	if err := tx.WithContext(ctx).Model(&entity.Customer{}).Count(&count).Error; err != nil {
		return dto.GetAllCustomerRepositoryResponse{}, err
	}

	if err := tx.WithContext(ctx).Scopes(Paginate(1, 10)).Find(&customers).Error; err != nil {
		return dto.GetAllCustomerRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(10)))

	return dto.GetAllCustomerRepositoryResponse{
		Customers: customers,
		PaginationResponse: dto.PaginationResponse{
			Page:    1,
			PerPage: 10,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
func (r *customerRepository) GetCustomerById(ctx context.Context, customerId string) (entity.Customer, error) {
	tx := r.db

	var customer entity.Customer
	if err := tx.WithContext(ctx).Where("id = ?", customerId).Take(&customer).Error; err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}
func (r *customerRepository) UpdateCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	tx := r.db

	var existingCustomer entity.Customer
	if err := tx.WithContext(ctx).Where("id = ?", customer.ID).Take(&existingCustomer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Customer{}, fmt.Errorf("Customer with ID %s not found", customer.ID)
		}
		return entity.Customer{}, err
	}

	if err := tx.WithContext(ctx).Model(&existingCustomer).Updates(customer).Error; err != nil {
		return entity.Customer{}, err
	}

	return existingCustomer, nil
}
func (r *customerRepository) DeleteCustomer(ctx context.Context, customerId string) error {
	tx := r.db

	if err := tx.WithContext(ctx).Delete(&entity.Customer{}, "id = ?", customerId).Error; err != nil {
		return err
	}

	return nil
}
