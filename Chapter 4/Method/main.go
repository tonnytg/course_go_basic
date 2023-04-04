package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func main() {
	fmt.Println("OOP")

	p := Person{}
	p.SetName("John")
	p.SetAge(30)

	fmt.Println(p)
}