package main

import (
	"fmt"
)

func main() {
	var jika bool = false
	if jika == true {
		var namaDepan string = "Junio"
		fmt.Printf("halo %s\n", namaDepan)
		fmt.Println("--------------------")
	} else {
		var nomorPunggung int = 19
		fmt.Printf("halo %d\n", nomorPunggung)
		fmt.Println("--------------------")
	}

	var nilai = (2 * 9) + (3 * 3)
	var nilaiSama = (nilai == 3)

	fmt.Printf("nilai sama %d (%t)", nilai, nilaiSama)
	fmt.Println("--------------------")

	const nilaiAwal = 9
	const nilaiAkhir = 15
	if nilaiAwal < 9 {
		fmt.Println("Nilai kurang dari batas")
	} else if nilaiAwal >= 9 && nilaiAwal <= 15 {
		fmt.Println("Nilai sesuai batas")
	}
	fmt.Println("--------------------")

	const startPoint = 0
	switch startPoint {
	case 1:
		fmt.Println("Good")
	case 2:
		fmt.Println("Nice")
	case 3:
		fmt.Println("Excellet")
	default:
		fmt.Println("You have to try harder")

	}
	fmt.Println("---------------------------")
	const class = 9
	switch {
	case class == 1:
		fmt.Println("Good")
	case (class > 2) && (class <= 5):
		fmt.Println("Nice")
	case (class > 5) && (class < 10):
		fmt.Println("Excellet")
	default:
		fmt.Println("You have to try harder")
	}
	fmt.Println("---------------------------")

	const count = 10
	for i := 0; i < count; i++ {
		fmt.Println("angka ke", i)
	}
	fmt.Println("---------------------------")

	const counts = 10
	for i := 1; i < counts; i++ {
		for j := i; j < counts; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

} //main
