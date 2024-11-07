package main

import "fmt"

func main() {
	var a int = 0
	var b *int = &a

	fmt.Println("Pointer Address : ", &a, ", Value : ", a)
	fmt.Println("Pointer Address : ", &b, ", Value : ", *b)
	fmt.Println("***")

	a = 1

	fmt.Println("a Variable Value is : ", a)
	fmt.Println("b Variable Value is : ", *b)

}
