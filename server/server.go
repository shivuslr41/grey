package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{"https://k7z632-3000.csb.app", "https://test.rebelscode.online", ""} // remove "" later
		origin := r.Header.Get("Origin")
		if !contains(allowedOrigins, origin) {
			log.Println("Invalid origin:", origin)
			http.Error(w, "Invalid origin", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/expiries/{symbol}/{date}", listExpiries).Methods("GET")
	myRouter.HandleFunc("/chain/{symbol}/{date}/{expiryDate}", chain).Methods("GET")
	myRouter.HandleFunc("/dates/{symbol}/{date}/{expiryDate}", dates).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", corsMiddleware(myRouter)))
}

func main() {
	handleRequests()
}

func contains(origins []string, origin string) bool {
	for i := range origins {
		if origins[i] == origin {
			return true
		}
	}
	return false
}
