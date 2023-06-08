package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDialector() gorm.Dialector {
	pgDialector := postgres.Open(os.Getenv("DATABASE_URL"))
	return pgDialector
}
