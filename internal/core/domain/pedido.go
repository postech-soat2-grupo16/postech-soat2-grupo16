package domain

import "gorm.io/gorm"

const (
	// Status of pedidos
	StatusPedidoCriado              = "CRIADO"
	StatusPedidoAguardandoPagamento = "AGUARDANDO_PAGAMENTO"
	StatusPedidoRecebido            = "RECEBIDO"
	StatusPedidoEmPreparacao        = "EM_PREPARACAO"
	StatusPedidoPronto              = "PRONTO"
	StatusPedidoEntregue            = "ENTREGUE"
)

type Pedido struct {
	ID        uint32       `gorm:"primary_key;auto_increment" json:"id"`
	Items     []PedidoItem `json:"items"`
	Status    string       `gorm:"not null" json:"status"`
	Notes     string       `gorm:"null" json:"notes"`
	ClienteID uint32       `json:"-" json:"cliente_id"`
	Cliente   Cliente      `gorm:"references:ID" json:"cliente"`
	// TODO: include gorm created_at and updated_at fields in the json with underscore.
	gorm.Model
}
