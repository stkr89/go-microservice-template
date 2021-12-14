package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDB() gorm.DB {
	dbURL := os.ExpandEnv("postgresql://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?&options=--cluster%3D$DB_CLUSTER")
	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		panic(err)
	}

	return *db
}

func InitialDBMigration(db gorm.DB) error {
	var m []interface{}

	for _, m := range m {
		err := db.AutoMigrate(m)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
