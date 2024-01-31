package main

// Any data type could be used as a Value but not every data type can be used as a map key
// Keys can be something comparable for equality eg: bool, int, string, structs
// What is not comparable for type equality is slices, maps and functions :- These are pointers and not values themselves

import (
	"errors"
	"fmt"
	"strconv"
)

func getCounts(userIDs []string) map[string]int {
	// create a data type for return -> in this case a new map
	counts := make(map[string]int)
	// loop over all of user Ids
	for _, userID := range userIDs {
		// check if the value for a map in the given user already exist
		count := counts[userID]
		// If Yes add it to the value
		count++
		// Now save it to userID
		counts[userID] = count
		// if not then add it as a new entry
	}
	return counts
}

// ///////////////////////////////////////
// Create a function that calculates the sum of all the digits
func calcTotalOfInts(n int) int {
	strNum := strconv.Itoa(n)
	var total int
	// fmt.Println("StrNum : ", strNum)
	fmt.Println("Number : ", n)

	//// If for checking the len of int is more than one
	if len(strNum) > 1 {
		// Make the int into a slice
		digitSlice := make([]int, len(strNum))
		// Iterate over the int and fill digits in slice
		for i, digitStr := range strNum {
			digit, _ := strconv.Atoi(string(digitStr))
			// Now put the values inside the newly created slice
			digitSlice[i] = digit
			total += digitSlice[i]

		}
		fmt.Println("Initial Total : ", total)
		// Now check if len(total) > 1
		// Now for to continue uptill the answer becomes a single digit
		for len(strconv.Itoa(total)) > 1 {
			var newTotal int
			// Iterate over the total and fill digits
			for _, digitStr := range strconv.Itoa(total) {
				digits, _ := strconv.Atoi(string(digitStr))
				newTotal += digits
			}
			total = newTotal
		}
		// repeat to reduce it to a total of 1 digit
		fmt.Println("Final Total :-->", total)
	}
	return total

}

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
	//////////////////////////////////
	names := []string{"Hameer", "Rajat", "Jabar", "Jabar"}
	fmt.Println(getCounts(names))

	////////////////////////////
	fmt.Println("------------------------")

	a := 123
	b := 123456789
	c := 9999999999999982
	// b := len(strconv.Itoa(a))
	// fmt.Println("Type of b : ", reflect.TypeOf(b))
	// if b > 1 {
	// 	fmt.Println("Success")
	// } else {
	// 	fmt.Println("Fail")
	// }
	calcTotalOfInts(a)
	fmt.Println("------------------------")
	calcTotalOfInts(b)
	fmt.Println("------------------------")
	calcTotalOfInts(c)

}
