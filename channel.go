package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		message <- "hello world"
	}()

	msg := <-message
	fmt.Println(msg)

}