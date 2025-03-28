package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	TransaksiCreateRequest struct {
		IdLoading string `json:"id_loading" form:"id_loading"`
		IdBarang  string `json:"id_barang" form:"id_barang"`
		Jumlah    int    `json:"jumlah" form:"jumlah"`
	}
	TransaksiResponse struct {
		ID        string          `json:"id"`
		IdLoading string          `json:"id_loading"`
		Loading   LoadingResponse `json:"loading"`
		IdBarang  string          `json:"id_barang"`
		Barang    BarangResponse  `json:"barang"`
		Jumlah    int             `json:"jumlah"`
	}
	TransaksiPaginationResponse struct {
		Data []TransaksiResponse `json:"data"`
		PaginationResponse
	}

	GetAllTransaksiRepositoryResponse struct {
		Transaksis []entity.Transaksi
		PaginationResponse
	}
	TransaksiUpdateRequest struct {
		IdLoading string `json:"id_loading" form:"id_loading"`
		IdBarang  string `json:"id_barang" form:"id_barang"`
		Jumlah    int    `json:"jumlah" form:"jumlah"`
	}
	TransaksiUpdateResponse struct {
		ID        string          `json:"id"`
		IdLoading string          `json:"id_loading"`
		Loading   LoadingResponse `json:"loading"`
		IdBarang  string          `json:"id_barang"`
		Barang    BarangResponse  `json:"barang"`
		Jumlah    int             `json:"jumlah"`
	}
)
