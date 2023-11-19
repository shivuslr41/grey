package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
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
	symbol := mux.Vars(r)["symbol"]
	if err != nil {
		log.Println("listExpiries: selected date parse error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("listExpiries: sending response error:", err)
		}
		return
	}
	// validate symbol
	if _, ok := availableExpiryDates[symbol]; !ok {
		log.Println("listExpiries: invalid symbol")
		w.WriteHeader(http.StatusBadRequest)
		_, err = fmt.Fprint(w, "BAD_REQUEST: ", err)
		if err != nil {
			log.Println("listExpiries: sending response error:", err)
		}
		return
	}
	monthlyMaxDate := selectedDate.AddDate(0, 0, 92)
	weeklyMaxDate := selectedDate.AddDate(0, 0, 32)
	beyondWeeklyDates := make(map[string]time.Time)
	var eligibleDates expiryDates
	// convert all available string dates to time.Time format
	// and consider only those eligible dates which are within
	// 90 days of current selectedDate
	for i := range availableExpiryDates[symbol].Dates {
		d, err := time.Parse("2006-01-02", availableExpiryDates[symbol].Dates[i])
		if err != nil {
			log.Println("listExpiries: available date parse error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
			if err != nil {
				log.Println("listExpiries: sending response error:", err)
			}
			return
		}
		if !d.After(monthlyMaxDate) && !d.Before(selectedDate) {
			if d.After(weeklyMaxDate) {
				beyondWeeklyDates[availableExpiryDates[symbol].Dates[i]] = d
				continue
			}
			eligibleDates.Dates = append(eligibleDates.Dates, availableExpiryDates[symbol].Dates[i])
		}
	}
	// get only monthly expiry dates among beyond weekly dates.
	eligibleDates.Dates = append(eligibleDates.Dates, getMonthlyDates(beyondWeeklyDates)...)
	// sort dates
	sort.Strings(eligibleDates.Dates)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&eligibleDates)
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

// get greatest date of each month among given dates
func getMonthlyDates(dates map[string]time.Time) []string {
	var monthlyExpiryDates []string
	monthlyExpiryDatesMap := make(map[time.Month]string)
	DatesByMonth := make(map[time.Month]time.Time)
	for strDate, date := range dates {
		month := date.Month()
		greatestDateOfMonth := DatesByMonth[month]
		if greatestDateOfMonth.Before(date) {
			DatesByMonth[month] = date
			monthlyExpiryDatesMap[month] = strDate
		}
	}
	for _, date := range monthlyExpiryDatesMap {
		monthlyExpiryDates = append(monthlyExpiryDates, date)
	}
	return monthlyExpiryDates
}
