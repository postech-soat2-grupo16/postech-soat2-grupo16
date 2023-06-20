package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresDialector() gorm.Dialector {
	connStr := os.Getenv("DATABASE_URL")
	pgDialector := postgres.Open(connStr)
	return pgDialector
}
