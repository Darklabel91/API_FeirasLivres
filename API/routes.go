package API

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/fairs", GetFairs).Methods("Get")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
