package main

import (
	"log"
	"os"
)

var availableExpiryDates expiryDates

func init() {
	entries, err := os.ReadDir("../../data/custom_option_chain/")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		availableExpiryDates.Dates = append(availableExpiryDates.Dates, e.Name())
	}
}
