package pedido

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PedidoRepository struct {
	orm *gorm.DB
}

func NewPedidoRepository(orm *gorm.DB) *PedidoRepository {
	return &PedidoRepository{orm: orm}
}

func (p *PedidoRepository) Save(pedido domain.Pedido) (*domain.Pedido, error) {
	result := p.orm.Create(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *PedidoRepository) Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error) {
	pedido.ID = pedidoID
	for i, _ := range pedido.Items {
		pedido.Items[i].PedidoID = pedidoID
	}
	result := p.orm.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pedido, nil
}

func (p *PedidoRepository) Delete(pedidoID uint32) error {
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

func (p *PedidoRepository) GetByID(pedidoID uint32) (*domain.Pedido, error) {
	pedido := domain.Pedido{
		ID: pedidoID,
	}
	result := p.orm.Preload(clause.Associations).Preload("Items.Item").First(&pedido)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *PedidoRepository) GetAll(conds ...interface{}) (pedidos []domain.Pedido, err error) {
	result := p.orm.Preload(clause.Associations).Preload("Items.Item").Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}}).Find(&pedidos, conds...)
	if result.Error != nil {
		log.Println(result.Error)
		return pedidos, result.Error
	}

	return pedidos, err
}
