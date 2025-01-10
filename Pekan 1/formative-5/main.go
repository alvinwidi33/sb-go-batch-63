package main

import (
	"fmt"
	"strings"
)

// soal 1
func luasPersegiPanjang(panjang int, lebar int) int{
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int, lebar int) int{
	return 2 * (panjang + lebar)
}
func volumeBalok(panjang int, lebar int, tinggi int) int{
	return panjang * lebar * tinggi
}

//soal 2
func introduce(name string, gender string, job string, age string) string {
	var title string
	if (strings.ToLower(gender) == "laki-laki"){
		title = "Pak"
	} else {
		title = "Bu"
	}
	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, job, age)
}

// soal 3
func buahFavorit(name string, buah ...string) string {
	list := ""
	for index, i := range buah {
		if (index == len(buah) - 1){
			list += "\""+i+"\""
		} else {
			list += "\""+i+"\", "
		}
	}
	return fmt.Sprintf("halo nama saya %s buah favorit saya adalah %s",name, list)
}

// soal 4
var dataFilm = []map[string]string{}
func tambahDataFilm(title, duration, genre, release string) {
	newFilm := map[string]string{
		"genre":  genre,
		"durasi": duration,
		"judul":  title,
		"tahun":  release,
	}
	dataFilm = append(dataFilm, newFilm)
}

func main(){
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Printf("Luas persegi panjang: %d\n", luas)
	fmt.Printf("Keliling persegi panjang: %d\n",keliling)
	fmt.Printf("Volume balok: %d\n\n", volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println()
	fmt.Println(buahFavoritJohn)

	fmt.Println()

	// soal 4
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
	for _, item := range dataFilm {
		fmt.Println(item)
	}
}