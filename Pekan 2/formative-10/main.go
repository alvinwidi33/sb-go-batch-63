package main

import (
	"fmt"
	"sort"
	"time"
	"sync"
)

// soal 1
func numberOne(kalimat string, tahun int){
	fmt.Printf("%s %d", kalimat, tahun)
}

func runNumberOne(){
	defer numberOne("Golang Backend Development", 2021)

}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, param bool) string{
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	if (param){
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, sisi * 3)
	} else if (!param){
		return fmt.Sprintf("%d", sisi * 3)
	} else if (sisi == 0 && param){
		return "Maaf anda belum menginput sisi dari segitiga sama sisi"
	}
	panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
}

// soal 3
func tambahAngka(angka1 int, angka2 *int){
	angka1 += *angka2
}

func cetakAngka(angka *int){
	fmt.Printf("\nNilai cetak angka %d",*angka)
}


//soal 4
func add(phones *[]string, data ...string) {
	*phones = append(*phones, data...)
}

// soal 6
func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}
	close(moviesChannel)
}

func main(){
	// soal 1
	fmt.Println("Program dimulai...")
	runNumberOne()
	fmt.Println("\nProses lainnya selesai!")

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	angka := 1

	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// soal 4
	var phones = []string{}

	add(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")

	sort.Strings(phones)
	fmt.Println()
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}

	// soal 5
	var phones2 = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	fmt.Println()
	sort.Strings(phones2)

	var wg sync.WaitGroup
	for i, phone := range phones2 {
		wg.Add(1)
		go func(index int, name string) {
			defer wg.Done()
			time.Sleep(time.Duration(index) * time.Second)
			fmt.Printf("%d. %s\n", index+1, name)
		}(i, phone)
	}

	wg.Wait()

	// soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)
	fmt.Println()
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}
}