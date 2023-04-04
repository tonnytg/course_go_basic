package main


func ShowMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Functions")

	x := func() {
		fmt.Println("Hello")
	}

	x()

	ShowMessage("Hello World")
}