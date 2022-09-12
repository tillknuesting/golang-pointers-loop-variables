package main

import "fmt"

func main() {
	example1()
}

func example1() {
	var intSlice []*int

	for _, p := range []int{10, 11, 12, 13} {
		intSlice = append(intSlice, &p) // want "taking a pointer for the loop variable p"
	}

	for _, p := range intSlice {
		fmt.Println(*p)
	}
}
