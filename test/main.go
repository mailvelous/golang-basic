package main

import (
	"errors"
	"fmt"
)

func main() {
	div, err := Divide(5, 1)

	fmt.Println(div, err)
}

func Divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}
	result = x / y
	return result, nil
}
