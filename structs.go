package main

import (
	"fmt"
)

// Create Struct -> define methods if any -> declare instance in main -> Then print the output

// Go is not an object based programming as it does not support Classes
// Classes and Inheritance exist only in OBP

// GUARD CLAUSE : These are early returns
// Instead of writing long if and elses, one use a if and if thats not true then return something else directly, without writing else

///////////// STRUCTS (key : value pair) {Objects + Class in JS} ///////////////
// It is a type containing other types

// //////////////// Struct ---- Nesting structs ////////////
// Good Practice is naming structs and not anonymous structs
type Car struct {
	Model      string
	Make       string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

//////////////////////// ANONYMOUS STRUCTS ////////////
// Generally used to define a certain JSON payload
// These are structs created only to make one instance of a certain struct
// Generally structs are created and then later n number of instances can be created out of it

// Alternative way of creating this struct

// type Car struct {
// 	Model      string
// 	Make       string
// 	Height     int
// 	Width      int
// 	// Wheel is a field containing anonymous struct
// 	Wheel struct{
// 		Radius   int
// 		Material string
// 	}
// }

////////////// EMBEDED STRUCT //////////////
// Since it is an embedded struct so we can access the sub-property directly by . notation
// Unlike nested struct where we hav to write the name of the whole struct
// Type of instance for Nested and embedded structs are the same

type Bus struct {
	make  string
	model string
}

type truck struct {
	Bus
	bedSize int
}

// ////////////// Methods on Structs //////////
type Rectangle struct {
	height int
	width  int
}

// area method for struct
// normal functions a name is defined and then inputs or output details are given
// while in Method functions a paramenter is defined forst -> then function name -> output type
func (r Rectangle) area() int {
	return r.width * r.height
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
	myCar.FrontWheel.Radius = 10

	// Instance of a Embedded / Nested structs are the same
	myTruck := truck{
		Bus: Bus{
			make:  "Volvo",
			model: "XYZ",
		},
		bedSize: 20,
	}

	// Only defined fileds need to be filled, methods give output on their own
	r := Rectangle{
		height: 10,
		width:  20,
	}

	// Print Statements
	fmt.Println("Car Model -> ", myCar.Model)
	fmt.Println("Car Make -> ", myCar.Make)
	fmt.Println("Car Height ->", myCar.Height)
	fmt.Println("Car Width -> ", myCar.Width)
	fmt.Println("Car FrontWheel -> ", myCar.FrontWheel)
	fmt.Println("Car BackWheel -> ", myCar.BackWheel)
	//
	fmt.Println("------------------------------")
	fmt.Println("MyTruck Bus make--> ", myTruck.Bus.make)
	// Since it is an embedded struct so we can access the sub-property directly by . notation
	// Unlike nested struct where we hav to write the name of the whole struct
	fmt.Println("MyTruck Bus model --> ", myTruck.model)
	fmt.Println("MyTruck bedsize --> ", myTruck.bedSize)
	// Print the Method
	fmt.Println("Area of the Rectangle is", r.area())

}
