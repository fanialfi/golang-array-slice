package main

import "fmt"

func contohPenerapan() {
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "Water"
	names[3] = "Law"

	fmt.Println(names[0], names[1], names[2], names[3])
}

func inisialisasiNilaiAwal() {
	var fruits = [4]string{"Apple", "Banana", "Grape", "Melon"}

	var gayaVertikal = [4]string{
		"Apple",
		"Banana",
		"Grape",
		"Melon"}

	fmt.Println("Jumplah element \t\t", len(fruits)) // fungsi len() digunakan untuk menghitung jumplah element dalam array
	fmt.Println("Isi semua element \t", fruits)      // penggunaan fungsi fmt.Println() pada array tanpa menyebutkan index element, akan menghasilkan output dalam bentuk string dari semua array yang ada.

	fmt.Println("==========")

	fmt.Println("Jumplah element \t\t", len(gayaVertikal))
	fmt.Println("Isi semua element \t", gayaVertikal)
}

func tanpaJumplahElement() {
	var numbers = [...]int{1, 2, 3, 4, 5}

	fmt.Println("Panjang element  :", len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("element ke-%d : %d\n", i, numbers[i])
	}
	// numbers[5] = 6 // tidak bisa menambahkan element baru ke dalam array meskipun inisialisasi jumplah element dalam array telah diganti dengan [...]
}

func multidimensi() {
	var numbers1 = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	var numbers2 = [3][2]int{{1, 2}, {1, 2}, {1, 2}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}

func arrayFor() {
	var numbers = [4][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	for i := 0; i < len(numbers); i++ {
		fmt.Println()
		fmt.Printf("Data index ke-%d : %d\n", i, numbers[i])

		for j := 0; j < len(numbers[i]); j++ {
			fmt.Printf("Data dengan index luar ke-%d dan index dalam ke-%d : %d\n", i, j, numbers[i][j])
		}
	}
}

func arrayForRange() {
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	for i, month := range months {
		fmt.Println()
		fmt.Printf("Data index ke-%d berisi : %s\n", i, month)
		fmt.Printf("Panjang %s adalah : %d huruf\n", month, len(month))

		// melakukan perulangan setiap huruf-nya
		for _, key := range month {
			fmt.Printf("bulan %s memiliki isi : %c\n", months[i], key)
		}
	}
}

func alokasiMake() {
	var weeks = make([]string, 7)

	weeks[0] = "SU"
	weeks[1] = "MO"
	weeks[2] = "TU"
	weeks[3] = "WE"
	weeks[4] = "TH"
	weeks[5] = "FR"
	weeks[6] = "SA"

	for i, week := range weeks {
		fmt.Printf("%s\t", week)

		if i == len(weeks)-1 {
			fmt.Printf(" \n")
		}
	}

	var step = len(weeks)
	for i := 1; i <= 30; i++ {
		fmt.Printf("%d\t", i)

		if step == i {
			fmt.Println()
			step = step + len(weeks)
		} else if i == 30 {
			fmt.Println()
		}
	}
}

func main() {
	contohPenerapan()
	fmt.Println("--------------------")
	inisialisasiNilaiAwal()
	fmt.Println("--------------------")
	tanpaJumplahElement()
	fmt.Println("--------------------")
	multidimensi()
	fmt.Println("--------------------")
	arrayFor()
	fmt.Println("--------------------")
	arrayForRange()
	fmt.Println("--------------------")
	alokasiMake()
}
