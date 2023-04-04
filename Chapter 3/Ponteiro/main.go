package main

import "fmt"

func ChangeValue1(b int) {
	b = 20
}

func ChangeValue2(b *int) {
	*b = 20
}

func main() {
	fmt.Println("Ponteiros")

	var a int = 10
	ChangeValue1(a)
	fmt.Println("Change Value without address:", a)

	ChangeValue2(&a)
	fmt.Println("Change Value with address:", a)

}