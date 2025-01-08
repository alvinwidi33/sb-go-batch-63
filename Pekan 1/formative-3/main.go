package main
import (
	"fmt"
	"strconv"
)

func main(){
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjangPersegiPanjang2, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjang2, _ = strconv.Atoi(lebarPersegiPanjang)
	var alasSegitiga2, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitiga2, _ = strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = panjangPersegiPanjang2 * lebarPersegiPanjang2
	kelilingPersegiPanjang = (2 * panjangPersegiPanjang2) + (2 * lebarPersegiPanjang2)
	luasSegitiga = (alasSegitiga2 * tinggiSegitiga2) / 2

	fmt.Println("Luas persegi panjang :", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang :", kelilingPersegiPanjang)
	fmt.Println("Luas segitiga :", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	getIndeks := func(nilai int) string {
		if (nilai >= 80){
			return "A"
		} else if (nilai >= 70 && nilai < 80){
			return "B"
		} else if (nilai >= 60 && nilai < 70){
			return "C"
		} else if (nilai >= 50 && nilai < 70){
			return "D"
		} else if (nilai < 50){
			return "E"
		}
		return "Nilai tidak valid"
	}
	fmt.Println("Indeksnya John", getIndeks(nilaiJohn))
	fmt.Println("Indeksnya Doe", getIndeks(nilaiDoe))

	// soal 3
	var tanggal = 3;
	var bulan = 3;
	var tahun = 2003;

	getBulan := func(a int) string {
		switch a {
		case 1:
			return "Januari"
		case 2:
			return "Februari"
		case 3:
			return "Maret"
		case 4 :
			return "April"
		case 5:
			return "Mei"
		case 6:
			return "Juni"
		case 7:
			return "July"
		case 8:
			return "Agustus"
		case 9:
			return "September"
		case 10:
			return "Oktober"
		case 11:
			return "November"
		case 12:
			return "Desember"
		default:
			return "Bulan tidak valid"
		}
	}
	fmt.Println(tanggal,getBulan(bulan),tahun)

	//soal 4
	if (tahun >= 1944 && tahun <= 1964){
		fmt.Println("Baby boomer")
	} else if (tahun >= 1965 && tahun <= 1979){
		fmt.Println("Generasi X")
	} else if (tahun >= 1980 && tahun <= 1994){
		fmt.Println("Generasi Y (Millenials)")
	} else if (tahun >= 1995 && tahun <= 2015){
		fmt.Println("Generasi Z")
	}
}