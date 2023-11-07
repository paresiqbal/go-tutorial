package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsen bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir && lulusAbsen

	fmt.Println(lulus)
}
