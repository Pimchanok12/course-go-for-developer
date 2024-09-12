package main

import "fmt"

// Factory pattern
func NewAirAsia(number int) flight {
	f := flight{
		number:      number,
		airlineCode: "FD",
	}

	return f
}

type flight struct {
	number      int
	airlineCode string
	destination string
}

// reciever
func (f flight) getFlightNumber() string {
	code := fmt.Sprintf("%s %d", f.airlineCode, f.number)
	return code
}

// public String getFlightNumber() {
// 	return this.airlineCode + this.number;
// }

func main() {
	// Flight f = new Flight();
	f := Flight{}
	// var f Flight
	fmt.Printf("%#v\n", f)
	fmt.Println(f.airlineCode)
	fmt.Println(f.number)

	// Flight f = new Flight("AA", "123");
	fl := Flight{123, "AA", "DMK"}
	fmt.Printf("%#v\n", fl)
	ff := Flight{number: 123, airlineCode: "AA"}
	fmt.Printf("%#v\n", ff)

	code := fl.getFlightNumber()
	fmt.Println("code in main: ", code)
}
