package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Barang struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaBarang string    `json:"nama_barang"`
	KodeBarang string    `json:"kode_barang"`
	HargaBeli  string    `json:"harga_beli"`
	HargaJual  string    `json:"harga_jual"`
	IdSatuan   uuid.UUID `json:"id_satuan"`
	Satuan     Satuan    `gorm:"foreignKey:IdSatuan" json:"satuan"`
	Stok       int       `json:"stok"`

	Timestamp
}

func (u *Barang) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
