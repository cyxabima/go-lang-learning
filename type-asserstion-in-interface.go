package main

import "fmt"

type rec struct {
	width  float64
	height float64
}

type Shape interface {
	area() float64
}

type cir struct {
	radius float64
}

func (c cir) area() float64 {
	return 3.142 * c.radius * c.radius
}

func (rec rec) area() float64 {
	return rec.width * rec.height
}

func shapeAreaWithShapeName(s Shape) string {
	c, ok := s.(cir)
	if ok {
		return fmt.Sprintf("The area of circle is %.2f", c.area())
	}
	r, ok := s.(rec)
	if ok {
		return fmt.Sprintf("The area of circle is %.2f", r.area())
	}

	return fmt.Sprintf("The area of unknown shape is %.2f", s.area())

}

func main() {
	circle := cir{radius: 22.12231}

	fmt.Println(shapeAreaWithShapeName(circle))

}


// Successive assertion is being done over here 
// the syntax here is interface name or we can say a variable of type interface here is "S"
// s.(rect) `--> rect :: type of struct which is implementing the interface`
// there are two variable we can use one for the actual oject and other for checked wheather that object of the given type or not 
// r,ok := s.(rec) if ok the s "shape" is rectangle