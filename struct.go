package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
	Maried       bool
}

func main() {
	var pares Customer
	pares.Name = "Pares Iqbal"
	pares.Adress = "Bengkulu"
	pares.Age = 23
	pares.Maried = false

	fmt.Println(pares.Name)
}
