package main

import "fmt"

// calculate the total coast of bulk messages
func totalCoast(messages int) float64 {
	total := 0.0
	for i := 0; i < messages; i++ {
		total += 1.0 + (0.01 * float64(i))
	}
	return total
}

// calculate the total number of message to be send under certain threshold

func totalMessages(thresh float64) int {
	amount := 0.0
	for i := 0; ; i++ {
		if amount > thresh {
			return i
		}
		amount += 1.0 + (0.01 * float64(i))
	}
}

func main() {
	amount := totalCoast(50)
	messages := totalMessages(62.00)
	fmt.Println("Total coast to send 50 messages is: ")
	fmt.Println(amount)
	fmt.Println("Total message to be send under 63: ")
	fmt.Println(messages)
}

// what we get is that we can omit any of the expression of for loop
/// just like is total messages we have not used comparison in the loop instead in the body of loop there is condition which
// directly return i from the function otherwise we have to create a global to the function varaiable while be return after
// after loop ends
