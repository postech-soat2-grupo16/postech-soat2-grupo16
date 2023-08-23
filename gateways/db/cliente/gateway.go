package cliente

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"gorm.io/gorm"
)

type Gateway struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Gateway {
	return &Gateway{repository: repository}
}

func (g *Gateway) Save(cliente entities.Cliente) (*entities.Cliente, error) {
	result := g.repository.Create(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &cliente, nil
}

func (g *Gateway) Update(cliente entities.Cliente) (*entities.Cliente, error) {
	result := g.repository.Updates(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &cliente, nil
}

func (g *Gateway) Delete(clienteID uint32) error {
	cliente := entities.Cliente{
		ID: clienteID,
	}
	result := g.repository.Delete(&cliente)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (g *Gateway) GetByID(clienteID uint32) (*entities.Cliente, error) {
	cliente := entities.Cliente{
		ID: clienteID,
	}
	result := g.repository.First(&cliente)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cliente, nil
}

func (g *Gateway) GetAll(conds ...interface{}) (clientes []entities.Cliente, err error) {
	result := g.repository.Find(&clientes, conds...)
	if result.Error != nil {
		log.Println(result.Error)
		return clientes, result.Error
	}

	return clientes, err
}
