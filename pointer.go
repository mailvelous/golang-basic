package main

import (
    "fmt"
)


func main() {
	var a string = "black"
	fmt.Println(a)

	var b *string
	b = &a
	fmt.Println(b)

	// variabel pointer tidak bisa menampung nilai yang bukan pointer
	// (&) is for access address memory
	// (*) is for access value

	x := 5
	y := 10
	calculate(&x, &y)

	//

	var number = 4
	fmt.Println("Number before : ", number)

	changeValue(&number, 10)
	fmt.Println("Number after : ", number)

}

func calculate(x, y *int) {
	fmt.Println(*x + *y)
}

func changeValue (original *int, value int) {
	*original = value
} 