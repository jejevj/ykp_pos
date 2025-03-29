package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	CustomerCreateRequest struct {
		NamaToko    string `json:"nama_toko" form:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik" form:"nama_pemilik"`
		Alamat      int    `json:"alamat" form:"alamat"`
		HP          int    `json:"hp" form:"hp"`
	}
	GetCustomerByIdRequest struct {
		ID string `json:"id" form:"id"`
	}

	CustomerResponse struct {
		ID          string `json:"id"`
		NamaToko    string `json:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik"`
		Alamat      int    `json:"alamat"`
		HP          int    `json:"hp"`
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
		Alamat      int    `json:"alamat"`
		HP          int    `json:"hp"`
	}

	CustomerUpdateResponse struct {
		NamaToko    string `json:"nama_toko"`
		NamaPemilik string `json:"nama_pemilik"`
		Alamat      int    `json:"alamat"`
		HP          int    `json:"hp"`
	}
)
