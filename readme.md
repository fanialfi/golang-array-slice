# tipe data array

array didalam golang adalah kumpulan data bertipe data sama yang disimpan didalam sebuah variabel. Default nilai tiep element array pada awalnya adalah zero value tergantung dari tipe data-nya, jika `int` maka nilai default-nya adalah `0`, jika `bool` maka adalah `false` dst. Setiap element pada array memiliki indeks berupa angka yang dimulai dari angka `0`.

deklarasi array dalam golang memiliki format sebagai berikut :

```go
var namaArray [panjang]tipeData
```

pengisian element array pada indeks yang tidak sesuai dengan alokasi atau panjang array akan menghasilkan error, contoh sederhana-nya adalah jika sebuah array yang panjangnya sebanyak 4 element, maka pengisian element ke 5 dst akan menghasilkan error. Solusi dari masalah tersebut adalah dengan menggunakan keyword `append` yang akan dibahas lain waktu.

```go
func main() {
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "Water"
	names[3] = "Law"

	fmt.Println(names[0], names[1], names[2], names[3])
}
```

cara pengisian element array diatas yaitu dilakukan dengan langsung mengakses element dengan nomor index-nya lalu mengisinya.

selain cara pengisian element array seperti diatas, pengisian element array bisa dilakukan pada saat pendeklarasian variabel, caranya dengan menuliskan data element didalam tanda kurung kurawal (`{}`) setelah tipe data kemudian untuk membatasi antara element satu dengan element yang lain dengan menyisipkan tanda koma (`,`).

```go
func main(){
	var fruits = [4]string{"Apple","banana","Grape","Melon"}

	fmt.Println("Jumplah element \t\t",len(fruits))
	fmt.Println("Isi semua element\t",fruits)
}
```

penulisan element array diatas juga bisa disebut dengan dituliskan dalam bentuk horizontal, selain penulisan element array dalam bentuk horizontal, bisa ditulis dalam bentuk vertikal, namun khusus untuk vertikal tanda koma (`,`) wajib dituliskan setelah element array meskipun itu element terakhir, jika tidak maka akan terjadi error, jika tidak ingin menuliskan tanda koma di element terakhir maka untuk penempatan kurung kurawal tutup ditempatkan setelah element terakhir.

```go
var one = [3]string{
	"Satu",
	"Dua",
	"tiga",
}

var two = [3]string{
	"satu",
	"dua",
	"tiga"}
```

penulisan kedua variabel diatas dianggap benar.

Selain itu saat pendeklarasian array yang nilai-nya di set di awal, bisa tidak dituliskan jumpah element yang dapat ditampung oleh variabel-nya, cukup ganti dengan tanda titik tiga (`[...]`), maka jumplah element akan di kalkulasi secara otomatis menyesuaikan data element yang di masukkan kedalam array.

```go
func main(){
	var numbers = [...]int{1,2,3,4,5}

	fmt.Println("Panjang element :",len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element ke-%d isinya %d\n",i,numbers[i])
	}
}
```

meskipun penulisan jumplah element pada saat pendeklarsian array sudah di ganti dengan tanda titik tiga (`[...]`) namun tetap tidak bisa menambahkan element baru kedalam array tersebut

```go
var numbers = [...]int{1,2,3,4,5}

// numbers[5] = 6 // error => invalid argument, index 5 out of bounds
```

selain itu sebuah array juga bisa disimpan didalam array atau yang biasa disebut dengan array multi dimensi. Cara deklarasi-nya secara umum sama dengan array biasa, namun khusus untuk array yang merupakan sub dimensi-nya boleh tidak dituliskan tipe data dan jumplah element-nya.

```go
func main(){
	var numbers1 = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	var numbers2 = [3][2]int{{1, 2}, {1, 2}, {1, 2}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
```

untuk array multidimensi, element terluar semuanya juga harus berupa array lagi

```go
func main(){
	var numbers = [2][2]int{{1,2},{1,2}} // true
	var numbers2 = [2][2]int{{1,2},2} // false
}
```

untuk mendapatkan isi dari array yang sudah didefinisikan, bisa dengan menggunakan perulangan `for` seperti berikut :

```go
func main(){
	var numbers = [4][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	for i := 0; i < len(numbers); i++ {
		fmt.Println()
		fmt.Printf("Data index ke-%d : %d\n", i, numbers[i])

		for j := 0; j < len(numbers[i]); j++ {
			fmt.Printf("Data dengan index luar ke-%d dan index dalam ke-%d : %d\n", i, j, numbers[i][j])
		}
	}
}
```

selain dengan menggunakan cara di atas untuk mengambil nilai dari array, bisa dengan menggunakan keyword `range`, keyword tersebut akan mengembalikan 2 data yaitu indeks dan element.

kadang kala ketika melakukan looping menggunakan `range` ada kemungkinan data yang ingi digunakan hanyalah element-nya saja, index nya tidak diperlukan, seperti yang sudah di ketahui di GO tidak diperbolehkan adanya variabel yang nganggur, jika di paksa maka akan terjadi error, untuk mengatasi hal tersebut data index yang dikembalikan oleh `range` bisa ditampung menggunakan variabel _underscore_ 

contoh penggunaan variabel _underscore_ :

```go
func main(){
	var month string = "Januari"

	for _, key := range month {
			fmt.Printf("bulan %s memiliki isi : %c\n", months[i], key)
		}
}
```

untuk mendeklarasikan sekaligus alokasi data array bisa dilakukan lewat keyword `make`

```go
func main(){
	var month = make([]string,12)
}
```
 parameter pertama keyword `make` ditulis dengan tipe data element array yang diinginkan, dan parameter ke-dua adalah jumplah elementnya.