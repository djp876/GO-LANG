package main

import "fmt"

func main() {
	fmt.Print("ENTER YOUR NAME:")
	var c string
	fmt.Scanln(&c)
	fmt.Print("WELCOME TO THE CLUB",c)
}
