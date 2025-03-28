package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	BarangCreateRequest struct {
		NamaBarang string `json:"nama_barang" form:"nama_barang"`
		KodeBarang string `json:"kode_barang" form:"kode_barang"`
		HargaBeli  int    `json:"harga_beli" form:"harga_beli"`
		HargaJual  int    `json:"harga_jual" form:"harga_jual"`
		IdSatuan   string `json:"id_satuan" form:"id_satuan"`
		Stok       int    `json:"stok" form:"stok"`
	}

	BarangResponse struct {
		ID         string         `json:"id"`
		NamaBarang string         `json:"nama_barang"`
		KodeBarang string         `json:"kode_barang"`
		HargaBeli  int            `json:"harga_beli"`
		HargaJual  int            `json:"harga_jual"`
		IdSatuan   string         `json:"id_satuan"`
		Satuan     SatuanResponse `json:"satuan"`
		Stok       int            `json:"stok"`
	}

	BarangPaginationResponse struct {
		Data []BarangResponse `json:"data"`
		PaginationResponse
	}

	GetAllBarangRepositoryResponse struct {
		Barangs []entity.Barang
		PaginationResponse
	}

	BarangUpdateRequest struct {
		NamaBarang string `json:"nama_barang" form:"nama_barang"`
		KodeBarang string `json:"kode_barang" form:"kode_barang"`
		HargaBeli  int    `json:"harga_beli" form:"harga_beli"`
		HargaJual  int    `json:"harga_jual" form:"harga_jual"`
		IdSatuan   string `json:"id_satuan" form:"id_satuan"`
		Stok       int    `json:"stok" form:"stok"`
	}

	BarangUpdateResponse struct {
		ID         string         `json:"id"`
		NamaBarang string         `json:"nama_barang"`
		KodeBarang string         `json:"kode_barang"`
		HargaBeli  int            `json:"harga_beli"`
		HargaJual  int            `json:"harga_jual"`
		IdSatuan   string         `json:"id_satuan"`
		Satuan     SatuanResponse `json:"satuan"`
		Stok       int            `json:"stok"`
	}
)
