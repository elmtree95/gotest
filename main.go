package main

import (
	"fmt"
)


type  person struct{
	name string
	age int
	occupation string
}

func main() {
	p := person{"test", 32, "programmer"}

	fmt.Println("Testing is going well......")
	fmt.Println("First object " + p.name)
}
