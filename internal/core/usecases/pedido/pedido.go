package pedido

import (
	"errors"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
)

type PedidoUseCase struct {
	pedidoRepository ports.PedidoRepository
}

func NewPedidoUseCase(repo ports.PedidoRepository) PedidoUseCase {
	return PedidoUseCase{pedidoRepository: repo}
}

func (p PedidoUseCase) List() (pedidos []domain.Pedido, err error) {
	return p.pedidoRepository.GetAll()
}

func (p PedidoUseCase) Create(pedido domain.Pedido) (*domain.Pedido, error) {
	return p.pedidoRepository.Save(pedido)
}

func (p PedidoUseCase) GetById(pedidoID uint32) (*domain.Pedido, error) {
	pedido, err := p.pedidoRepository.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return pedido, nil
}

func (p PedidoUseCase) Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error) {
	return p.pedidoRepository.Update(pedidoID, pedido)
}

func (p PedidoUseCase) Delete(pedidoID uint32) error {
	return p.pedidoRepository.Delete(pedidoID)
}
