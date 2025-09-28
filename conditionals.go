package main

import "fmt"

func main() {

	height := 6

	if height >= 6 {
		fmt.Println("You are very tall")
	} else if height > 5 {
		fmt.Println("You are okay")
	} else {
		fmt.Println("Not tall enough")
	}

	// there also a possible syntax is go lang

	// if INITIALIZATION ; CONDITION {}
	if size := 32; size > 18 {
		fmt.Println("you are allowed")
	}

	// keep in mind now the scope of size is limited to if block only
}
