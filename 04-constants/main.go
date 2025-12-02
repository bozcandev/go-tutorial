package main

import "fmt"

// constant values can not change like a variable
const version = 0.1

func main() {
	fmt.Println("Application Version is : ", version)
	// we can not change constant values
	// version = 0.2
}
