package main

import "fmt"

func add(s1 string, s2 string) string {
	return s1 + s2
}

// syntactic sugar here if all the parameter have
// same types then you can declare types at the end only
func add_int(i1, i2 int) int {
	return i1 + i2
}

// in go normally arguments are passed by value not by reference.
// its variable are like c style not python or javascript binding
func inc(i int) int {
	i++
	return i
}


func main() {
	fmt.Println(add("Hello", "World"))
	fmt.Println(add_int(52, 12))
	//-----------------------------------//
	// working of incrementing function
	//-----------------------------------//
	i := 5
	i = inc(i)
	fmt.Println(i)
	//-----------------------------------//
	// Go does not allow un used variable
	//-----------------------------------//
	x, _ := getPoint() // as it does not allow to have un used variable what if i have to use only one variable other variable will be left unused so here come _ {underscore} compiler completely ignores or remove it, it is like python or js convention
	fmt.Printf("The value of x is : %d", x)
	//-----------------------------------//

}

// func add_int(i1, i2 int) int  is called function signature btw go idiom
// as you know the concept of call back in javascript same in go function acn be passed to another function
// so another function can  call them back just like a will call you back. on phone call

// --- DO YOU Know instead of garbage value if variable is declared it is assigned to zero
