package main

import "fmt"

func counterGenerator(initialVal int) func() int {
	count := initialVal

	return func() int {
		count++
		return count
	}
}

func main() {
	counterA := counterGenerator(1)
	counterB := counterGenerator(10)

	fmt.Println("Counter A:")
	fmt.Println(counterA()) // so where the function is called count is mutated and see that count as 1 as a value it is updated to 2 
	fmt.Println(counterA()) // updated 3 
	fmt.Println(counterA()) // updated 4 

	fmt.Println("Counter B:")
	fmt.Println(counterB())
	fmt.Println(counterB())
	fmt.Println(counterB())
	fmt.Println(counterB())
}

// closure in go is similar to closure in js or python

// if in a inner function a variable of outer function is used and we are returning that inner function as a return value

// every instance of that function will have the mutable reference of that variable
