package main

import "fmt"

// you already know about them
// lambda in python
// or call back in js "arrow functions"

func math_(bin_op func(a, b int) int, val int) int {
	return bin_op(val, val)
}

func main() {

	sq := math_(func(a, b int) int {
		return a * b
	},
		45)

	fmt.Println(sq)
}
