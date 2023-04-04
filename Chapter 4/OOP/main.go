package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("OOP")

	p := Person{}
	p.Name = "John"
	p.Age = 30

	fmt.Println(p)
}