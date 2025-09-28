package main

import "fmt"

// the place where you will actually going to use this in nested struct

type car struct {
	name      string
	brand     string
	modelYear int
	price     float64
	supplier  struct {
		name    string
		number  int
		address string
	}
}

func main() {

	// this is the anonymous struct it is declare and use at the same time by the way and we do not need them any where else
	wheel := struct {
		make  string
		model int
	}{
		make:  "tesla",
		model: 12,
	}
	fmt.Println(wheel.make)

	my_car := car{
		name:      "Tesla",
		brand:     "Tesla",
		modelYear: 2022,
		price:     2300.56,
		supplier: struct {
			name    string
			number  int
			address string
		}{
			name:    "Inshal",
			number:  923363438230,
			address: "Lahore Pakistan",
		}, // as it was anonymous struct
	}
	my_car.supplier.name = "Ukasha"
	my_car.supplier.number = 923363148906
	my_car.supplier.address = "Karachi Pakistan"

	fmt.Println("My car is " + my_car.name)

}

// key points just like c
// go also have
// - positional initialization of structs like wheel:= {"Tesla",23}
// - designated or named initialization
