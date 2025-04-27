package main

import "fmt"

func main() {
	var num1, num2 int
	var op string
	flag := true
	for flag {
		fmt.Println("Enter '+' for Addition")
		fmt.Println("Enter '-' for Subtraction")
		fmt.Println("Enter '*' for Multiplication")
		fmt.Println("Enter '/' for Division")
		fmt.Println("Enter '!' for Exit")

		fmt.Print("Enter Your Choice:")
		fmt.Scan(&op)
		//enter 2 Elements
		fmt.Print("Enter 1st Number:")
		fmt.Scan(&num1)
		fmt.Print("Enter 2nd Number:")
		fmt.Scan(&num2)

		switch op {
		case "+":
			result := Addition(num1, num2)
			fmt.Println("Result of Addition is:", result)
		case "-":
			result := Subtraction(num1, num2)
			fmt.Println("Result of Subtraction is:", result)
		case "*":
			result := Multiply(num1, num2)
			fmt.Println("Result of Multiplication is:", result)
		case "/":
			result, err := Division(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result of Division is:", result)

			}
		case "!":
			fmt.Println("Exiting the program.")
			flag = false
		}

	}

}
