package main

import "fmt"

func main() {

	// variables
	// ====================================

	// full declaration
	var josh string = "josh"

	// short declaration
	josh2 := "josh2"

	fmt.Println(josh)
	fmt.Println(josh2)

	// Constants
	// ====================================

	// full declaration
	const myConst int = 42
	// Short declaration
	const a = 43

	// Array (fixed size must be known at declaration )
	// ====================================

	// An arrray of integers that can hold 3 integers
	grades := [3]int{89, 97, 76}

	fmt.Println(grades)

	// Slices
	// ====================================

	a := []int{1, 2, 3}
	b := make([]int, 3)

	// Structs (objects)
	// ====================================

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

}

// Methods
// ====================================

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
