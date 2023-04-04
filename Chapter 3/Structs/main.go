package main

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Structs")

	p := Person{}
	p.Name = "John"
	p.Age = 30

	fmt.Println(p)
}