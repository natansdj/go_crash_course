package main

import "fmt"

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d : FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d : Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d : Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// Long method
	i := 1
	until := 5
	for i <= until {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 1; i <= until; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	fizzBuzz()
}
