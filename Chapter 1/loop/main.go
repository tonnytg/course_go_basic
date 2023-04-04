package main

import "fmt"

var indice int = 0

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("infinite loop")
		if indice == 10 {
			break
		}
		indice = indice + 1
	}

	var books = [3]string{"book1", "book2", "book3"}
	for i, v := range books {
		fmt.Println(i, v)
	}
}