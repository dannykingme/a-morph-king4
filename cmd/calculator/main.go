// Package main is the entry point for the calculator CLI application.
// It orchestrates the main control loop, integrating calculator operations,
// input validation, and UI display functions to create an interactive
// command-line calculator with robust error handling.
package main

import (
	"fmt"

	"github.com/example/calculator/internal/calculator"
	"github.com/example/calculator/internal/input"
	"github.com/example/calculator/internal/ui"
)

func main() {
	var choice int
	var num1, num2 float64
	var err error

	// Display welcome message
	ui.DisplayWelcome()

	// Main control loop - runs until user selects exit option
	for {
		// Display menu and read user's choice using the input validation module
		ui.DisplayMenu()

		// Read and validate menu choice from stdin
		// The input module handles both type validation (non-numeric) and range validation (1-5)
		choice, err = input.ReadMenuChoice()
		if err != nil {
			// Display the error message from the input module
			fmt.Println(err.Error())
			continue
		}

		// Check if user wants to exit
		if choice == 5 {
			ui.DisplayGoodbye()
			break
		}

		// Read two numbers from user using the input validation module
		// This replaces the inline validation logic with the robust input package
		num1, num2, err = input.ReadTwoNumbers()
		if err != nil {
			// Display the error message from the input module
			fmt.Println(err.Error())
			continue
		}

		// Execute the selected operation
		var result float64

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
		ui.WaitForEnter()
	}
}
