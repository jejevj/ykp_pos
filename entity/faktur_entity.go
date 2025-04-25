package entity

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Faktur struct {
	ID            uuid.UUID             `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NoFaktur      string                `json:"no_faktur"`
	TanggalFaktur *time.Time            `gorm: json:"tanggal_faktur"`
	TanggalTempo  *time.Time            `gorm: json:"tanggal_tempo"`
	CaraBayar     string                `json:"cara_bayar"`
	IdUser        string                `json:"id_user"`
	Driver        User                  `gorm:"foreignKey:IdDriver" json:"driver"`
	BuktiBayar    *multipart.FileHeader `json:"bukti_bayar" form:"bukti_bayar"`
	Status        string                `gorm:"default:belum_bayar" json:"status"`

	Timestamp
}

func (u *Faktur) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return nil
}
