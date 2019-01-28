package main

import "fmt"

func inc(x *int) {
	*x++
}

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("Type : %T\n", b)

	//  Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	//Increment Func
	fmt.Println("-- Increment Func --")
	inc(&a)
	fmt.Println(a)
}
