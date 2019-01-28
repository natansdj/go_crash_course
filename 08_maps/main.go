package main

import "fmt"

func setGmail(a string) string {
	return a + "@gmail.com"
}

func main() {
	// Define map (key-value pair)
	emails := make(map[string]string)

	//  Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	// Decalre map and add kv
	emails2 := map[string]string{"Bob": setGmail("bob"), "Sharon": setGmail("sharon")}

	emails2["Mike"] = setGmail("mike")

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	fmt.Println("-")
	fmt.Println(emails2)

	delete(emails2, "Bob")
	fmt.Println(emails2)
}
