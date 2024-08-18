package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var secret interface{}

	secret = 42

	var value int
	value = secret.(int) * 10
	fmt.Println(value)

	var info interface{} = &Person{Name: "John Doe", Age: 42}
	var infoPerson = info.(*Person)
	fmt.Println(infoPerson.Name)

	// Dengan menggunakan interface{} kita dapat membuat data yang dapat menampung apa saja

	var fruits = []interface{}{
		[]string{"Apple", "Orange", "Banana"},
		map[string]interface{}{"name":"Date", "total": 1},
		"Melon",
	}

	for _,each := range fruits {
		fmt.Println(each)
	}
		


}