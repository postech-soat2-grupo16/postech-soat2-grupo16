package item

import (
	"log"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

type Repository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) *Repository {
	return &Repository{orm: orm}
}

func (c *Repository) Save(item domain.Item) (*domain.Item, error) {
	result := c.orm.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (c *Repository) Update(item domain.Item) (*domain.Item, error) {
	result := c.orm.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (c *Repository) Delete(itemID uint32) error {
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

func (c *Repository) GetByID(itemID uint32) (*domain.Item, error) {
	item := domain.Item{
		ID: itemID,
	}
	result := c.orm.First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (c *Repository) GetByCategory(category string) ([]domain.Item, error) {
	item := domain.Item{
		Category: category,
	}
	var items []domain.Item
	result := c.orm.Find(&items, item)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (c *Repository) GetAll() (items []domain.Item, err error) {
	result := c.orm.Find(&items)
	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}
