package main

import "fmt"

func main() {

	var carNames = []string{"Alfa Romeo", "BMW"}
	originCountry := []string{"Italy", "Germany"}

	fmt.Println(carNames)
	fmt.Println(originCountry)

	fmt.Println("******************")

	carNames = append(carNames, "Maserati")
	fmt.Println(carNames)

	//for-range loop for slices..
	for i, v := range carNames {
		fmt.Println(i, v)

	}
}
