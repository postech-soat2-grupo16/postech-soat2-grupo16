package domain

import "gorm.io/gorm"

type Pedido struct {
	gorm.Model
	ID          uint32 `gorm:"primary_key;auto_increment"`
	PedidoItems []Item `gorm:"many2many:pedido_items;"`
	Status      string `gorm:"not null;"`
	Notes       string `gorm:"null;"`
	ClienteID   uint32
	Cliente     Cliente
}
