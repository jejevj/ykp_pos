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
	CustomerService interface {
		AddCustomer(ctx context.Context, req dto.CustomerCreateRequest) (dto.CustomerResponse, error)
		GetAllCustomerWithPagination(ctx context.Context) (dto.CustomerPaginationResponse, error)
		GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error)
		UpdateCustomer(ctx context.Context, req dto.CustomerUpdateRequest, customerId string) (dto.CustomerUpdateResponse, error)
		DeleteCustomer(ctx context.Context, customerId string) error
	}
	customerService struct {
		customerRepo repository.CustomerRepository
		jwtService   JWTService
	}
)

func NewCustomerService(customerRepo repository.CustomerRepository, jwtService JWTService) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
		jwtService:   jwtService,
	}
}
func (s *customerService) AddCustomer(ctx context.Context, req dto.CustomerCreateRequest) (dto.CustomerResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	customer := entity.Customer{
		NamaToko:    req.NamaToko,
		NamaPemilik: req.NamaPemilik,
		Alamat:      req.Alamat,
		HP:          req.HP,
	}

	customerAdd, err := s.customerRepo.AddCustomer(ctx, customer)
	if err != nil {
		return dto.CustomerResponse{}, dto.ErrCreateCustomer
	}

	return dto.CustomerResponse{
		ID:          customerAdd.ID.String(),
		NamaToko:    customerAdd.NamaToko,
		NamaPemilik: customerAdd.NamaPemilik,
		Alamat:      customerAdd.Alamat,
		HP:          customerAdd.HP,
	}, nil
}
func (s *customerService) GetAllCustomerWithPagination(ctx context.Context) (dto.CustomerPaginationResponse, error) {
	dataWithPaginate, err := s.customerRepo.GetAllCustomerWithPagination(ctx)
	if err != nil {
		return dto.CustomerPaginationResponse{}, err
	}

	var datas []dto.CustomerResponse
	for _, customer := range dataWithPaginate.Customers {
		data := dto.CustomerResponse{
			ID:          customer.ID.String(),
			NamaToko:    customer.NamaToko,
			NamaPemilik: customer.NamaPemilik,
			Alamat:      customer.Alamat,
			HP:          customer.HP,
		}

		datas = append(datas, data)
	}

	return dto.CustomerPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (s *customerService) GetCustomerById(ctx context.Context, customerId string) (dto.CustomerResponse, error) {
	customer, err := s.customerRepo.GetCustomerById(ctx, customerId)
	if err != nil {
		return dto.CustomerResponse{}, dto.ErrGetCustomerById
	}

	return dto.CustomerResponse{
		ID:          customer.ID.String(),
		NamaToko:    customer.NamaToko,
		NamaPemilik: customer.NamaPemilik,
		Alamat:      customer.Alamat,
		HP:          customer.HP,
	}, nil
}
func (s *customerService) UpdateCustomer(ctx context.Context, req dto.CustomerUpdateRequest, customerId string) (dto.CustomerUpdateResponse, error) {
	// Convert string ID to uuid.UUID (if needed)
	id, err := uuid.Parse(customerId)
	if err != nil {
		return dto.CustomerUpdateResponse{}, fmt.Errorf("invalid ID format: %v", err)
	}

	// Prepare the entity to be updated
	data := entity.Customer{
		ID:          id,
		NamaToko:    req.NamaToko,
		NamaPemilik: req.NamaPemilik,
		Alamat:      req.Alamat,
		HP:          req.HP,
	}

	// Call the repository to update
	customerUpdate, err := s.customerRepo.UpdateCustomer(ctx, data)
	if err != nil {
		return dto.CustomerUpdateResponse{}, fmt.Errorf("failed to update Customer: %v", err)
	} // Convert Satuan entity to SatuanResponse DTO

	return dto.CustomerUpdateResponse{
		ID:          customerUpdate.ID.String(),
		NamaToko:    customerUpdate.NamaToko,
		NamaPemilik: customerUpdate.NamaPemilik,
		Alamat:      customerUpdate.Alamat,
		HP:          customerUpdate.HP,
	}, nil
}

func (s *customerService) DeleteCustomer(ctx context.Context, customerId string) error {
	customer, err := s.customerRepo.GetCustomerById(ctx, customerId)
	if err != nil {
		return dto.ErrCustomerNotFound
	}

	err = s.customerRepo.DeleteCustomer(ctx, customer.ID.String())
	if err != nil {
		return dto.ErrDeleteCustomer
	}

	return nil
}
