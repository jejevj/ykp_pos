package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiFaktur struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	IdFaktur string    `json:"id_faktur"`
	IdBarang string    `json:"id_barang"`
	Barang   Barang    `gorm:"foreignKey:IdBarang" json:"barang"`
	Krat     int       `json:"krat"`
	Lusin    int       `json:"lusin"`
	Satuan   int       `json:"satuan"`
	JumlahRP int       `json:"jumlah_rp"`
	Diskon   int       `json:"diskon"`
	DiskonP  float32   `json:"diskon_p"`
	Ket      string    `json:"keterangan"`

	Timestamp
}

func (u *TransaksiFaktur) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return nil
}
