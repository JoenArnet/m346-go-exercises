package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	// TODO: correct up to index 4 using direct element access
	fibs[3] = fibs[3-1] + fibs[3-2]
	fibs[4] = fibs[4-1] + fibs[4-2]

	fibs = append(fibs, fibs[5-1]+fibs[5-2]) // TODO: replace 0 with the next Fibonacci number
	// TODO: compute three more Fibonacci numbers and append them

	fibs = append(fibs, fibs[6-1]+fibs[6-2])
	fibs = append(fibs, fibs[7-1]+fibs[7-2])
	fibs = append(fibs, fibs[8-1]+fibs[8-2])

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
