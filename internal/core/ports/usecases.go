package ports

import "github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"

type ItemUseCase interface {
	List() ([]domain.Item, error)
	Create(name, category, description string, price float32) (*domain.Item, error)
	GetByID(itemID uint32) (*domain.Item, error)
	GetByCategory(category string) ([]domain.Item, error)
	Update(itemID uint32, name, category, description string, price float32) (*domain.Item, error)
	Delete(itemID uint32) (*domain.Item, error)
}

type PedidoUseCase interface {
	List(status string) ([]domain.Pedido, error)
	Create(pedido domain.Pedido) (*domain.Pedido, error)
	GetByID(pedidoID uint32) (*domain.Pedido, error)
	GetLastPaymentStatus(pedidoID uint32) (*domain.Pagamento, error)
	Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error)
	UpdatePaymentStatusByPaymentID(pagamentoID string) (*domain.Pedido, error)
	Delete(pedidoID uint32) error
	CreateQRCode(pedidoID uint32) (*string, error)
}

type ClienteUseCase interface {
	List(cpf string) ([]domain.Cliente, error)
	Create(email, cpf, nome string) (*domain.Cliente, error)
	GetByID(clienteID uint32) (*domain.Cliente, error)
	Update(clienteID uint32, email, cpf, nome string) (*domain.Cliente, error)
	Delete(clienteID uint32) error
}
