package main

import (
	"errors"
	"fmt"
)

type users struct {
	name            string
	phoneNumber     int
	ready_to_delete bool
}

func map_maker(names []string, phones []int) (map[string]users, error) {
	if len(names) != len(phones) {
		return nil, errors.New("invalid size")
	}

	user_map := make(map[string]users)
	for i := 0; i < len(names); i++ {
		name := names[i]
		phone := phones[i]

		user_map[name] = users{
			name:        name,
			phoneNumber: phone,
		}
	}

	return user_map, nil
}

func delete_user(user_map map[string]users, name string) (bool, error) {
	// search if user exits
	user, ok := user_map[name]

	if !ok {
		return false, errors.New("user not found")
	}

	if user.ready_to_delete {
		print(user.name)
		delete(user_map, name)
		return true, nil
	}
	return false, nil
}

func main() {
	fmt.Println("here are maps")

	names := []string{"ukasha", "inshal", "urwah"}
	phones := []int{1, 2, 3}

	user_map, err := map_maker(names, phones)

	if err != nil {
		fmt.Print("Sorry failed to make a map")
	}

	urwah := user_map["urwah"]
	urwah.ready_to_delete = true
	user_map["urwah"] = urwah

	_, err = delete_user(user_map, "urwah")
	if err != nil {
		println("failed to delete")
	}

	for _, name := range names {

		// why cuz one of the key is deleted from map but if we access if we get falsy values instead of error therefor
		user, ok := user_map[name]

		if !ok {
			continue
		}
		fmt.Printf("Key : %v, value:\n", name)
		fmt.Println("- name", user.name)
		fmt.Println("- phone", user.phoneNumber)

	}

}
