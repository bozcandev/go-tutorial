package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() int {
	return r.height * r.width
}

type Cars struct {
	id   int
	name string
}

func (c Cars) GetCars() (int, string) {
	return c.id, c.name
}

type Avarage struct {
	numbers []int
}

func (a Avarage) GetAvarage() int {
	total := 0
	for _, k := range a.numbers {
		total += k

	}
	return total / len(a.numbers)
}

func main() {
	calculation := Rectangle{width: 5, height: 20}
	fmt.Println(calculation.Area())

	fmt.Println("***")

	names := Cars{
		id:   1,
		name: "Alfa Romeo",
	}

	fmt.Println(names.GetCars())

	fmt.Println("***")

	numbers := Avarage{[]int{10, 20, 30}}
	fmt.Println(numbers.GetAvarage())
}
