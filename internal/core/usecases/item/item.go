package item

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

func NewItemUseCase(itemRepo *gorm.DB) ItemUseCase {
	return ItemUseCase{
		itemRepo: itemRepo,
	}
}

type ItemUseCase struct {
	itemRepo *gorm.DB
}

func (p ItemUseCase) List() (items []domain.Item, err error) {
	result := p.itemRepo.Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
		return items, result.Error
	}

	return items, err
}

func (p ItemUseCase) GetById(itemID uint32) (*domain.Item, error) {
	item := domain.Item{
		ID: itemID,
	}
	result := p.itemRepo.First(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (p ItemUseCase) Create(name, category, description string, price float32) (*domain.Item, error) {
	item := domain.Item{
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
	}
	result := p.itemRepo.Create(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (p ItemUseCase) Update(itemID uint32, name, category, description string, price float32) (*domain.Item, error) {
	item := domain.Item{
		ID:          itemID,
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
	}
	result := p.itemRepo.Updates(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (p ItemUseCase) Delete(itemID uint32) error {
	item := domain.Item{
		ID: itemID,
	}
	result := p.itemRepo.Delete(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil
}
