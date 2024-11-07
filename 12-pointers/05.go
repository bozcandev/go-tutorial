package main

import "fmt"

func main() {
	var number int = 10
	change(&number)
	fmt.Println(number)

	calc_number := 10
	calc(&calc_number)
	fmt.Println(calc_number)

}

func change(variable *int) {
	//change the comming variable value to default 50
	*variable = 50
}

func calc(variable *int) {
	*variable *= 10
}
