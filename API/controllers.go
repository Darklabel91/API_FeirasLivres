package API

import (
	"encoding/json"
	"github.com/Darklabel91/API_FeirasLivres/database"
	"github.com/Darklabel91/API_FeirasLivres/models"
	"log"
	"net/http"
)

//GetFairs json encode all fairs
func GetFairs(w http.ResponseWriter, r *http.Request) {
	var fairs []models.Fair
	database.DB.Find(&fairs)

	err := json.NewEncoder(w).Encode(fairs)
	if err != nil {
		log.Println(err)
	}

	log.Println("get request to all fairs")
}
