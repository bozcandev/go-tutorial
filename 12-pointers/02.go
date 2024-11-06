package main

import "fmt"

func main() {

	var number1 int = 0
	// declare a "p" pointer varible and set it "number1" 's value to it
	var p *int = &number1

	fmt.Println("Number 1 variable value is : ", number1)
	fmt.Println("Number 1 variable pointer address is : ", &number1)
	fmt.Println("P variable's pointer address is : ", &p)
	fmt.Println("Access to number1's pointer address and get it's data : ", *p)
	fmt.Println("***")
	//change number1 variable's value
	number1 = 10
	fmt.Println("Accessing to x variable's value", number1)
	*p = 202
	fmt.Println("Again accessing to x variable's value", number1)

}
