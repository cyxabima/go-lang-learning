package main

import "fmt"

func main() {
	number := 64

	for number != 0 {
		reminder := number % 2
		fmt.Println(reminder)
		number = number / 2
	}
}

// there no while keyword in go
// the syntax of while loop look likes this
// for condition{
// 	body;
// }
