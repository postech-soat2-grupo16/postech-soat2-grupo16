package item

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"gorm.io/gorm"
)

type Gateway struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Gateway {
	return &Gateway{repository: repository}
}

func (g *Gateway) Save(item entities.Item) (*entities.Item, error) {
	result := g.repository.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (g *Gateway) Update(item entities.Item) (*entities.Item, error) {
	result := g.repository.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (g *Gateway) Delete(itemID uint32) error {
	item := entities.Item{
		ID: itemID,
	}
	result := g.repository.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (g *Gateway) GetByID(itemID uint32) (*entities.Item, error) {
	item := entities.Item{
		ID: itemID,
	}
	result := g.repository.First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (g *Gateway) GetByCategory(category string) ([]entities.Item, error) {
	item := entities.Item{
		Category: category,
	}
	var items []entities.Item
	result := g.repository.Find(&items, item)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (g *Gateway) GetAll() (items []entities.Item, err error) {
	result := g.repository.Find(&items)
	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}
