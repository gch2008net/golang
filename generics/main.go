package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func main() {
	fmt.Println("this is a generics example")

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// int64 intssum:=SumInts(ints)
	var intsum int64 = SumInts(ints)
	var floatsum float64 = SumFloats(floats)

	printMsg := fmt.Sprintf("N-Neneric sum2222222:  %v  and  %v ", intsum, floatsum)

	fmt.Println("printMsg: " + printMsg)

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		intsum,
		floatsum)

	var intsum2 int64 = SumIntsOrFloats(ints)
	var floatsum2 float64 = SumIntsOrFloats(floats)

	fmt.Printf("generics  intsum: %v  floatsum:  %v \n", intsum2, floatsum2)

	var intsum3 int64 = SumNumbers(ints)
	var floatsum3 float64 = SumNumbers(floats)

	fmt.Printf("generics333333333  intsum: %v  floatsum:  %v \n", intsum3, floatsum3)
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
