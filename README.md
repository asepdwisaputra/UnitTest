# Pengenalan Software Testing
- Disiplin ilmu dalam software eng.
- tujuan memastikan kualitas kode dan aplikasi baik(minim bug/kesalahan)
- Kala ini fokus UNIT TESTING
- Test Pyramid:
	```
	^ UI Test/EndtoEnd Test	   = Lambat - Mahal --> Fokus ke Seluruh sistem
	| Service/Integration Test = Sedang 	    --> Fokus ke 1 App(Backend)
	v Unit Test	      	   	   = Cepat  - Murah --> Fokus ke layer terkecil(DB ga)
	```

## UNIT TEST
- berfokus menguji kode terkecil(ex. method)
- kode Unit Test biasanya lebih banyak dari program biasanya
- bisa digunakan untuk meningkatkan kualitas kode


# Testing Package
- di GO unit test udah ada package <testing> dan bisa dijalankan di GO || https://golang.org/pkg/testing/
	+ `testing.T` --> struct unit test di GO
	+ `testing.M` --> mengatur life cycle testing, dibahas di chapter MAIN
	+ `testing.B` --> melakuksn Benchmark(mengukur kecepatan program)

# Membuat Unit Test
1. Buat Projek >> go mod init <UNITTEST> >> kadang <go mod tidy>
2. bikin package(ex helper) >> file hello_world.go >> bikin func Hello dengan parameter name
	```go
	package helper

	func Hello(name string) string {
		return "Hello " + name
	}
	```
3. lalu kita akan buat Unit Test func tadi dengan ATURAN:
	- kita WAJIB! menggunakan akhiran <_test> di file
	- BAGUSnya dan KONSISTEN <namaFile>_test.go (ex. hello_world_test.go)
	- ATURAN FUNC:
		+ func test diawali <Test> (ex. TestHello) biar KONSISTEN
		+ memiliki parameter (t *testing.T) dan tidak mengembalikan return value
			```go
			package helper
			import "testing"

			func TestHello(t *testing.T) {
				result := Hello("Asep")
				if result != "Hello Asep" {
					// Test Eror
					panic("Program Eror")
				}
			}
			```
4. buat hello_world_test.go di package yg sama(hal ini helper)
5. buat func test untuk func Hello
6. jalankan dengan cara:
	- `go test`			 		  --> me-run semua _test dalam package // `kalo di run di module tidak akan melihat _test di dalam pkg-pkg`
	- `go test -v`		 		  --> me run dengan melihat detail, (v = verbose)
	- `go test -v -run TestHello` --> cuma run func Hello === go test -v -run=TestHello // bisa juga `go test -v -run=TestHello`
	- `go test -v ./...` 		  --> run semua _test bahkan di dalam pkg
	- `cd ..`			 		  --> change directory  keluar folder
	- `cd helper`		 		  --> masuk directory/folder helper
	- `Note` Pastikan kamu me-run di dalam package ya, untuk beberapa kasus.
	
## Menggagalkan Unit Test selain <panic()>
Karena menggunakan panic() kurang baik Golang sudah menyediakan alternatif yang berada dalan package `t *testing.T` tadi:
+ `Fail()` 			==> menggagalkan namun tetap melanjutkan func, dan jika sudah selesai akan lanjut TestFunc lainnya
+ `FailNow()`		==> menggagalkan saat itu juga func yang berjalan, kemudian lanjut TestFunc selanjutnya
+ `Eror(args..)` 	==> setelah eror akan mencetak argumen, kemudian lanjut Fail()
+ `Fatal(args..)` 	==> setelah eror akan mencetak argumen, kemudian lanjut FailNow()
+ Contoh:
	```go
	func TestHelloGuys(t *testing.T) {
		result := Hello("Semua")
		if result != "Hello Aja" { // Disini kita jadikan eror
			// Test Eror
			t.Error("Program Eror") // Jika eror akan dicetak pesannya
		}
	}
	```

# Assertion by Testify
## Assert
- Pengecekan tanpa menggunakan if else yang menyebalkan
- karena GO tidak menyediakan assertion maka kita perlu menginput library/framework dari luar.
	+ TESTIFY || https://github.com/stretchr/testify
	+ Kita bisa menggunakan di Module --> go get github.com/stretchr/testify
	+ Silahkan eksplor sendiri, perhatikan bagian `import`, file `go.mod`, file `go.sum`
	
## Assert vs Require
- Testify menyediakan 2 package, Assert dan Require
- Assert 	=> mengembalikan `Fail()` jika eror
- Require 	=> mengembalikan `FailNow()` jika eror

## Skip Test(Membatalkan Tes)
- Terkadang ada keperluan untuk membatalkan sebuah test
- untuk membatalkan unit test kita bisa mengggunakan function `Skip()`
- jadi akan otomatis di hentikan, break
- ada di dalam struct `testing.T`
- Contoh:
	```go
	func TestHelloSkip(t *testing.T) {
		result := Hello("Hello2")
		if result == "Hello Hello2" {
			t.Skip("Tes ini dilewati")
		}
	}
	```

## Before-After Test
- di go terdapat fungsi `testing.M`
- fitur ini bernama Main, untuk mengatur eksekusi Unit Test.
- Namun, bisa digunakan untuk Before-After di unit test(mengakali aja)
	+ untuk mengatur eksekusi unit test, cukup buat func bernama `TestMain` dengan parameter `testing.M`
	+ Func hanya akan di eksekusi 1x per package!!
	+ di eksekusi pertama sebelum unit test
	+ Contoh:
		```go
		func TestMain(m *testing.M) {
			fmt.Println("Aplikasi Dimulai")
			m.Run() // Eksekusi semua UnitTest
			fmt.Println("Aplikasi Selesai")
		}
		```

## SubTest
- Go mendukung func unit test di dalam func unit test(seperti recursive yaa)
- Ini memang sedikit aneh dan jarang ada di bahasa lainnya
- kita bisa gunakan function `Run()`
- Contoh
	```go
	func TestSubtest(t *testing.T) {
		t.Run("Asep", func(t *testing.T) {
			result := Hello("Asep")
			assert.Equal(t, "Hello Asep", result)
		})
		t.Run("Putra", func(t *testing.T) {
			result := Hello("Putra")
			assert.Equal(t, "Hello Putra", result)
		})
	}
	```

### Hanya menjalankan SubTest
- Ketika ingin menjalankan sebuah unit Test
	```
	go test -run TestNamaFunction
	```
- Ketika ingin menjalankan HANYA salah 1 sub Test
	```
	go test -run TestNamaFunction/NamaSubTest
	```
- Ketika ingin menjalankan Semua UnitTest yang ber-SubTest khusus
	```
	go test -run /NamaSubTest
	```

## TableTest
- menggunakan table untuk menglakukan subTest
- TableTest sebenarnya menggunakan sebuah slice, yang berisi parameter dan ekspektasi hasil yang dijalankan
- nanti slice itu di implementasikan menggunakan SubTest
- Contoh
	```go
	func TestHelloTableTest(t *testing.T) {
		tes := []struct {
			nama       string
			permintaan string
			ekspektasi string
		}{
			{
				nama:       "HelloAsep",
				permintaan: "Asep",
				ekspektasi: "Hello Asep",
			},
			{
				nama:       "HelloPutra",
				permintaan: "Putra",
				ekspektasi: "Hello Putra",
			},
		}
		for _, v := range tes {
			t.Run(v.nama, func(t *testing.T) {
				result := Hello(v.permintaan)
				assert.Equal(t, v.ekspektasi, result)
			})
		}
	}
	```

# Mock
- adalah sebuah objek yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data sesuai yang kita program di awal
- Mock = salah 1 teknik UnitTest, kita bisa membuat MockObject dari suatu objek yang memang sulit di Testing
- EX: kita membuat UnitTest ternyata harus memanggil database ataupun API-Call ke 3rd party service(Lazada memanggil API Dana, tentu Lazada perlu terus-terus memanggil API Dana yang belum tentu sesuai hasilnya). Pada case ini, baiknya pake Mock Object.

## Testify Mock
- MockObject tidak ada di fitur bawaan GO.
- Alternatifnya kita pake library Testifyyang tadi untuk Assertion.
- Namun apabila desain code-nya jelek, sulit untuk di Mock.

### Contoh Aplikasi Query ke database
- case aplikasi Go yang melakukan Query ke database.
- Dimana kita akan membuat:
	+ Layer Service 	= sebagai Business Logic
	+ Layer Repository  = sebagai jembatan ke Database
- Agar code mudah di test, SARAN pake kontrak interface.
	1. Buat struct Category(ini entity=ibarat representasi tabel di DB), di package entity
	2. di file Category.go ya struct-nya
	3. Bikin Package Repository, Bikin Category_Repository.go, Bikin interface
	4. Isinya bisa dilihat di code-ya
	5. Bikin Package Service, Bikin Category_Service.go, Bikin struct CategoryService, yang butuh repo CategoryRepository
	6. Isinya bisa dilihat di code-ya
	7. selanjutnya kita buat mock-nya
	8. Ke Repository>>bikin Category_Repository_Mock.go>>bikin struct
	9. Isinya bisa dilihat di code-ya
	10. Ke service>>bikin unit test Category_Service_Test.go
	11. Isinya bisa dilihat di code-ya
	12. `Note` jadi, ada package entity>>didalamnya ada category.go || ada package repository>>didalamnya ada category_repository.go **Dan seterusnya - paham lah ya😂**

# Benchmark
- Package testing Go, juga mendukung Benchmark.
- Benchmark adalah menghitung kecepatan performa program kita.
- diatur otomatis oleh testing.B dari package testing
	+ testing.B adalah struct untuk melakukan benchmark
	+ hampir sama dengan testing.T(mengandung Fail(), FailNow(), Error(), Fatal(), dll namun mungkin akan jarang digunakan)
	+ testing.B memiliki atribut dan fungsi tambahan untuk benchmark
	+ salah satunya atribut N, untuk melakukan total iterasi

## Membuat Benchmark
- Nama func harus diawali 'Benchmark'(BenchmarkHelloWorld)
- harus memiliki parameter (b *testing.B)
- Tidak mengembalikan value
- Nama file diakhiri '_test' (hello_world_test.go)
- bisa digabung dengan file testing.T sebenarnya

## Menjalankan Benchnark
- Menjalankan seluruh benchmark dalam 1 module
	```
	go test -v -bench=.
	```
- Menjalankan Bench tanpa Unit-Test
	```
	go test -v - run=IsiNamaTestYangGaAda -bench=.
	```
- Menjalankan Benchmark tertentu
	```
	go test -v - run=IsiNamaTestYangGaAda -bench=BenchmarkNamaYangKamuIngin
	```
- Menjalankan Bench di Root Module, dan ingin menjalankan semua
	```
	go test -v -bench=. ./...
	```
## Sub Benchmark
- hampir sama dengan sub test
- menjalankan hanya sub bench
	```
	go test -v -bench=BenxhmarkNama/NamaSub
	```
## Table Benchmark
- hampir sama dengan table test
- agar bisa melakukan performace test dengan kombinasi data berbeda-beda tanpa harus bnyak membuat benchmark func