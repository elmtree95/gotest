package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name       string
	age        int
	occupation string
	orientation string
}

var a [100]person

func populateArray() {
	for t := 0; t < 100; t++ {
		a[t] = person{"test" + strconv.Itoa(t), 32 + t, "programmer" + strconv.Itoa(t)}
	}
}

func main() {

	//call func to build array with contents
	populateArray()

	for t := 0; t < 100; t++ {
		fmt.Println("Testing is going well......")
		fmt.Println("Object " + strconv.Itoa(t) + " " + a[t].name + " " + strconv.Itoa(a[t].age))

		fmt.Println("")
	}
}
