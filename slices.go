package main

import "fmt"

func main() {
	coast := make([]int, 3, 5)
	for i := 0; i < len(coast); i++ {
		coast[i] = i + 1
	}

	fmt.Println(coast)
}

// make is use to create a slice directly without caring about the underlying array
// coast := make([]int, 3, 5)
// here array is of capacity 5 and len of slice is 3
// generally we  do not declare the capacity of array the compiler infer the capacity as len
// coast := make([]int, 3)

// there are two functions cap() and len()
// so  len return the current size of slice
// cap() returns the max length of slice before reallocation of array is necessary

// A slice is not the data itself.
// It's a small descriptor (header) that contains:

// Pointer → points to the first element of the underlying array

// Length → number of accessible elements

// Capacity → max elements until reallocation is needed
