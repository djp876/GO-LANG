package main

import "fmt"

func main() {
	fmt.Print("Enter First String: ")
	var first int
	fmt.Scanln(&first)
	fmt.Print("Enter Second String: ")
	var second int
	fmt.Scanln(&second)
	fmt.Print(first + second)
}
