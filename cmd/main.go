package main

import (
	"github.com/afritzler/covid-skill"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cases", covid.Cases)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
