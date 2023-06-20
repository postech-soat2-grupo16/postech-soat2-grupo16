package cliente

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClienteRepository struct {
	orm *gorm.DB
}

func NewClienteRepository(orm *gorm.DB) *ClienteRepository {
	return &ClienteRepository{orm: orm}
}

func (p ClienteRepository) Save(cliente domain.Cliente) (*domain.Cliente, error) {
	result := p.orm.Create(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &cliente, nil
}

func (p ClienteRepository) Update(cliente domain.Cliente) (*domain.Cliente, error) {
	result := p.orm.Updates(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &cliente, nil
}

func (p ClienteRepository) Delete(clienteID uint32) error {
	cliente := domain.Cliente{
		ID: clienteID,
	}
	result := p.orm.Delete(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (p ClienteRepository) GetByID(clienteID uint32) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		ID: clienteID,
	}
	result := p.orm.Preload(clause.Associations).Preload("Items.Item").First(&cliente)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cliente, nil
}

func (p ClienteRepository) GetAll() (clientes []domain.Cliente, err error) {
	result := p.orm.Find(&clientes)
	if result.Error != nil {
		log.Println(result.Error)
		return clientes, result.Error
	}

	return clientes, err
}
