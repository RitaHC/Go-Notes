package main

// RAM is where all teh memory is stored
// Pointers is the position in RAM where the said data is stored
// & -> Creates a pointer i.e. a reference
// * ->  (1) Defines the type of pointer(int, string etc) & (2) Also to deference a pointer (create new reference to the value)
// Deference means chaning the underlying value (refer to the z in the example())

import (
	"fmt"
	"strings"
)

// Example
func example() {
	// x stores value 5
	x := 5
	// y stores value of x
	y := x
	// z stores as value, location of x : z is a pointer
	z := &x
	// *  - Deference operator
	// Now without altering the value of x directly, we are changing its value
	// by changing the value of z, we've changed the value of x
	// So now x = 6
	*z = 6
	fmt.Println(x, y, z)
}

// ////////////// Practice 1 /////////////
type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// ////////////// Practice 2 /////////////
func removeProfanity(message *string) {
	// create a reference to messages
	messageVal := *message
	// now change the value of message using pointer reference created above
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	*message = messageVal
}

// ////////////// Pointers in Methods /////////////
// pointers are used a lot in methods as they alter the values of thestruct themselves

type Car struct {
	color string
}

func (c *Car) setColor(color string) {
	c.color = color
}

func main() {
	fmt.Println("----- Example -------")
	example()
	fmt.Println("------ Practice 1 ------")
	m := Message{
		Recipient: "Hameer Chauhan",
		Text:      "Hammu, please come home soon",
	}
	sendMessage(m)
	fmt.Println("------ Pointers in Methods ------")
	car := Car{
		color: "white",
	}
	car.setColor("blue")
	fmt.Println(car.color)

}
