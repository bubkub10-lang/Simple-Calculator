package main

import "fmt"

func askNumber() float64 {
	var number float64
	fmt.Printf("Give a First Number:")
	fmt.Scan(&number)
	return number
}

func askSecondNumber() float64 {
	var secondNumber float64
	fmt.Println("")
	fmt.Printf("Give a Second Number:")
	fmt.Scan(&secondNumber)
	return secondNumber
}

func main() {
	firstNum := askNumber()
	secondNum := askSecondNumber()
	fmt.Println("")
	fmt.Println("addition:", firstNum+secondNum)
	fmt.Println("subtraction:", firstNum-secondNum)
	fmt.Println("multiplication:", firstNum*secondNum)
	fmt.Println("division:", firstNum/secondNum)
}
