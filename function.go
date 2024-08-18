package main

import "fmt"

func main() {
	var names = []string{"Humble", "God"}

	printMessage("Hello", names[:1])

	sum := add(1, 2)
	fmt.Println("Sum some numbers:", sum)
}

func printMessage(message string, names []string) {
	nameString := names
	fmt.Println(message, nameString)
}

// function return value

func add(x int, y int) int {

	return x + y
}

