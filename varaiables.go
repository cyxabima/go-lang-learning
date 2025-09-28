package main

import "fmt"

func main() {
	var name string = "ukasha"
	var adult bool = true
	var age uint8 = 20
	birthdate := 24 // it is still static type just like auto in cpp

	// multiple declaration on same line
	brother,b_age := "Inshal",18

	
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(adult)
	fmt.Println(birthdate)
	fmt.Println(brother)
	fmt.Println(b_age)
}


// there are multiple variant of int uint floats complex:
	// int int8 int16 int32 int64
	// uint uint8 uint16 int32 uint64 uintptr
	//float32 float64
	//complex64 complex128

	// there are other types as well for
	//~ byte ;; alias of uint8
	// rune ;; represent a unicode point under the hood it is alias of int32


// unless you have a specific reason stick to default types