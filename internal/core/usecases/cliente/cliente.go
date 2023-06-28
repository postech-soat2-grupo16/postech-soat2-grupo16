package cliente

import (
	"errors"
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
)

type UseCase struct {
	clienteRepo ports.ClienteRepository
}

func NewUseCase(clienteRepo ports.ClienteRepository) *UseCase {
	return &UseCase{
		clienteRepo: clienteRepo,
	}
}

func (p *UseCase) List() ([]domain.Cliente, error) {
	result, err := p.clienteRepo.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, err
}

func (p *UseCase) GetByID(clienteID uint32) (*domain.Cliente, error) {
	result, err := p.clienteRepo.GetByID(clienteID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Create(email, cpf, nome string) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result, err := p.clienteRepo.Save(cliente)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Update(clienteID uint32, email, cpf, nome string) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		ID:    clienteID,
		Email: email,
		CPF:   cpf,
		Name:  nome,
	}

	result, err := p.clienteRepo.Update(cliente)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (p *UseCase) Delete(clienteID uint32) error {
	err := p.clienteRepo.Delete(clienteID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
