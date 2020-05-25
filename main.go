package main

import (
	"encoding/json"
	ghs_api "github.com/gabielfemi/ghs-api/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/ghs", Index)
	router.HandleFunc("/api/ghs/{id}", GetAGhs)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:7000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	allGhs := ghs_api.GetAllGhsFromDb()
	_ = json.NewEncoder(w).Encode(allGhs)

}

func GetAGhs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(vars["id"])
}
