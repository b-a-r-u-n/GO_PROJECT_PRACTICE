package main

import (
	"errors"
	"fmt"
)

func calculate(num1, num2 float64, operation string) (float64, error) {
	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			// fmt.Println("Second number must be greater than 0")
			return result, errors.New("Cannot divide by zero")
		}
		result = num1 / num2
	default:
		// fmt.Println("Invalid operatin")
		return result, errors.New("Invalid operatin")
	}
	return result, nil
}

func main() {
	fmt.Println("========= Go Calculator =========")
	var num1, num2 float64

	//                              METHOD - 1
	// fmt.Printf("Enter the first number: ")
	// fmt.Scan(&num1)
	// fmt.Printf("You entered: %.2f\n", num1)

	// fmt.Printf("Enter the second number: ")
	// fmt.Scan(&num2)
	// fmt.Printf("You entered: %.2f\n", num2)

	// var operation string
	// fmt.Printf("Enter the operation you want to do (+, -, *, /): ")
	// fmt.Scan(&operation)

	// result, err := calculate(num1, num2, operation)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }

	//                              METHOD - 2
	// choice := "Y"
	// for {
	// 	if choice == "Y" || choice == "y" {
	// 		fmt.Printf("Enter the first number: ")
	// 		fmt.Scan(&num1)
	// 		fmt.Printf("You entered: %.2f\n", num1)

	// 		fmt.Printf("Enter the second number: ")
	// 		fmt.Scan(&num2)
	// 		fmt.Printf("You entered: %.2f\n", num2)

	// 		var operation string
	// 		fmt.Printf("Enter the operation you want to do (+, -, *, /): ")
	// 		fmt.Scan(&operation)

	// 		result, err := calculate(num1, num2, operation)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("Result =", result)
	// 		}

	// 		fmt.Println("Do you want to continue? (y/n)")
	// 		fmt.Scan(&choice)
	// 	} else if choice == "N" || choice == "n" {
	// 		break
	// 	} else {
	// 		fmt.Println("Please enter a valid choice (y/n)")
	// 		fmt.Scan(&choice)
	// 	}
	// }

	//                              METHOD - 3
	// fmt.Println("1. Calculate")
	// fmt.Println("2. Exit")

	// var menuChoice string
	// fmt.Printf("Enter your choice: ")
	// fmt.Scan(&menuChoice)

	// switch menuChoice {
	// case "1":
	// 	choice := "Y"
	// 	for {
	// 		if choice == "Y" || choice == "y" {
	// 			fmt.Printf("Enter the first number: ")
	// 			fmt.Scan(&num1)
	// 			fmt.Printf("You entered: %.2f\n", num1)

	// 			fmt.Printf("Enter the second number: ")
	// 			fmt.Scan(&num2)
	// 			fmt.Printf("You entered: %.2f\n", num2)

	// 			var operation string
	// 			fmt.Printf("Enter the operation you want to do (+, -, *, /): ")
	// 			fmt.Scan(&operation)

	// 			result, err := calculate(num1, num2, operation)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			} else {
	// 				fmt.Println("Result =", result)
	// 			}

	// 			fmt.Println("Do you want to continue? (y/n)")
	// 			fmt.Scan(&choice)
	// 		} else if choice == "N" || choice == "n" {
	// 			break
	// 		} else {
	// 			fmt.Println("Please enter a valid choice (y/n)")
	// 			fmt.Scan(&choice)
	// 		}
	// 	}
	// case "2":
	// 	return
	// default:
	// 	fmt.Println("Choose again..")
	// 	fmt.Scan(&menuChoice)
	// }

	//                              METHOD - 4
	fmt.Println("1. Calculate")
	fmt.Println("2. Exit")

	var menuChoice string
	fmt.Printf("Enter your choice: ")
	fmt.Scan(&menuChoice)

	for {
		switch menuChoice {
		case "1":
			fmt.Printf("Enter the first number: ")
			fmt.Scan(&num1)
			fmt.Printf("You entered: %.2f\n", num1)

			fmt.Printf("Enter the second number: ")
			fmt.Scan(&num2)
			fmt.Printf("You entered: %.2f\n", num2)

			var operation string
			fmt.Printf("Enter the operation you want to do (+, -, *, /): ")
			fmt.Scan(&operation)

			result, err := calculate(num1, num2, operation)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result =", result)
				fmt.Printf("\n\n")
			}

			fmt.Println("1. Calculate")
			fmt.Println("2. Exit")

			fmt.Printf("Enter your choice: ")
			fmt.Scan(&menuChoice)
		case "2":
			fmt.Printf("\nGoodbye!!")
			return
		default:
			fmt.Println("Choose again...")
			fmt.Scan(&menuChoice)
		}
	}
}
