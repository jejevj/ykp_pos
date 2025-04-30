package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MainSetting struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaUsaha  string    `json:"nama_usaha"`
	JenisUsaha string    `json:"jenis_usaha"`
	Alamat     string    `json:"alamat"`
	Hp         string    `json:"hp"`
	LogoUrl    string    `json:"logo_url"`

	Timestamp
}

func (u *MainSetting) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
