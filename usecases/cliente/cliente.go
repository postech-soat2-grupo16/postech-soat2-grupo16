package cliente

import (
	"errors"
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"gorm.io/gorm"
)

type UseCase struct {
	clienteGateway interfaces.ClienteGatewayI
}

func NewUseCase(clienteGateway interfaces.ClienteGatewayI) *UseCase {
	return &UseCase{
		clienteGateway: clienteGateway,
	}
}

func (p *UseCase) List(CPF string) ([]entities.Cliente, error) {
	if CPF != "" {
		client := entities.Cliente{CPF: CPF}
		return p.clienteGateway.GetAll(client)
	}
	return p.clienteGateway.GetAll()
}

func (p *UseCase) GetByID(clienteID uint32) (*entities.Cliente, error) {
	result, err := p.clienteGateway.GetByID(clienteID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Create(email, cpf, nome string) (*entities.Cliente, error) {
	cliente := entities.Cliente{
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result, err := p.clienteGateway.Save(cliente)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Update(clienteID uint32, email, cpf, nome string) (*entities.Cliente, error) {
	cliente := entities.Cliente{
		ID:    clienteID,
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result, err := p.clienteGateway.Update(cliente)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Delete(clienteID uint32) error {
	err := p.clienteGateway.Delete(clienteID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
