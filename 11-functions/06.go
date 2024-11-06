package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 20, 40, -70))
}

func Sum(numbers ...int) int {
	totally := 0
	//in this line also we can use range function to get values from "numbers"
	for i := 0; i < len(numbers); i++ {
		totally += numbers[i]
	}
	return totally
}
