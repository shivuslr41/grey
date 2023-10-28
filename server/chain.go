package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	date = iota
	cvega
	cgamma
	ctheta
	cdelta
	ce
	strike
	pe
	pdelta
	ptheta
	pgamma
	pvega
	spot
	fut
	vix
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
	// maybe for v2, allow client to select open, high and low
	// current default is close prices
	f, err := os.Open("../../data/custom_option_chain/" + expiryDate + "/" + symbol + "/close/" + selectedDate + ".csv")
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
			CE:     records[i][ce],
			PE:     records[i][pe],
			Cdelta: records[i][cdelta],
			Pdelta: records[i][pdelta],
			Ctheta: records[i][ctheta],
			Ptheta: records[i][ptheta],
			Vega:   records[i][cvega],
			Gamma:  records[i][cgamma],
			Strike: records[i][strike],
			Spot:   records[i][spot],
			Fut:    records[i][fut],
			Vix:    records[i][vix],
		})
	}
	return ochain, nil
}
