package pedido

import (
	"errors"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"gorm.io/gorm"
)

type UseCase struct {
	pedidoGateway      interfaces.PedidoGatewayI
	mercadoPagoGateway interfaces.MercadoPagoGatewayI
}

func NewUseCase(pedidoGateway interfaces.PedidoGatewayI, mercadoPagoGateway interfaces.MercadoPagoGatewayI) UseCase {
	return UseCase{pedidoGateway: pedidoGateway, mercadoPagoGateway: mercadoPagoGateway}
}

func (p UseCase) List(status string) (pedidos []entities.Pedido, err error) {
	if status != "" {
		pedido := entities.Pedido{
			Status: status,
		}
		return p.pedidoGateway.GetAll(pedido)
	}

	return p.pedidoGateway.GetAll()
}

func (p UseCase) Create(pedido entities.Pedido) (*entities.Pedido, error) {
	return p.pedidoGateway.Save(pedido)
}

func (p UseCase) CreateQRCode(pedidoID uint32) (*string, error) {
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	qrCode, err := p.mercadoPagoGateway.CreateQRCodeForPedido(*pedido)
	if err != nil {
		return nil, err
	}

	return &qrCode, nil
}

func (p UseCase) GetByID(pedidoID uint32) (*entities.Pedido, error) {
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return pedido, nil
}

func (p UseCase) GetLastPaymentStatus(pedidoID uint32) (*entities.Pagamento, error) {
	lastPayment, err := p.pedidoGateway.GetLastPaymentStatus(pedidoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return lastPayment, nil
}

func (p UseCase) UpdatePaymentStatusByPaymentID(pagamentoID string) (*entities.Pedido, error) {
	pedidoID, err := p.mercadoPagoGateway.GetPedidoIDByPaymentID(pagamentoID)
	if err != nil {
		return nil, err
	}
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	pedido.Status = entities.StatusPedidoEmPreparacao

	pedido.Pagamentos = append(pedido.Pagamentos, entities.Pagamento{
		Amount: pedido.GetAmount(),
		Status: entities.StatusPagamentoAprovado,
	})

	pedido, err = p.pedidoGateway.Update(pedidoID, *pedido)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return p.pedidoGateway.GetByID(pedidoID)
}

func (p UseCase) Update(pedidoID uint32, pedido entities.Pedido) (*entities.Pedido, error) {
	if _, err := p.pedidoGateway.GetByID(pedidoID); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return p.pedidoGateway.Update(pedidoID, pedido)
}

func (p UseCase) UpdatePedidoStatus(pedidoID uint32, pedidoStatus string) (*entities.Pedido, error) {
	pedido, err := p.pedidoGateway.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	pedido.Status = pedidoStatus
	return p.pedidoGateway.Update(pedidoID, *pedido)
}

func (p UseCase) Delete(pedidoID uint32) error {
	return p.pedidoGateway.Delete(pedidoID)
}
