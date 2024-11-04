package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "Name"
	name[1] = "Surname"
	fmt.Println(name)

	var devs = [4]string{"C", "C++", "GO", "PASCAL"}
	fmt.Println(devs)

	//short array decleration
	codes := [3]string{"0", "1", "2"}
	fmt.Println(codes)

	//change the value of an array for specific index - it must be the same type
	codes[0] = "-1"
	fmt.Println(codes)

	//get special array's values
	//example first 2 array's values
	fmt.Println(devs[0:2])

	var arr = [...]string{"negative", "positive", "zero"}
	fmt.Println(arr)

	//len command - gives us the lengh of the arrays

	var comments = [2]string{"go", "lang"}
	fmt.Println(len(comments))
}
