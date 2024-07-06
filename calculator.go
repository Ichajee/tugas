// calculator
package main

import "fmt"

func main() {
    fmt.Println("Simple Calculator")
    fmt.Println("----------------")

    var num1 float64
    var num2 float64
    var op string

    fmt.Print("Enter number 1: ")
    fmt.Scanln(&num1)

    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&op)

    fmt.Print("Enter number 2: ")
    fmt.Scanln(&num2)

    var result float64

    switch op {
    case "+":
        result = num1 + num2
    case "-":
        result = num1 - num2
    case "*":
        result = num1 * num2
    case "/":
        if num2 != 0 {
            result = num1 / num2
        } else {
            fmt.Println("Error: Division by zero!")
            return
        }
    default:
        fmt.Println("Error: Invalid operator!")
        return
    }

    fmt.Printf("Result: %.2f\n", result)
}