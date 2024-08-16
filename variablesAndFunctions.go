package main

import (
    "fmt"

)

func main() {
	var sayHello string
	sayHello = "Hello"

	fmt.Println(sayHello)

	sayHelloToTheWorld, sayHellotoAnotherWorld := sayHelloWorld()
	fmt.Println(sayHelloToTheWorld, sayHellotoAnotherWorld)
}

func sayHelloWorld() (string, string) {
	return "Hello World", "Hello another World"
}