package main

import "fmt"

func main() {
	var myName = "Pahreza Iqbal"
	fmt.Println(myName)

	// Menggungkan ":" sebelum "=" dapat menggantikan var
	wifeName := "Risti Yolanda"
	fmt.Println(wifeName)

	var (
		firstName = "Pares"
		lastName  = "Iqbal"
	)

	fmt.Println(firstName + lastName)
	fmt.Println(lastName)
}
