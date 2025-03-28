package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	SatuanCreateRequest struct {
		NamaSatuan string `json:"nama_satuan" form:"nama_satuan"`
		Value      int    `json:"value" form:"value"`
	}

	SatuanResponse struct {
		ID         string `json:"id"`
		NamaSatuan string `json:"nama_satuan"`
		Value      int    `json:"value"`
	}

	SatuanPaginationResponse struct {
		Data []SatuanResponse `json:"data"`
		PaginationResponse
	}

	GetAllSatuanRepositoryResponse struct {
		Satuans []entity.Satuan
		PaginationResponse
	}

	GetSatuanByIdRequest struct {
		ID string `json:"id"`
	}

	SatuanUpdateRequest struct {
		ID         string `json:"id" form:"id"`
		NamaSatuan string `json:"nama_satuan" form:"nama_satuan"`
		Value      int    `json:"value" form:"value"`
	}

	SatuanUpdateResponse struct {
		ID         string `json:"id"`
		NamaSatuan string `json:"nama_satuan"`
		Value      int    `json:"value"`
	}
)
