package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	repeatTo100()

}
func repeatTo100() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		if i == 50 {
			//break to loop and finish the application
			break
		}

	}
}
