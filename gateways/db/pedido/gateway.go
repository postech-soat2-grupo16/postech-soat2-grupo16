package pedido

import (
	"fmt"
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (p *Repository) Save(pedido entities.Pedido) (*entities.Pedido, error) {
	result := p.repository.Create(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *Repository) Update(pedidoID uint32, pedido entities.Pedido) (*entities.Pedido, error) {
	pedido.ID = pedidoID
	for i := range pedido.Items {
		pedido.Items[i].PedidoID = pedidoID
	}

	for pa := range pedido.Pagamentos {
		pedido.Pagamentos[pa].PedidoID = pedidoID
	}
	result := p.repository.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pedido, nil
}

func (p *Repository) Delete(pedidoID uint32) error {
	pedido := entities.Pedido{
		ID: pedidoID,
	}
	result := p.repository.Preload("Items.Item").Delete(&pedido)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (p *Repository) GetByID(pedidoID uint32) (*entities.Pedido, error) {
	pedido := entities.Pedido{
		ID: pedidoID,
	}
	result := p.repository.Preload(clause.Associations).Preload("Items.Item").Preload("Pagamentos").First(&pedido)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pedido, nil
}

func (p *Repository) GetLastPaymentStatus(pedidoID uint32) (*entities.Pagamento, error) {
	pagamento := entities.Pagamento{
		PedidoID: pedidoID,
	}
	result := p.repository.Preload(clause.Associations).Where("pedido_id = ?", pedidoID).Last(&pagamento)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pagamento, nil
}

func (p *Repository) GetAll(conds ...interface{}) (pedidos []entities.Pedido, err error) {
	expressionOrderBy := fmt.Sprintf(
		"created_at, CASE status WHEN '%s' THEN 1 WHEN '%s' THEN 2 WHEN '%s' THEN 3 ELSE 4 END",
		entities.StatusPedidoPronto,
		entities.StatusPedidoEmPreparacao,
		entities.StatusPedidoRecebido,
	)

	result := p.repository.Debug().Preload(clause.Associations).Preload("Items.Item").
		Preload("Pagamentos").
		Where("status != ?", entities.StatusPedidoFinalizado).
		Clauses(clause.OrderBy{Expression: clause.Expr{
			SQL:                expressionOrderBy,
			WithoutParentheses: true,
		}}).
		Find(&pedidos, conds...)
	if result.Error != nil {
		log.Println(result.Error)
		return pedidos, result.Error
	}

	return pedidos, err
}
