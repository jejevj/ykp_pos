package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaksi struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	IdLoading string    `json:"id_loading"`
	Loading   Loading   `gorm:"foreignKey:IdLoading" json:"loading"`
	IdBarang  string    `json:"id_barang"`
	Barang    Barang    `gorm:"foreignKey:IdBarang" json:"barang"`
	Jumlah    int       `json:"jumlah"`

	Timestamp
}

func (u *Transaksi) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return nil
}
