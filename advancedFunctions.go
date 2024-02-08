package main

// Higher order functions OR First class functions
// calling a function inside another function as a parameter
// Using functions as data

// 1) Used for HTTP API handlers
// 2) Pub/Sib handlers
// 3) Onclick handlers

// Calling of Higher order function
// FIRST CALL ::::: the Func1() which is the main function and store the value in a VARIABLE,
// WITH PARAMETER : if their an input parameter
// WITHOUT PARAMETER : if there is no input parameter
// SECOND CALL :::: call the VARIABLE from the 1st Call with the return value of the first function || or the input parameter for the 2nd function

import (
	"fmt"
)

// /////////// Practice 1 ////////////
// STEP 1
func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}

// STEP 2
func aggregate(a, b, c int, arithmetic func(x, y int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// /////////// Practice 2 ////////////
// CURRYING: Takes a function as input and returns a new function
func selfMath(mathFunc func(a, b int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

//////////// Practice 3 : Custom Logger //////

func getLogger(formatter func(string, string) string) func(string, string) {
	// First return the main function
	return func(a string, b string) {
		// Now return the output for the inner function
		fmt.Println(formatter(a, b))

	}
}

// Defer Keyword
// execute some function only at the end of some other function

// Closures
// A closure function references variables from outside its own function body.
// The function will access and assign to the reference variables

// //////////////// Practice 4 ///////////////
func closureExample() func(string) string {
	// closure variable declared
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

// //////////////// Practice 5 ///////////////

func sumClosure() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	fmt.Println("-----Practice 1 -------")
	fmt.Println("ADD Aggregate Function Called : ", aggregate(2, 2, 3, add))
	fmt.Println("MULTIPLY Aggregate Function Called : ", aggregate(1, 2, 3, multiply))
	fmt.Println("------Practice 2 ------")
	squareFunc := selfMath(multiply)
	fmt.Println(squareFunc(5))
	addFunc := selfMath(add)
	fmt.Println(addFunc(5))
	fmt.Println("-----PRACTICE 3 -------")
	// Create an input function
	logger := getLogger(func(a, b string) string {
		return fmt.Sprintf("Formatted: %s %s", a, b)
	})
	// Now call the hoghre order function
	logger("Hello", "World")
	fmt.Println("-----PRACTICE 4 -------")
	concatterFunc := closureExample()
	concatterFunc("Rita")
	concatterFunc("and")
	concatterFunc("Hameer")
	concatterFunc("are")
	concatterFunc("meant")
	concatterFunc("for")
	concatterFunc("each")
	concatterFunc("other")

	fmt.Println(concatterFunc("!"))

	fmt.Println("-----PRACTICE 5 -------")
	calculateTotal := sumClosure()
	calculateTotal(1)
	calculateTotal(1)
	calculateTotal(1)
	calculateTotal(1)
	calculateTotal(1)

	fmt.Println("Total of the numbers : ", calculateTotal(0))

}
