package main

import "fmt"

type shape interface {
	area() float32
	parameter() float32
}

// now any struct which has same method we called that its interface is shape meaning it is shape
// like in java we do not need to explicitly used the implements shape
// as long as methods have same parameter and return types it will be called circle is shape like that circle implements shape a kind of this

// ------------------------------------//
// interface are implemented implicitly
// ------------------------------------//

type circle struct {
	radius float32
}
type rectangle struct {
	length, width float32
}
type square struct {
	width float32
}

func (r rectangle) area() float32 {
	area := r.length * r.width
	return area
}
func (s square) area() float32 {
	area := s.width * s.width
	return area
}
func (c circle) area() float32 {
	area := 3.142 * c.radius * c.radius
	return area
}

func (r rectangle) parameter() float32 {
	area := 2 * (r.length + r.width)
	return area
}

func (s square) parameter() float32 {
	area := 2 * (s.width + s.width)
	return area
}

func (c circle) parameter() float32 {
	area := 2 * 3.142 * c.radius
	return area
}

// now the actual use of interface
// .. so this function is using interface shape in parameter meaning we can pass any struct which has same interface
func printShapeDetails(s shape) {
	area := s.area()
	parameter := s.parameter()
	fmt.Println("The area of shape is : ", area)
	fmt.Println("The parameter of shape is : ", parameter)
}

func main() {
	my_circle := circle{radius: 12.45}
	my_rect := rectangle{length: 13.5, width: 2.1}
	my_sq := square{width: 24}
	printShapeDetails(my_circle)
	printShapeDetails(my_rect)
	printShapeDetails(my_sq)
}

// also remember
// you can have named argument/parameter or un named arguments or parameter is function signature in interface for each method
// for eg
// we can have
type myInterface interface {
	concat_para(p1, p2 string) (detail string) // more readable
	// or
	concat_para2(string, string) string
}
