package main

import "fmt"

func main() {
	var operator string
	var out float32
	fmt.Print("Please enter First number: ")
	var num1 float32
	fmt.Scanln(&num1)
	fmt.Print("Please enter Second number: ")
	var num2 float32
	fmt.Scanln(&num2)
	fmt.Print("Please enter Operator (+,-,/,*):")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		out = num1 + num2
	case "-":
		out = num1 - num2
	case "*":
		out = num1 * num2
	case "/":
		out = num1 / num2
	default:
		fmt.Println("Invalid Operation")
	}
	fmt.Println(num1, operator, num2, "=", out)
}
