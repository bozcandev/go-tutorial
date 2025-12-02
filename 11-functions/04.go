//named return functions

package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 20))

}

func Sum(x, y int) (sum int) {
	// sum is the variable name
	sum = x + y
	return sum
}
