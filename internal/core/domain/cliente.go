package domain

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	ID    uint32 `gorm:"primary_key;auto_increment"`
	Email string `gorm:"null;"`
	CPF   string `gorm:"null;"`
	Name  string `gorm:"not null;"`
}
