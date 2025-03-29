package migrations

import (
	"fmt"

	"github.com/jejevj/ykp_pos/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	queries := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}

	for _, query := range queries {
		result := db.Exec(query)
		if result.Error != nil {
			fmt.Println("Error executing query:", result.Error)
		} else {
			fmt.Println("Executed query successfully:", query)
		}
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Satuan{},
		&entity.Barang{},
		&entity.Loading{},
		&entity.Transaksi{},
		&entity.Customer{},
	); err != nil {
		return err
	}

	return nil
}
