package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func (t *Talk) AddProperty(name string, value interface{}) *Talk {
	t.Properties[name] = value

	return t
}

func NewTalk(name, description string) *Talk {
	return &Talk{
		Name: name,
		Description: description,
		Properties: make(map[string]interface{}),
	}
}

func main() {
	talks := []*Talk{
		NewTalk("Malo", "Ekino"),
		NewTalk("Thomas", "golang"),
		NewTalk("Leila", "Le « RUN » (ou la Tierce Maintenance Applicative)"),
		NewTalk("Thomas Carsuzan", "Cassandra"),
		NewTalk("Guillaume ", "REX Atelier"),
		NewTalk("Mickaël", "Panorama des solutions mobiles hybrides"),
		NewTalk("All", "Apéro"),
	}

	http.HandleFunc("/talks/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		e.Encode(talks[1])
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		e := json.NewEncoder(w)
		e.Encode(talks)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

