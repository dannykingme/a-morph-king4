# Makefile for Simple Calculator

# Compiler
CXX = g++

# Compiler flags
CXXFLAGS = -std=c++11 -Wall -Wextra -O2

# Target executable
TARGET = calculator

# Source files
SOURCES = calculator.cpp

# Object files
OBJECTS = $(SOURCES:.cpp=.o)

# Default target
all: $(TARGET)

# Build the executable
$(TARGET): $(OBJECTS)
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(OBJECTS)

# Compile source files to object files
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Clean up build files
clean:
	rm -f $(OBJECTS) $(TARGET)

# Run the calculator
run: $(TARGET)
	./$(TARGET)

# Install dependencies (if needed)
install:
	@echo "No dependencies to install for this simple calculator."

# Help target
help:
	@echo "Available targets:"
	@echo "  all     - Build the calculator (default)"
	@echo "  clean   - Remove object files and executable"
	@echo "  run     - Build and run the calculator"
	@echo "  help    - Show this help message"

.PHONY: all clean run install help
