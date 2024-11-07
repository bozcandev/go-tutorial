package main

import "fmt"

type Students struct {
	id      int
	name    string
	surname string
	age     int
}

func main() {

	var Student1 = Students{
		id:      1,
		name:    "Burak",
		surname: "Ozcan",
		age:     38,
	}

	Student2 := Students{2, "Oyku", "Ozcan", 9}

	fmt.Println(Student1)
	fmt.Println("***")
	fmt.Println("id: ", Student1.id, "name: ", Student1.name, "surname: ", Student1.surname, "age :", Student1.age)
	fmt.Println("***")
	fmt.Println(Student2)
	fmt.Println(Student2.name, Student2.surname)
}
