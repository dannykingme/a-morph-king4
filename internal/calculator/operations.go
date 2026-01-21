// Package calculator provides core arithmetic operations for the calculator application.
// It implements pure, stateless functions for addition, subtraction, multiplication,
// and division with proper error handling for division by zero.
package calculator

import "errors"

// Add returns the sum of two float64 values.
// This function corresponds to the add method in the C++ Calculator class.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two float64 values (a - b).
// This function corresponds to the subtract method in the C++ Calculator class.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two float64 values.
// This function corresponds to the multiply method in the C++ Calculator class.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two float64 values (a / b).
// It returns an error if b is zero, as division by zero is not allowed.
// This function corresponds to the divide method in the C++ Calculator class,
// translating the exception-based error handling to Go's idiomatic error return values.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is not allowed!")
	}
	return a / b, nil
}
