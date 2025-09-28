package main

import "fmt"

type message struct {
	len   int
	title string
	date  string
}

// nested structs

type user struct {
	userName string
	number   int
}

type messageWithDirection struct {
	content  string
	sender   user
	receiver user
}

func isSent(msg messageWithDirection) bool {
	if msg.sender.userName == "" {
		return false
	}
	if msg.receiver.userName == "" {
		return false
	}

	if msg.sender.number == 0 {
		return false
	}

	if msg.receiver.number == 0 {
		return false
	}

	return true
}

func main() {
	my_message := message{}
	my_message.len = 12
	my_message.title = "Hello World!"
	my_message.date = "10/09/2025"

	details := fmt.Sprintf("The message details is : \n title: %s \n len: %d \n date: %s \n ",
		my_message.title, my_message.len, my_message.date)

	fmt.Printf(details)
}
