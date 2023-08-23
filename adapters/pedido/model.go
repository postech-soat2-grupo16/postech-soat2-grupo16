package pedido

import (
	"time"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
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

type QRCode struct {
	QRCode string `json:"qr_code"`
}

type Pagamento struct {
	ID        uint32    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Pedido) PedidoItemToDomain() (list []entities.PedidoItem) {
	for _, pi := range p.Items {
		list = append(list, entities.PedidoItem{
			PedidoID: p.ID,
			ItemID:   pi.ItemID,
			Item: entities.Item{
				ID: pi.ItemID,
			},
			Quantity: pi.Quantity,
		})
	}
	return list
}

func (p *Pedido) ToDomain() entities.Pedido {
	return entities.Pedido{
		ID:        p.ID,
		Items:     p.PedidoItemToDomain(),
		Status:    entities.StatusPedidoCriado,
		Notes:     p.Notes,
		ClienteID: p.ClienteID,
		Cliente: entities.Cliente{
			ID: p.ClienteID,
		},
	}
}

type PaymentCallback struct {
	Id          int       `json:"id"`
	LiveMode    bool      `json:"live_mode"`
	Type        string    `json:"type"`
	DateCreated time.Time `json:"date_created"`
	UserId      int       `json:"user_id"`
	ApiVersion  string    `json:"api_version"`
	Action      string    `json:"action"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
}
