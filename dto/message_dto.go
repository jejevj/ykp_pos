package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed add data"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list data"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get data token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get data"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update data"
	MESSAGE_FAILED_DELETE_USER             = "failed delete data"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL            = "failed verify email"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success add data"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list data"
	MESSAGE_SUCCESS_GET_USER                = "success get data"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update data"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete data"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
)

var (
	ErrCreateUser             = errors.New("failed to create data")
	ErrGetAllUser             = errors.New("failed to get all data")
	ErrGetUserById            = errors.New("failed to get data by id")
	ErrGetUserByEmail         = errors.New("failed to get data by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update data")
	ErrUserNotAdmin           = errors.New("data not admin")
	ErrUserNotFound           = errors.New("data not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete data")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")

	// Satuan Error
	ErrCreateSatuan   = errors.New("failed to create satuan")
	ErrGetSatuanById  = errors.New("failed to get satuan by id")
	ErrUpdateSatuan   = errors.New("failed to update satuan")
	ErrSatuanNotFound = errors.New("data not found")
	ErrDeleteSatuan   = errors.New("failed to delete satuan")

	// Barang Error
	ErrCreateBarang   = errors.New("failed to create barang")
	ErrGetBarangById  = errors.New("failed to get barang by id")
	ErrUpdateBarang   = errors.New("failed to update barang")
	ErrBarangNotFound = errors.New("data not found")
	ErrDeleteBarang   = errors.New("failed to delete barang")

	// Loading Error
	ErrCreateLoading   = errors.New("failed to create loading")
	ErrGetLoadingById  = errors.New("failed to get loading by id")
	ErrUpdateLoading   = errors.New("failed to update loading")
	ErrLoadingNotFound = errors.New("data not found")
	ErrDeleteLoading   = errors.New("failed to delete loading")
	// Transaksi Error
	ErrCreateTransaksi   = errors.New("failed to create transaksi")
	ErrGetTransaksiById  = errors.New("failed to get transaksi by id")
	ErrUpdateTransaksi   = errors.New("failed to update transaksi")
	ErrTransaksiNotFound = errors.New("data not found")
	ErrDeleteTransaksi   = errors.New("failed to delete transaksi")
)
