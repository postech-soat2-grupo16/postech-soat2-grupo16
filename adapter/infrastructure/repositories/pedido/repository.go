package pedido

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) *Repository {
	return &Repository{orm: orm}
}

func (p *Repository) Save(pedido domain.Pedido) (*domain.Pedido, error) {
	result := p.orm.Create(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *Repository) Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error) {
	pedido.ID = pedidoID
	for i := range pedido.Items {
		pedido.Items[i].PedidoID = pedidoID
	}
	result := p.orm.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pedido, nil
}

func (p *Repository) Delete(pedidoID uint32) error {
	pedido := domain.Pedido{
		ID: pedidoID,
	}
	result := p.orm.Delete(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (p *Repository) GetByID(pedidoID uint32) (*domain.Pedido, error) {
	pedido := domain.Pedido{
		ID: pedidoID,
	}
	result := p.orm.Preload(clause.Associations).Preload("Items.Item").First(&pedido)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *Repository) GetAll(conds ...interface{}) (pedidos []domain.Pedido, err error) {
	result := p.orm.Preload(clause.Associations).Preload("Items.Item").Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}}).Find(&pedidos, conds...)
	if result.Error != nil {
		log.Println(result.Error)
		return pedidos, result.Error
	}

	return pedidos, err
}
