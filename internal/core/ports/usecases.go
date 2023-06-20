package ports

import "github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"

type ItemUseCase interface {
	List() ([]domain.Item, error)
	Create(name, category, description string, price float32) (*domain.Item, error)
	GetByID(itemID uint32) (*domain.Item, error)
	GetByCategory(category string) ([]domain.Item, error)
	Update(itemID uint32, name, category, description string, price float32) (*domain.Item, error)
	Delete(itemID uint32) error
}

type PedidoUseCase interface {
	List() ([]domain.Pedido, error)
	Create(pedido domain.Pedido) (*domain.Pedido, error)
	GetById(pedidoID uint32) (*domain.Pedido, error)
	Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error)
	Delete(pedidoID uint32) error
}

type ClienteUseCase interface {
	List() ([]domain.Cliente, error)
	Create(email, cpf, nome string) (*domain.Cliente, error)
	GetByID(clienteID uint32) (*domain.Cliente, error)
	Update(clienteID uint32, email, cpf, nome string) (*domain.Cliente, error)
	Delete(clienteID uint32) error
}
