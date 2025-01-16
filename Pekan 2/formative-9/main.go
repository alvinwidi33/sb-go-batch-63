package main

import (
	"fmt"
	"formative-9/variabel" )
func main(){
	// Soal 1
	var hitungBangunDatar variabel.HitungBangunDatar
	var hitungBangunRuang variabel.HitungBangunRuang

	var panjangPP, lebarPP, alasSG, tinggiSG int
	fmt.Print("Masukan panjang persegi panjang: ")
	fmt.Scan(&panjangPP)

	fmt.Print("Masukan lebar persegi panjang: ")
	fmt.Scan(&lebarPP)

	fmt.Print("Masukan alas segitiga: ")
	fmt.Scan(&alasSG)

	fmt.Print("Masukan tinggi segitiga: ")
	fmt.Scan(&tinggiSG)

	var panjangBalok, lebarBalok, tinggiBalok int
	fmt.Print("Masukan panjang balok: ")
	fmt.Scan(&panjangBalok)

	fmt.Print("Masukan lebar balok: ")
	fmt.Scan(&lebarBalok)

	fmt.Print("Masukan tinggi balok: ")
	fmt.Scan(&tinggiBalok)

	var jariJari, tinggiTb float64
	fmt.Print("Masukan jari-jari tabung: ")
	fmt.Scan(&jariJari)

	fmt.Print("Masukan tinggi tabung: ")
	fmt.Scan(&tinggiTb)

	hitungBangunDatar = variabel.PersegiPanjang{Panjang: panjangPP, Lebar: lebarPP}
	fmt.Println("\n===== Bangun Datar =====")
	fmt.Printf("Luas Persegi Panjang : %d\n", hitungBangunDatar.Luas())
	fmt.Printf("Keliling Persegi Panjang: %d\n", hitungBangunDatar.Keliling())

	hitungBangunDatar = variabel.SegitigaSamaSisi{Alas: alasSG, Tinggi: tinggiSG}
	fmt.Printf("Luas Segitiga : %d\n", hitungBangunDatar.Luas())
	fmt.Printf("Keliling Segtiga: %d\n\n", hitungBangunDatar.Keliling())

	hitungBangunRuang = variabel.Balok{Panjang: panjangBalok, Lebar: lebarBalok, Tinggi: tinggiBalok}
	fmt.Println("===== Bangun Ruang =====")
	fmt.Printf("Luas Permukaan Balok: %.2f\n", hitungBangunRuang.LuasPermukaan())
	fmt.Printf("Volume Balok: %.2f\n", hitungBangunRuang.Volume())

	hitungBangunRuang = variabel.Tabung{JariJari: jariJari, Tinggi: tinggiTb}
	fmt.Printf("Luas Permukaan Tabung: %.2f\n", hitungBangunRuang.LuasPermukaan())
	fmt.Printf("Volume Tabung: %.2f\n", hitungBangunRuang.Volume())

	// Soal 2
	myPhone := variabel.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	var phoneInterface variabel.PhoneInterface = myPhone
	fmt.Println(phoneInterface.Detail())

	// Soal 3
	fmt.Println()
	fmt.Println(variabel.LuasPersegi(4, true))
	fmt.Println(variabel.LuasPersegi(8, false))
	fmt.Println(variabel.LuasPersegi(0, true))
	fmt.Println(variabel.LuasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

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
