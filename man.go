package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
)

// ---------------------------------------------------
// menu
func arithmeticMenu() int {
	clearScreen()
	var chooseArith int
	fmt.Println("|||----------ARITHMETIC MENU----------|||")
	fmt.Println("|||------------choose one...----------|||")
	fmt.Println("[0]Go to Trigonometric Menu")
	fmt.Println("[1]Addition")
	fmt.Println("[2]Subtraction")
	fmt.Println("[3]Multiplication")
	fmt.Println("[4]Division")
	fmt.Println("----------------[5]exit")
	fmt.Printf("------|")
	fmt.Scan(&chooseArith)
	return chooseArith
}

// ---------------------------------------------------
func trigonometricMenu() int {
	clearScreen()
	var chooseTrigo int
	fmt.Println("|||----------TRIGONOMETRIC MENU----------|||")
	fmt.Println("|||------------choose one...----------|||")
	fmt.Println("[0]Go to Mathematic Menu")
	fmt.Println("[1]Cosinus of Degree")
	fmt.Println("[2]Sinus of Degree")
	fmt.Println("[3]Tangens of Degree")
	fmt.Println("[4]Degrees to Radians")
	fmt.Println("[5]Radians to Degrees")
	fmt.Println("-----------------[6]exit")
	fmt.Printf("------|")
	fmt.Scan(&chooseTrigo)
	return chooseTrigo
}

// ---------------------------------------------------
func mathematicMenu() int {
	clearScreen()
	var chooseMat int
	fmt.Println("|||----------MATHEMATIC MENU----------|||")
	fmt.Println("|||------------choose one...----------|||")
	fmt.Println("[0]Back to Arithmetic Menu")
	fmt.Println("[1]Exponentiation")
	fmt.Println("[2]Square Root")
	fmt.Println("[3]Logarithm")
	fmt.Println("[4]Find Maximum")
	fmt.Println("[5]Find Minimum")
	fmt.Println("----------------[6]exit")
	fmt.Printf("------|")
	fmt.Scan(&chooseMat)
	return chooseMat
}

// ---------------------------------------------------
// asks
func askNumber() float64 {
	var number float64
	fmt.Printf("Give a First Operand:")
	fmt.Scan(&number)
	return number
}

// ---------------------------------------------------
func askSecondNumber() float64 {
	var secondNumber float64
	fmt.Println("")
	fmt.Printf("Give a Second Operand:")
	fmt.Scan(&secondNumber)
	return secondNumber
}

// ---------------------------------------------------
func askForSinKosTan() float64 {
	var ask float64
	fmt.Printf("Give Degree:")
	fmt.Scan(&ask)
	return ask
}

// ---------------------------------------------------
func askForMath() float64 {
	var askMath float64
	fmt.Printf("Give Operand:")
	fmt.Scan(&askMath)
	return askMath
}

// ---------------------------------------------------
func askForDegrees() float64 {
	var ask float64
	fmt.Printf("Give Degree:")
	fmt.Scan(&ask)
	return ask
}

// ---------------------------------------------------
func askForRadians() float64 {
	var ask float64
	fmt.Printf("Give Radians:")
	fmt.Scan(&ask)
	return ask
}

// ---------------------------------------------------
// arithmetic menu
func addition(a float64, b float64) float64 {
	return a + b
}
func subtraction(a float64, b float64) float64 {
	return a - b
}
func multiplication(a float64, b float64) float64 {
	return a * b
}
func division(a float64, b float64) float64 {
	return a / b
}

// ---------------------------------------------------
// trigonometric menu
func sinus() float64 {
	degrees := askForSinKosTan()
	radians := degrees * (math.Pi / 180)
	resultSinus := math.Sin(radians)
	return resultSinus
}
func cosinus() float64 {
	degrees := askForSinKosTan()
	radians := degrees * (math.Pi / 180)
	resultkosinus := math.Cos(radians)
	return resultkosinus
}
func tangens() float64 {
	degrees := askForSinKosTan()
	radians := degrees * (math.Pi / 180)
	resultTangens := math.Tan(radians)
	return resultTangens
}
func degreesToRadians() float64 {
	degrees := askForDegrees()
	radians := degrees * (math.Pi / 180)
	return radians
}

// ---------------------------------------------------
func radiansToDegrees() float64 {
	radians := askForRadians()
	degrees := radians * (180 / math.Pi)
	return degrees
}

// ---------------------------------------------------
// mathematic menu
func squareRoot(a float64) float64 {
	squareResult := math.Sqrt(a)
	return squareResult
}
func logarithm(a float64) float64 {
	log := math.Log(a)
	return log
}
func exponentiation(a float64, b float64) float64 {
	powResult := math.Pow(a, b)
	return powResult
}
func compMax(a float64, b float64) float64 {
	compMaxResult := math.Max(a, b)
	return compMaxResult
}
func compMin(a float64, b float64) float64 {
	compMinResult := math.Min(a, b)
	return compMinResult
}

// ---------------------------------------------------
// another functions
func clearScreen() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func waitForEnter() {
	fmt.Println("\nPress Enter to Continue...")
	fmt.Scanln()
}

// ---------------------------------------------------
// main function
func mainCalculator() {
	currentMenu := 1
MainLoop:
	for {
		switch currentMenu {
		case 1:
			choose := arithmeticMenu()
			switch choose {
			case 1:
				num1 := askNumber()
				num2 := askSecondNumber()
				additionResult := addition(num1, num2)
				fmt.Printf("Result: %.2f \n \n", additionResult)
				waitForEnter()
			case 2:
				num1 := askNumber()
				num2 := askSecondNumber()
				subtractionResult := subtraction(num1, num2)
				fmt.Printf("Result: %.2f \n \n", subtractionResult)
				waitForEnter()
			case 3:
				num1 := askNumber()
				num2 := askSecondNumber()
				multiplicationResult := multiplication(num1, num2)
				fmt.Printf("Result: %.2f \n \n", multiplicationResult)
				waitForEnter()
			case 4:
				num1 := askNumber()
				num2 := askSecondNumber()
				divisionResult := division(num1, num2)
				fmt.Printf("Result: %.2f \n \n", divisionResult)
				waitForEnter()
			case 0:
				currentMenu = currentMenu + 1
			}
			if choose == 5 {
				break MainLoop
			}
		case 2:
			choose := trigonometricMenu()
			switch choose {
			case 1:
				sinusResult := sinus()
				fmt.Printf("Result: %.3f \n \n", sinusResult)
				waitForEnter()
			case 2:
				cosinusResult := cosinus()
				fmt.Printf("Result: %.3f \n \n", cosinusResult)
				waitForEnter()
			case 3:
				tangensResult := tangens()
				fmt.Printf("Result: %.3f \n \n", tangensResult)
				waitForEnter()
			case 4:
				degreesToRadiansResult := degreesToRadians()
				fmt.Printf("Result: %.2f \n \n", degreesToRadiansResult)
				waitForEnter()
			case 5:
				radiansToDegreesResult := radiansToDegrees()
				fmt.Printf("Result: %.2f", radiansToDegreesResult)
				waitForEnter()
			case 0:
				currentMenu = currentMenu + 1
			}
			if choose == 6 {
				break MainLoop
			}
		case 3:
			choose := mathematicMenu()
			switch choose {
			case 1:
				num1 := askNumber()
				num2 := askSecondNumber()
				powResult := exponentiation(num1, num2)
				fmt.Printf("Result: %.2f \n \n", powResult)
				waitForEnter()
			case 2:
				num := askForMath()
				rootResult := squareRoot(num)
				fmt.Printf("Result: %.3f \n \n", rootResult)
				waitForEnter()
			case 3:
				num := askForMath()
				logResult := logarithm(num)
				fmt.Printf("Result: %.2f \n \n", logResult)
				waitForEnter()
			case 4:
				num1 := askNumber()
				num2 := askSecondNumber()
				maxResult := compMax(num1, num2)
				fmt.Println("Result: \n", maxResult)
				waitForEnter()
			case 5:
				num1 := askNumber()
				num2 := askSecondNumber()
				minResult := compMin(num1, num2)
				fmt.Println("Result: \n", minResult)
				waitForEnter()
			case 0:
				currentMenu = currentMenu - 2
			}
			if choose == 6 {
				break MainLoop
			}
		}
	}
}

// фигня какая-то
func main() {
	mainCalculator()
}
