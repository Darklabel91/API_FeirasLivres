package API

import (
	"encoding/json"
	"github.com/Darklabel91/API_FeirasLivres/database"
	"github.com/Darklabel91/API_FeirasLivres/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//CreateFair create a new Fair on database
func CreateFair(w http.ResponseWriter, r *http.Request) {
	var newFair models.Fair
	err := json.NewDecoder(r.Body).Decode(&newFair)
	if err != nil {
		log.Println(err)
	}

	//because of the mass import made with the csv is nice to assure that the id sequence is ok
	database.DB.Exec("SELECT setval('fairs_id_seq', (SELECT MAX(id) FROM fairs));")

	database.DB.Create(&newFair)
	err = json.NewEncoder(w).Encode(newFair)
	if err != nil {
		log.Println(err)
	}

	log.Println("create new fair")
}

//DeleteFair  deletes a fair by given id
func DeleteFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var personality models.Fair
	database.DB.Delete(&personality, id)
	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		log.Println(err)
	}

	log.Println("delete fair id " + id)
}

//UpdateFair edit fair by given id
func UpdateFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var personality models.Fair
	database.DB.First(&personality, id)
	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		log.Println(err)
	}

	database.DB.Save(&personality)
	err = json.NewEncoder(w).Encode(personality)
	if err != nil {
		log.Println(err)
	}

	log.Println("update fair id " + id)
}

//GetFair read and find fair
func GetFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchType := vars["searchType"]
	searchParam := vars["searchParam"]

	var err error
	switch searchType {
	case "district":
		err = searchFair(w, searchParam, "district = ?")
		if err != nil {
			log.Println(err)
		}
		log.Println("get request to \"/api/fairs/" + searchType + "/" + searchParam + "\"")
	case "region":
		err = searchFair(w, searchParam, "region_five = ?")
		if err != nil {
			log.Println(err)
		}
		log.Println("get request to \"/api/fairs/" + searchType + "/" + searchParam + "\"")
	case "name":
		err = searchFair(w, searchParam, "name_fair = ?")
		if err != nil {
			log.Println(err)
		}
		log.Println("get request to \"/api/fairs/" + searchType + "/" + searchParam + "\"")
	case "neighbourhood":
		err = searchFair(w, searchParam, "neighbourhood = ?")
		if err != nil {
			log.Println(err)
		}
		log.Println("get request to \"/api/fairs/" + searchType + "/" + searchParam + "\"")
	case "id":
		err = searchFair(w, searchParam, "id = ?")
		if err != nil {
			log.Println(err)
		}
		log.Println("get request to \"/api/fairs/" + searchType + "/" + searchParam + "\"")
	default:
		log.Println("searchType " + searchType + " unexpected. searchType must be [district,region,name or neighbourhood]")
	}

	log.Println("get request to  http://localhost:8000/api/fairs/" + searchType + "/" + searchParam)

}

//searchFair simple function for json return
func searchFair(w http.ResponseWriter, searchParam string, query string) error {
	var foundFairs []models.Fair
	var totalFairs int64
	database.DB.Select("*").Where(query, searchParam).Table("fairs").Find(&foundFairs).Count(&totalFairs)
	if totalFairs != 0 {
		err := json.NewEncoder(w).Encode(foundFairs)
		if err != nil {
			return err
		}
	}
	return nil
}

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
