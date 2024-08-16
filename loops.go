package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	animals := []string{
		"Elephant",
		"Cat",
		"Dog",
	}

	number := 5

	for i,a := range animals {
		fmt.Println(i, a)
	}

	for number < 10 {
		fmt.Println(number)
		number++
	}

	value := 1
	for {
		fmt.Println("Count that", value)
		value++

		if value == 8 {
			break
		}
	}

	var people []Person
	people = append(people, Person{ Name: "Humble God", Age: 1337 })
	people = append(people, Person{ Name: "Diabolic Addict", Age: 102 })
	people = append(people, Person{ Name: "Taksonomi", Age: 11 })

	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}
} 