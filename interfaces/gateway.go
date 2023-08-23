package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type PedidoGatewayI interface {
	Save(pedido entities.Pedido) (*entities.Pedido, error)
	Update(pedidoID uint32, pedido entities.Pedido) (*entities.Pedido, error)
	Delete(pedidoID uint32) error
	GetByID(pedidoID uint32) (*entities.Pedido, error)
	GetLastPaymentStatus(pedidoID uint32) (*entities.Pagamento, error)
	GetAll(conds ...interface{}) ([]entities.Pedido, error)
}

type ClienteGatewayI interface {
	Save(cliente entities.Cliente) (*entities.Cliente, error)
	Update(cliente entities.Cliente) (*entities.Cliente, error)
	Delete(clienteID uint32) error
	GetByID(clienteID uint32) (*entities.Cliente, error)
	GetAll(conds ...interface{}) ([]entities.Cliente, error)
}

type ItemGatewayI interface {
	Save(item entities.Item) (*entities.Item, error)
	Update(item entities.Item) (*entities.Item, error)
	Delete(itemID uint32) error
	GetByID(itemID uint32) (*entities.Item, error)
	GetAll() ([]entities.Item, error)
	GetByCategory(category string) ([]entities.Item, error)
}

type MercadoPagoGatewayI interface {
	GetPedidoIDByPaymentID(paymentID string) (uint32, error)
	CreateQRCodeForPedido(pedido entities.Pedido) (string, error)
}
