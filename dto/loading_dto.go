package dto

import (
	"github.com/jejevj/ykp_pos/entity"
)

type (
	LoadingCreateRequest struct {
		IdUser string `json:"id_user" form:"id_user"`
	}

	GetLoadingByIdRequest struct {
		ID string `json:"id" form:"id"`
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
		ID         string `json:"id" form:"id"`
		IsApproved bool   `json:"is_approved" form:"is_approved"`
	}

	LoadingUpdateResponse struct {
		ID         string       `json:"id"`
		IdUser     string       `json:"id_user"`
		User       UserResponse `json:"user"`
		IsApproved bool         `json:"is_approved"`
	}
)
