package database

import (
	"log"

	"gorm.io/gorm"
)

func NewORM(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector)
	if err != nil {
		log.Println(err)
	}
	return db
}
