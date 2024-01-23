package main

import (
	"errors"
	"fmt"
)

/////////// ARRAYS == SLICES ///////

var myInts [10]int
var myStr []string

func getMessages() [3]string {
	return [3]string{
		"Click Sign up",
		"Better please click here",
		"We beg you to sign up",
	}
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessagesWithRetriesForPlan(plan string) ([]string, error) {
	// allMessages is an Array and not slice

	allMessages := getMessages()

	if plan == planPro {
		// so we need to make them all a slice to return the same dtat type
		return allMessages[:], nil
	} else if plan == planFree {
		return allMessages[0:2], nil
	} else {
		return nil, errors.New("unsupported plan")
	}
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	a := primes[1:5]

	fmt.Println("==== SLICES =====")
	fmt.Println(primes)
	fmt.Println("-----------------")
	messages := getMessages()
	for _, msg := range messages {
		fmt.Println(msg)
	}
	fmt.Println("-----------------")
	fmt.Println(a)
	fmt.Println("-----------------")

	// Struct of Person
	type Person struct {
		Name string
		Plan string
	}
	// Slice of type Struct laid above
	people := []Person{
		{Name: "Rita", Plan: planFree},
		{Name: "Hameer", Plan: planPro},
		{Name: "Rajat", Plan: "basic"},
	}

	// Iterate over the slice to analyse all the people with their plans
	for _, person := range people {
		// Now call the function in a manner that it can catch error if any
		// Call the function by calling the string having plan (use . Notation to call)
		result, err := getMessagesWithRetriesForPlan(person.Plan)
		if err != nil {
			fmt.Println(person.Name, " --> ")
			fmt.Println("Error: ", err)
			return
		}
		// Now that the error has been handled and the result is a slice
		// Now iterate over the Slice of result to display the messages in different row
		fmt.Println(person.Name, "-->")
		for _, ping := range result {
			fmt.Println(ping)
		}
	}

}
