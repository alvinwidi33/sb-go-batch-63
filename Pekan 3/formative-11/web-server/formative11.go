package webserver

import (
	"fmt"
	"math"
	"net/http"
)

// soal 4
func luasAlas(jariJari int) string {
	return fmt.Sprintf("luas alas: %.2f", math.Pi*math.Pow(float64(jariJari), 2))
}

func kelilingAlas(jariJari int) string {
	return fmt.Sprintf("keliling alas: %.2f", math.Pi*2*float64(jariJari))
}

func volume(jariJari int, tinggi int) string {
	return fmt.Sprintf("volume: %.2f", math.Pi*math.Pow(float64(jariJari), 2)*float64(tinggi))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jariJari := 7
	tinggi := 10

	fmt.Fprintf(w, "jariJari: %d, tinggi: %d, volume: %s, luas alas: %s, keliling alas: %s\n",
		jariJari, tinggi, volume(jariJari, tinggi), luasAlas(jariJari), kelilingAlas(jariJari))
}


