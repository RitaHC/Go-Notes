package main

import "fmt"

// Creating new functions and minizing the use of our main

// 1 Subtract
func sub(x int, y int) int {
	// Declare the function name (with the type of variable) type of result
	return x - y
	// Then return the output
}
func concat(str1 string, str2 string) string {
	return str1 + str2
}

////////////////////// POINTERS ////////////////
// The concept of pointers exist in Go because we pass variables by value and not by reference in go
// A pointer takes a code to a specific place in the memory
// Multiple variables mayclaim to be at the same location in the memory
// Also sometimes we may have copies of the same value

func incrementSends(sendSoFar int, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar
}

// IGNORING RETURN VALUES
func getNames() (string, string) {
	return "John", "Doe"
}

// Prefer sending Explicit return instead of inplicit returns
// ///////////// Naming the return values ://///////////////
func returns() (x, y string) {

	return "Return", "Return Again"
}

func main() {
	fmt.Println("----- FUNCTIONS ------")
	fmt.Println("Func Sub : ", sub(20, 10))
	fmt.Println("Func Concat", concat("hello", "world"))
	/////// Pointer ///////
	sendSoFar := 20
	const sendToAdd = 30
	sendSoFar = incrementSends(sendSoFar, sendToAdd)
	fmt.Println("Func Pointer Increment Sends", sendSoFar)
	// Ignore Return value
	firstName, _ := getNames()
	fmt.Println("Func GetName", firstName)
	return1, _ := returns()
	fmt.Println("Func Returns : --> ", return1, " --> ")

}
