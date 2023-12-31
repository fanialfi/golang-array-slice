# tipe data array

array didalam golang adalah kumpulan data bertipe data sama yang disimpan didalam sebuah variabel. Default nilai tiep element array pada awalnya adalah zero value tergantung dari tipe data-nya, jika `int` maka nilai default-nya adalah `0`, jika `bool` maka adalah `false` dst. Setiap element pada array memiliki indeks berupa angka yang dimulai dari angka `0`.

deklarasi array dalam golang memiliki format sebagai berikut :

```go
var namaArray [panjang]tipeData
```

pengisian element array pada indeks yang tidak sesuai dengan alokasi atau panjang array akan menghasilkan error, contoh sederhana-nya adalah jika sebuah array yang panjangnya sebanyak 4 element, maka pengisian element ke 5 dst akan menghasilkan error. Solusi dari masalah tersebut adalah dengan menggunakan keyword `append` yang akan dibahas lain waktu.
<details>
	<summary>example</summary>

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
</details>
<br>
cara pengisian element array diatas yaitu dilakukan dengan langsung mengakses element dengan nomor index-nya lalu mengisinya.

selain cara pengisian element array seperti diatas, pengisian element array bisa dilakukan pada saat pendeklarasian variabel, caranya dengan menuliskan data element didalam tanda kurung kurawal (`{}`) setelah tipe data kemudian untuk membatasi antara element satu dengan element yang lain dengan menyisipkan tanda koma (`,`).
<details>
	<summary>example</summary>

```go
func main(){
	var fruits = [4]string{"Apple","banana","Grape","Melon"}

	fmt.Println("Jumplah element \t\t",len(fruits))
	fmt.Println("Isi semua element\t",fruits)
}
```
</details>
<br>

penulisan element array diatas juga bisa disebut dengan dituliskan dalam bentuk horizontal, selain penulisan element array dalam bentuk horizontal, bisa ditulis dalam bentuk vertikal, namun khusus untuk vertikal tanda koma (`,`) wajib dituliskan setelah element array meskipun itu element terakhir, jika tidak maka akan terjadi error, jika tidak ingin menuliskan tanda koma di element terakhir maka untuk penempatan kurung kurawal tutup ditempatkan setelah element terakhir.
<details>
	<summary>example</summary>

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
</details>
<br>

penulisan kedua variabel diatas dianggap benar.

Selain itu saat pendeklarasian array yang nilai-nya di set di awal, bisa tidak dituliskan jumpah element yang dapat ditampung oleh variabel-nya, cukup ganti dengan tanda titik tiga (`[...]`), maka jumplah element akan di kalkulasi secara otomatis menyesuaikan data element yang di masukkan kedalam array.

<details>
	<summary>example</summary>

```go
func main(){
	var numbers = [...]int{1,2,3,4,5}

	fmt.Println("Panjang element :",len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element ke-%d isinya %d\n",i,numbers[i])
	}
}
```
</details><br>

meskipun penulisan jumplah element pada saat pendeklarsian array sudah di ganti dengan tanda titik tiga (`[...]`) namun tetap tidak bisa menambahkan element baru kedalam array tersebut

```go
var numbers = [...]int{1,2,3,4,5}

// numbers[5] = 6 // error => invalid argument, index 5 out of bounds
```

selain itu sebuah array juga bisa disimpan didalam array atau yang biasa disebut dengan array multi dimensi. Cara deklarasi-nya secara umum sama dengan array biasa, namun khusus untuk array yang merupakan sub dimensi-nya boleh tidak dituliskan tipe data dan jumplah element-nya.

<details>
	<summary>example</summary>

```go
func main(){
	var numbers1 = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	var numbers2 = [3][2]int{{1, 2}, {1, 2}, {1, 2}}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
```
</details><br>

untuk array multidimensi, element terluar semuanya juga harus berupa array lagi

```go
func main(){
	var numbers = [2][2]int{{1,2},{1,2}} // true
	var numbers2 = [2][2]int{{1,2},2} // false
}
```

untuk mendapatkan isi dari array yang sudah didefinisikan, bisa dengan menggunakan perulangan `for` seperti berikut :

<details>
	<summary>example</summary>

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
</details><br>

selain dengan menggunakan cara di atas untuk mengambil nilai dari array, bisa dengan menggunakan keyword `range`, keyword tersebut akan mengembalikan 2 data yaitu indeks dan element.

kadang kala ketika melakukan looping menggunakan `range` ada kemungkinan data yang ingi digunakan hanyalah element-nya saja, index nya tidak diperlukan, seperti yang sudah di ketahui di GO tidak diperbolehkan adanya variabel yang nganggur, jika di paksa maka akan terjadi error, untuk mengatasi hal tersebut data index yang dikembalikan oleh `range` bisa ditampung menggunakan variabel _underscore_ 

<details>
	<summary>contoh penggunaan variabel <i>underscore</i></summary>

```go
func main(){
	var month string = "Januari"

	for _, key := range month {
			fmt.Printf("bulan %s memiliki isi : %c\n", months[i], key)
		}
}
```
</details><br>

untuk mendeklarasikan sekaligus alokasi data array bisa dilakukan lewat keyword `make`

```go
func main(){
	var month = make([]string,12)
}
```

 parameter pertama keyword `make` ditulis dengan tipe data element array yang diinginkan, dan parameter ke-dua adalah jumplah elementnya.

# slice

merupakan referensi element array, bisa dibuat atau juga dihasilkan dari manipulasi sebuah array atau sclice lainnya. Karena merupakan referensi, menjadikan perubahan data di tiap element akan berdampak pada slice yang lain yang memiliki alamat memori yang sama.

cara pembuatan slice sama seperti cara pembuatan array pada umumnya, bedanya tidak perlu mendefinisikan jumplah element yang bisa di tampung ketika awal deklarasi.

```go
func main(){
	var fruits = []string{"apple","banana","grape","melon"}
}
```

slice adalah referensi dari tiap tiap element yang ada di sebuah array, sedangkan array adalah kumpulan dari sebuah element.

```go
var fruits = []string{"apple","banana","grape","melon"}
var newFruits = fruits[0:3]

fmt.Println(newFruits)
```

`fruits[0:3]` maksudnya adalah pengaksesan element dalam fruits dimulai dari index ke-0 sampai index sebelum ke-3 dan kemudian disimpan didalam variabel `newFruits`.

Pada contoh di atas `newFruits` adalah slice baru yang tercetak atau referensi dari slice `fruits`.

list operasi operasi yang menggunakan teknik 2 index yang bisa dilakukan :

```go
var fruits = []string{"apple","banana","grape","melon"}
```

|operasi|output|keterangan|
|-------|------|----------|
|`fruits[0:2]`|`["apple","banana"]`|semua element dari index ke-0 sampai sebelum index ke-2|
|`fruits[0:4]`|`["apple","banana","grape","melon"]`|semua element dari index ke-0 sampai sebelum index ke-4|
|`fruits[0:0]`|`[]`|menghasilkan slice kosong|
|`fruits[4:4]`|`[]`|menghasilkan sclie kosong|
|`fruits[4:0]`|`[]`|error pada saat penulisan `fruits[a:b]` karena `a` harus lebih kecil dari `b`|
|`fruits[:]`|`["apple","banana","grape","melon"]`|semua element|
|`fruits[2:]`|`["grape","melon"]`|semua element yang dimulai dari index ke-2|
|`fruits[:2]`|`["apple","banana"]`|semua element sebelum index ke-2 dari akhir|

jika sebuah element pada slice diubah, maka element element yang memiliki referensi yang sama dengan element slice yang diubah tersebut, maka semuanya akan ikut berubah.

# built-in function bawaan untuk keperluan slice

## `len()`

digunakan untuk menghitung jumplah element yang ada, sebagai contoh jika sebuah variabel adalah slice dengan data 4 element, maka fungsi ini akan mengembalikan angka 4. selain digunakan untuk slice, bisa juga digunakan untuk array dan string.

## `cap()`

digunakan untuk menghitung lebar atau kapasitas maksimum slice.
Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan `len()`, tapi bisa berubah seiring dengan operasi yang dilakukan.

jika membuat sebuah slice baru dengan menggunakan metode 2 index slice[a:b], jika `a` lebih besar dari 0, maka kapasitas atau lebar maksimum dihitung dari panjang slice sumber dikurangi `a`, tapi jika `a` sama dengan 0, maka kapasitas atau lebar maksimum sama dengan slice sumber.

Tapi jika ingin melakukan slicing element yang sekaligus menentukan kapasitasnya, maka bisa dengan menggunakan metode 3 index `slice[a:b:c]`, dimana `c` adalah angka kapasitas, angka tersebut tidak boleh melebihi kapasitas slice yang akan di slicing.

## `append()`

fungsi ini digunakan untuk menambahkan element ke dalam slice, element baru tersebut diposisikan setelah element paling akhir.
fungsi ini akan mengembalikan slice baru yang sudah ditambahkan nilai baru-nya.

Jika jumplah element dan ukuran element sama, maka element baru tersebut merupakan referensi baru dan dimana ukuran atau size dari slice hasil dari fungsi `append()` tersebut akan dilipatkan 2x dari ukuran slice aslinya, maka jika terdapat penggunaan fungsi `append()` pada slice yang sama GO tidak akan mengkalkulasi memori lagi, yang membuat penggunaan memori menjadi lebih sedikit.

Tapi jika jumplah element lebih kecil dari ukuran element, maka element baru tersebut ditempatkan ke dalam cangkupan kapasitas-nya, membuat element slice dengan referensi yang sama akan berubah.


## `copy()`

fungsinya digunakan untuk men-copy element slice pada src (parameter ke-2), ke dst (parameter ke-1)

```go
copy(dst,src)
```

jumplah element yang akan di copy adalah sejumplah lebar dari `dst`, jika jumplah slice dari `src` lebih kecil dari `dst` maka akan tercopy semua.

jika `dst` sudah ada isinya maka isi dari `dst` akan di ganti dengan `src` sesuai dengan jumplah `dst`.

