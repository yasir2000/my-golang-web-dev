package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Endpoint     string `json: "endpoint"`
	Year         int    `json: "year"`
	Organization string `json: "organization"`
}

func cache(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("cachable", "ttl=30")

	response := Response{
		Endpoint:     "Cachable",
		Year:         2021,
		Organization: "GoYasir",
	}

	b, _ := json.Marshal(response)
	w.Write(b)
}

func nonCache(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Endpoint:     "NOT cachable",
		Year:         2021,
		Organization: "GoYasir",
	}

	b, _ := json.Marshal(response)
	w.Write(b)
}

func main() {
	http.HandleFunc("/cache", cache)
	http.HandleFunc("/no-cache", nonCache)

	http.ListenAndServe(":8083", nil)
}
