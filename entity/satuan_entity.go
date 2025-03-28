package entity

import (
	"github.com/google/uuid"
	"github.com/jejevj/ykp_pos/helpers"
	"gorm.io/gorm"
)

type Satuan struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	NamaSatuan string    `json:"nama_satuan"`
	Value      int       `json:"value"`

	Timestamp
}

func (u *Satuan) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
