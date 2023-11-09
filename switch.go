package main

import "fmt"

func main() {
	name := "Pares"

	switch name {
	case "Pares":
		fmt.Println("Hello Pares")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Halo Nama")

	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")

	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
