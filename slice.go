// package main

// import "fmt"

// func main() {
// 	names := [...]string{"Pahreza", "Iqbal", "Prastowo", "Risti", "Yolanda"}
// 	slice1 := names[1:3]
// 	slice2 := names[:2]

// 	fmt.Println(slice1)
// 	fmt.Println(slice2)

// 	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

// 	daySlice1 := days[5:] // sabtu & minggu
// 	fmt.Println(daySlice1)

// 	daySlice1[0] = "saturday"
// 	daySlice1[1] = "sunday"
// 	fmt.Println(daySlice1)
// 	fmt.Println(days)

// 	daySlice2 := append(daySlice1, "hari baru")
// 	fmt.Println(daySlice1)
// 	fmt.Println(daySlice2)

// 	newSlice := make([]string, 2, 5)
// 	newSlice[0] = "Pares"
// 	newSlice[1] = "Thejs"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	newSliceApen := append(newSlice, "099")
// 	fmt.Println(newSliceApen)

// 	fromSlice := days[:]
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))
// 	copy(toSlice, fromSlice)

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)
// }
