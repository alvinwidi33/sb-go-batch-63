package main

import ("fmt"
		"math"
		"strings")

// soal 1
type segitigaSamaSisi struct{
  alas, tinggi int
}

type persegiPanjang struct{
  panjang, lebar int
}

type tabung struct{
  jariJari, tinggi float64
}

type balok struct{
  panjang, lebar, tinggi int
}

type hitungBangunDatar interface{
  luas() int
  keliling() int
}

type hitungBangunRuang interface{
  volume() float64
  luasPermukaan() float64
}

func (p persegiPanjang) luas() int{
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) /2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (b balok) luasPermukaan() float64 {
	panjang := float64(b.panjang)
	lebar := float64(b.lebar)
	tinggi := float64(b.tinggi)
	return 2 * ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi))
}

func (b balok) volume() float64 {
	panjang := float64(b.panjang)
	lebar := float64(b.lebar)
	tinggi := float64(b.tinggi)
	return panjang * lebar * tinggi
}

// soal 2
type phone struct{
   name, brand string
   year int
   colors []string
}
type phoneInterface interface{
	detail() string
}

func (p phone) detail() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf(
		"name: %s\nbrand: %s\nyear: %d\ncolors: %s",
		p.name, p.brand, p.year, colors,
	)
}

//soal 3
func luasPersegi(sisi int, kalimat bool) interface{} {
	if kalimat {
		if sisi == 0 {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return fmt.Sprintf("Luas persegi adalah %d", sisi*sisi)
	}

	if sisi == 0 {
		return nil
	}
	return sisi * sisi
}

func main(){
	// soal 1
	var hitungBangunDatar hitungBangunDatar
	var hitungBangunRuang hitungBangunRuang

	var panjangPP, lebarPP, alasSG, tinggiSG int
	fmt.Print("Masukan panjang persegi panjang:")
	fmt.Scan(&panjangPP)

	fmt.Print("Masukan lebar persegi panjang:")
	fmt.Scan(&lebarPP)

	fmt.Print("Masukan alas segitiga:")
	fmt.Scan(&alasSG)

	fmt.Print("Masukan tinggi segitiga:")
	fmt.Scan(&tinggiSG)

	var panjangBalok, lebarBalok, tinggiBalok int
	fmt.Print("Masukan panjang balok:")
	fmt.Scan(&panjangBalok)

	fmt.Print("Masukan lebar balok:")
	fmt.Scan(&lebarBalok)

	fmt.Print("Masukan tinggi balok:")
	fmt.Scan(&tinggiBalok)

	var jariJari, tinggiTb float64
	fmt.Print("Masukan jari-jari tabung:")
	fmt.Scan(&jariJari)

	fmt.Print("Masukan tinggi tabung:")
	fmt.Scan(&tinggiTb)

	hitungBangunDatar = persegiPanjang{panjangPP, lebarPP}
	fmt.Println("\n===== Bangun Datar =====")
	fmt.Printf("Luas Persegi Panjang : %d\n", hitungBangunDatar.luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", hitungBangunDatar.keliling())

	hitungBangunDatar = segitigaSamaSisi{alasSG, tinggiSG}
	fmt.Printf("Luas Segitiga : %d\n", hitungBangunDatar.luas())
	fmt.Printf("Keliling Segtiga: %d\n\n", hitungBangunDatar.keliling())

	hitungBangunRuang = balok{panjangBalok, lebarBalok, tinggiBalok}
	fmt.Println("===== Bangun Ruang =====")
	fmt.Printf("Luas Permukaan Balok: %.2f\n", hitungBangunRuang.luasPermukaan())
	fmt.Printf("Volume Balok: %.2f\n", hitungBangunRuang.volume())

	hitungBangunRuang = tabung{jariJari, tinggiTb}
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", hitungBangunRuang.luasPermukaan())
	fmt.Printf("Volume Tabung: %.2f\n", hitungBangunRuang.volume())

	//soal 2
	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	var phoneInterface phoneInterface = myPhone
	fmt.Println(phoneInterface.detail())

	//soal 3
	fmt.Println()
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//soal 4
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	var semuaAngka []int
	semuaAngka = append(semuaAngka, kumpulanAngkaPertama.([]int)...)
	semuaAngka = append(semuaAngka, kumpulanAngkaKedua.([]int)...)

	var hasilPenjumlahan int
	output := fmt.Sprintf("\n%s", prefix)
	for i, angka := range semuaAngka {
		hasilPenjumlahan += angka
		if i == len(semuaAngka)-1 {
			output += fmt.Sprintf(" %d", angka)
		} else {
			output += fmt.Sprintf(" %d +", angka)
		}
	}
	output += fmt.Sprintf(" = %d", hasilPenjumlahan)

	fmt.Println(output)
}