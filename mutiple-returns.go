package main

import "fmt"

// In go function can have multiple return value
// for multiple return value we wrap then in parenthesis
func getPoint() (x int, y int) { // what named return val??
	return 2, 14
}

// without named return value
func getPoint2() (int, int) { // not named return val
	return 2, 14
}

// ... some more thing about named return if you used named return you do not need to create variable in the block
// it will be declare by function signature it self just in parameter

func getCoordinates() (x int, y int) {
	//-----------------------------------//
	// Implicit / Naked return
	//-----------------------------------//
	x = 32
	y = 12
	return // this will automatically return x and y even if you can tells which arguments to return

}

func main() {
	//-----------------------------------//
	// Go does not allow un used variable
	//-----------------------------------//
	x, _ := getPoint() // as it does not allow to have un used variable what if i have to use only one variable other variable will be left unused so here come _ {underscore} compiler completely ignores or remove it, it is like python or js convention
	fmt.Printf("The value of x is : %d \n", x)
	//-----------------------------------//
	
	//-----------------------------------//
	_, y := getCoordinates() // in this we have used naked / implicit return
	fmt.Printf("The value of y is : %d \n", y)
}


//-----------------//
//---KEY POINTS---//
//----------------//

// - only used implicit return in smaller function
// having  named return value is good but avoid naked return 