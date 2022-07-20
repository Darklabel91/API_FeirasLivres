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

//CreateFair create new models.Fair on database
func CreateFair(w http.ResponseWriter, r *http.Request) {
	var newFair models.Fair
	err := json.NewDecoder(r.Body).Decode(&newFair)
	if err != nil {
		log.Println("ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on POST request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = models.ValidateFair(&newFair)
	if err != nil {
		log.Println("ERROR: JSON file error, all fields must have a value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	database.DB.Create(&newFair)
	err = json.NewEncoder(w).Encode(newFair)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on POST request")
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println("INFO: post request http://localhost:8000/api/fairs new fair created")
	w.WriteHeader(http.StatusCreated)
}

//DeleteFair  delete models.Fair by given id
func DeleteFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var fair models.Fair
	database.DB.Delete(&fair, id)

	err := json.NewEncoder(w).Encode(fair)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on DELETE request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("INFO: fair with id:" + id + " DELETED")
	w.WriteHeader(http.StatusAccepted)
}

//UpdateFair update models.Fair by given id
func UpdateFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var fair models.Fair
	database.DB.Select("*").Where("id = ?", id).Table("fairs").Find(&fair)

	err := json.NewDecoder(r.Body).Decode(&fair)
	if err != nil {
		log.Println("ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on PUT request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = models.ValidateFair(&fair)
	if err != nil {
		log.Println("ERROR: JSON file error, all fields must have a value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Probably a way to do it simpler
	sameID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("ERROR: could not convert string to int in: http://localhost:8000/api/fairs on PUT request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	updateFair := models.Fair{Id: sameID, Longitude: fair.Longitude, Latitude: fair.Latitude, SetCen: fair.SetCen, AreaP: fair.AreaP, CodDist: fair.CodDist, District: fair.District, CodSubPref: fair.CodSubPref, SubPref: fair.SubPref, RegionFive: fair.RegionFive, RegionEight: fair.RegionEight, NameFair: fair.NameFair, Record: fair.Record, Street: fair.Street, Number: fair.Number, Neighbourhood: fair.Neighbourhood, Reference: fair.Reference}

	database.DB.Save(&updateFair)
	err = json.NewEncoder(w).Encode(updateFair)
	if err != nil {
		log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on PUT request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("INFO: put request http://localhost:8000/api/fairs fair updated with ID:" + id)
	w.WriteHeader(http.StatusAccepted)
}

//GetFair read and find models.Fair using SearchFair
func GetFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchType := vars["searchType"]
	searchParam := vars["searchParam"]

	switch searchType {
	case "district":
		_ = SearchFair(w, searchParam, searchType, "district = ?")
	case "region":
		_ = SearchFair(w, searchParam, searchType, "region_five = ?")
	case "name":
		_ = SearchFair(w, searchParam, searchType, "name_fair = ?")
	case "neighbourhood":
		_ = SearchFair(w, searchParam, searchType, "neighbourhood = ?")
	case "id":
		_ = SearchFair(w, searchParam, searchType, "id = ?")
	default:
		log.Println("WARNING: searchType: " + searchType + " unexpected. searchType must be [district,region,name or neighbourhood]	http://localhost:8000/api/fairs/" + searchType + "/" + searchParam)
		w.WriteHeader(http.StatusBadRequest)
	}
}

//SearchFair simple function for json return using searchType with a searchParam
func SearchFair(w http.ResponseWriter, searchParam, searchType, query string) error {
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

//GetFairs read and find all models.Fair in database
func GetFairs(w http.ResponseWriter, r *http.Request) {
	var fairs []models.Fair
	var totalFairs int64
	database.DB.Find(&fairs).Count(&totalFairs)

	if totalFairs < 1 {
		log.Println("WARNING: get request to http://localhost:8000/api/fairs	have return 0")
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		err := json.NewEncoder(w).Encode(fairs)
		if err != nil {
			log.Println("ERROR: could not encode JSON file in: http://localhost:8000/api/fairs")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println("INFO: get request to http://localhost:8000/api/fairs" + "	fairs found:" + strconv.Itoa(int(totalFairs)))
	}
}
