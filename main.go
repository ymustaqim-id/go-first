package main

import (
	"fmt"
	"math"
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

	fmt.Println("---------------------------")
	var fruits = [4]string{
		"anggur",
		"bananan",
		"grape",
		"apple",
	}
	fmt.Println("array buah \t", fruits)
	fmt.Println("array jumlah \t", len(fruits))
	fmt.Println("---------------------------")

	var angka = [...]int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println("angka bebas \t", angka)
	fmt.Println("angka bebas \t", len(angka))
	fmt.Println("---------------------------")

	var nilaiAngka = [...]string{"aku", "kamu", "dan", "mereka"}
	for _, hasil := range nilaiAngka {
		fmt.Println("hasil nya adalah", hasil)
	}
	fmt.Println("---------------------------")
	var alokasi = make([]int, 2)
	alokasi[0] = 1
	alokasi[1] = 2
	fmt.Println("nilai", alokasi)
	fmt.Println("---------------------------")

	var data = map[string]int{"nilai": 2, "kelas": 3, "tempat": 4}
	for _, v := range data {
		fmt.Println("nilai data", data)
		fmt.Println("nilai v", v)
	}
	fmt.Println("---------------------------")
	var randomValue = randomNilai(2, 3)
	fmt.Println("nilai random 3K, 5K, 10K \t", randomValue)

	fmt.Println("---------------------------")
	var nilaiData float64 = 15
	var number1, number2 = mathHere(nilaiData)
	fmt.Println("hasil math \t:", number1)
	fmt.Println("hasil \t\t:", number2)
	fmt.Println("---------------------------")

	var hitungData = hitungNilaiHere(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Hasil hitung", hitungData)
	fmt.Println("---------------------------")

	var numbersHere = []int{2, 3, 4, 5, 6, 7, 8, 9, 9, 9}
	var rerata = hitungNilaiHere(numbersHere...)
	fmt.Printf("hasil rerata %f \n", rerata)
	fmt.Println("---------------------------")

	var closure = []int{2, 3, 3, 4, 4, 3, 3, 2, 2, 3}
	var getMinMax = func(input int) []int {
		var dataR []int
		for _, val := range closure {
			if val < input {
				continue
			}
			dataR = append(dataR, val)

		}
		return dataR
	}(3)
	fmt.Println("Nilai min max", getMinMax)
	fmt.Println("---------------------------")

	// POINTER
	var pointer1 int = 5
	var pointer2 *int = &pointer1

	fmt.Println("value pointer1", pointer1)
	fmt.Println("memory pointer1", &pointer1)

	fmt.Println("value pointer2", *pointer2)
	fmt.Println("memory pointer2", pointer2)
	fmt.Println("---------------------------")

	var smk Pelajar
	smk.nama = "andre"
	smk.kelas = 4

	fmt.Printf("nama %v", smk.nama)
	fmt.Printf("kelas %v", smk.kelas)
	fmt.Println("--------------------")

	var rumahKu = Rumah{
		genteng:     "gentengKu",
		jendela:     "kusen",
		pintu:       "pintuAjaib",
		jumlahPintu: 4,
		jumlahOrang: 5,
	}

	fmt.Println("merk genteng", rumahKu.genteng)
	fmt.Println("merk jendela", rumahKu.jendela)
	fmt.Println("merk pintu", rumahKu.pintu)
	fmt.Println("jumlah Pintu", rumahKu.jumlahPintu)
	fmt.Println("di dalam nya ada", rumahKu.jumlahOrang, "orang")
	fmt.Println("--------------------")

} //main

func randomNilai(data, value int) int {
	var nilaiRand = (data * value) + 12
	return nilaiRand
}
func mathHere(d float64) (float64, float64) {
	var number1 = math.Pi * math.Pow(d/2, 2)
	var number2 = math.Pi * d

	return number1, number2
}

func hitungNilaiHere(value ...int) float64 {
	var total int = 0
	for _, val := range value {
		total += val
	}
	var rata2 = float64(total) / float64(len(value))
	return rata2
}

type Pelajar struct {
	nama  string
	kelas int
}

type Rumah struct {
	genteng, jendela, pintu  string
	jumlahPintu, jumlahOrang int
}
