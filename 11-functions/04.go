//named return functions

package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 20))

}

func Sum(x, y int) (sum int) {
	sum = x + y
	return sum
}
