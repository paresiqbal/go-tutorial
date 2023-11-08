package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Pahreza"
	names[1] = "Iqbal"
	names[2] = "Prastowo"

	fmt.Println(names[1])

	var value = [3]int{1, 2, 3}
	fmt.Println(value)
}
