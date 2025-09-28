package main

import "fmt"

func main() {
	const PI = 3.142

	// one thing to remember const in go is not javascript or typescript
	// its value must be know in the compiled time or if you are computing the value it must
	// be computed on compiled time

	// btw walrus operator syntax doesn't exits for constants
	// const PI float32 = 3.142 also correct

	fmt.Println("Hey There is PI:", PI)
}
