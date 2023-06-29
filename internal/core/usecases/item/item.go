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

func NewUseCase(itemRepo ports.ItemRepository) UseCase {
	return UseCase{
		itemRepo: itemRepo,
	}
}

type UseCase struct {
	itemRepo ports.ItemRepository
}

func (p UseCase) List() ([]domain.Item, error) {
	items, err := p.itemRepo.GetAll()
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, err
}

func (p UseCase) GetByID(itemID uint32) (*domain.Item, error) {
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

func (p UseCase) GetByCategory(category string) ([]domain.Item, error) {
	item := domain.Item{
		Category: strings.ToUpper(category),
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	result, err := p.itemRepo.GetByCategory(item.Category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []domain.Item{}, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p UseCase) Create(name, category, description string, price float32) (*domain.Item, error) {
	item := domain.Item{
		Name:        name,
		Category:    strings.ToUpper(category),
		Description: description,
		Price:       price,
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	if item.IsNameNull() {
		return nil, util.NewErrorDomain("Nome vazio")
	}

	if !item.IsPriceValid() {
		return nil, util.NewErrorDomain("Preço negativo")
	}

	result, err := p.itemRepo.Save(item)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Update(itemID uint32, name, category, description string, price float32) (*domain.Item, error) {
	item, err := p.itemRepo.GetByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	updatedItem := item.CopyItemWithNewValues(name, strings.ToUpper(category), description, price)

	if !updatedItem.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	if updatedItem.IsNameNull() {
		return nil, util.NewErrorDomain("Nome vazio")
	}

	if !updatedItem.IsPriceValid() {
		return nil, util.NewErrorDomain("Preço inválido")
	}

	result, err := p.itemRepo.Update(updatedItem)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Delete(itemID uint32) (*domain.Item, error) {
	item, err := p.itemRepo.GetByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	err = p.itemRepo.Delete(item.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return item, nil
}
