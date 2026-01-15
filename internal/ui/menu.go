// Package ui provides user interface display functions for the calculator application.
// It handles all formatted output to stdout, including welcome messages, menus,
// calculation results, error messages, and goodbye messages.
package ui

import (
	"bufio"
	"fmt"
	"os"
)

// DisplayWelcome shows the welcome message when the calculator starts.
func DisplayWelcome() {
	fmt.Println("Welcome to the Simple Calculator!")
}

// DisplayMenu shows the interactive menu with ASCII borders and operation options.
// The menu displays options 1-5 for calculator operations and exit.
func DisplayMenu() {
	fmt.Println("\n========== Simple Calculator ==========")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exit")
	fmt.Println("=======================================")
	fmt.Print("Enter your choice (1-5): ")
}

// DisplayResult shows the calculation result in the format "a op b = result".
// The operation parameter should be one of: "+", "-", "*", "/"
func DisplayResult(operation string, a, b, result float64) {
	fmt.Printf("%g %s %g = %g\n", a, operation, b, result)
}

// DisplayError shows an error message with "Error: " prefix.
func DisplayError(err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

// DisplayGoodbye shows the goodbye message when the user exits the calculator.
func DisplayGoodbye() {
	fmt.Println("Thank you for using the calculator! Goodbye!")
}

// WaitForEnter displays a prompt and waits for the user to press Enter.
// This is used to pause between operations so users can see results before
// the menu is displayed again.
func WaitForEnter() {
	fmt.Print("\nPress Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
