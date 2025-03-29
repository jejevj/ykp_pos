package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	CustomerCreateRequest struct {
		NamaToko    string `json:"nama_toko" form:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik" form:"nama_pemilik"`
		Alamat      string `json:"alamat" form:"alamat"`
		HP          string `json:"hp" form:"hp"`
	}
	GetCustomerByIdRequest struct {
		ID string `json:"id" form:"id"`
	}

	CustomerResponse struct {
		ID          string `json:"id"`
		NamaToko    string `json:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik"`
		Alamat      string `json:"alamat"`
		HP          string `json:"hp"`
	}

	CustomerPaginationResponse struct {
		Data []CustomerResponse `json:"data"`
		PaginationResponse
	}

	GetAllCustomerRepositoryResponse struct {
		Customers []entity.Customer
		PaginationResponse
	}

	CustomerUpdateRequest struct {
		ID          string `json:"id" form:"id"`
		NamaToko    string `json:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik"`
		Alamat      string `json:"alamat"`
		HP          string `json:"hp"`
	}

	CustomerUpdateResponse struct {
		ID          string `json:"id"`
		NamaToko    string `json:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik"`
		Alamat      string `json:"alamat"`
		HP          string `json:"hp"`
	}
)
