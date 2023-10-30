package main

import "fmt"

func main() {
	var myName = "Pahreza Iqbal"
	fmt.Println(myName)

	// Menggungkan ":" sebelum "=" dapat menggantikan var
	wifeName := "Risti Yolanda"
	fmt.Println(wifeName)

	var (
		firstName = "USS"
		lastName  = "KIDD"
	)
	fmt.Println(firstName + " " + lastName)

	// const variabel yang datanya tidak dapat dibubah lagi
	const shipName = "USS IOWA"
	fmt.Println(shipName)

	const (
		nation  = "KMS"
		name    = "Hanover"
		caliber = "381mm"
	)
	fmt.Println(nation + " " + name + " " + caliber)
}
