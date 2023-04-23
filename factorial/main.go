package main

import "fmt"

func faktorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * faktorial(number-1)
}
func main() {
	fmt.Print(faktorial(5))
}
