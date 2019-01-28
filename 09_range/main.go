package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}
	//Print ids
	fmt.Println(ids)

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)

		//another way to print index and value
		fmt.Println("index", i, "value", id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Println(k + " " + v)
		fmt.Printf("%v: %T\n", k, v)
	}

	//for k := range emails {
	//	fmt.Println("Name: " + k)
	//}
}
