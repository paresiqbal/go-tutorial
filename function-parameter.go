// package main

// import "fmt"

// func sayHello(name string, filter func(string) string) {
// 	flteredName := filter(name)
// 	fmt.Println("Hello", flteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "***"
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHello("Pares", spamFilter)

// 	filter := spamFilter
// 	sayHello("Anjing", filter)
// }
