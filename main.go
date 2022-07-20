package main

import (
	"github.com/Darklabel91/API_FeirasLivres/API"
	"github.com/Darklabel91/API_FeirasLivres/database"
	"github.com/Darklabel91/API_FeirasLivres/writeLog"
	"log"
)

const path = "database/migration/DEINFO_AB_FEIRASLIVRES_2014.csv"

//main
//	The project starts setting up the log  writeLog.WriteLog
//	Connect to  database.Connect
//	Load the .csv with database.LoadCSV
//	Start API with API.HandleRequest
func main() {
	writeLog.WriteLog()
	log.Println("INFO: Start API")

	database.Connect()
	database.LoadCSV(path)

	API.HandleRequest()
}
