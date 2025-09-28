package main

import "fmt"

func sum(num ...int) (total int) {

	for i := 0; i < len(num); i++ {
		total += num[i]
	}
	return total
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	// or  we can use spread operator by spreading the slice
	fmt.Println(total)
	numb := []int{1,2,3,4}
	total2:= sum(numb...)
	fmt.Println(total2)
}
