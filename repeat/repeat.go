package main

import "fmt"

func main() {
	/*
		x := 10

		if( x < 5){
			fmt.Println("x is less than 5")
		}else {
			fmt.Println("x is greater than or equal to 5")
		}

		if y:=10; y > 5 {
			fmt.Println("y is greater than 5")
		}

		var num int
		fmt.Println("Enter a number:")
		fmt.Scanln(&num)
		if(num < 0){
			fmt.Println("The number is negative")
		}else if (num == 0){
			fmt.Println("The number is zero")
		}else {
			fmt.Println("The number is positive")
		}

		var age int
		fmt.Println("Enter your age:")
		fmt.Scanln(&age)
		if(age < 0){
			panic("Age cannot be negative")
		}
		if ( age < 13){
			fmt.Println("You are a child")
		}else if ( age >= 13 && age < 18){
			fmt.Println("You are a teenager")
		}else{
			fmt.Println("You are an adult")
		}
	*/

	var num1, num2 int
	fmt.Println("Enter two numbers: ")
	_, err := fmt.Scanln(&num1, &num2)
	if err != nil {
		fmt.Println("Invalid inputs")
		return
	}
	if num2 > num1 {
		fmt.Println(num2)
	} else if num1 > num2 {
		fmt.Println(num1)
	} else {
		fmt.Println("Both numbers are equal")
	}

	var num int
	fmt.Println("Enter a number: ")
	_, err = fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	var grade int
	fmt.Println("Enter your grade: ")
	_, err = fmt.Scanln(&grade)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if grade < 0 || grade > 100 {
		fmt.Println("Invalid grade")
		return
	}
	if grade >= 90 && grade <= 100 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else if grade >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	var operator string
	var num11, num22 int
	fmt.Println("Enter a first number: ")
	_, err = fmt.Scanln(&num11)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Println("Enter an operator (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println("Enter a second number: ")
	_, err = fmt.Scanln(&num22)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	switch operator {
	case "+":
		fmt.Println(num11 + num22)
	case "-":
		fmt.Println(num11 - num22)
	case "*":
		fmt.Println(num11 * num22)
	case "/":
		if num22 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Println(num11 / num22)
	default:
		fmt.Println("Invalid operator")
		return
	}

}
