package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Define command-line flags
    num1 := flag.String("num1", "", "First number")
    num2 := flag.String("num2", "", "Second number")
    operation := flag.String("op", "add", "Operation to perform: add or subtract")

    // Parse the flags
    flag.Parse()

    // Validate and convert the input numbers
    n1, err := strconv.ParseFloat(*num1, 64)
    if err != nil {
        fmt.Println("Invalid value for num1")
        os.Exit(1)
    }

    n2, err := strconv.ParseFloat(*num2, 64)
    if err != nil {
        fmt.Println("Invalid value for num2")
        os.Exit(1)
    }

    // Perform the operation
    var result float64
    switch *operation {
    case "add":
        result = n1 + n2
    case "subtract":
        result = n1 - n2
    default:
        fmt.Println("Invalid operation. Use 'add' or 'subtract'")
        os.Exit(1)
    }

    // Display the result
    fmt.Printf("Result: %f\n", result)
}
