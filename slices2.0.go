package main

// Slices can hold other slices
//

import (
	"fmt"
	// "reflect"
)

type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	// Create a new slice to store the total cost per day
	costByDay := []float64{}
	// Iterate over the slice
	for i := 0; i < len(costs); i++ {
		// cost is the cost per day
		cost := costs[i]
		// Now nest over the current iterated object
		// Since the result is an object so we can use . notation to access the day/values
		for cost.day >= len(costByDay) {
			// Append to the new slice
			// if the day does not exist then just add 0.0
			costByDay = append(costByDay, 0.0)
		}
		// If the day already exist, then go back to the index and add the value at that index
		costByDay[cost.day] += cost.value

	}

	for _, eachValue := range costByDay {
		fmt.Println(eachValue)
	}
	return costByDay

}

// ///////////// SLICES INSIDE SLICES ///////////////
func createMatrix(rows, cols int) [][]int {
	// rows and cols determine the length of rows and columns in form of int
	// create a return slice
	matrix := make([][]int, 0)
	// Create a nested loop for rows and columns
	for i := 0; i < rows; i++ {
		// Create a slice of row
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			// create col by appending rows
			row = append(row, i*j)
		}
		// Append row to matrix, the main return func
		matrix = append(matrix, row)
	}
	return matrix
}

func main() {
	fmt.Println("---------------")
	// Create a instance of the struct cost as a slice
	costSlice := []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}
	getCostByDay(costSlice)
	fmt.Println("---------------")
	// Print each row on a different line
	matrixSlice := createMatrix(10, 10)
	for _, result := range matrixSlice {
		fmt.Println(result)
	}

}
