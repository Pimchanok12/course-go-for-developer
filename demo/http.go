package main

import (
	"fmt"
	"log"
	"net/http"
)

func flightHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("flightHandler Method: ", r.Method)
	if r.Method == "GET" {
		raw := `{
			"name": "anuchitO",
			"age": 25
			}`

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(raw))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/flights", flightHandler)

	h := http.HandlerFunc(flightHandler)

	http.Handle("/anuchito", h)

	log.Println("start server at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
