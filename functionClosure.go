package main

import (
	"fmt"
)

func main() {
	incrementByTen := makeIncrement(10)

	fmt.Println(incrementByTen())
	fmt.Println(incrementByTen())

	// function clposure as variable

	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)


	// function closure fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//function closure as return value

func makeIncrement(increment int) func() int {
	total := 0

	return func() int {
		total += increment
		return total
	}
}

//function closure as variable

//function closure fibonacci

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b

		result := a
		return result
	}
}

