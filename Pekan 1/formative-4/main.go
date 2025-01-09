package main

import (
	"fmt"
)

func main(){
	// soal 1
	for i := 1; i < 21; i++ {
		if  i % 2 == 1 {
			if i % 3 == 0 {
				fmt.Println(i, "- I Love Coding")
			} else {
				fmt.Println(i, "- Santai")
			}
		} else {
			fmt.Println(i, "- Berkualitas")
		}
	}

	// soal 2
	for i := 0; i < 8 ; i++ {
		for j := 0 ; j < i ; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println()
	fmt.Println(kalimat[2:6])

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	fmt.Println()
	for i, sayuran2 := range sayuran{
		fmt.Printf("%d. %s\n",i + 1 , sayuran2)
	}

	// soal 5
	var satuan = map[string]int{
	"panjang": 7,
	"lebar":   4,
	"tinggi":  6,
	}
	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Println()
	fmt.Printf("Panjang = %d\nLebar = %d\nTinggi = %d\nVolume = %d", satuan["panjang"], satuan["lebar"], satuan["tinggi"], volume)
}