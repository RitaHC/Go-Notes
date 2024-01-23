package main

import (
	"errors"
	"fmt"
)

/////////// ARRAYS == SLICES ///////
// Slices references to arrays : They are build on top of arrays
// When slices are passed in the function means, slice affects the array and pass changes into it
// Multiple slices can point to the same array
// a function that only had access to a slice can modify the underlying array

// RAM : Random Access Memory --> This is an address to some data
// eg: address - 0 , data 6; address - 1, data 7
// Slice has [2] int {6,7}

// ADDING SLICE :
//  Copying over an array and then adding/ deleting to another item

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

// /////////////// Get Message Cost ////////////////
// The function takes in a slice of string and return a slice of float64 which is the cost
func getMessageCost(messages []string) []float64 {
	// Create a slice which is a collection of float64 of the length messages
	costs := make([]float64, len(messages))
	// Iterate over each message and calculate cost
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		// Calculate the cost by getting the length of message and then make it to float64 and multiply it by cost
		cost := float64(len(message)) * 0.01
		// Insert the cost in the slice of costs
		costs[i] = cost
	}
	return costs
}

// //////////////// Variadic /////////////
func sum(nums ...int) int {
	// nums is just a slice
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

////////////////// APPEND /////////////////
// Add elements to an existing slice

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	a := primes[1:5]

	fmt.Println("==== SLICES =====")
	fmt.Println(primes)
	fmt.Println("-----------------")
	messages := getMessages()
	for _, msg := range messages {
		fmt.Println(msg)
	}
	fmt.Println("-------VARIADIC----------")
	// The variadic function makes the len and cap of the slice open
	// Any number of input could be send using the variadic
	// WAY 1
	fmt.Println(sum(1, 2, 3))
	// WAY 2 : Spread Operator
	fmt.Println(sum(primes...))

	////////////////// APPEND /////////////////
	// Add elements to an existing slice
	fmt.Println("-------APPEND----------")
	primes = append(primes, 19, 23)
	fmt.Println(primes)

	fmt.Println("-----------------")
	fmt.Println(a)

	fmt.Println("-------MAKE----------")
	// SLICES - MAKE
	// creating a slice without thinking abou the underlying arrays
	// func make([]Int, len, cap)
	// cap : It is the expandable limit or the length of the underlying array
	mySlice := make([]int, 5, 10)
	// Capacity is generally omitted to keep the code simple and it does not have a great impact on the memory
	mySlice2 := make([]int, 3)
	// Initializing a new slice
	myStrings := []string{"I", "Love", "You"}
	fmt.Println(mySlice)
	fmt.Println(mySlice2)
	fmt.Println(myStrings)

	fmt.Println("-------COST OF SENDING MESSAGES----------")
	costPerMessage := getMessageCost(myStrings)
	fmt.Println("Message : Cost")
	for i, msg := range myStrings {
		fmt.Println(msg, " : ", costPerMessage[i])
	}

	fmt.Println("--------ITERATE OVER MESSAGES BASED ON PLAN OF A PERSON---------")
	// Struct of Person
	type Person struct {
		Name string
		Plan string
	}
	// Slice of type Person laid above
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
