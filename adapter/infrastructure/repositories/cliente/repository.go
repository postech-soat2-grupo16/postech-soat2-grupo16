package cliente

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

type Repository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) *Repository {
	return &Repository{orm: orm}
}

func (c *Repository) Save(cliente domain.Cliente) (*domain.Cliente, error) {
	result := c.orm.Create(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &cliente, nil
}

func (c *Repository) Update(cliente domain.Cliente) (*domain.Cliente, error) {
	result := c.orm.Updates(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &cliente, nil
}

func (c *Repository) Delete(clienteID uint32) error {
	cliente := domain.Cliente{
		ID: clienteID,
	}
	result := c.orm.Delete(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (c *Repository) GetByID(clienteID uint32) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		ID: clienteID,
	}
	result := c.orm.First(&cliente)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cliente, nil
}

func (c *Repository) GetAll() (clientes []domain.Cliente, err error) {
	result := c.orm.Find(&clientes)
	if result.Error != nil {
		log.Println(result.Error)
		return clientes, result.Error
	}

	return clientes, err
}
