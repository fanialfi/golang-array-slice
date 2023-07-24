package main

import "fmt"

func cetakHasil(s []string, b bool) {
	if b {
		fmt.Printf("isi : %s, panjang %d, lebar %d\n", s, len(s), cap(s))
	} else {
		fmt.Println(s)
	}
}

func cetakHasilInt(i []int, b bool) {
	if b {
		fmt.Printf("isi %d, panjang %d, lebar %d\n", i, len(i), cap(i))
	} else {
		fmt.Println(i)
	}
}

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

func slice() {
	var fruits = []string{"apple", "banana", "grape", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("isi index ke-%d : %s\n", i, fruit)
	}

	{
		var fruits = [4]string{"apple", "banana", "grape", "melon"}

		var newFruits = fruits[1:4] // diambil dari element ke-1 hingga element sebelum index ke-4

		cetakHasil(newFruits, false)
	}
}

func sliceReferensi() {
	var fruits = []string{
		"apel",
		"grape",
		"banana",
		"melon",
	}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bbFruits[0] = "pinnaple" // buah "grape" diubah menjadi "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
}

func lenFunc() {
	var names = []string{"fani", "alfi", "fanialfi"}

	fmt.Println(len(names)) // len() untuk mengecek panjang sebuah slice

	for i, name := range names {
		fmt.Printf("index ke-%d berisi : %s panjangnya %d\n", i, name, len(name)) // len() untuk mengecek panjang string

		for _, alphabet := range name {
			fmt.Printf("%c\n", alphabet)
		}

		fmt.Println()
	}
}

func capFunc() {
	var fruits = []string{
		"Pisang",
		"Pepaya",
		"Alpukat",
		"Mangga",
		"Apple",
		"Lemon",
		"Delima",
		"Kiwi",
	}

	fmt.Println("panjang slice :", len(fruits))
	fmt.Println("kapasitas slice :", cap(fruits))
	fmt.Println(fruits)

	fmt.Println()

	var aFruits = fruits[0:5]
	fmt.Println("panjang slice :", len(aFruits))
	fmt.Println("kapasitas slice:", cap(aFruits))
	fmt.Println(aFruits)

	fmt.Println()

	var bFruits = fruits[3:7]
	fmt.Println("panjang slice :", len(bFruits))
	fmt.Println("kapasitas slice :", cap(bFruits))
	fmt.Println(bFruits)
}

func appendFunc() {
	var fruits = []string{"Apel", "Pisang", "Anggur"}

	var newFruits = append(fruits, "Melon", "Pepaya") // fungsi append() mengembalikan slice baru

	cetakHasil(fruits, false)
	cetakHasil(newFruits, false)

	fmt.Println()

	{
		var fruits = []string{"apel", "pisang", "anggur", "melon", "pepaya"}
		var aFruits = fruits[0:3]

		cetakHasil(fruits, true)
		cetakHasil(aFruits, true)

		fmt.Println()

		var bFruits = append(aFruits, "jagung") // jika size lebih besar dari panjang slice, maka element baru ditempatkan dalam cangkupan size-nya, dan menjadikan slice dengan referensi yang sama akan ikut berubah.
		cetakHasil(fruits, true)
		cetakHasil(aFruits, true)
		cetakHasil(bFruits, true)
	}
}

func copyFunc() {
	dst := make([]string, 4)
	var src = []string{"a", "b", "c"}

	fmt.Println(dst)
	cetakHasil(src, true)
	copy(dst, src)
	cetakHasil(dst, true)

	fmt.Println()

	{
		var dst = []string{"melon", "melon", "melon"}
		var src = []string{"apel", "anggur"}

		fmt.Println(dst, src)
		copy(dst, src)
		fmt.Println(dst, src)
	}
}

func tigaIndex() {
	var A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var B = A[0:5:5]
	// var C = A[3:8:19] // tidak bisa, karena ukuran slice A = 10, dan tidak bisa menggunakan 19
	C := A[3:8:8]

	cetakHasilInt(A, true)
	cetakHasilInt(B, true)
	cetakHasilInt(C, true)

	for i, num := range C {
		fmt.Printf("index ke-%d : %d\n", i, num)
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
	fmt.Println("--------------------")
	fmt.Println("===== SLICE =====")
	fmt.Println("--------------------")
	slice()
	fmt.Println("--------------------")
	sliceReferensi()
	fmt.Println("--------------------")
	fmt.Println("===== Built-In Function =====")
	fmt.Println("--------------------")
	lenFunc()
	fmt.Println("--------------------")
	capFunc()
	fmt.Println("--------------------")
	appendFunc()
	fmt.Println("--------------------")
	copyFunc()
	fmt.Println("--------------------")
	tigaIndex()
}
