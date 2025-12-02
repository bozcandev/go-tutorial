package main

import "fmt"

func main() {
	var a int = 10
	add(a)
	fmt.Println(a)
	fmt.Println("***")
	addPointer(&a)
	fmt.Println(a)
}

func add(x int) {
	x = x + 10
}

func addPointer(x *int) {
	*x = *x + 10
}
