package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	fmt.Printf("%T\n", sum)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : %d\n", i, sum(i))
	}

	fmt.Printf("Last : %d\n", sum(100))
}
