package item

import (
	"errors"
	"log"
	"strings"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"github.com/joaocampari/postech-soat2-grupo16/internal/util"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

func NewItemUseCase(itemRepo ports.ItemRepository) ItemUseCase {
	return ItemUseCase{
		itemRepo: itemRepo,
	}
}

type ItemUseCase struct {
	itemRepo ports.ItemRepository
}

func (p ItemUseCase) List() ([]domain.Item, error) {
	items, err := p.itemRepo.GetAll()
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, err
}

func (p ItemUseCase) GetByID(itemID uint32) (*domain.Item, error) {
	result, err := p.itemRepo.GetByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p ItemUseCase) GetByCategory(category string) ([]domain.Item, error) {
	item := domain.Item{
		Category: category,
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	result, err := p.itemRepo.GetByCategory(category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domain.Item{}, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
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

	result, err := p.itemRepo.Save(item)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
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

	result, err := p.itemRepo.Update(item)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p ItemUseCase) Delete(itemID uint32) error {
	err := p.itemRepo.Delete(itemID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
