package ports

import "github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"

type ItemUseCase interface {
	List() ([]domain.Item, error)
	Create(name, category, description string, price float32) (*domain.Item, error)
	GetById(itemID uint32) (*domain.Item, error)
	Update(itemID uint32, name, category, description string, price float32) (*domain.Item, error)
	Delete(itemID uint32) error
}

type PedidoUseCase interface {
	List() ([]domain.Pedido, error)
	Create(pedido domain.Pedido) (*domain.Pedido, error)
	GetById(pedidoID uint32) (*domain.Pedido, error)
	Update(pedido domain.Pedido) (*domain.Pedido, error)
	Delete(pedidoID uint32) error
}
