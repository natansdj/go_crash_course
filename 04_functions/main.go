package main

import (
	"errors"
	"fmt"
	"math"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(greeting("Test"))
	fmt.Println(getSum(3, 4))

	result, err := sqrt(16)

	//Do the sqrt function
	fmt.Println("Sqrt : ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
