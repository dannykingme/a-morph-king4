// Package input provides input validation functions for the calculator application.
// It handles reading and validating user input from stdin, including menu choices
// and numeric operands, with proper error handling and buffer management.
package input

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// Error messages that match the semantic meaning of the C++ version
const (
	ErrInvalidMenuInput = "Invalid input! Please enter a number between 1-5."
	ErrMenuOutOfRange   = "Invalid choice! Please select a number between 1-5."
	ErrInvalidNumber    = "Invalid input! Please enter a valid number."
)

// reader is a package-level buffered reader for stdin
// This provides full control over buffer state and ensures reliable input reading
var reader = bufio.NewReader(os.Stdin)

// ReadMenuChoice reads and validates a menu choice from stdin.
// It returns the validated choice (1-5) or an error if the input is invalid.
// This function handles both type errors (non-numeric input) and range errors (out of range).
//
// Error types:
// - ErrInvalidMenuInput: Input is not a valid integer
// - ErrMenuOutOfRange: Input is an integer but not in range 1-5
//
// This function corresponds to the C++ input validation logic at lines 72-86.
func ReadMenuChoice() (int, error) {
	// Read a line of input from stdin
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New(ErrInvalidMenuInput)
	}

	// Trim whitespace and newline characters
	input = strings.TrimSpace(input)

	// Parse the input as an integer
	choice, err := strconv.Atoi(input)
	if err != nil {
		// Input is not a valid integer
		return 0, errors.New(ErrInvalidMenuInput)
	}

	// Validate that the choice is in the valid range (1-5)
	if choice < 1 || choice > 5 {
		return 0, errors.New(ErrMenuOutOfRange)
	}

	return choice, nil
}

// ReadOperand reads and validates a numeric operand from stdin.
// It displays the given prompt, then reads and parses a float64 value.
// Returns the validated number or an error if the input is invalid.
//
// This function corresponds to the C++ getNumbers function logic at lines 44-60,
// adapted to read a single operand at a time for better reusability.
func ReadOperand(prompt string) (float64, error) {
	// Display the prompt
	print(prompt)

	// Read a line of input from stdin
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New(ErrInvalidNumber)
	}

	// Trim whitespace and newline characters
	input = strings.TrimSpace(input)

	// Parse the input as a float64
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// Input is not a valid number
		return 0, errors.New(ErrInvalidNumber)
	}

	return num, nil
}

// ReadTwoNumbers reads and validates two numeric operands from stdin.
// It prompts for "first number" and "second number", returning both values
// or an error if either input is invalid.
//
// This is a convenience function that directly corresponds to the C++ getNumbers
// function at lines 44-60, maintaining the same two-operand reading pattern.
func ReadTwoNumbers() (float64, float64, error) {
	// Read first operand
	num1, err := ReadOperand("Enter first number: ")
	if err != nil {
		return 0, 0, err
	}

	// Read second operand
	num2, err := ReadOperand("Enter second number: ")
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}

// ClearInputBuffer clears any remaining input in the buffer.
// This function corresponds to the C++ clearInput function at lines 28-31.
//
// Note: In our implementation using bufio.Reader.ReadString('\n'), the buffer
// is automatically cleared after each successful read (we read until newline).
// This function is provided for API completeness and future extensibility,
// but in practice, our line-based reading approach already ensures clean buffer state.
func ClearInputBuffer() {
	// With bufio.Reader.ReadString('\n'), we already read until newline,
	// so the buffer is automatically cleared after each read.
	// This function is a no-op in our implementation but maintained for
	// API consistency with the C++ clearInput pattern.
	//
	// If we ever need to clear the buffer explicitly (e.g., after a partial read),
	// we can use reader.ReadString('\n') here.
}
