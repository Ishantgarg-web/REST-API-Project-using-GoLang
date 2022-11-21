package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Problems struct {
	Category         string   `json:"category"`
	CorrectAnswer    string   `json:"correctAnswer"`
	IncorrectAnswers []string `json:"incorrectAnswers"`
	Question         string   `json:"question"`
	Difficulty       string   `json:"difficulty"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloWorld).Methods("GET")
	r.HandleFunc("/{category}/{difficulty}/{limit}", getProblems).Methods("GET")

	http.ListenAndServe(":8010", r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is API Project using GoLang!! Please Follow me, if you like the content")
}

func getProblems(w http.ResponseWriter, r *http.Request) {
	//getting variables passed through endpoint
	params := mux.Vars(r)

	url := "https://the-trivia-api.com/api/questions?categories=" + params["category"] + "&limit=" + params["limit"] + "&difficulty=" + params["difficulty"]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var allProblems []Problems
	json.NewDecoder(resp.Body).Decode(&allProblems)

	json.NewEncoder(w).Encode(allProblems)
}
