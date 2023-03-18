package main

import "fmt"

func noteForConcept(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 7 && n <= 8 {
		return "B"
	} else if n >= 5 && n <= 6 {
		return "C"
	} else if n >= 3 && n <= 4 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(noteForConcept(10))
	fmt.Println(noteForConcept(8))
	fmt.Println(noteForConcept(6))
	fmt.Println(noteForConcept(4))
	fmt.Println(noteForConcept(2))
}
