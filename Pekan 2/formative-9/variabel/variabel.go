package variabel

import ("fmt"
		"math"
		"strings")

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (b Balok) LuasPermukaan() float64 {
	panjang := float64(b.Panjang)
	lebar := float64(b.Lebar)
	tinggi := float64(b.Tinggi)
	return 2 * ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi))
}

func (b Balok) Volume() float64 {
	panjang := float64(b.Panjang)
	lebar := float64(b.Lebar)
	tinggi := float64(b.Tinggi)
	return panjang * lebar * tinggi
}

// Soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneInterface interface {
	Detail() string
}

func (p Phone) Detail() string {
	colors := strings.Join(p.Colors, ", ")
	return fmt.Sprintf(
		"name: %s\nbrand: %s\nyear: %d\ncolors: %s",
		p.Name, p.Brand, p.Year, colors,
	)
}

// Soal 3
func LuasPersegi(sisi int, kalimat bool) interface{} {
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