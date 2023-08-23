package item

import (
	"errors"
	"log"
	"strings"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"github.com/joaocampari/postech-soat2-grupo16/util"

	"gorm.io/gorm"
)

type UseCase struct {
	itemGateway interfaces.ItemGatewayI
}

func NewUseCase(itemGateway interfaces.ItemGatewayI) UseCase {
	return UseCase{
		itemGateway: itemGateway,
	}
}

func (p UseCase) List() ([]entities.Item, error) {
	items, err := p.itemGateway.GetAll()
	if err != nil {
		log.Println(err)
		return items, err
	}

	return items, err
}

func (p UseCase) GetByID(itemID uint32) (*entities.Item, error) {
	result, err := p.itemGateway.GetByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p UseCase) GetByCategory(category string) ([]entities.Item, error) {
	item := entities.Item{
		Category: strings.ToUpper(category),
	}

	if !item.IsCategoryValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	result, err := p.itemGateway.GetByCategory(item.Category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.Item{}, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p UseCase) Create(name, category, description string, price float32) (*entities.Item, error) {
	item := entities.Item{
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

	result, err := p.itemGateway.Save(item)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Update(itemID uint32, name, category, description string, price float32) (*entities.Item, error) {
	item, err := p.itemGateway.GetByID(itemID)
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

	result, err := p.itemGateway.Update(updatedItem)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Delete(itemID uint32) (*entities.Item, error) {
	item, err := p.itemGateway.GetByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	err = p.itemGateway.Delete(item.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return item, nil
}
