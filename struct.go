package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
	Maried       bool
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var pares Customer
	pares.Name = "Pares Iqbal"
	pares.Adress = "Bengkulu"
	pares.Age = 23
	pares.Maried = false

	pahreza := Customer{
		Name:   "Pahreza Iqbal Prastowo",
		Adress: "Curup",
		Age:    23,
		Maried: true,
	}

	risti := Customer{"Risti Yolanda", "Bengkulu", 22, false}

	fmt.Println(risti)
	fmt.Println(pares)
	pahreza.sayHello()
}
