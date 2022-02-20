package main

import (
	"fmt"
	"math"
	"os"
)

func add() (number1 float64, number2 float64, answer float64) {
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	answer = number1 + number2
	fmt.Printf("answer: %.2f\n", answer)
	return number1, number2, answer
}

func mult() (number1 float64, number2 float64, answer float64) {
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	answer = number1 * number2
	fmt.Printf("answer: %.2f\n", answer)
	return number1, number2, answer
}

func div() (number1 float64, number2 float64, answer float64) {
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	answer = number1 / number2
	fmt.Printf("answer: %.2f\n", answer)
	return number1, number2, answer
}

func sub() (number1 float64, number2 float64, answer float64) {
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	answer = number1 - number2
	fmt.Printf("answer: %.2f\n", answer)
	return number1, number2, answer
}

func sqrt() (number1 float64, answer float64) {
	fmt.Scan(&number1)
	answer = math.Sqrt(number1)
	fmt.Printf("answer: %.3f\n", answer)
	return number1, answer
}
func main() {

	var mathOperator string
	fmt.Println("birdcalc v1.0")
	fmt.Println("type 'help' for usage!")
	fmt.Println("type 'exit' to exit!")
	for {

		fmt.Print(">> ")
		fmt.Scan(&mathOperator)

		switch mathOperator {

		case "add":
			add()

		case "sub":
			sub()

		case "mult":
			mult()

		case "div":
			div()

		case "sqrt":
			sqrt()

		case "clear":
			fmt.Print("\033[H\033[2J")

		case "help":
			fmt.Println("Mathematical Operations: add, sub, mult, div, sqrt")
			fmt.Println("How to use:")
			fmt.Println("eg. 'add 19 20'")
			fmt.Println("is going to output 39")
			fmt.Println("currently you can only pass in 2 numbers into a mathematical operation")

		case "exit":
			os.Exit(0)

		default:
			fmt.Println("Invalid Operation!")
		}

	}

}
