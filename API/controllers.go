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
		msgErr := "ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on POST request"
		jsonErrMsg(w, msgErr)
		return
	}

	err = models.ValidateFair(&newFair)
	if err != nil {
		msgErr := "ERROR: JSON file error, all fields must have a value"
		jsonErrMsg(w, msgErr)
		return
	}

	createFair := models.Fair{Longitude: newFair.Longitude, Latitude: newFair.Latitude, SetCen: newFair.SetCen, AreaP: newFair.AreaP, CodDist: newFair.CodDist, District: newFair.District, CodSubPref: newFair.CodSubPref, SubPref: newFair.SubPref, RegionFive: newFair.RegionFive, RegionEight: newFair.RegionEight, NameFair: newFair.NameFair, Record: newFair.Record, Street: newFair.Street, Number: newFair.Number, Neighbourhood: newFair.Neighbourhood, Reference: newFair.Reference}
	database.DB.Create(&createFair)

	err = json.NewEncoder(w).Encode(createFair)
	if err != nil {
		msgErr := "ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on POST request"
		jsonErrMsg(w, msgErr)
		return
	}

	log.Println("INFO: post request http://localhost:8000/api/fairs new fair created")
}

//DeleteFair  delete models.Fair by given id
func DeleteFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var totalFairs int64
	var fair models.Fair
	database.DB.Select("*").Where("id = ?", id).Table("fairs").Find(&fair).Count(&totalFairs)
	if totalFairs < 1 {
		msgErr := "ERROR: impossible to delete a fair that does not exist!"
		jsonErrMsg(w, msgErr)
		return
	}

	database.DB.Delete(&fair, id)

	msg := "INFO: fair with id:" + id + " DELETED"
	log.Println(msg)
	_ = json.NewEncoder(w).Encode(msg)

	err := json.NewEncoder(w).Encode(fair)
	if err != nil {
		msgErr := "ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on DELETE request"
		jsonErrMsg(w, msgErr)
		return
	}
}

//UpdateFair update models.Fair by given id
func UpdateFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["searchParam"]

	var totalFairs int64
	var fair models.Fair
	database.DB.Select("*").Where("id = ?", id).Table("fairs").Find(&fair).Count(&totalFairs)
	if totalFairs < 1 {
		msgErr := "ERROR: impossible to update a fair that does not exist!"
		jsonErrMsg(w, msgErr)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&fair)
	if err != nil {
		msgErr := "ERROR: could not decode JSON file in: http://localhost:8000/api/fairs on PUT request"
		jsonErrMsg(w, msgErr)
		return
	}

	err = models.ValidateFair(&fair)
	if err != nil {
		msgErr := "ERROR: JSON file error, all fields must have a value"
		jsonErrMsg(w, msgErr)
		return
	}

	//Probably a way to do it simpler
	sameID, err := strconv.Atoi(id)
	if err != nil {
		msgErr := "ERROR: could not convert string to int in: http://localhost:8000/api/fairs on PUT request"
		jsonErrMsg(w, msgErr)
		return
	}

	updateFair := models.Fair{Id: sameID, Longitude: fair.Longitude, Latitude: fair.Latitude, SetCen: fair.SetCen, AreaP: fair.AreaP, CodDist: fair.CodDist, District: fair.District, CodSubPref: fair.CodSubPref, SubPref: fair.SubPref, RegionFive: fair.RegionFive, RegionEight: fair.RegionEight, NameFair: fair.NameFair, Record: fair.Record, Street: fair.Street, Number: fair.Number, Neighbourhood: fair.Neighbourhood, Reference: fair.Reference}
	database.DB.Save(&updateFair)

	err = json.NewEncoder(w).Encode(updateFair)
	if err != nil {
		msgErr := "ERROR: could not encode JSON file in: http://localhost:8000/api/fairs on PUT request"
		jsonErrMsg(w, msgErr)
		return
	}

	log.Println("INFO: put request http://localhost:8000/api/fairs fair updated with ID:" + id)
}

//GetFair read and find models.Fair using SearchFair
func GetFair(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchType := vars["searchType"]
	searchParam := vars["searchParam"]

	switch searchType {
	case "district":
		err := SearchFair(w, searchParam, searchType, "district = ?")
		if err != nil {
			return
		}
	case "region":
		err := SearchFair(w, searchParam, searchType, "region_five = ?")
		if err != nil {
			return
		}
	case "name":
		err := SearchFair(w, searchParam, searchType, "name_fair = ?")
		if err != nil {
			return
		}
	case "neighbourhood":
		err := SearchFair(w, searchParam, searchType, "neighbourhood = ?")
		if err != nil {
			return
		}
	case "id":
		err := SearchFair(w, searchParam, searchType, "id = ?")
		if err != nil {
			return
		}
	default:
		msgErr := "WARNING: searchType: " + searchType + " unexpected. searchType must be [district,region,name,neighbourhood or id]"
		jsonErrMsg(w, msgErr)
		return
	}
}

//SearchFair simple function for json return using searchType with a searchParam
func SearchFair(w http.ResponseWriter, searchParam, searchType, query string) error {
	var foundFairs []models.Fair
	var totalFairs int64

	database.DB.Select("*").Where(query, searchParam).Table("fairs").Find(&foundFairs).Count(&totalFairs)
	if totalFairs < 1 {
		msgErr := "WARNING: http://localhost:8000/api/fairs/" + searchType + "/" + searchParam + " have return 0"
		jsonErrMsg(w, msgErr)
		return errors.New(searchParam + " not found")
	}

	err := json.NewEncoder(w).Encode(foundFairs)
	if err != nil {
		msgErr := "ERROR: could not encode JSON file in: http://localhost:8000/api/fairs/" + searchType + "/" + searchParam
		jsonErrMsg(w, msgErr)
		return err
	}

	log.Println("INFO: get request to http://localhost:8000/api/fairs/" + searchType + "/" + searchParam)
	return nil
}

//GetFairs read and find all models.Fair in database
func GetFairs(w http.ResponseWriter, r *http.Request) {
	var fairs []models.Fair
	var totalFairs int64
	database.DB.Find(&fairs).Count(&totalFairs)

	if totalFairs < 1 {
		msgErr := "WARNING: get request to http://localhost:8000/api/fairs	have return 0"
		jsonErrMsg(w, msgErr)
		return
	} else {
		err := json.NewEncoder(w).Encode(fairs)
		if err != nil {
			msgErr := "ERROR: could not encode JSON file in: http://localhost:8000/api/fairs"
			jsonErrMsg(w, msgErr)
			return
		}

		log.Println("INFO: get request to http://localhost:8000/api/fairs" + "	fairs found:" + strconv.Itoa(int(totalFairs)))
	}
}

func jsonErrMsg(w http.ResponseWriter, msgErr string) {
	log.Println(msgErr)
	_ = json.NewEncoder(w).Encode(msgErr)
}
