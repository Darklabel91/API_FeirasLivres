package main

import (
	"github.com/Darklabel91/API_FeirasLivres/API"
	"github.com/Darklabel91/API_FeirasLivres/database"
)

const path = "database/migration/DEINFO_AB_FEIRASLIVRES_2014.csv"

func main() {
	//Load Database
	database.Connect()

	//Script for .csv file
	_ = database.LoadCSV(path)

	//API - Handle Requests
	API.HandleRequest()
}
