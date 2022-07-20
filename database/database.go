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
	//because of the mass import made with the csv is nice to assure that the id sequence is ok
	DB.Exec("SELECT setval('fairs_id_seq', (SELECT MAX(id) FROM fairs));")
}

//LoadCSV used for migrate the csv file to the database
//	Recommend to ignore the error because the database should only be imported the first time the project is run
func LoadCSV(path string) {
	_ = migrateCSV(path)
}
