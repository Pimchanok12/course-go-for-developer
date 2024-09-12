package main

import (
	"encoding/json"
	"log"
)

type Flight struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	AirlineCode string `json:"airlineCode"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

func main() {
	// using encoding/json
	// struct -> JSON ([]byte of text)   ==> Marshal (Serialize)

	f := Flight{ID: 1, Number: 3250, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"}

	log.Printf("%#v\n", f)

	b, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))

	// JSON ([]byte of text) -> struct 	 ==> Unmarshal (Deserialize)
	raw := `{"id":1,"number":3250,"AirlineCode":"FD","destination":"DMK","arrival":"KKC"}`
	var f2 Flight
	err = json.Unmarshal([]byte(raw), &f2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", f2)
}
