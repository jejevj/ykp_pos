package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NoPelanggan   string    `json:"no+pelanggan"`
	NamaPelanggan string    `gorm: json:"nama_pelanggan"`
	Alamat        string    `gorm: json:"alamat"`
	Telp          string    `json:"telp"`
	Npwp          string    `json:"npwp" `

	Timestamp
}

func (u *Pelanggan) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return nil
}
