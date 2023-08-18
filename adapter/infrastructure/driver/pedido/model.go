package pedido

import (
	"time"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

type Pedido struct {
	ID        uint32 `json:"id"`
	Items     []Item `json:"items"`
	Notes     string `json:"notes"`
	ClienteID uint32 `json:"clienteId"`
}

type Item struct {
	ItemID   uint32 `json:"itemId"`
	Quantity int    `json:"quantity"`
}

type Pagamento struct {
	ID        uint32    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
		Status:    domain.StatusPedidoCriado,
		Notes:     p.Notes,
		ClienteID: p.ClienteID,
		Cliente: domain.Cliente{
			ID: p.ClienteID,
		},
	}
}
