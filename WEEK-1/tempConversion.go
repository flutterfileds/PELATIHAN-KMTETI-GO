package main

import "fmt"

func main() {

	var choice int
	var cels float64

	fmt.Println("Choose the target unit for conversion:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Print("Enter your choice : ")
	fmt.Scanln(&choice)

	fmt.Print("Enter the temperature in Celsius: ")
	fmt.Scanln(&cels)

	switch choice {
	case 1:
		// Celsius to Fahrenheit
		fahr := (cels * 9 / 5) + 32
		fmt.Printf("Temperature in F : %.2f", fahr)
	case 2:
		// Celsius to Kelvin
		kelv := cels + 273.15
		fmt.Printf("Temperature in K : %.2f", kelv)
	case 3:
		// Celsius to Reamur
		ream := cels * 0.8
		fmt.Printf("Temperature in R : %.2f", ream)
	default:
		fmt.Println("Invalid. Please choose a number between 1 and 3.")
	}
}
