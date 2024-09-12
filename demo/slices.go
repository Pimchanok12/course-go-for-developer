package main

import "fmt"

func main() {
	//     index   0    1    2
	sa := [3]string{"a", "b", "c"}
	nong := sa[0:2] // "a", "b"
	fmt.Printf("origin: %#v\n", nong)
	nong = append(nong, "d")
	fmt.Printf("s: %#v\n", nong)

	water := sa[1:len(sa)] // "b", "c"
	c := cap(water)
	fmt.Println("water cap: ", c)
	fmt.Printf("water: %#v\n", water)
	water = append(water, "e")
	water[0] = "hello"
	fmt.Printf("water: %#v\n", water)
	fmt.Printf("nong: %#v\n", nong)
	fmt.Printf("sa: %#v\n", sa)

	xs := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("xs: ", xs[0:len(xs)])
	fmt.Println("xs: ", xs[:len(xs)])
	fmt.Println("xs: ", xs[2:])
	fmt.Println("xs: ", xs[:3])

	Hello(xs[:])
	Hi5(xs)
}

func Hello(xs []string) {
	fmt.Println("xs: ", xs)
}

func Hi5(xs [5]string) {
	fmt.Println("xs: ", xs)
}
