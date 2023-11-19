package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	date = iota
	symbol
	expiry
	strike
	type_CE
	open_CE
	high_CE
	low_CE
	close_CE
	ltp_CE
	pclose_CE
	oi_CE
	coi_CE
	lot
	spot
	type_PE
	open_PE
	high_PE
	low_PE
	close_PE
	ltp_PE
	pclose_PE
	oi_PE
	coi_PE
	expiry_fut
	open
	high
	low
	close
	ltp
	pclose
	oi
	coi
	lot_fut
	symbol_vix
	open_vix
	close_vix
	high_vix
	low_vix
	pclose_vix
	iv_CE
	iv_PE
	delta_CE
	delta_PE
	theta_CE
	theta_PE
	gamma_CE
	gamma_PE
	vega_CE
	vega_PE
)

type optionChain struct {
	CE     string `json:"ce"`
	PE     string `json:"pe"`
	Cdelta string `json:"cdelta"`
	Pdelta string `json:"pdelta"`
	Ctheta string `json:"ctheta"`
	Ptheta string `json:"ptheta"`
	Vega   string `json:"vega"`
	Gamma  string `json:"gamma"`
	Strike string `json:"strike"`
	Spot   string `json:"spot"`
	Fut    string `json:"fut"`
	Vix    string `json:"vix"`
}

// chain respond with option chain data of current selected date and expiry date
func chain(w http.ResponseWriter, r *http.Request) {
	expiryDate := mux.Vars(r)["expiryDate"]
	selectedDate := mux.Vars(r)["date"]
	symbol := mux.Vars(r)["symbol"]
	oc, err := getChain(symbol, expiryDate, selectedDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("chain: sending response error:", err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&oc)
	if err != nil {
		log.Println("chain: sending/encoding error response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprint(w, "INTERNAL_ERROR: ", err)
		if err != nil {
			log.Println("chain: sending response error:", err)
		}
		return
	}
}

func getChain(symbol, expiryDate, selectedDate string) ([]*optionChain, error) {
	// TODO for v2, allow client to select open, high and low
	// current default is close prices
	ed, err := time.Parse("2006-01-02", expiryDate)
	if err != nil {
		log.Println("getChain: date convertion error:", err)
		return nil, err
	}
	f, err := os.Open("../../data/option_chain/" + fmt.Sprint(ed.Year(), "/") + expiryDate + "/" + symbol + "/" + selectedDate + ".csv")
	if err != nil {
		log.Println("getChain: read from file error:", err)
		return nil, err
	}
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("getChain: read csv error:", err)
		return nil, err
	}
	var ochain []*optionChain
	for i := range records {
		ochain = append(ochain, &optionChain{
			CE:     records[i][close_CE],
			PE:     records[i][close_PE],
			Cdelta: records[i][delta_CE],
			Pdelta: records[i][delta_PE],
			Ctheta: records[i][theta_CE],
			Ptheta: records[i][theta_PE],
			Vega:   records[i][vega_CE],
			Gamma:  records[i][vega_PE],
			Strike: records[i][strike],
			Spot:   records[i][spot],
			Fut:    records[i][close],
			Vix:    records[i][close_vix],
		})
	}
	return ochain, nil
}
