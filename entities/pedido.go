package entities

import (
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

const (
	// Status of pedidos
	StatusPedidoCriado       = "CRIADO"
	StatusPedidoRecebido     = "RECEBIDO"
	StatusPedidoEmPreparacao = "EM_PREPARACAO"
	StatusPedidoPronto       = "PRONTO"
	StatusPedidoEntregue     = "ENTREGUE"
	StatusPedidoFinalizado   = "FINALIZADO"
	StatusPagamentoAprovado  = "APROVADO"
	StatusPagamentoNegado    = "NEGADO"
)

type Pedido struct {
	ID         uint32       `gorm:"primary_key;auto_increment" json:"id"`
	Items      []PedidoItem `json:"items"`
	Status     string       `gorm:"not null" json:"status"`
	Notes      string       `gorm:"null" json:"notes"`
	ClienteID  uint32       `json:"-" json:"cliente_id"`
	Cliente    Cliente      `gorm:"references:ID" json:"cliente"`
	Pagamentos []Pagamento  `gorm:"constraint:OnDelete:CASCADE" json:"pagamentos"`
	gorm.Model
}

func (p *Pedido) IsStatusValid() bool {
	status := []string{StatusPedidoCriado, StatusPedidoRecebido, StatusPedidoEmPreparacao, StatusPedidoPronto, StatusPedidoEntregue, StatusPedidoFinalizado}
	return slices.Contains(status, p.Status)
}

func (p *Pedido) GetAmount() float64 {
	var amount float64
	for _, item := range p.Items {
		amount += float64(item.Item.Price) * float64(item.Quantity)
	}
	return amount
}
