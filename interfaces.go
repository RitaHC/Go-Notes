package main

import "fmt"

// Collection of Method signatures
// An interface is comprised of methods made with reference to some structs
// Further to call the interface a function is needd where the struct is called as a parameter
// Create Inference -> make a Calling function -> Print the function(Pass a struct as a parameter)
// Implemented implicitly
type shape interface {
	area() float64
	perimeter() float64
}

// To print the inference a function needs to be created and then the function needs to be called
func printShapeInfo(s shape) {
	fmt.Println("Area: --> ", s.area())
	fmt.Println("Perimeter: --> ", s.perimeter())
}

type Rect struct {
	height, width float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}
func (r Rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{
		height: 2.0,
		width:  3.0,
	}

	/////////////// PRINT //////////////////////
	fmt.Println("---------- INTERFACES -------")
	fmt.Println("Area of Rectangle", r.area())
	fmt.Println("Perimeter of Rectangle", r.perimeter())
	fmt.Println("------- Inference Function Called -------")
	printShapeInfo(r)
}
