package main

import (
	"fmt"
	"net/http"
)

func formData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Halaman Tidak Ditemukan.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "formTugas.html")

	case "POST":
		var kesalahan = r.ParseForm()
		if kesalahan != nil {
			fmt.Fprintln(w, "Ada Kesalahan : ", kesalahan)
			return
		}

		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		var prodi = r.FormValue("prodi")
		var kelas = r.FormValue("kelas")
		fmt.Fprintln(w, "NIM = ", nim)
		fmt.Fprintln(w, "NAMA = ", nama)
		fmt.Fprintln(w, "PRODI = ", prodi)
		fmt.Fprintln(w, "KELAS = ", kelas)

	default:
		fmt.Fprint(w, "Maaf. Method yang didukung hanya GET dan POST.")
	}
}

func main() {
	http.HandleFunc("/", formData)

	fmt.Println("Web berjalan di alamat http:/localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
