package pedido

import (
	"errors"
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ ports.PedidoUseCase = (*PedidoUseCase)(nil)

type PedidoUseCase struct {
	repo *gorm.DB
}

func NewPedidoUseCase(repo *gorm.DB) PedidoUseCase {
	return PedidoUseCase{repo: repo}
}

func (p PedidoUseCase) List() (pedidos []domain.Pedido, err error) {
	result := p.repo.Preload(clause.Associations).Preload("Items.Item").Find(&pedidos)
	if result.Error != nil {
		log.Println(result.Error)
		return pedidos, result.Error
	}

	return pedidos, err
}

func (p PedidoUseCase) Create(pedido domain.Pedido) (*domain.Pedido, error) {
	result := p.repo.Create(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pedido, nil
}

func (p PedidoUseCase) GetById(pedidoID uint32) (*domain.Pedido, error) {
	pedido := domain.Pedido{
		ID: pedidoID,
	}
	result := p.repo.Preload(clause.Associations).Preload("Items.Item").First(&pedido)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(result.Error)
		return nil, result.Error
	}

	return &pedido, nil
}

func (p PedidoUseCase) Update(pedido domain.Pedido) (*domain.Pedido, error) {
	result := p.repo.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pedido, nil
}

func (p PedidoUseCase) Delete(pedidoID uint32) error {
	pedido := domain.Pedido{
		ID: pedidoID,
	}
	result := p.repo.Delete(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
