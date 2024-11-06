package main

import "fmt"

func main() {

	var cars = map[int]string{
		0: "Alfa Romeo",
		1: "BMW",
		2: "Mercedes",
	}
	fmt.Println(cars)
	fmt.Println("***")
	fmt.Println(cars[0])
	fmt.Println("***")
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	fmt.Println("***")
}
