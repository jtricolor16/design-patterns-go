package main

import "fmt"

func init() {
	fmt.Println("START DECORATOR")
}

func main() {
	sumInstance := NewSum()
	multiplyInstance := NewMultiply()
	divideInstance := NewDivide()
	subtractInstance := NewSubtract()
	calc := NewCalculator([]MathOperation{sumInstance, divideInstance, multiplyInstance, multiplyInstance, subtractInstance})
	fmt.Println(calc.Execute(3.0, 5.0))
}
