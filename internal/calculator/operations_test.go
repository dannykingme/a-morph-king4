package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed signs", -5.0, 3.0, -2.0},
		{"with zero", 5.0, 0.0, 5.0},
		{"large numbers", 1e10, 2e10, 3e10},
		{"decimal numbers", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%g, %g) = %g; expected %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 2.0},
		{"negative numbers", -5.0, -3.0, -2.0},
		{"mixed signs", -5.0, 3.0, -8.0},
		{"with zero", 5.0, 0.0, 5.0},
		{"zero minus positive", 0.0, 5.0, -5.0},
		{"large numbers", 2e10, 1e10, 1e10},
		{"decimal numbers", 5.5, 2.3, 3.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%g, %g) = %g; expected %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 15.0},
		{"negative numbers", -5.0, -3.0, 15.0},
		{"mixed signs", -5.0, 3.0, -15.0},
		{"multiply by zero", 5.0, 0.0, 0.0},
		{"multiply by one", 5.0, 1.0, 5.0},
		{"large numbers", 1e5, 1e5, 1e10},
		{"decimal numbers", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%g, %g) = %g; expected %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
		errorMsg    string
	}{
		{"positive numbers", 6.0, 3.0, 2.0, false, ""},
		{"negative numbers", -6.0, -3.0, 2.0, false, ""},
		{"mixed signs", -6.0, 3.0, -2.0, false, ""},
		{"divide by one", 5.0, 1.0, 5.0, false, ""},
		{"zero divided by positive", 0.0, 5.0, 0.0, false, ""},
		{"decimal result", 5.0, 2.0, 2.5, false, ""},
		{"large numbers", 1e10, 1e5, 1e5, false, ""},
		{"division by zero", 5.0, 0.0, 0.0, true, "Division by zero is not allowed!"},
		{"zero divided by zero", 0.0, 0.0, 0.0, true, "Division by zero is not allowed!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%g, %g) expected error but got none", tt.a, tt.b)
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Divide(%g, %g) error = %q; expected %q", tt.a, tt.b, err.Error(), tt.errorMsg)
				}
				if result != 0 {
					t.Errorf("Divide(%g, %g) returned %g on error; expected 0", tt.a, tt.b, result)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%g, %g) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%g, %g) = %g; expected %g", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}
