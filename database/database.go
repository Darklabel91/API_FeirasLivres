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
		log.Println("ERROR: connection error")
	}
}

//LoadCSV used for migrate the csv file to the database
//	Recommend to ignore the error because the database should only be imported the first time the project is run
func LoadCSV(path string) {
	_ = migrateCSV(path)
}
