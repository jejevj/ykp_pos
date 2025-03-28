package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Loading struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	IdUser     uuid.UUID `json:"id_user"`
	User       User      `gorm:"foreignKey:IdUser" json:"user"`
	IsApproved bool      `json:"is_approved"`

	Timestamp
}

func (u *Loading) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return nil
}
