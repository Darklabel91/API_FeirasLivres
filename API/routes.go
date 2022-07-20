package API

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(contentType)
	r.HandleFunc("/api/fairs/{searchType}/{searchParam}", GetFair).Methods("Get")
	r.HandleFunc("/api/fairs/id/{searchParam}", DeleteFair).Methods("Delete")
	r.HandleFunc("/api/fairs/id/{searchParam}", UpdateFair).Methods("Put")
	r.HandleFunc("/api/fairs", CreateFair).Methods("Post")
	r.HandleFunc("/api/fairs", GetFairs).Methods("Get")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func contentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
