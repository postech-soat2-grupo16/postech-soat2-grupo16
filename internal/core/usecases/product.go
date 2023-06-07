package usecases

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"gorm.io/gorm"
)

var _ ports.ProductUseCase = (*ProductUseCase)(nil)

func NewProductUseCase(productRepo *gorm.DB) ProductUseCase {
	return ProductUseCase{
		productRepo: productRepo,
	}
}

type ProductUseCase struct {
	productRepo *gorm.DB
}

func (p ProductUseCase) List() (products []domain.Product, err error) {
	result := p.productRepo.Find(&products)
	if result.Error != nil {
		log.Fatal(result.Error)
		return products, result.Error
	}

	return products, err
}

func (p ProductUseCase) GetById(productID uint32) (*domain.Product, error) {
	product := domain.Product{
		ID: productID,
	}
	result := p.productRepo.First(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &product, nil
}

func (p ProductUseCase) Create(name, category, description string) (*domain.Product, error) {
	product := domain.Product{
		Name:        name,
		Category:    category,
		Description: description,
	}
	result := p.productRepo.Create(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (p ProductUseCase) Update(productID uint32, name, category, description string) (*domain.Product, error) {
	product := domain.Product{
		ID:          productID,
		Name:        name,
		Category:    category,
		Description: description,
	}
	result := p.productRepo.Updates(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (p ProductUseCase) Delete(productID uint32) error {
	product := domain.Product{
		ID: productID,
	}
	result := p.productRepo.Delete(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil
}
