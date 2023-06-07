package ports

import "github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"

type ProductUseCase interface {
	List() ([]domain.Product, error)
	Create(name, category, description string) (*domain.Product, error)
}
