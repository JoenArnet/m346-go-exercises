package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := make(map[int]string)

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[320] = "Modul 320:"
	// TODO: replace one
	modules[117] = "Neues Modul 117:"

	fmt.Println(modules)
}
