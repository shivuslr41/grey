package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/expiries/{date}", listExpiries).Methods("GET")
	myRouter.HandleFunc("/chain/{symbol}/{date}/{expiryDate}", chain).Methods("GET")
	myRouter.HandleFunc("/payoff/{symbol}/{date}/{expiryDate}", payoff).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	handleRequests()
}
