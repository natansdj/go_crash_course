package main

import "fmt"

var name = "Nath"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var age int32 = 37
	var isCool = true
	var size float32 = 2.3
	var email = "nth@gmail.com"

	fmt.Println(name, age, isCool, size, email)

	// Shorthand Declaration
	name2, email2 := "Test", "test@test.com"
	isCool = false

	fmt.Println(name2, age, isCool, size, email2)
	fmt.Printf("Size : %T\n", size)
	fmt.Printf("Age : %T\n", age)
}
