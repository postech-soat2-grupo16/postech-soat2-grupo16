package domain

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          uint32  `gorm:"primary_key;auto_increment"`
	Name        string  `gorm:"size:255;not null;"`
	Category    string  `gorm:"size:100;not null;"`
	Description string  `gorm:"size:255;not null;"`
	Price       float32 `gorm:"type: numeric"`
}

func (i *Item) IsCategoryValid() bool {
	return i.Category == ""
}
