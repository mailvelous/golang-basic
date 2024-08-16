package main

import (
    "fmt"
)

func main(){
    myIq := 129
    averageIndonesianIq := 78
    livesInIndonesia := true

    if myIq > averageIndonesianIq && livesInIndonesia == true {
        fmt.Println("My iq is higher than average Indonesian")
    } else if myIq < averageIndonesianIq && livesInIndonesia == true {
        fmt.Println("My iq is lower than average Indonesian")
    } else {
        fmt.Println("My iq is average Indonesian")
    }

    switchAnimals("Elephant")
}

func switchAnimals(myAnimals string){

    switch myAnimals {
        case "Elephant":
            fmt.Println("I love eating grass")
        case "Cat":
            fmt.Println("I love eating fish")
        case "Dog":
            fmt.Println("I love eating bones")
        default:
            fmt.Println("does not animals eating")
    }

}

