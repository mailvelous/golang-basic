package main

import (
	"fmt"
)

func main() {
	area, perimeter := calculateRectangle(10, 20)

	fmt.Printf("Area of rectangle is %d and perimeter is %d", area, perimeter)
	
}

func calculateRectangle(length, width int) (int, int) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}