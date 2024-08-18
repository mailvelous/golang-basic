package helpers

import (
	"fmt"
	"math/rand"
)

func Hello(name string) string {

	return fmt.Sprintf("Hello %s", name)
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}