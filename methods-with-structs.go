package main

import (
	"errors"
	"fmt"
)

type rect struct {
	length, width float32 // again syntactic sugar
}

// here i have use almost all the possible function symbol construct
// this (r rect ) is the used to defined this method for any struct having type rect to access its value inside the function we have used a variable or parameter named r
func (r rect) area() (area float32, error error) {

	if r.length == 0 || r.width == 0 {
		return 0, errors.New("Rect is not defined yet")
	}

	area = r.length * r.width
	return // naked return
}

func main() {
	rectangle := rect{length: 22, width: 23.3}
	area, _ := rectangle.area()

	fmt.Printf("area is : %.2f", area)
}
