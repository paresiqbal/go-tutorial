package main

import "fmt"

func main() {
	for counter := 1; counter <= 20; counter++ {
		fmt.Println("Perulangan ke", counter)

	}

	fmt.Println("Selesai")

	names := []string{"Pares", "Iqbal", "Prastowo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
