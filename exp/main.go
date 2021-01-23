package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	type Dog struct {
		Name   string
		Breed  string
		DogAge int
	}

	type User struct {
		Name string
		Dog  Dog
		Age  int
		Favs []string
		Map  map[int]string
	}

	data := User{
		Name: "Adam Woolhether",
		Dog: Dog{
			Name:   "Max",
			Breed:  "Labrador",
			DogAge: 3,
		},
		Age:  32,
		Favs: []string{"chess", "checkers", "go"},
		Map: map[int]string{
			1: "One",
			2: "Two",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Joey Montana"
	data.Dog.Name = "Choco"
	data.Dog.DogAge = 5
	data.Age = 23
	data.Favs = []string{"baseball", "basketball"}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
