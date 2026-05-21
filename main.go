package main

import "fmt"

func main() {

	var a, b float64
	var symbol string

	fmt.Println("=== Continuous Calculator ===")
	fmt.Println("Enter 'q' at any prompt to quit.")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("\nEnter first number (or 'q' to quit): ")
		_, err := fmt.Scan(&symbol)
		if err != nil {
			fmt.Println("Invalid input.")
			break
		}
		if symbol == "q" {
			break
		}

		_, err = fmt.Sscan(symbol, &a)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			break
		}

		fmt.Println("Enter operator (+, -, *, /) or 'q' to quit: ")

		_, err = fmt.Scan(&symbol)
		if err != nil {
			fmt.Println("Invalid operator input.")
			break
		}
		if symbol == "q" {
			break
		}

		var input2 string
		fmt.Println("Enter second number (or 'q' to quit): ")
		_, err = fmt.Scan(&input2)
		if err != nil {
			fmt.Println("Please enter a valid number or 'q' to quit.")
			break
		}
		if input2 == "q" {
			break
		}

		_, err = fmt.Sscan(input2, &b)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			break
		}

		switch symbol {
		case "+":
			add := a + b
			fmt.Println("Result is: ", add)
		case "-":
			min := a - b
			fmt.Println("Result is: ", min)
		case "*":
			mul := a * b
			fmt.Println("Result is: ", mul)
		case "/":
			if b == 0 {
				fmt.Println("Cannot divide by zero")
			} else {
				div := a / b
				fmt.Println("Result is: ", div)
			}
		default:
			fmt.Println("Invalid operator")
		}
	}
	fmt.Println("Goodbye!")

}
