package main

import ("fmt"
		"strings"
		)

// soal 1
func updateLingkaran(luas *float64, keliling *float64, radius float64) {
    *luas = 3.14 * radius * radius
    *keliling = 2 * 3.14 * radius
}

// soal 2
func introduce(sentence *string, name string, gender string, job string, age string){
	var title string
	if (strings.ToLower(gender) == "laki-laki"){
		title = "Pak"
	} else {
		title = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, job, age)
}

// soal 3
func tambahBuah(buah *[]string) {
	dataBuah := []string{
		"Jeruk",
		"Semangka",
		"Mangga",
		"Strawberry",
		"Durian",
		"Manggis",
		"Alpukat",
	}
	*buah = append(*buah, dataBuah...)
}

func buahPrint(buah *[]string) string {
	var result []string
	for index, i := range *buah {
		result = append(result, fmt.Sprintf("%d. %s", index+1, i))
	}
	return strings.Join(result, "\n")
}

// soal 4
var dataFilm = []map[string]string{}
func tambahDataFilm(data *[]map[string]string, title, duration, genre, release string) {
	newFilm := map[string]string{
		"genre":  genre,
		"durasi": duration,
		"judul":  title,
		"tahun":  release,
	}

	*data = append(*data, newFilm)
}

func tampilkanDataFilm(data *[]map[string]string) {
	for _, film := range *data {
		fmt.Println(film)
	}
}

func main() {
	// soal 1
    var luasLingkaran float64
    var kelilingLingkaran float64
    var radius float64

    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scanln(&radius)

    updateLingkaran(&luasLingkaran, &kelilingLingkaran, radius)

    fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
    fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	// soal 2
	fmt.Println()
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence)

	// soal 3
	var buah []string
	tambahBuah(&buah)
	fmt.Println()
	fmt.Println(buahPrint(&buah))
	fmt.Println()

	// soal 4
	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")
	tampilkanDataFilm(&dataFilm)

}