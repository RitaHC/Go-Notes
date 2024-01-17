package main

import (
	"errors"
	"fmt"
)

// This is equivalent of Try/Catch or .then in JavaScript
// Basic Syntax:
// user, err := getUser()
// if err != nil{
// 	fmt.Println(err)
//	return
// } use user here for the aforesaid purpose

func sendSms(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Cant send text over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil

}

func sendSmsToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSms(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSms(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpouse, nil
}

// /////////// Error Interface ////////////
type error interface {
	Error() string
}
type userError struct {
	name string
}

// This is a method based on userError() struct
func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

///////// How to use the error interface writtem above
// func sendSMS (msg, userName string) error {
// 	if !canSendToUser(userName) {
// 		return userError{name: userName}
// 	}
// 	...
// }

// //////////// ERRORS PACKAGE ///////////
// This makes it easy to handle new errors/
// There is no need to create an Error interface or struct now
var err error = errors.New("something went wrong")

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New(("Problem: Y == 0"))
	}
	return x / y, nil
}
func main() {
	// Give the parameters a value
	msgToCustomer := "Hello Customer!"
	msgToSpouse := "Hi Spouse"
	// Call the funstion now with the said parameters
	cost, err := sendSmsToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	/////////////// error package usage ////////
	x := 10.0
	y := 0.0
	result, err := divide(x, y)
	if err != nil {
		fmt.Println("Error message -->", err)
		return
	}

	////////////////// PRINT STATEMENTS ////////
	fmt.Println("------- Errors -------")
	fmt.Printf("Total cost for the couple: $%.4f\n", cost)
	fmt.Println("Error Package", err)
	/////////////// error package print /////////
	fmt.Println("The result of the function divide() is : ", result)

}
