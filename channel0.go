package main

import (
	"fmt"
	"golang-basic/helpers"
)

const a = 10

func main() {
	intChan := make(chan int)

	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(a)
	intChan <- randomNumber


	
}