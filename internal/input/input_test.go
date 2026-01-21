package input

import (
	"bufio"
	"strings"
	"testing"
)

// TestReadMenuChoice_Valid tests that valid menu choices (1-5) are accepted
func TestReadMenuChoice_Valid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Choice 1", "1\n", 1},
		{"Choice 2", "2\n", 2},
		{"Choice 3", "3\n", 3},
		{"Choice 4", "4\n", 4},
		{"Choice 5", "5\n", 5},
		{"Choice with spaces", "  3  \n", 3},
		{"Choice with tabs", "\t2\t\n", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			choice, err := ReadMenuChoice()
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
			if choice != tt.expected {
				t.Errorf("Expected choice %d, got %d", tt.expected, choice)
			}
		})
	}
}

// TestReadMenuChoice_InvalidType tests that non-numeric input is rejected with appropriate error
func TestReadMenuChoice_InvalidType(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Letters", "abc\n"},
		{"Mixed alphanumeric", "1a\n"},
		{"Special characters", "!@#\n"},
		{"Empty input", "\n"},
		{"Just spaces", "   \n"},
		{"Decimal number", "2.5\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			choice, err := ReadMenuChoice()
			if err == nil {
				t.Errorf("Expected error for input %q, got choice %d", tt.input, choice)
			}
			if err.Error() != ErrInvalidMenuInput {
				t.Errorf("Expected error %q, got %q", ErrInvalidMenuInput, err.Error())
			}
		})
	}
}

// TestReadMenuChoice_OutOfRange tests that out-of-range values are rejected with appropriate error
func TestReadMenuChoice_OutOfRange(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Zero", "0\n"},
		{"Six", "6\n"},
		{"Negative", "-1\n"},
		{"Large positive", "100\n"},
		{"Large negative", "-999\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			choice, err := ReadMenuChoice()
			if err == nil {
				t.Errorf("Expected error for input %q, got choice %d", tt.input, choice)
			}
			if err.Error() != ErrMenuOutOfRange {
				t.Errorf("Expected error %q, got %q", ErrMenuOutOfRange, err.Error())
			}
		})
	}
}

// TestReadOperand_Valid tests that valid numeric operands are accepted
func TestReadOperand_Valid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Positive integer", "42\n", 42.0},
		{"Negative integer", "-17\n", -17.0},
		{"Positive decimal", "3.14\n", 3.14},
		{"Negative decimal", "-2.5\n", -2.5},
		{"Zero", "0\n", 0.0},
		{"Zero decimal", "0.0\n", 0.0},
		{"Leading zeros", "007\n", 7.0},
		{"Scientific notation", "1.5e2\n", 150.0},
		{"With whitespace", "  123.45  \n", 123.45},
		{"Very small number", "0.0001\n", 0.0001},
		{"Very large number", "999999.99\n", 999999.99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			num, err := ReadOperand("")
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
			if num != tt.expected {
				t.Errorf("Expected %g, got %g", tt.expected, num)
			}
		})
	}
}

// TestReadOperand_Invalid tests that invalid operand input is rejected
func TestReadOperand_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Letters", "abc\n"},
		{"Mixed alphanumeric", "12.3abc\n"},
		{"Special characters", "!@#$\n"},
		{"Empty input", "\n"},
		{"Just spaces", "   \n"},
		{"Multiple numbers", "1 2\n"},
		{"Expression", "1+2\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			num, err := ReadOperand("")
			if err == nil {
				t.Errorf("Expected error for input %q, got number %g", tt.input, num)
			}
			if err.Error() != ErrInvalidNumber {
				t.Errorf("Expected error %q, got %q", ErrInvalidNumber, err.Error())
			}
		})
	}
}

// TestReadTwoNumbers_Valid tests that two valid operands can be read in sequence
func TestReadTwoNumbers_Valid(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedNum1  float64
		expectedNum2  float64
	}{
		{"Two positive integers", "10\n20\n", 10.0, 20.0},
		{"Positive and negative", "5\n-3\n", 5.0, -3.0},
		{"Two decimals", "1.5\n2.7\n", 1.5, 2.7},
		{"Zero and non-zero", "0\n42\n", 0.0, 42.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			num1, num2, err := ReadTwoNumbers()
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
			if num1 != tt.expectedNum1 {
				t.Errorf("Expected first number %g, got %g", tt.expectedNum1, num1)
			}
			if num2 != tt.expectedNum2 {
				t.Errorf("Expected second number %g, got %g", tt.expectedNum2, num2)
			}
		})
	}
}

// TestReadTwoNumbers_FirstInvalid tests error handling when first operand is invalid
func TestReadTwoNumbers_FirstInvalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"First is letters", "abc\n20\n"},
		{"First is empty", "\n15\n"},
		{"First is special chars", "!@#\n10\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			_, _, err := ReadTwoNumbers()
			if err == nil {
				t.Errorf("Expected error for input %q, got none", tt.input)
			}
			if err.Error() != ErrInvalidNumber {
				t.Errorf("Expected error %q, got %q", ErrInvalidNumber, err.Error())
			}
		})
	}
}

// TestReadTwoNumbers_SecondInvalid tests error handling when second operand is invalid
func TestReadTwoNumbers_SecondInvalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Second is letters", "10\nabc\n"},
		{"Second is empty", "15\n\n"},
		{"Second is special chars", "20\n!@#\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace the package-level reader with a test reader
			reader = bufio.NewReader(strings.NewReader(tt.input))

			_, _, err := ReadTwoNumbers()
			if err == nil {
				t.Errorf("Expected error for input %q, got none", tt.input)
			}
			if err.Error() != ErrInvalidNumber {
				t.Errorf("Expected error %q, got %q", ErrInvalidNumber, err.Error())
			}
		})
	}
}

// TestBufferClearing tests that the buffer is properly cleared after errors
// This test simulates the scenario where multiple invalid inputs are entered in sequence
func TestBufferClearing(t *testing.T) {
	// Simulate a sequence of inputs: invalid menu choice -> valid menu choice
	input := "abc\n3\n"
	reader = bufio.NewReader(strings.NewReader(input))

	// First call should fail with invalid input
	_, err1 := ReadMenuChoice()
	if err1 == nil {
		t.Error("Expected first call to return error for invalid input")
	}
	if err1.Error() != ErrInvalidMenuInput {
		t.Errorf("Expected error %q, got %q", ErrInvalidMenuInput, err1.Error())
	}

	// Second call should succeed with valid input
	// This verifies that the buffer was properly cleared after the first error
	choice, err2 := ReadMenuChoice()
	if err2 != nil {
		t.Errorf("Expected second call to succeed, got error: %v", err2)
	}
	if choice != 3 {
		t.Errorf("Expected choice 3, got %d", choice)
	}
}

// TestBufferClearingWithOperands tests buffer clearing with operand reading
func TestBufferClearingWithOperands(t *testing.T) {
	// Simulate a sequence: invalid operand -> valid operand
	input := "invalid\n42.5\n"
	reader = bufio.NewReader(strings.NewReader(input))

	// First call should fail
	_, err1 := ReadOperand("")
	if err1 == nil {
		t.Error("Expected first call to return error for invalid input")
	}

	// Second call should succeed
	num, err2 := ReadOperand("")
	if err2 != nil {
		t.Errorf("Expected second call to succeed, got error: %v", err2)
	}
	if num != 42.5 {
		t.Errorf("Expected 42.5, got %g", num)
	}
}
