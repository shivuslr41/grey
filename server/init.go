package main

import (
	"log"
	"os"
)

var availableExpiryDates = make(map[string]*expiryDates)

func init() {
	basePath := "../../data/option_chain/"
	entries, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		expDates, err := os.ReadDir(basePath + e.Name())
		if err != nil {
			log.Fatal(err)
		}
		for _, ed := range expDates {
			symbols, err := os.ReadDir(basePath + e.Name() + "/" + ed.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, sym := range symbols {
				if _, ok := availableExpiryDates[sym.Name()]; !ok {
					availableExpiryDates[sym.Name()] = &expiryDates{Dates: []string{ed.Name()}}
				} else {
					availableExpiryDates[sym.Name()].Dates = append(availableExpiryDates[sym.Name()].Dates, ed.Name())
				}
			}
		}
	}
}
