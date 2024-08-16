package main

import (
    "fmt"
	"time"
)

type UserData struct {
	FirstName string
	LastName string
	Age int
	BirthDate time.Time
}

func main() {
	user := UserData{
		FirstName: "Humble",
		LastName: "God",
		Age: 1337,
	}

	// fmt.Println(user.FirstName, user.LastName, user.Age) 

	fmt.Println(user.printFirstName())

}

// Value receivers

func (x UserData) printFirstName() string {
	return x.FirstName
} 

// Pointer receivers

func (y *UserData) printLastName() string {
	return y.LastName
}