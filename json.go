package main

import (
	"fmt"
	"encoding/json"
)

type LiverpoolPlayer struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Position string `json:"position"`
	HasDog bool `json:"has_dog"`

}

type Superhero struct {
	Name string
	Power string
	Age int
}

func main() {
	//Decode JSON ke array object

	myJsonLiverpool := `
[
    {
        "first_name": "Luis",
        "last_name": "Diaz",
        "position": "Attacker",
        "has_dog": true
    },
    {
        "first_name": "Alexis",
        "last_name": "MacAllister",
        "position": "Midfielder",
        "has_dog": false
    }
]`



	var unmarshalledLiverpoolPlayer []LiverpoolPlayer

	err := json.Unmarshal([]byte(myJsonLiverpool), &unmarshalledLiverpoolPlayer)
	if err != nil {
		fmt.Println(err.Error())
	}


	fmt.Println(unmarshalledLiverpoolPlayer[0].FirstName)
	
//Encode Object ke JSON stringt
	var superhero []Superhero

	s1 := Superhero{
		Name: "Spiderman",
		Power: "Web",
		Age: 30,
	}

	superhero = append(superhero, s1)

	s2 := Superhero{
		Name: "Batman",
		Power: "Money",
		Age: 40,
	}

	
	superhero = append(superhero, s2)

	
	var jsonDataSuperhero, err1 = json.MarshalIndent(superhero, "", "		")
	if err1 != nil {
        fmt.Println(err1.Error())
		return
    }

	fmt.Println(string(jsonDataSuperhero))


}

