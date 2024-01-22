package main

import "fmt"

// There is NO WHILE LOOP in Go
// The For loop can at times function like the while loop

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}

	return totalCost
}

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func printPrime(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

func main() {
	fmt.Println("-------- Loops -----")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("*******************")
	numMessages := 4
	totalCost := bulkSend(numMessages)
	fmt.Println("Total cost from Bulkmsg() -->", totalCost)

	fmt.Println("*******************")
	// For loop working as While loop
	totalNum := 5
	for totalNum < 10 {
		totalNum++
		fmt.Println(totalNum)
	}
	max := 13

	fizzbuzz()
	fmt.Println("==================")
	printPrime(max)
}
