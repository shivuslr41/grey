package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
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
	CE     float64 `json:"ce"`
	PE     float64 `json:"pe"`
	Cdelta float64 `json:"cdelta"`
	Pdelta float64 `json:"pdelta"`
	Ctheta float64 `json:"ctheta"`
	Ptheta float64 `json:"ptheta"`
	Cvega  float64 `json:"cvega"`
	Pvega  float64 `json:"pvega"`
	Cgamma float64 `json:"cgamma"`
	Pgamma float64 `json:"pgamma"`
	Strike float64 `json:"strike"`
	Spot   float64 `json:"spot"`
	Fut    float64 `json:"fut"`
	Vix    float64 `json:"vix"`
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
	f, err := os.Open(basePath + fmt.Sprint(ed.Year(), "/") + expiryDate + "/" + symbol + "/" + selectedDate + ".csv")
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
			CE:     parseFloat(records[i][close_CE], false),
			PE:     parseFloat(records[i][close_PE], false),
			Cdelta: parseFloat(records[i][delta_CE], true),
			Pdelta: parseFloat(records[i][delta_PE], true),
			Ctheta: parseFloat(records[i][theta_CE], true),
			Ptheta: parseFloat(records[i][theta_PE], true),
			Cvega:  parseFloat(records[i][vega_CE], true),
			Pvega:  parseFloat(records[i][vega_PE], true),
			Cgamma: parseFloat(records[i][gamma_CE], true),
			Pgamma: parseFloat(records[i][gamma_PE], true),
			Strike: parseFloat(records[i][strike], false),
			Spot:   parseFloat(records[i][spot], false),
			Fut:    parseFloat(records[i][close], false),
			Vix:    parseFloat(records[i][close_vix], true),
		})
	}
	return ochain, nil
}

func parseFloat(f string, r bool) float64 {
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		if r {
			return math.Round(s*100) / 100
		}
		return s
	} else {
		return 0
	}
}
