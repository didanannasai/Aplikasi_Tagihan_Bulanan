package main

import (
	"fmt"
	"code/functions"
	"code/model"
)

func main() {
	var status bool = true
	var no, jumlah_tagihan, total int
	var data = make([]model.Tagihan, 100)
	var kategori []string
	fmt.Println("\n                     PROGRAM TAGIHAN (SIMTAB) \n")
	fmt.Println(">>> INFORMASI PENGGUNAAN APLIKASI <<<")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Ketik Angka:")
	fmt.Println("1. untuk menambahkan data        5. untuk menampilkan tagihan")
	fmt.Println("2. untuk mengubah data           6. untuk mengurutkan jatuh tempo terdekat")
	fmt.Println("3. untuk delete data             7. untuk menampilkan statistik")
	fmt.Println("4. untuk mencari nama tagihan    8. untuk menghentikan program")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

	for status {
		fmt.Print("Ketik : ")
		fmt.Scanln(&no)
		if no == 1 {
			functions.TambahData(&kategori, data, &jumlah_tagihan, &total)
		}else if no == 2 {
			functions.UbahData(data, jumlah_tagihan, &total)
		}else if no == 3 {
			functions.HapusData(data, &jumlah_tagihan, &total)
		}else if no == 4 {
			var tipe string
			fmt.Println("Mencari berdasarkan? (Nama tagihan / Kategori)")
			fmt.Scanln(&tipe)
			if tipe == "Nama_tagihan" {
				functions.SearchNamaTagihan(data, jumlah_tagihan)
			}else if tipe == "Kategori" {
				functions.SearchKategori(data, kategori, jumlah_tagihan)
			}
		}else if no == 5 {
			functions.Menampilkan(data, kategori)
		}else if no == 6 {
			functions.Sorting(data, jumlah_tagihan)
		}else if no == 7 {
			functions.Statistik(data, jumlah_tagihan, total)
		}else if no == 8 {
			status = false
			fmt.Println("TERIMA KASIH :)")
		} else {
			fmt.Println("MASUKKAN NOMOR SESUAI KETENTUAN")
		}
		fmt.Println()
		no = 0
	}
}
