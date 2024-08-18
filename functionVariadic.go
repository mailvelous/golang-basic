package main

import (
	"fmt"
)

func main() {
	avg := calculateAverage(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(avg)
}


func calculateAverage(numbers ...float64) float64 {
	var total float64 = 0
	for _, number := range numbers {
		total += number
	}

	var average float64
	average = total / float64(len(numbers))
	return average
	
}