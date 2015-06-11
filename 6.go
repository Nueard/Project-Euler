package main

import (
	"fmt"
	"math"
)

func main() {
    sumTo := float64(100)
    sumOne := 1.0/4*(math.Pow(sumTo,2)*math.Pow(sumTo+1,2))
    sumTwo := math.Pow(sumTo,3)/3+math.Pow(sumTo,2)/2+sumTo/6

	fmt.Printf("%f\n",sumOne-sumTwo)
}
