package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	SatuanCreateRequest struct {
		NameSatuan string `json:"nama_satuan" form:"nama_satuan"`
		Value      int    `json:"value" form:"value"`
	}

	SatuanResponse struct {
		ID         string `json:"id"`
		NamaSatuan string `json:"nama_satuan"`
		Value      string `json:"value"`
	}

	SatuanPaginationResponse struct {
		Data []SatuanResponse `json:"data"`
		PaginationResponse
	}

	GetAllSatuanRepositoryResponse struct {
		Satuans []entity.Satuan
		PaginationResponse
	}

	SatuanUpdateRequest struct {
		NameSatuan string `json:"nama_satuan" form:"nama_satuan"`
		Value      int    `json:"value" form:"value"`
	}

	SatuanUpdateResponse struct {
		ID         string `json:"id"`
		NamaSatuan string `json:"nama_satuan"`
		Value      string `json:"value"`
	}
)
