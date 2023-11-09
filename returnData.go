package main

import "fmt"

func getHello(name string) string {
	return "Hello" + name
}

func getFullname() (string, string) {
	return "Pahreza", "Iqbal"
}

func main() {
	result := getHello("Pares")
	fmt.Println(result)
	fmt.Println(getHello("Risti"))

	firstName, lastName := getFullname() // gunakan _ untuk ignor value
	fmt.Println(firstName, lastName)

}
