// send slice to func
package main

import "fmt"

func main() {
	var values = []int{1, 2, 3}
	// add new value for slice
	values = append(values, 10)
	fmt.Println(Sum(values))
}

func Sum(numbers []int) int {
	totally := 0
	for _, v := range numbers {
		totally += v
	}
	return totally
}
