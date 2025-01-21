package webserver

import ("net/http"
		"encoding/json"
		"strconv"
		"log")

type NilaiMahasiswa struct{
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID uint
}
func Nilai()[]NilaiMahasiswa{
	nilais := []NilaiMahasiswa {

	}
	return nilais
} 
var nilaiNilaiMahasiswa = []NilaiMahasiswa{} 


// soal 1
func TentukanIndeksNilai(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	}
	return "E"
}
func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Nil NilaiMahasiswa
	if r.Method == "POST" {
	  if r.Header.Get("Content-Type") == "application/json" {
		decodeJSON := json.NewDecoder(r.Body)
		if err := decodeJSON.Decode(&Nil); err != nil {
		  log.Fatal(err)
		}
	  } else {
		id, _ := strconv.Atoi(r.PostFormValue("id"))
        nama := r.PostFormValue("nama")
        mataKuliah := r.PostFormValue("mata_kuliah")
        nilai, _ := strconv.Atoi(r.PostFormValue("nilai"))
		
		Nil = NilaiMahasiswa{
		   ID:    uint(id),
		   Nama: nama,
		   MataKuliah: mataKuliah,
		   Nilai: uint(nilai),
		   IndeksNilai: TentukanIndeksNilai(nilai),
		}
	  }
	  nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Nil)
	  dataNilai, _ := json.Marshal(Nil) 
	  w.Write(dataNilai)                
	  return
	}
  
	http.Error(w, "NOT FOUND", http.StatusNotFound)
}


// soal 2
func GetNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	  dataNilais, err := json.Marshal(nilaiNilaiMahasiswa)
	  
	  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	  }
  
	  w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
	  w.Write(dataNilais)
	  return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
  }