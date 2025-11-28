package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		firstName string
		lastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}

	classA := Class{
		Students: []Student{
			{firstName: "Max", lastName: "Mustermann"},
			{firstName: "Anna", lastName: "Meyer"},
			{firstName: "Lukas", lastName: "Schmidt"},
		},
	}

	classB := Class{
		Students: []Student{
			{firstName: "Laura", lastName: "Fischer"},
			{firstName: "Tim", lastName: "Becker"},
			{firstName: "Sophie", lastName: "Klein"},
		},
	}
	// TODO: declare a map of modules being attended by multiple classes
	modules := make(map[int][]Class)
	modules[114] = []Class{classA, classB}
	modules[117] = []Class{classA}
	modules[346] = []Class{classB}

	// TODO: output everything using fmt.Println()
	for moduleNumber, classes := range modules {
		fmt.Printf("Modul %d:\n", moduleNumber)
		for _, class := range classes {
			for _, student := range class.Students {
				fmt.Println(" -", student.firstName, student.lastName)
			}
		}
	}
}
