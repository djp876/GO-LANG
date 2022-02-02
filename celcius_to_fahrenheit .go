package main

import "fmt"

func main() {
	var conversion int
	var c int
	var f int
	fmt.Print("THE CONVERSION TO BE DONE\n")
	fmt.Print("1.CELCIUS\t 2.FAHERINHEIT\n")
	fmt.Scanln(&conversion)
	switch conversion {
	case 1:
		fmt.Print("ENTER THE TEMPERATURE IN CELCIUS:")
		fmt.Scanln(&c)
		var out1 = (float32(c) * 1.8) + 32
		fmt.Print(out1)
	case 2:
		fmt.Print("ENTER THE TEMPERATURE IN FAREHINHEIT")
		fmt.Scanln(&f)
		var out = (float32(f) - 32) / 1.8
		fmt.Print(out)
	default:
		fmt.Println("Invalid Operation")

	}
}
