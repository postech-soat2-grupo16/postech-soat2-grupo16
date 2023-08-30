package interfaces

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

type ItemUseCase interface {
	List() ([]entities.Item, error)
	Create(name, category, description string, price float32) (*entities.Item, error)
	GetByID(itemID uint32) (*entities.Item, error)
	GetByCategory(category string) ([]entities.Item, error)
	Update(itemID uint32, name, category, description string, price float32) (*entities.Item, error)
	Delete(itemID uint32) (*entities.Item, error)
}

type PedidoUseCase interface {
	List(status string) ([]entities.Pedido, error)
	Create(pedido entities.Pedido) (*entities.Pedido, error)
	GetByID(pedidoID uint32) (*entities.Pedido, error)
	GetLastPaymentStatus(pedidoID uint32) (*entities.Pagamento, error)
	Update(pedidoID uint32, pedido entities.Pedido) (*entities.Pedido, error)
	UpdatePedidoStatus(pedidoID uint32, pedidoStatus string) (*entities.Pedido, error)
	UpdatePaymentStatusByPaymentID(pagamentoID string) (*entities.Pedido, error)
	Delete(pedidoID uint32) error
	CreateQRCode(pedidoID uint32) (*string, error)
}

type ClienteUseCase interface {
	List(cpf string) ([]entities.Cliente, error)
	Create(email, cpf, nome string) (*entities.Cliente, error)
	GetByID(clienteID uint32) (*entities.Cliente, error)
	Update(clienteID uint32, email, cpf, nome string) (*entities.Cliente, error)
	Delete(clienteID uint32) error
}
