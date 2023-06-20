package ports

import (
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

type GenericRepository interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	Session(config *gorm.Session) *gorm.DB
}

type PedidoRepository interface {
	Save(pedido domain.Pedido) (*domain.Pedido, error)
	Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error)
	Delete(pedidoID uint32) error
	GetByID(pedidoID uint32) (*domain.Pedido, error)
	GetAll() ([]domain.Pedido, error)
}

type ClienteRepository interface {
	Save(cliente domain.Cliente) (*domain.Cliente, error)
	Update(cliente domain.Cliente) (*domain.Cliente, error)
	Delete(clienteID uint32) error
	GetByID(clienteID uint32) (*domain.Cliente, error)
	GetAll() ([]domain.Cliente, error)
}

type ItemRepository interface {
	Save(item domain.Item) (*domain.Item, error)
	Update(item domain.Item) (*domain.Item, error)
	Delete(itemID uint32) error
	GetByID(itemID uint32) (*domain.Item, error)
	GetAll() ([]domain.Item, error)
	GetByCategory(category string) ([]domain.Item, error)
}
