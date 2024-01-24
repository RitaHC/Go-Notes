package main

import (
	"errors"
	"fmt"
)

type user struct {
	name                string
	number              int
	scheduleForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	// Use a if else syntax to check if the key exist
	if !ok {
		return false, errors.New("not found")
	}
	if existingUser.scheduleForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

func main() {
	// CHECK IF A KEY EXIST
	// SYNTAX:  elem, ok := m[key]
	// if key exist in m --> ok is true; if not then its false
	// Delete an element
	// SYNTAX:  delete(m, key)
	// delete (fromTheMap, Key)
	fmt.Println("------------------")
	// Create a map for testing with type user as value
	users := map[string]user{
		"Alice": {name: "Alice", number: 376, scheduleForDeletion: true},
		"Bob":   {name: "Bob", number: 264, scheduleForDeletion: false},
		"Chloe": {name: "Charlie", number: 376, scheduleForDeletion: false},
	}

	result1, err1 := deleteIfNecessary(users, "Alice")
	// check for error
	if err1 != nil {
		fmt.Println("ERROR: ", err1)
		return
	}
	fmt.Println(result1)
	////////////////////////////
	result2, err2 := deleteIfNecessary(users, "Bob")
	if err2 != nil {
		fmt.Println("ERROR: ", err2)
		return
	}
	fmt.Println(result2)
	////////////////////////////
	result3, err3 := deleteIfNecessary(users, "Chloe")
	if err3 != nil {
		fmt.Println("ERROR: ", err3)
		return
	}
	fmt.Println(result3)
	fmt.Println("-------------------------")
	for key, value := range users {
		fmt.Println(key, " : ", value)
	}

}
