package pedido

import (
	"errors"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
)

type UseCase struct {
	pedidoRepository ports.PedidoRepository
}

func NewUseCase(repo ports.PedidoRepository) UseCase {
	return UseCase{pedidoRepository: repo}
}

func (p UseCase) List(status string) (pedidos []domain.Pedido, err error) {
	if status != "" {
		pedido := domain.Pedido{
			Status: status,
		}
		return p.pedidoRepository.GetAll(pedido)
	}

	return p.pedidoRepository.GetAll()
}

func (p UseCase) Create(pedido domain.Pedido) (*domain.Pedido, error) {
	return p.pedidoRepository.Save(pedido)
}

func (p UseCase) GetByID(pedidoID uint32) (*domain.Pedido, error) {
	pedido, err := p.pedidoRepository.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return pedido, nil
}

func (p UseCase) Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error) {
	return p.pedidoRepository.Update(pedidoID, pedido)
}

func (p UseCase) Delete(pedidoID uint32) error {
	return p.pedidoRepository.Delete(pedidoID)
}
