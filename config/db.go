package config

import (
	"fmt"
	"github.com/shopr-org/grpc-service-template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=UTC",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return *db
}

func InitialDBMigration(db gorm.DB) error {
	m := []interface{}{
		&models.User{},
	}

	for _, m := range m {
		err := db.AutoMigrate(m)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
