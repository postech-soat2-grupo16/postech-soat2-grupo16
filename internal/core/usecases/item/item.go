package item

import (
	"errors"
	"github.com/joaocampari/postech-soat2-grupo16/internal/util"
	"log"
	"strings"

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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (p ItemUseCase) Create(name, category, description string, price float32) (*domain.Item, error) {
	item := domain.Item{
		Name:        name,
		Category:    strings.ToUpper(category),
		Description: description,
		Price:       price,
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
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
		Category:    strings.ToUpper(category),
		Description: description,
		Price:       price,
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
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
