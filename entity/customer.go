package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaToko    string    `json:"nama_toko"`
	NamaPemilik string    `json:"nama_pemilik"`
	Alamat      string    `json:"alamat"`
	HP          string    `json:"HP"`

	Timestamp
}

func (u *Customer) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
