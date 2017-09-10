package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("** GoPrime **\nPlease type an integer: ")
	var number int
	fmt.Scanln(&number)
	fmt.Printf("Performing prime factorization on number %v...\n", number)
	factors := performPrimeFactorization(number)
	factorsString := prettyFactors(number, factors)
	fmt.Println(factorsString)
}

// prettyFactors returns a string with the format "number = factor1 * factor2 ..."
func prettyFactors(number int, factors []int) string {
	factorsString := strconv.Itoa(number) + " = "
	for i, v := range factors {
		factorsString += strconv.Itoa(v)
		if i != len(factors)-1 {
			factorsString += " * "
		}
	}
	return factorsString
}

// performPrimeFactorization factors an integer using trial division
func performPrimeFactorization(number int) []int {
	currentNumber := number
	factors := []int{}
	fmt.Print(currentNumber)
	for i := 2; i <= number; i++ {
		if !isPrimeNumber(i) {
			continue
		}
		for isNumberDivisible(currentNumber, i) {
			fmt.Println(" |", i)
			factors = append(factors, i)
			currentNumber /= i
			fmt.Print(currentNumber)
		}
	}
	fmt.Println() // Empty line on purpose
	return factors
}

func isNumberDivisible(number1, number2 int) bool {
	return number1%number2 == 0
}

func isPrimeNumber(number int) bool {
	// Prime numbers only have two divisors in the Natural range: 1 and the number itself
	for i := 2; i < number; i++ {
		if isNumberDivisible(number, i) {
			return false
		}
	}
	return true
}
