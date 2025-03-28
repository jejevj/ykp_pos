package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	LoadingCreateRequest struct {
		IdUser string `json:"id_user" form:"id_user"`
	}

	LoadingResponse struct {
		ID         string       `json:"id"`
		IdUser     string       `json:"id_user"`
		User       UserResponse `json:"user"`
		IsApproved bool         `json:"is_approved"`
	}

	LoadingPaginationResponse struct {
		Data []LoadingResponse `json:"data"`
		PaginationResponse
	}

	GetAllLoadingRepositoryResponse struct {
		Loadings []entity.Loading
		PaginationResponse
	}

	LoadingUpdateRequest struct {
		NamaBarang string `json:"nama_barang" form:"nama_barang"`
		KodeBarang string `json:"kode_barang" form:"kode_barang"`
		HargaBeli  int    `json:"harga_beli" form:"harga_beli"`
		HargaJual  int    `json:"harga_jual" form:"harga_jual"`
		IdSatuan   string `json:"id_satuan" form:"id_satuan"`
		Stok       int    `json:"stok" form:"stok"`
	}

	LoadingUpdateResponse struct {
		ID         string       `json:"id"`
		IdUser     string       `json:"id_user"`
		User       UserResponse `json:"user"`
		IsApproved bool         `json:"is_approved"`
	}
)
