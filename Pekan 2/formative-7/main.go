package main

import "fmt"

// soal 1
type Buah struct {
	nama, warna string
	adaBijinya bool
	harga int
}

// soal 2
type segitiga struct{
  alas, tinggi int
}

type persegi struct{
  sisi int
}

type persegiPanjang struct{
  panjang, lebar int
}

func (s segitiga) luasSegitiga(){
	fmt.Printf("Luas segitiga : %d\n", s.alas * s.tinggi)
}

func (p persegi) luasPersegi(){
	fmt.Printf("Luas persegi : %d\n", p.sisi * p.sisi)
}

func (pp persegiPanjang) luasPersegiPanjang(){
	fmt.Printf("Luas persegi panjang : %d\n\n", pp.panjang * pp.lebar)
}

// soal 3
type phone struct{
   name, brand string
   year int
   colors []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// soal 4
type movie struct{
	title, genre string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, data *[]movie) {
	newFilm := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}

	*data = append(*data, newFilm)
}

func tampilkanDataFilm(data []movie) {
	for i, film := range data {
		jam := film.duration / 60
		fmt.Printf("%d. Title: %s\n   Duration:%d jam\n   Genre: %s\n   Year: %d\n",
			i+1, film.title, jam, film.genre, film.year)
	}
}

func main(){
	// soal 1
	var nanas = Buah{"Nanas", "Kuning", false, 9000}
	var jeruk = Buah{"Jeruk", "Oranye", true, 8000}
	var semangka = Buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = Buah{"Pisang", "Kuning", false, 5000}

	fmt.Printf("Buah 1 adalah: %s berwarna %s ada bijinya %t seharga %d\n", nanas.nama, nanas.warna, nanas.adaBijinya, nanas.harga)
	fmt.Printf("Buah 2 adalah: %s berwarna %s ada bijinya %t seharga %d\n", jeruk.nama, jeruk.warna, jeruk.adaBijinya, jeruk.harga)
	fmt.Printf("Buah 3 adalah: %s berwarna %s ada bijinya %t seharga %d\n", semangka.nama, semangka.warna, semangka.adaBijinya, semangka.harga)
	fmt.Printf("Buah 4 adalah: %s berwarna %s ada bijinya %t seharga %d\n\n", pisang.nama, pisang.warna, pisang.adaBijinya, pisang.harga)

	// soal 2
	var alas int
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scanln(&alas)

	var tinggi int
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)
	var segitiga = segitiga{alas, tinggi}
	segitiga.luasSegitiga()

	var sisi int
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)
	var persegi = persegi{sisi}
	persegi.luasPersegi()

	var panjang int
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scanln(&panjang)
	var lebar int
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scanln(&lebar)
	var persegiPanjang = persegiPanjang{panjang, lebar}
	persegiPanjang.luasPersegiPanjang()

	// soal 3
	var name string
	fmt.Print("Masukkan name: ")
	fmt.Scanln(&name)

	var brand string
	fmt.Print("Masukkan brand: ")
	fmt.Scanln(&brand)

	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&year)
	myPhone := phone{
		name:   name,
		brand:  brand,
		year:   year,
		colors: []string{"Black", "White"},
	}

	var newColor string
	fmt.Print("Masukkan warna baru: ")
	fmt.Scanln(&newColor)

	myPhone.addColor(newColor)

	fmt.Printf("\nDetail handphone:\nNama: %s\nBrand: %s\nTahun: %d\nWarna: %v\n\n", myPhone.name, myPhone.brand, myPhone.year, myPhone.colors)

	// soal 4
	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	tampilkanDataFilm(dataFilm)
}