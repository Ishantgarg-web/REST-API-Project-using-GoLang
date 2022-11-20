package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestCountry struct {
	Name struct {
		Common   string
		Official string
	} `json:"name"`
	Independenet bool     `json:"independent"`
	Capital      []string `json:"capital"`
	Region       string   `json:"region"`
	SubRegion    string   `json:"subregion"`
	Area         float64  `json:"area"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloWorld).Methods("GET")
	r.HandleFunc("/{name}", searchByCountryName).Methods("GET")

	http.ListenAndServe(":8010", r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is API Project using GoLang!!\n Please Follow me, if you like the content")
}

func searchByCountryName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //getting parameters that passes through endpoint

	url := "https://restcountries.com/v3.1/name/" + params["name"] + "?fullText=true"
	// This is REST API that we will use to get the data

	data, err := http.Get(url)
	if err != nil { // checking error
		log.Fatal(err)
	}
	var countryData []RestCountry
	json.NewDecoder(data.Body).Decode(&countryData) // decode the reponse as our struct

	// Now encode the response as our struct
	json.NewEncoder(w).Encode(countryData)
}
