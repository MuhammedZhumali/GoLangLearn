package main

import "fmt"


func main() {
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
}