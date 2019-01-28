package main

import "fmt"

func checkColor(color string) string {
	// Switch
	ret :=  "switch : "
	switch color {
	case "red":
		ret += "color is red"
	case "blue":
		ret += "color is blue"
	default:
		ret += "color is NOT blue or red"
	}

	return ret
}

func main() {
	x := 10
	y := 10

	//  If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "reds"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	//Switch
	fmt.Printf("%v\n", checkColor(color))

	color = "blue"

	fmt.Printf("%v\n", checkColor(color))
}
