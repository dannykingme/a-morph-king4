// Package main is the entry point for the calculator CLI application.
// It orchestrates the main control loop, integrating calculator operations
// and UI display functions to create an interactive command-line calculator.
package main

import (
	"fmt"

	"github.com/example/calculator/internal/calculator"
	"github.com/example/calculator/internal/ui"
)

func main() {
	var choice int
	var num1, num2 float64

	// Display welcome message
	ui.DisplayWelcome()

	// Main control loop - runs until user selects exit option
	for {
		// Display menu and read user's choice
		ui.DisplayMenu()

		// Read menu choice from stdin
		// If input is invalid (not an integer), fmt.Scan returns error
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Invalid input! Please enter a number between 1-5.")
			// Clear the invalid input from the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		// Check if user wants to exit
		if choice == 5 {
			ui.DisplayGoodbye()
			break
		}

		// Validate choice is in valid range (1-4 for operations)
		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice! Please select a number between 1-5.")
			continue
		}

		// Read two numbers from user
		fmt.Print("Enter first number: ")
		if _, err := fmt.Scan(&num1); err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			// Clear the invalid input from the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		fmt.Print("Enter second number: ")
		if _, err := fmt.Scan(&num2); err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			// Clear the invalid input from the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		// Execute the selected operation
		var result float64
		var err error

		switch choice {
		case 1:
			result = calculator.Add(num1, num2)
			ui.DisplayResult("+", num1, num2, result)
		case 2:
			result = calculator.Subtract(num1, num2)
			ui.DisplayResult("-", num1, num2, result)
		case 3:
			result = calculator.Multiply(num1, num2)
			ui.DisplayResult("*", num1, num2, result)
		case 4:
			result, err = calculator.Divide(num1, num2)
			if err != nil {
				// Handle division by zero error
				ui.DisplayError(err)
			} else {
				ui.DisplayResult("/", num1, num2, result)
			}
		}

		// Wait for user to press Enter before showing menu again
		// Note: In Milestone 1, the interaction between fmt.Scan and WaitForEnter
		// may not be perfect - this will be refined in Milestone 2
		ui.WaitForEnter()
	}
}
