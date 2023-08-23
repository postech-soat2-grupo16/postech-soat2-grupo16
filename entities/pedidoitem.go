package entities

import "gorm.io/gorm"

type PedidoItem struct {
	PedidoID uint32 `json:"-"`
	ItemID   uint32 `json:"-"`
	Item     Item   `gorm:"references:ID"`
	Quantity int
	gorm.Model
}
