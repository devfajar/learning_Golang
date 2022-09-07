package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"John", "Wick"}
	printMessage("halo", names)

	// example
	rand.Seed(time.Now().Unix())
	var randomValue int
}
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
