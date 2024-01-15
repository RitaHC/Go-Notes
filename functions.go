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

func main() {
	fmt.Println("----- FUNCTIONS ------")
	fmt.Println("Func Sub : ", sub(20, 10))
	fmt.Println("Func Concat", concat("hello", "world"))
}
