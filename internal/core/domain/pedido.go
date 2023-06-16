package domain

import "gorm.io/gorm"

const (
	PedidoCreated = "CREATED"
)

type Pedido struct {
	ID        uint32 `gorm:"primary_key;auto_increment"`
	Items     []PedidoItem
	Status    string  `gorm:"not null;"`
	Notes     string  `gorm:"null;"`
	ClienteID uint32  `json:"-"`
	Cliente   Cliente `gorm:"references:ID"`
	gorm.Model
}
