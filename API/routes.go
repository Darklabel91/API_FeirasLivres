package API

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/fairs/id/{searchParam}", UpdateFair).Methods("Put")
	r.HandleFunc("/api/fairs/{searchType}/{searchParam}", GetFair).Methods("Get")
	r.HandleFunc("/api/fairs", CreateFair).Methods("Post")
	r.HandleFunc("/api/fairs", GetFairs).Methods("Get")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
