package main

import "fmt"

func getNthFib(number int) int {

	if number <= 1 || number == 0 {
		return 0
	}

	previous := 0
	current := 1
	next := 1

	for number > 2 {
		next = previous + current
		previous = current
		current = next
		number--
	}
	return next
}

func main () {
	fmt.Println(getNthFib(6))
}
