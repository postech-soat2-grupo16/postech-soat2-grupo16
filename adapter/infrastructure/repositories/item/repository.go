package item

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

type ItemRepository struct {
	orm *gorm.DB
}

func NewItemRepository(orm *gorm.DB) *ItemRepository {
	return &ItemRepository{orm: orm}
}

func (c *ItemRepository) Save(item domain.Item) (*domain.Item, error) {
	result := c.orm.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (c *ItemRepository) Update(item domain.Item) (*domain.Item, error) {
	result := c.orm.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (c *ItemRepository) Delete(itemID uint32) error {
	item := domain.Item{
		ID: itemID,
	}
	result := c.orm.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (c *ItemRepository) GetByID(itemID uint32) (*domain.Item, error) {
	item := domain.Item{
		ID: itemID,
	}
	result := c.orm.First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (c *ItemRepository) GetByCategory(category string) ([]domain.Item, error) {
	item := domain.Item{
		Category: category,
	}
	var items []domain.Item
	result := c.orm.First(&items, item)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (c *ItemRepository) GetAll() (items []domain.Item, err error) {
	result := c.orm.Find(&items)
	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}
