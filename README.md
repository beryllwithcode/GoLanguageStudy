<div align="center">
    <a href="https://github.com/beryllwithcode/golang-study">
    <img src="https://cdn.icon-icons.com/icons2/2699/PNG/512/golang_logo_icon_171073.png" alt="Logo" width="150" height="150">
    </a>
    <h1 align="center">Study Golang</h1>
    <p align="center">This repository contains about learning Go programming language</p>
</div>

## Fundamental Golang
<h3>Hello World</h3>
Sample code Hello World in Golang:

  ```go
    package main

    // Import Package fmt
    import "fmt"

    func main() {
    //Print "Hello, World!" on console
    fmt.println("Hello, World!")
    }
   ```
- Di Go, setiap program adalah bagian dari paket. kata kunci ``package``. Dalam contoh ini, Program milik ``main`` package. 
- ``import ("fmt")`` bertujuan untuk mengimport file yang termasuk dalam ``fmt`` package.
- ``fmt.println()`` adalah fungsi yang tersedia dari ``fmt`` package. Ini digunakan untuk menampilkan/mencetak text. Dalam contoh kita, hasil output nya akan menghasilkan Hello, World!

<h3>Compile File Go-lang</h3>
<h4>(Go source) -> (Compiler) -> (Executable Binary)</h4>

Untuk compile suatu file dapat menjalankan perintah: <br>
``go build nama_file.go`` <br>
atau jika ingin menjalankan suatu file tanpa compile, jalankan perintah: <br>
``go run nama_file.go``

<h3>Tipe data</h3> 

- Number (Integer, Float)
```go
    func main() {
	    fmt.Println("one", 1) //output: one, 1
	    fmt.Println("two", 2) //output: two, 2
	    fmt.Println("one point six", 1.6) //output: one point six, 1.6
    }
```
- Boolean (True or False)
```go
    func main() {
        fmt.Println(true) //output: true
        fmt.Println(false) //output: false
    }
```
- Alias
<table>
    <tr>
        <th>Tipe</th>
        <th>Alias</th>
    </tr>
    <tr>
        <td>byte</td>
        <td>uint8</td>
    </tr>
    <tr>
        <td>rune</td>
        <td>int32</td>
    </tr>
    <tr>
        <td>any</td>
        <td>interface{}</td>
    </tr>
</table>

- String 
    <p>1. String adalahh tipe data dari kumpulan karakter</p>
    <p>2. Tipe data String di Go-Lang diwakili dengan kata kunci string</p>
    <p>3. Nilai data string di Go-Lang selalu dimulai dan di akhiri dengan (tanda kutip ganda)-> "Ini string" <-(Tanda kutip ganda)</p>
```go
    func main() {
        fmt.Println("Beryll") //output: Beryll
        fmt.Println("Pradana") //output: Pradana
        fmt.Println("Ramadhan") //output: Ramadhan
    }
```

<h3>Variable</h3>

- Variable adalah tempat untuk menyimpan suatu data. <br>
Untuk membuat variable, kita dapat menggunakan keyword ``var``, kemudian di ikuti dengan nama variable dan tipe data.
```go
// Example 1

    func main() {
        //Declaration Variable
        var name string
        var age int

        //Assign value to variable name & age
        name = "Beryll Pradana Ramadhan"
        age = 21

        //Print value of variable
        fmt.Println(name) //output: Beryll Pradana Ramadhan
        fmt.Println(age) //output: 21
    }
```
```go
// Example 2

    func main() {
        //Assign value to variable name & age
        name := "Beryll Pradana Ramadhan"
        age := 21

        //Print value of variable
        fmt.Println(name) //output: Beryll Pradana Ramadhan
        fmt.Println(age) //output: 21
```
```go
// Example 3

    func main() {
        //Declaration Variable
	var born, now

	//Assign value to Variable born & now
	born = 2003
	now = 2024

	//Print value of Variable
	fmt.Println("I was born on ", born) //output: I was born on 2003
	fmt.Println("My age is ", now - born) //output: My age is 21
```

- Multi Variable
Go mendukung metode deklarasi banyak secara bersamaan, caranya adalah dengan menuliskan variable nya dengan pembatas tanda koma ( , ). Untuk pengisian value nya juga bisa secara bersamaan.
```go
// Example 1

	var first, second, third string
	first, second, third = "satu", "dua", "tiga" 
```
Dengan menggunakan teknik type inference, deklarasi multi variable bisa dilakukan untuk variable-variable yang memiliki tipe data yang berbeda.
```go
// Example 2

	var one, isFriday, twoPointTwo, say
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello!"	
```

- Reserved Variable
Variable _(Underscore), Go punya aturan yang cukup unik yang dimana mengharuskan harus menggunakan semua variable yang telah di deklarasikan, jika tidak maka akan terjadi error. Reserved Varianle biasa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa di bilang variable ini merupakan keranjang sampah.
```go
_ = "Study Jam Golang"
_ = "Belajar Golang itu Mudah!"
name, _ = "Beryll", "Belajar Golang"
```

- Variable Declaration Using Keyword ``new``
Fungsi ``new()`` digunakan untuk membuat variable pointer dengan tipe data tertentu. Nilai data default nya akan menyesuaikan dengan tipe datanya.
```go
name := new(string)

fmt.Println(name) //output: 0x2081a220
fmt.Pritnln(*name) //output: ""
```
Variable ``name`` menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilai melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal). Untuk menampilkan nilai aslinya, variable tersebut perlu di Dereference terlebih dahulu, caranya dengan menuliskan tanda asterisk ( * ) sebelum nama variabel.

<h3>Konstanta</h3>

Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah di deklarasikan. Inisialisasi nilai konstanta hanya di lakukan sekali saja di awal, setelah itu variable tidak bisa di ubah nilai nya. <br>
Cara membuat sebuah variable konstanta sama seperti mendeklarasikan variable biasa, perbedaan nya hanya pada keyword yang di gunakan, yaitu ``const`` (bukan ``var``).
```go
const firstName string = Beryll
fmt.Print("Halo nama saya ", firstName, "!\n") //output: Halo nama saya Beryll!
```
Teknik type inference juga bisa di terapkan pada konstanta, cara nya cukup dengan menghilangkan tipe data pada variable saat di deklarasikan.
```go
const firstName = Beryll
const midName = Pradana
const lastName = Ramadhan
fmt.Print("Halo nama lengkap saya "), firstName, " ", midName, " ", lastName, "!\n" //output: Halo nama lengkap saya Beryll Pradana Ramadhan!
```
<br>

Penggunaan fungsi ``fmt.Print()`` <br>
Fungsi ini memiliki peran yang sama seperti fungsi ``fmt.Println()``, perbedaan fungsi ``fmt.Print()`` tidak menghasilkan baris baru di akhir output nya. <br>
Perbedaan lainnya: nilai argument parameter yang ditulis saat pemanggilan fungsi akan di print tanpa pemisah. Tidak seperti pada fungsi ``fmt.Println()`` yang nilai argument parameter nya di pisah menggunakan karakter spasi.

```go
fmt.Println("Beryll Pradana Ramadhan")
fmt.Println("Beryll", "Pradana", "Ramadhan")

fmt.Print("Beryll Pradana Ramadhan\n")
fmt.Print("Beryll ", "Pradana ", "Ramadhan\n")
fmt.Print("Beryll", " ", "Ramadhan"\n)
```

<h3>Looping</h3>
Looping adalah proses mengulang-ulang suatu eksekusi blok tanpa henti, selama kondisi yang di jadikan acuan terpenuhi. Biasanya disiapkan variable untuk iterasi atau variable penanda kapan perulangan akan diberhentikan.
<br>

Di Go keyword perulangan hanya ``for`` saja, tetapi meski demikian, kemampuannya merupakan gabungan dari ``for``, ``foreach``, dan ``while`` ibarat bahasa pemrograman lain.
- Looping menggunakan keyword ``for``

```go
// Example 1

	for i := i = 0; i < 5; i++ {
	    fmt.Println("Angka", i) //output: Angka 1, Angka 2, Angka 3, Angka 4, Angka 5
	}
```
- Penggunaan Keyword ``for`` dengan Argumen hanya Kondisi
Cara ke-2 adalah dengan menuliskan kondisi setelah keyword ``for`` (hanya kondisi). Deklarasi dan iterasi variable counter tidak di tuliskan setelah keyword, hanya kondisi perulangan saja. Konsep nya mirip seperti ``while`` milik bahasa lain.
<br>

Kode berikut adalah contoh ``for`` dengan argumen hanya kondisi (seperti ``if``), output yang dihasilkan akan sama seperti penerapan for pertama.
```go
// Example 2

	var i = 0

	for i < 5 {
	   fmt.Println("Angka", 1) //output: Angka 1, Angka 2, Angka 3, Angka 4, Angka 5
	   i++	
	}
```

- Penggunaan Keyword ``for`` Tanpa Argumen
Cara ke-3 adalah ``for`` di tulis tanpa kondisi. Dengan ini akan dihasilkan perulangan tanpa henti (sama dengan ``for true``). Pemberentian perulangan dilakukan dengan menggunakan keyword ``break``.
```go
// Example 3

	var i = 0

	for {
	   fmt.Println("Angka", i)

	   i++
	   if i == 5 {
		break
	  }
	}
```
Dalam perulangan tanpa henti di atas, variable i yang nilai awalnya 0 di-ingkrementasi. Ketika nilai ``i`` sudah mencapai ``5``, keyword ``break`` akan digunakan, dan Looping akan berhenti. 

- Penggunaan Keyword ``for - range``
cara ke-4 adalah perulangan dengan menggunakan kombinasi keyword ``for`` dan ``range``. Cara ini biasanya digunakan untuk melooping data gabungan (misalnya string, array, slice, map).
```go
// Example 4

	var xs = "123" // String
	for i, v := range xs {
	   fmt.Println("Index=", i, "value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} //Array
	for _, v := range ys {
	   fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // Slice
	for _, v := range zs {
	   fmt.println("Value", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2,} // map
	for k, v := range kvs {
	   fmt.Println("key=", k, "Value", v)
	}

	// Boleh juga baik k atau v nya di abaikan
	for range kvs {
	   fmt.Println("Done!")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik bilangan
	for i := range 5 {
	   fmt.Println(i) //output: 0, 1, 2, 3, 4
	}
```
- Penggunaan keyword ``break`` & ``continue``
Keyword ``break`` digunakan untuk memberhentikan paksa sebuah perulangan, sedangkan ``continue`` di pakai untuk memaksa maju perulangan berikutnya.
<br>
Berikut merupakan contoh penerapan ``continue`` dan ``break``. Kedua keyword tersebut dimanfaatkan untuk menampilkan angka genap berurutan yang lebih besar dari 0 dan kurang dari atau sama dengan 8.
```go
	for i := 1; 1 <= 10; i++ {
	  if i % 2 == 1 
		continue
	  }

	  if i > 8 {
		break
	  }

	  fmt.Println("Angka", i) //output: Angka 2, Angka 4, Angka 6, Angka 8
```
Penjelasan kode di atas:
1. Melakukan perulangan mulai angka 1 hingga 10 dengan ``i`` sebagai variable iterasi.
2. Ketika ``i`` adalah ganjil (dapat diketahui dari ``1 % 2``, jika hasil nya ``1``, berati ganjil), maka akan dipaksa lanjut ke perulangan berikutnya.
3. Ketika ``i`` lebih besar dari 8, maka perulangan akan berhenti.
4. Nilai ``i`` ditampilkan.

- Perulangan Bersarang
Tidak hanya kondisi yang bisa bersarang, perulangan juga bisa. Cara pengaplikasiannnya kurang lebih sama, tinggal tulis blok statement perulangannya di dalam perulangan.
```go
// Example

	for i:=0; i<5; i++ {
	   for j:=i; j<5; j++ {
		fmt.print(j, " ")
	   }

	   fmt.Println() //output: 	0 1 2 3 4
					1 2 3 4
					2 3 4
					3 4
					4
					
	}
```
Pada kode diatas, untuk pertama kalinya fungsi ``fmt.Println`` di panggil tanpa disisipkan parameter. Cara seperti ini bisa digunakan untuk menampilkan baris baru. Kegunaannya sama seperti output dari statement ``fmt.Print("\n")``.


<a href="https://dasarpemrogramangolang.novalagung.com/A-tipe-data.html">Click untuk penjelasan lengkap nya!</a>
