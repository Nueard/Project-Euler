package main

import (
	"fmt"
	"math"
)

func main() {
	maxNumber := 999 * 999
	found := false
	for i := maxNumber; i > 0 && !found; i-- {
		if checkPalindrome(i) {
			for j := 999; j >= i/999 && !found; j-- {
				if i%j == 0 {
					fmt.Println(i)
					found = true
				}
			}
		}
	}
}

func getNumberOfDigits(number int) int {
	counter := 0
	for {
		counter++
		number /= 10
		if number < 10 {
			counter++
			break
		}
	}
	return counter
}

func getDigit(number int, digit int) int {
	if digit < 0 {
		digit += getNumberOfDigits(number)
	}
	return (number / int(math.Pow(10, float64(getNumberOfDigits(number)-digit-1)))) % 10
}

func checkPalindrome(number int) bool {
	size := getNumberOfDigits(number)
	for i := 0; i < size/2+1; i++ {
		if getDigit(number, i) != getDigit(number, -i-1) {
			return false
		}
	}
	return true
}
