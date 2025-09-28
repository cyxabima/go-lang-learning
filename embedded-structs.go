package main

import "fmt"

type Supplier struct {
	Name    string
	Number  int64
	Address string
}

// type Car struct {
// 	Name      string
// 	Brand     string
// 	ModelYear int
// 	Price     float64
// 	SupplierDetails  Supplier // nested
// }

// Embedded

type Car struct {
	Name      string
	Brand     string
	ModelYear int
	Price     float64
	Supplier  // embedded  as there is no name is just like an inheritance in oops but it is actually not. It’s just that the compiler lets you skip writing .Supplier.
}

func main() {
	// positional initialization  both are same
	myCar := Car{"Tesla", "Tesla", 2022, 2300.56, Supplier{"Ukasha", 923363148906, "Karachi"}}

	myCar = Car{
		Name:      "Tesla 34Ax",
		ModelYear: 2021,
		Brand:     "tesla",
		Price:     3267676,
		Supplier: Supplier{ // 👈 type name acts like the key the only difference is that this key will be called supplierDetails in nested as its name is this in nested
			Name:    "ukasha",
			Number:  923363148906,
			Address: "Karachi Pakistan"},
	}
	fmt.Println(myCar.Name) // preferred the actual which is not in another struct
	fmt.Println(myCar.ModelYear)
	fmt.Println(myCar.Supplier.Name)
	fmt.Println(myCar.Number)
}
