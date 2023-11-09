package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Pares",
		"address": "Curup",
	}

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku golang"
	book["author"] = "Pares"
	book["ups"] = "ups"

	fmt.Println(book)
	delete(book, "ups")

	fmt.Println(book)
}
