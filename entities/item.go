package entities

import (
	"strings"

	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

const (
	bebida         = "BEBIDA"
	lanche         = "LANCHE"
	sobremesa      = "SOBREMESA"
	acompanhamento = "ACOMPANHAMENTO"
)

type Item struct {
	ID          uint32  `gorm:"primary_key;auto_increment"`
	Name        string  `gorm:"size:255;not null;"`
	Category    string  `gorm:"size:100;not null;"`
	Description string  `gorm:"size:255;not null;"`
	Price       float32 `gorm:"type: numeric"`
	gorm.Model
}

func (i *Item) IsCategoryValid() bool {
	categories := []string{bebida, lanche, sobremesa, acompanhamento}
	return slices.Contains(categories, i.Category)
}

func (i *Item) IsNameNull() bool {
	return len(strings.TrimSpace(i.Name)) == 0
}

func (i *Item) IsPriceValid() bool {
	return i.Price >= 0
}

func (i *Item) CopyItemWithNewValues(name, category, description string, price float32) Item {
	return Item{
		ID:          i.ID,
		Name:        name,
		Category:    category,
		Description: description,
		Price:       price,
	}
}
