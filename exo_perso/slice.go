package main

import "fmt"

func main() {
	notes := make([]int, 3, 5)

	notes[0] = 12
	notes[1] = 15
	notes[2] = 8

	fmt.Println("slice:", notes)
	fmt.Println("len:", len(notes))
	fmt.Println("cap:", cap(notes))

	notes = append(notes, 20)

	fmt.Println("après append:", notes)
	fmt.Println("len:", len(notes))
	fmt.Println("cap:", cap(notes))
}
