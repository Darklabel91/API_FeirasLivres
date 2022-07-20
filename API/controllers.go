package API

import (
	"encoding/json"
	"errors"
	"github.com/Darklabel91/API_FeirasLivres/database"
	"github.com/Darklabel91/API_FeirasLivres/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//CreateFair create a new Fair on database
func CreateFair(w http.ResponseWriter, r *http.Request) {
	var newFair models.Fair
	err := json.NewDecoder(r.Body).Decode(&newFair)
	if err != nil {
		log.Println("ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on POST request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	//because of the mass import made with the csv is nice to assure that the id sequence is ok
	database.DB.Exec("SELECT setval('fairs_id_seq', (SELECT MAX(id) FROM fairs));")
	database.DB.Create(&newFair)

	err = json.NewEncoder(w).Encode(newFair)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on POST request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("INFO: post request http://localhost:8000/api/fairs new fair created")
	w.WriteHeader(http.StatusCreated)
}

//DeleteFair  deletes a fair by given id
func DeleteFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var personality models.Fair
	database.DB.Delete(&personality, id)

	err := json.NewEncoder(w).Encode(personality)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on DELETE request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("INFO: fair with id:" + id + " DELETED")
	w.WriteHeader(http.StatusAccepted)
}

//UpdateFair edit fair by given id
func UpdateFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var personality models.Fair
	database.DB.First(&personality, id)

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		log.Println("ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on PUT request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	database.DB.Save(&personality)
	err = json.NewEncoder(w).Encode(personality)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on PUT request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("INFO: put request http://localhost:8000/api/fairs fair updated with ID:" + strconv.Itoa(personality.Id))
	w.WriteHeader(http.StatusAccepted)
}

//GetFair read and find fair
func GetFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchType := vars["searchType"]
	searchParam := vars["searchParam"]

	switch searchType {
	case "district":
		_ = searchFair(w, searchParam, searchType, "district = ?")
	case "region":
		_ = searchFair(w, searchParam, searchType, "region_five = ?")
	case "name":
		_ = searchFair(w, searchParam, searchType, "name_fair = ?")
	case "neighbourhood":
		_ = searchFair(w, searchParam, searchType, "neighbourhood = ?")
	case "id":
		_ = searchFair(w, searchParam, searchType, "id = ?")
	default:
		log.Println("WARNING: searchType: " + searchType + " unexpected. searchType must be [district,region,name or neighbourhood]	http://localhost:8000/api/fairs/" + searchType + "/" + searchParam)
		w.WriteHeader(http.StatusBadRequest)
	}
}

//searchFair simple function for json return
func searchFair(w http.ResponseWriter, searchParam, searchType, query string) error {
	var foundFairs []models.Fair
	var totalFairs int64

	database.DB.Select("*").Where(query, searchParam).Table("fairs").Find(&foundFairs).Count(&totalFairs)
	if totalFairs < 1 {
		log.Println("WARNING: http://localhost:8000/api/fairs/" + searchType + "/" + searchParam + " have return 0")
		w.WriteHeader(http.StatusNoContent)
		return errors.New(searchParam + " not found")
	}

	err := json.NewEncoder(w).Encode(foundFairs)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs/" + searchType + "/" + searchParam)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	log.Println("INFO: get request to http://localhost:8000/api/fairs" + searchType + "/" + searchParam)
	return nil
}

//GetFairs json encode all fairs
func GetFairs(w http.ResponseWriter, r *http.Request) {
	var fairs []models.Fair
	var totalFairs int64
	database.DB.Find(&fairs).Count(&totalFairs)

	if totalFairs < 1 {
		log.Println("WARNING: get request to http://localhost:8000/api/fairs	have return 0")
		w.WriteHeader(http.StatusNoContent)
	} else {
		err := json.NewEncoder(w).Encode(fairs)
		if err != nil {
			log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs")
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Println("INFO: get request to http://localhost:8000/api/fairs" + "	fairs found:" + strconv.Itoa(int(totalFairs)))
	}
}
