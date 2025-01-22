package main

import (
	"fmt"
	"log"
	"net/http"
)

var stats map[string]map[string]int

func init() {
	stats = make(map[string]map[string]int)
	stats["morning"] = make(map[string]int)
	stats["evening"] = make(map[string]int)
}

type morningHandler struct {
	requestCount int
}

func (mh *morningHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mh.requestCount++

	beverage := r.PathValue("beverage")
	temperature := r.PathValue("temperature")

	if beverage == "" || temperature == "" {
		http.Error(w, "beverage and temperature are required", http.StatusBadRequest)
		return
	}

	stats["morning"][beverage]++
	stats["morning"][temperature]++

	fmt.Fprintf(w, "Good Morning! Have some %s %s.", temperature, beverage)
}

type eveningHandler struct {
	requestCount int
}

func (eh *eveningHandler) HandleEvening(w http.ResponseWriter, r *http.Request) {
	eh.requestCount++
	fmt.Fprintf(w, "Good Evening!")
}

func (eh *eveningHandler) HandleEveningBeverage(w http.ResponseWriter, r *http.Request) {
	eh.requestCount++

	beverage := r.PathValue("beverage")
	if beverage == "" {
		http.Error(w, "beverage is required", http.StatusBadRequest)
		return
	}

	stats["evening"][beverage]++

	fmt.Fprintf(w, "Good Evening! Have some %s.", beverage)
}

func main() {
	eh := &eveningHandler{}
	mh := &morningHandler{}

	mux := http.NewServeMux()
	mux.Handle("GET /evening", http.HandlerFunc(eh.HandleEvening))
	mux.Handle("GET /evening/{beverage}", http.HandlerFunc(eh.HandleEveningBeverage))
	mux.Handle("GET /morning/{beverage}/{temperature}", mh)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings!")
	})
	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Morning: %v\n", stats["morning"])
		fmt.Fprintf(w, "Evening: %v\n", stats["evening"])
	})

	log.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
