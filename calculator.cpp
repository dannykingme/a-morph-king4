#include <iostream>
#include <limits>

using namespace std;

class Calculator {
public:
    double add(double a, double b) {
        return a + b;
    }

    double subtract(double a, double b) {
        return a - b;
    }

    double multiply(double a, double b) {
        return a * b;
    }

    double divide(double a, double b) {
        if (b == 0) {
            throw invalid_argument("Division by zero is not allowed!");
        }
        return a / b;
    }
};

void clearInput() {
    cin.clear();
    cin.ignore(numeric_limits<streamsize>::max(), '\n');
}

void displayMenu() {
    cout << "\n========== Simple Calculator ==========" << endl;
    cout << "1. Addition (+)" << endl;
    cout << "2. Subtraction (-)" << endl;
    cout << "3. Multiplication (*)" << endl;
    cout << "4. Division (/)" << endl;
    cout << "5. Exit" << endl;
    cout << "=======================================" << endl;
    cout << "Enter your choice (1-5): ";
}

bool getNumbers(double& num1, double& num2) {
    cout << "Enter first number: ";
    if (!(cin >> num1)) {
        cout << "Invalid input! Please enter a valid number." << endl;
        clearInput();
        return false;
    }

    cout << "Enter second number: ";
    if (!(cin >> num2)) {
        cout << "Invalid input! Please enter a valid number." << endl;
        clearInput();
        return false;
    }

    return true;
}

int main() {
    Calculator calc;
    int choice;
    double num1, num2, result;

    cout << "Welcome to the Simple Calculator!" << endl;

    while (true) {
        displayMenu();

        if (!(cin >> choice)) {
            cout << "Invalid input! Please enter a number between 1-5." << endl;
            clearInput();
            continue;
        }

        if (choice == 5) {
            cout << "Thank you for using the calculator! Goodbye!" << endl;
            break;
        }

        if (choice < 1 || choice > 5) {
            cout << "Invalid choice! Please select a number between 1-5." << endl;
            continue;
        }

        if (!getNumbers(num1, num2)) {
            continue;
        }

        try {
            switch (choice) {
                case 1:
                    result = calc.add(num1, num2);
                    cout << num1 << " + " << num2 << " = " << result << endl;
                    break;
                case 2:
                    result = calc.subtract(num1, num2);
                    cout << num1 << " - " << num2 << " = " << result << endl;
                    break;
                case 3:
                    result = calc.multiply(num1, num2);
                    cout << num1 << " * " << num2 << " = " << result << endl;
                    break;
                case 4:
                    result = calc.divide(num1, num2);
                    cout << num1 << " / " << num2 << " = " << result << endl;
                    break;
            }
        } catch (const exception& e) {
            cout << "Error: " << e.what() << endl;
        }

        cout << "\nPress Enter to continue...";
        clearInput();
        cin.get();
    }

    return 0;
}
