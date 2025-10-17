package main

import "fmt"

func self_math(binary_func func(int, int) int) func(int) int {
	return func(val int) int {
		return binary_func(val, val)
	}
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Currying means a function which takes function as an argument and return function as return value")
	sq_func := self_math(multiply)
	fmt.Println(sq_func(5))
}
