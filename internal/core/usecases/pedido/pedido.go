package pedido

import (
	"errors"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
)

type UseCase struct {
	pedidoRepository      ports.PedidoRepository
	mercadoPagoRepository ports.MercadoPagoRepository
}

func NewUseCase(repo ports.PedidoRepository, mercadoPagoRepository ports.MercadoPagoRepository) UseCase {
	return UseCase{pedidoRepository: repo, mercadoPagoRepository: mercadoPagoRepository}
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

func (p UseCase) GetLastPaymentStatus(pedidoID uint32) (*domain.Pagamento, error) {
	lastPayment, err := p.pedidoRepository.GetLastPaymentStatus(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return lastPayment, nil
}

func (p UseCase) UpdatePaymentStatusByPaymentID(pagamentoID string) (*domain.Pedido, error) {
	pedidoID, err := p.mercadoPagoRepository.GetPedidoIDByPaymentID(pagamentoID)
	if err != nil {
		return nil, err
	}
	pedido, err := p.pedidoRepository.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	pedido.Status = domain.StatusPedidoEmPreparacao

	pedido.Pagamentos = append(pedido.Pagamentos, domain.Pagamento{
		Amount: pedido.GetAmount(),
		Status: domain.StatusPagamentoAprovado,
	})

	pedido, err = p.pedidoRepository.Update(pedidoID, *pedido)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return p.pedidoRepository.GetByID(pedidoID)
}

func (p UseCase) Update(pedidoID uint32, pedido domain.Pedido) (*domain.Pedido, error) {
	if _, err := p.pedidoRepository.GetByID(pedidoID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return p.pedidoRepository.Update(pedidoID, pedido)
}

func (p UseCase) Delete(pedidoID uint32) error {
	return p.pedidoRepository.Delete(pedidoID)
}
