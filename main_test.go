package main

import (
	"bytes"
	"encoding/json"
	"github.com/Darklabel91/API_FeirasLivres/API"
	"github.com/Darklabel91/API_FeirasLivres/database"
	"github.com/Darklabel91/API_FeirasLivres/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//Test for "/api/fairs/district/{searchParam}" GET
func TestGetFairByDistrictName(t *testing.T) {
	err := searchFair("district", "mockDistrict")
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/fairs/region/{searchParam}" GET
func TestGetFairByRegion(t *testing.T) {
	err := searchFair("region", "mockRegion5")
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/fairs/name/{searchParam}" GET
func TestGetFairByNameFair(t *testing.T) {
	err := searchFair("name", "mockNameFair")
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/fairs/neighbourhood/{searchParam}" GET
func TestGetFairByNeighbourhood(t *testing.T) {
	err := searchFair("neighbourhood", "mockNeighbourhood")
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/fairs/id/{searchParam}" GET
func TestGetFairByID(t *testing.T) {
	err := searchFair("id", strconv.Itoa(ID))
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/fairs" GET
func TestGetAllFairs(t *testing.T) {
	database.Connect()
	createFairMock()
	defer deleteFairMock()

	r := routsTestSetup()
	r.HandleFunc("/api/fairs", API.GetFairs).Methods("Get")
	req, _ := http.NewRequest("GET", "/api/fairs", nil)
	response := executeRequest(req, r)

	err := checkResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/id/{searchParam}" DELETE
func TestDeleteFairByID(t *testing.T) {
	database.Connect()
	createFairMock()

	r := routsTestSetup()
	r.HandleFunc("/api/fairs/id/{searchParam}", API.DeleteFair).Methods("Delete")
	req, _ := http.NewRequest("DELETE", "/api/fairs/id/"+strconv.Itoa(ID), nil)
	response := executeRequest(req, r)

	err := checkResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err)
	}
}

//Test for "/api/id/{searchParam}" PUT
func TestUpdateFair(t *testing.T) {
	database.Connect()
	createFairMock()
	defer deleteFairMock()

	fairMockUpdate := models.Fair{Id: ID, Longitude: "UpdateMockLongitude", Latitude: "UpdateMockLatitude", SetCen: "UpdateMockSetCen", AreaP: "UpdateMockAreaP", CodDist: "UpdateMockCodDist", District: "UpdateMockDistrict", CodSubPref: "UpdateMockCodSubPref", SubPref: "UpdateMockSubPref", RegionFive: "UpdateMockRegion5", RegionEight: "UpdateMockRegion8", NameFair: "UpdateMockNameFair", Record: "UpdateMockRecord", Street: "UpdateMockStreet", Number: "UpdateMockNumber", Neighbourhood: "UpdateMockNeighbourhood", Reference: "UpdateMockReference"}
	jsonMockUpdate, _ := json.Marshal(fairMockUpdate)

	r := routsTestSetup()
	r.HandleFunc("/api/fairs/{searchType}/{searchParam}", API.UpdateFair).Methods("Put")
	req, _ := http.NewRequest("PUT", "/api/fairs/id/"+strconv.Itoa(ID), bytes.NewBuffer(jsonMockUpdate))
	response := executeRequest(req, r)

	err := checkResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err)
	}

}

//Test for "/api/fairs" POST
func TestCreateFair(t *testing.T) {
	database.Connect()

	fairMockUpdate := models.Fair{Longitude: "CreateMockLongitude", Latitude: "CreateMockLatitude", SetCen: "CreateMockSetCen", AreaP: "CreateMockAreaP", CodDist: "CreateMockCodDist", District: "CreateMockDistrict", CodSubPref: "CreateMockCodSubPref", SubPref: "CreateMockSubPref", RegionFive: "CreateMockRegion5", RegionEight: "CreateMockRegion8", NameFair: "CreateMockNameFair", Record: "CreateMockRecord", Street: "CreateMockStreet", Number: "CreateMockNumber", Neighbourhood: "CreateMockNeighbourhood", Reference: "CreateMockReference"}
	jsonMockUpdate, _ := json.Marshal(fairMockUpdate)

	r := routsTestSetup()
	r.HandleFunc("/api/fairs", API.CreateFair).Methods("Post")
	req, _ := http.NewRequest("POST", "/api/fairs", bytes.NewBuffer(jsonMockUpdate))
	response := executeRequest(req, r)

	err := checkResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err)
	}

}

//-----------------TestFunctionBegin-----------------//

//Mocking a Fair for testing purposes
var ID int

func createFairMock() {
	fairMock := models.Fair{
		Longitude:     "mockLongitude",
		Latitude:      "mockLatitude",
		SetCen:        "mockSetCen",
		AreaP:         "mockAreaP",
		CodDist:       "mockCodDist",
		District:      "mockDistrict",
		CodSubPref:    "mockCodSubPref",
		SubPref:       "mockSubPref",
		RegionFive:    "mockRegion5",
		RegionEight:   "mockRegion8",
		NameFair:      "mockNameFair",
		Record:        "mockRecord",
		Street:        "mockStreet",
		Number:        "mockNumber",
		Neighbourhood: "mockNeighbourhood",
		Reference:     "mockReference",
	}
	database.DB.Exec("SELECT setval('fairs_id_seq', (SELECT MAX(id) FROM fairs));")
	database.DB.Create(&fairMock)
	ID = fairMock.Id
}
func deleteFairMock() {
	var fairMock models.Fair
	database.DB.Delete(&fairMock, ID)
}

//Use search type and search param to find a specific fair "/api/fairs/{searchType}/{searchParam}"
func searchFair(searchType string, searchParam string) error {
	database.Connect()
	createFairMock()
	defer deleteFairMock()

	r := routsTestSetup()
	r.HandleFunc("/api/fairs/{searchType}/{searchParam}", API.GetFair).Methods("Get")

	req, _ := http.NewRequest("GET", "/api/fairs/"+searchType+"/"+searchParam, nil)
	response := executeRequest(req, r)

	err := checkResponseCode(http.StatusOK, response.Code)
	if err != nil {
		return err
	}

	err = checkResponseBodyLength(response)
	if err != nil {
		return err
	}

	return nil
}

//Creates a router
func routsTestSetup() *mux.Router {
	routes := mux.NewRouter()
	return routes
}

//Get response record
func executeRequest(req *http.Request, r *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

//Verify status code
func checkResponseCode(expected, actual int) error {
	if expected != actual {
		return errors.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
	return nil
}

//Verify response body length
func checkResponseBodyLength(response *httptest.ResponseRecorder) error {
	if response.Body.Len() < 1 {
		return errors.New("response body is not expected")
	}

	return nil
}

//-----------------TestFunctionsEnd-----------------//
