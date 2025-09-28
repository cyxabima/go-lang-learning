package main

import (
	"fmt"
)

type email struct {
	senderEmail   string
	receiverEmail string
	message       string
	coast         float64
}

type sms struct {
	senderPhone   string
	receiverPhone string
	message       string
	coast         float64
}

func showDetails(m Message) (string, float64) {

	switch v := m.(type) {
	case sms:
		return v.receiverPhone, v.coast
	case email:
		return v.receiverEmail, v.coast
	default:
		return "", 0.0
	}

}

type Message interface {
}

func main() {
	fmt.Println(showDetails(
		email{receiverEmail: "ukashaanwerali@gmail.com", coast: 22.12}))

	fmt.Println(showDetails(
		sms{receiverPhone: "03363148906", coast: 122.12}))

	fmt.Println(showDetails(12))
}

// an interface with no method is the interface of every object.
