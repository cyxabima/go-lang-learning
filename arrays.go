package main

import "fmt"

func main() {
	// var array [3]int
	array := [3]int{1, 2, 3}
	fmt.Println(array)
	fmt.Println(array[:2]) // this is slice
}

// array are of fixed size
//c := [...]string{"apple", "banana", "cherry"} // Compiler determines size based on elements
// slices are on top of array

// when we create slice all the values is copied from array to other memory location with more memory reserved
