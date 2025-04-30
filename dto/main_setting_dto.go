package dto

import (
	"mime/multipart"

	"github.com/jejevj/ykp_pos/entity"
)

type (
	MainSettingCreateRequest struct {
		NamaUsaha  string                `json:"nama_usaha" form:"nama_usaha"`
		JenisUsaha string                `json:"jenis_usaha" form:"jenis_usaha"`
		Alamat     string                `json:"alamat" form:"alamat"`
		Logo       *multipart.FileHeader `json:"logo" form:"logo"`
		Hp         string                `json:"hp" form:"hp"`
	}
	GetMainSettingByIdRequest struct {
		ID string `json:"id" form:"id"`
	}

	MainSettingResponse struct {
		ID         string `json:"id"`
		NamaUsaha  string `json:"nama_usaha"`
		JenisUsaha string `json:"jenis_usaha"`
		Alamat     string `json:"alamat"`
		LogoUrl    string `json:"logo" form:"logo"`
		Hp         string `json:"hp"`
	}

	MainSettingPaginationResponse struct {
		Data []MainSettingResponse `json:"data"`
		PaginationResponse
	}

	GetAllMainSettingRepositoryResponse struct {
		MainSettings []entity.MainSetting
		PaginationResponse
	}

	MainSettingUpdateRequest struct {
		ID         string                `json:"id" form:"id"`
		NamaUsaha  string                `json:"nama_usaha" form:"nama_usaha"`
		JenisUsaha string                `json:"jenis_usaha" form:"jenis_usaha"`
		Alamat     string                `json:"alamat" form:"alamat"`
		Logo       *multipart.FileHeader `json:"logo" form:"logo"`
		Hp         string                `json:"hp" form:"hp"`
	}

	MainSettingUpdateResponse struct {
		ID         string `json:"id"`
		NamaUsaha  string `json:"nama_usaha"`
		JenisUsaha string `json:"jenis_usaha"`
		Alamat     string `json:"alamat"`
		Logo       string `json:"logo" form:"logo"`
		Hp         string `json:"hp"`
	}
)
