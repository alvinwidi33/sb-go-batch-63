package main

import (
	"fmt"
	"strings"
	"strconv"
)


func main(){
	//soal 1
	var kataSatu = "Bootcamp"
	var kataDua = "Digital"
	var kataTiga = "Skill"
	var kataEmpat = "Sanbercode"
	var kataLima = "Golang"
	
	fmt.Println(kataSatu+" "+ kataDua+" "+ kataTiga+" "+kataEmpat+" "+kataLima)

	//soal 2
	halo := "Halo Dunia"
	var newText = strings.Replace(halo,"Dunia","Golang",-1)

	fmt.Println(newText)

	//soal 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	var kataKedua2 = strings.Replace(kataKedua,"s","S",1)
	var kataKetiga2 = strings.Replace(kataKetiga,"r","R",1)
	var kataKeempat2 = strings.ToUpper(kataKeempat)
	fmt.Println(kataPertama+" "+kataKedua2+" "+kataKetiga2+" "+kataKeempat2)

	//soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	var angkaPertama2, err1 = strconv.Atoi(angkaPertama)
	var angkaKedua2, err2 = strconv.Atoi(angkaKedua)
	var angkaKetiga2, err3 = strconv.Atoi(angkaKetiga)
	var angkaKeempat2, err4 = strconv.Atoi(angkaKeempat)

	if (err1 == nil && err2 == nil && err3 == nil && err4 == nil) {
		fmt.Println(angkaPertama2 + angkaKedua2 + angkaKetiga2 + angkaKeempat2)
	}

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var hi = strings.Replace(kalimat,"halo","Hi",2)
	var bandung = strings.Replace(hi,"b","B",1)

	newAngka := strconv.Itoa(angka)
	fmt.Println(bandung+" - "+newAngka)
}
