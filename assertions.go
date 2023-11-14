package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

}
