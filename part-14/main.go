package main

import (
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Talk struct {
	Name        string
	Description string
	Properties  map[string]interface{}
}

func encode(v interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	talks := []*Talk{
		&Talk{"Malo", "Ekino", map[string]interface{}{}},
		&Talk{"Thomas", "golang", map[string]interface{}{}},
	}

	goji.Get("/talks/:id", func(c web.C, w http.ResponseWriter, r *http.Request) {
		pos, _ := strconv.Atoi(c.URLParams["id"])
		encode(talks[pos], w)
	})

	goji.Get("/talks", func(w http.ResponseWriter, r *http.Request) {
		encode(talks, w)
	})

	goji.Serve()
}