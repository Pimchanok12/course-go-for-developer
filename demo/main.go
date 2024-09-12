package main

import "fmt"

// Signature of the function
// func(int, int) int
func sum(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

type myFunc func(int, int) int

func calculate(fn myFunc, a int, b int) int {
	return fn(a, b)
}

func getUser(name string, age int) (r string, err error) {
	r = fmt.Sprintf("Name: %s, Age: %d", name, age)
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {

	a, b := swap(1, 2)
	fmt.Println(a, b)

	var fn = func(i1, i2 int) int {
		return i1 + i2
	}

	fmt.Println(fn(1, 2))

	r1 := calculate(sum, 1, 2)
	fmt.Println(r1) // 3

	r2 := calculate(minus, 1, 2)
	fmt.Println(r2) // -1

	fmt.Println("Hello, World!")
}

/* Java

-- Flight.java --
import System.out;

// FD 3567
public class Flight {
	private String airlineCode;
	private String number;

	public static Flight NewAirAsia(String number) {
		return new Flight("FD", number);
	}

	private Flight() {
	}

	private Flight(String airlineCode, String number) {
		this.airlineCode = airlineCode;
		this.number = number;
	}

	public String getFlightNumber() {
		return this.airlineCode + this.number;
	}
}


--- Hello.java ---
import System.out;

public class Hello {
	public static void main(String[] args) {
		Integer number = 2;
		out.println("Hello, World!");
	}

	// lambda predicate method sum
	public static Integer calculate(predicate<Integer, Integer> sum, Integer a, Integer b) {
	}

	// signature of the function is Integer, Integer -> Integer
	public static Integer sum(Integer a, Integer b) {
		return a + b;
	}
}
*/
