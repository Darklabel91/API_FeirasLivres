package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

//Connect database connection
func Connect() {
	conStr := "host=localhost user=postgres password=postgres dbname=feira port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conStr))
	if err != nil {
		log.Panic("connection error")
	}
}

//LoadCSV used for migrate the csv file to the database
func LoadCSV(path string) error {
	err := migrateCSV(path)
	if err != nil {
		return err
	}
	return nil
}
