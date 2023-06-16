package cliente

import (
	"errors"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
	"log"
)

type ClienteUseCase struct {
	clienteRepo *gorm.DB
}

func NewClienteUseCase(clienteRepo *gorm.DB) *ClienteUseCase {
	return &ClienteUseCase{
		clienteRepo: clienteRepo,
	}
}

func (p *ClienteUseCase) List() (clientes []domain.Cliente, err error) {
	result := p.clienteRepo.Find(&clientes)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return clientes, err
}

func (p *ClienteUseCase) GetByID(clienteID uint32) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		ID: clienteID,
	}
	result := p.clienteRepo.First(&cliente)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &cliente, nil
}

func (p *ClienteUseCase) Create(email, cpf, nome string) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result := p.clienteRepo.Create(&cliente)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &cliente, nil
}

func (p *ClienteUseCase) Update(clienteID uint32, email, cpf, nome string) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		ID:    clienteID,
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result := p.clienteRepo.Updates(&cliente)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &cliente, nil
}

func (p *ClienteUseCase) Delete(clienteID uint32) error {
	cliente := domain.Cliente{
		ID: clienteID,
	}

	result := p.clienteRepo.Delete(&cliente)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil
}
