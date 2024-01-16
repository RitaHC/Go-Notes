package main

import "fmt"

// GUARD CLAUSE : These are early returns
// Instead of writing long if and elses, one use a if and if thats not true then return something else directly, without writing else

///////////// STRUCTS (key : value pair) {Objects + Class in JS} ///////////////
// It is a type containing other types

// Car Struct
type Car struct {
	Model  string
	Make   string
	Height int
	Width  int
}

func main() {
	fmt.Println("------ Structs -------")
	// making a car
	myCar := Car{
		Model:  "Sedan",
		Make:   "Toyota",
		Height: 150,
		Width:  100,
	}

	// Print Statements
	fmt.Println("Car Model -> ", myCar.Model)
	fmt.Println("Car Make -> ", myCar.Make)
	fmt.Println("Car Height ->", myCar.Height)
	fmt.Println("Car Width -> ", myCar.Width)

}
