package database

import (
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

func DoMigration(db *gorm.DB) {
	db.AutoMigrate(domain.Product{})
	db.AutoMigrate(domain.Cliente{})
	db.AutoMigrate(domain.Item{})
	db.AutoMigrate(domain.Pedido{})
}
