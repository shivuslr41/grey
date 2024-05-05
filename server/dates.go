package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// dates respond with list of available trading dates from current selected date to expiry date
func dates(w http.ResponseWriter, r *http.Request) {
	expiryDate := mux.Vars(r)["expiryDate"]
	selectedDate := mux.Vars(r)["date"]
	symbol := mux.Vars(r)["symbol"]
	ed, err := time.Parse("2006-01-02", expiryDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("dates: date convertion error:", err)
		}
		return
	}
	entries, err := os.ReadDir(basePath + fmt.Sprint(ed.Year(), "/") + expiryDate + "/" + symbol)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("dates: read from folder error:", err)
		}
		return
	}
	var availableDates []string
	for i := range entries {
		d := strings.TrimRight(entries[i].Name(), ".csv")
		if d >= selectedDate {
			availableDates = append(availableDates, d)
		}
	}
	sort.Strings(availableDates)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&availableDates)
	if err != nil {
		log.Println("dates: sending/encoding error response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("dates: sending response error:", err)
		}
		return
	}
}
