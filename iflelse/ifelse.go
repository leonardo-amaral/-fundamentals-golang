package main

import "fmt"

func displayResults(value float64) {
	if value > 7 {
		fmt.Println("Passed with value: ", value)
	} else {
		fmt.Println("Reproved with value: ", value)
	}
}

func main() {
	displayResults(7.1)
	displayResults(6.9)
}
