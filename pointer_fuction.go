package main

import "fmt"

type Adress struct {
	City, Province, Country string
}

func ChangeCountry(address Adress) {
	address.Country = "Indonesia"
}

func ChangeCity(address *Adress) {
	address.City = "Curup"
}

func main() {
	address := Adress{}
	ChangeCountry(address)

	var address1 Adress = Adress{}
	ChangeCity(&address1)

	fmt.Println(address)
	fmt.Println(address1)
}
