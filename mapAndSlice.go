package main

import (
    "fmt"
)

type FootballPlayer struct {
	Name string
	Club string
	Nationality string
}

func main() {

	// Map

	player := make(map[string]FootballPlayer)

	Cristiano := FootballPlayer {
		Name: "Cristiano Ronaldo",
        Club: "Al-Nassr",
        Nationality: "Portuguese",
	}

	player["Cristiano Ronaldo"] = Cristiano


	fmt.Println(player["Cristiano Ronaldo"].Name)

	//Slice


	animals := []string{"Monkey", "Cow", "Fish", "Chicken"}

	animals = append(animals, "Elephant")
 
	mammalsAnimals := animals[0:2]

	fmt.Println(mammalsAnimals)
	fmt.Println(animals)

	for i, animal := range animals {
		fmt.Printf("%d: %s\n", i, animal)
	}
}