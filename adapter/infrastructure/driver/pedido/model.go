package pedido

import (
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

type Pedido struct {
	ID        uint32       `json:"id"`
	Items     []PedidoItem `json:"items"`
	Notes     string       `json:"notes"`
	ClienteID uint32       `json:"clienteId"`
}

type PedidoItem struct {
	ItemID   uint32 `json:"itemId"`
	Quantity int    `json:"quantity"`
}

func (p *Pedido) PedidoItemToDomain() (list []domain.PedidoItem) {
	for _, pi := range p.Items {
		list = append(list, domain.PedidoItem{
			PedidoID: p.ID,
			ItemID:   pi.ItemID,
			Item: domain.Item{
				ID: pi.ItemID,
			},
			Quantity: pi.Quantity,
		})
	}
	return list
}

func (p *Pedido) ToDomain() domain.Pedido {
	return domain.Pedido{
		ID:        p.ID,
		Items:     p.PedidoItemToDomain(),
		Status:    domain.PedidoCreated,
		Notes:     p.Notes,
		ClienteID: p.ClienteID,
		Cliente: domain.Cliente{
			ID: p.ClienteID,
		},
	}
}
