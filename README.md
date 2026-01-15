# Simple Calculator (Go)

A command-line calculator application written in Go, migrated from C++.

## Features

- Basic arithmetic operations (addition, subtraction, multiplication, division)
- Interactive menu-driven interface
- Input validation and error handling
- Division by zero protection

## Requirements

- Go 1.21 or higher

## Installation

Clone the repository:

```bash
git clone https://github.com/example/calculator.git
cd calculator
```

## Building the Application

To build the calculator:

```bash
go build -o calculator cmd/calculator/main.go
```

## Running the Application

To run the calculator without building:

```bash
go run cmd/calculator/main.go
```

Or run the built binary:

```bash
./calculator
```

## Usage

1. When you start the application, you'll see a welcome message and menu
2. Choose an operation by entering a number (1-5):
   - 1: Addition
   - 2: Subtraction
   - 3: Multiplication
   - 4: Division
   - 5: Exit
3. Enter two numbers when prompted
4. View the result
5. Press Enter to continue and see the menu again

## Project Structure

```
.
├── cmd/
│   └── calculator/          # Main application entry point
│       └── main.go
├── internal/
│   ├── calculator/          # Core arithmetic operations
│   │   └── operations.go
│   ├── ui/                  # User interface and display functions
│   │   └── menu.go
│   └── input/               # Input handling and validation
│       └── input.go
├── go.mod                   # Go module file
└── README.md
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Dependencies

This project uses only the Go standard library with no external dependencies.

## Source

Migrated from C++ calculator application.
Original source: https://github.com/dannykingme/a-morph-king4.git
