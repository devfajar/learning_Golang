package main

import "fmt"

func main() {
	// Komentar one line
	/* Komentar Multiple Line

	 */
	fmt.Println("Hello World")
	// fmt.Println() dapat menampung tak terbatas parameternya
	fmt.Println("Hello", "World!", "how", "are", "you")

	// Deklarasi Variable
	var firstName string = "John"
	var lastName string
	lastName = "Doe"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)

}
