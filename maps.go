package main

// JavaScript Objects = Maps
// Key : value pair
// . Notation is used to capture the key and values

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// create a new map for return value
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid Slices")
	}
	// Now for the user having same length of name and phone numbers
	// create a new map
	// First iterate over the existing slices
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		// Now add the key and values to the new map
		// Key is the name : Value is the user Struct
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}

	}
	return userMap, nil

}

func main() {
	fmt.Println("-----------------")
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Manie"] = 21
	fmt.Println(ages)

	fmt.Println("-----------------")
	names := map[string]int{
		"Rita":   30,
		"Hameer": 32,
	}
	fmt.Println(names)
	fmt.Println("-----------------")
	fmt.Println(len(names))

	fmt.Println("-----------------")
	testNames := []string{"Alice", "Bob", "Charlie"}
	testPhoneNumbers := []int{123, 456, 7899}
	// store the reesult and catch errors
	result, err := getUserMap(testNames, testPhoneNumbers)
	// catch error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// If no error then iterate over the map and print the result on different lines
	for userName, userInfo := range result {
		fmt.Println(userName, " : ", userInfo.phoneNumber)
	}

	fmt.Println("-----------------")
	testNames2 := []string{"Alice", "Rita", "Bob", "Charlie"}
	testPhoneNumbers2 := []int{123, 456, 7899}
	result2, err2 := getUserMap(testNames2, testPhoneNumbers2)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}
	for userName, userInfo := range result2 {
		fmt.Println(userName, " : ", userInfo.phoneNumber)
	}

	fmt.Println("-----------------")
}
