// multiple return functions
package main

import "fmt"

func main() {
	sum, difference := SumMultiple(10, 20)
	fmt.Println(sum, "|", difference)
}

func SumMultiple(x, y int) (int, int) {
	return x + y, x - y
}
