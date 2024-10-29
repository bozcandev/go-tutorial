package main

import "fmt"

func main() {
	//variable decleration
	var name string
	name = "MyName"
	fmt.Println(name)

	//variable decleration and assing a value
	var surname string = "MySurname"
	fmt.Println(surname)

	//short variable decleration and assinging a value - Shortest Way
	age := 18 //I want to come to my 18 years :()
	fmt.Println(age)

	/*
		lets change the variable value - we can change it whenever we want, because it is a variable :)
	*/
	language := "C"
	fmt.Print("First Variable Value:", language)
	language = "GO"
	fmt.Println(",Second Variable Value:", language)
}