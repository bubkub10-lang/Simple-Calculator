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
func askWhatToDo() int {
	var choose int
	fmt.Println("What to Do?")
	fmt.Printf("||[1]addition,")
	fmt.Printf("[2]subtraction,")
	fmt.Printf("[3]multiplication,")
	fmt.Println("[4]division||")
	fmt.Printf("--------")
	fmt.Scan(&choose)
	return choose
}

func whatToDo() {
	thing := askWhatToDo()
	firstNum := askNumber()
	secondNum := askSecondNumber()
	if thing == 1 {
		fmt.Println("addittion", firstNum+secondNum)
	}
	if thing == 2 {
		fmt.Println("subtraction", firstNum-secondNum)
	}
	if thing == 3 {
		fmt.Println("multiplication", firstNum*secondNum)
	}
	if thing == 4 {
		fmt.Println("division", firstNum/secondNum)
	}
}

func main() {
	whatToDo()
}
