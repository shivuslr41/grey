package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// expiryDates holds expiry dates available for backtesting
type expiryDates struct {
	Dates []string `json:"dates"`
}

// listExpiries accepts none or selected expiries and list eligible expiries to trade.
func listExpiries(w http.ResponseWriter, r *http.Request) {
	selectedDate, err := time.Parse("2006-01-02", mux.Vars(r)["date"])
	if err != nil {
		log.Println("listExpiries: selected date parse error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("listExpiries: sending response error:", err)
		}
		return
	}
	maxDate := selectedDate.AddDate(0, 0, 90)
	var availableDates expiryDates
	// convert all available string dates to time.Time format
	// and consider only those availableDates which are within
	// 90 days of current selectedDate
	for i := range availableExpiryDates.Dates {
		d, err := time.Parse("2006-01-02", availableExpiryDates.Dates[i])
		if err != nil {
			log.Println("listExpiries: available date parse error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
			if err != nil {
				log.Println("listExpiries: sending response error:", err)
			}
			return
		}
		if !d.After(maxDate) && !d.Before(selectedDate) {
			availableDates.Dates = append(availableDates.Dates, availableExpiryDates.Dates[i])
		}
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&availableDates)
	if err != nil {
		log.Println("listExpiries: sending/encoding error response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("listExpiries: sending response error:", err)
		}
		return
	}
}
