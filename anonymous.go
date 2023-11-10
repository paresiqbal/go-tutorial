// package main

// import "fmt"

// type Blacklist func(string) bool

// func registerUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "Pares"
// 	}

// 	registerUser("Iqbal", blacklist)
// 	// Anonymous function
// 	registerUser("Pares", func(name string) bool {
// 		return name == "Pares"
// 	})
// }
