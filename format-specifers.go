package main

import "fmt"

func main() {
	name := "Inshal"
	age := 18
	u_name, u_age := "urwah", 16
	fmt.Printf("The age of %T is %v \n", name, age) // v is btw verb T is type
	fmt.Printf("The age of %s is %d \n", name, age)
	msg := fmt.Sprintf("The age of %s is %d", u_name, u_age)
	print(msg)
}

// ?? REMEMBER
// println is a built-in function in Go. This means it can be used directly without importing any package.
// It is primarily intended as a debugging tool and is not recommended for production code.
