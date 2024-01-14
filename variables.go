package main

import "fmt"

// ADD REPO TO GIT
// git remote set-url origin https://github.com/RitaHC/your-repository-name.git
// git push origin main

func main() {
	// Types of Variables
	// Most used DEFAULT TYPES (int, uint, float64, complex128)
	// Bool

	// string

	// INTEGERES:   int int8 int16 int32 int64
	// + or - Whole numbers
	// UNSIGNED INTEGERES: uint uint8 uint16 uint32 uint64 uintptr
	// ONLY + integers

	// byte // alias for uint8

	// rune // alias for int32
	// represents a Unicode code point

	// float32 float64 : Fractional numbers

	// complex64 complex128 : Imaginary numbers

	// USAGE
	// 1) Byte might be used a lot while marshalling a JSON object to be sent across a network connection.
	// 2) Rune a one character in a string

	// DECLARATION : Declaration and usage are both very important
	// (1)
	number := 4
	// var hasPermission bool
	// var username string
	// (2) Assignmaent operator : assigning an empty operator is a type
	congrats := "Happy Birthday"
	// Declare m ultiple values in the same line
	mileage, company := 9090, "Hameer"

	// Convert data types
	numberChange := float64(number)

	// PRINT
	fmt.Println("Welcome Back!")
	fmt.Println(number)
	fmt.Println(congrats)
	fmt.Println(mileage, company)
	// Printf and use %T/n which tells the type of variable being used
	fmt.Printf("The type of number is %T\n", number)
	fmt.Printf("The type of number is %T\n", numberChange)

	// INTERPOLATION FOR PRINT STATEMENTS
	// %v is a default operator which returns what ever is provided
	// %s Interpolates - a string
	// %d Interpolates - a Integer in Decimal form
	// %f Interpolates - a Decimal
	fmt.Printf("I am %v years old", 10)
	fmt.Printf("My name is %s\n", "Hameer")
	fmt.Printf("I am %d years old \n", 10)
	// Print on;y uptill the first value after the decimal
	fmt.Printf("I am %.1f years old \n", 10.00)

	// Sprintf : Printf prints the output on console while Sprintf stores the value in a given variable
	sprintExample := fmt.Sprintf("Hello to the new %s", "world of SPRINT")
	fmt.Println(sprintExample)

	// CONDITIONALS
	// SYNTAX 1
	if congrats == "Happy Birhday" {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}

	// SYNTAX 2 : initializing a onetime used variable in the if statement itself

	if length := 20; length < 20 {
		fmt.Println("Length is < 20")
	} else {
		fmt.Println("Length is > 20")
	}

	// HOW TO RUN A GO CODE
	// make the folder/file inside the go path
	//
}
