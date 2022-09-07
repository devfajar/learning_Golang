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

	// Tipe Data
	var positiveNumber uint8 = 89
	var negativeNumber = 123456789
	var message string = "hello"

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Println(message, "\n")

	// Perulangan
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Perulangan Bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisiasi isi array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah Element \t\t", len(fruits))
	fmt.Println("Isi Semua Element \t", fruits)

	// Slice
	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	// Map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// deteeksi keberadaan map
	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Slice on Map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
