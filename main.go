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
	fmt.Println("               ===== PROGRAM TAGIHAN (SIMTAB) =====\n")
	fmt.Println(">>> INFORMASI PENGGUNAAN APLIKASI <<<")
	fmt.Println("Ketik Angka:")
	fmt.Println("1. untuk menambahkan data        6. untuk mengurutkan jatuh tempo terdekat")
	fmt.Println("2. untuk mengubah data           7. untuk menampilkan statistik")
	fmt.Println("3. untuk delete data             8. untuk menampilkan tagihan berdasarkan Kategori")
	fmt.Println("4. untuk mencari nama tagihan    9. untuk menghentikan program")
	fmt.Println("5. untuk menampilkan tagihan\n")

	for status {
		fmt.Print("Ketik : ")
		fmt.Scanln(&no)
		if no == 1 {
			functions.TambahData(&kategori, data, &jumlah_tagihan, &total)
		}else if no == 2 {
			functions.UbahData(data, jumlah_tagihan, &total)
		}else if no == 3 {

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
			var nomor int = 1
			for i := 0; i < len(kategori); i++ {
				fmt.Println("======================================")
				fmt.Printf("kategori %s:\n", kategori[i])
				fmt.Println("--------------------------------------")
				for j := 0; j < len(data); j++ {
					if data[j].Kategori == kategori[i] {
						fmt.Printf("%d. Tagihan          : %s\n", nomor, data[j].Nama_tagihan)
						fmt.Printf("   Nominal          : Rp%d\n", data[j].Nominal)
						fmt.Printf("   Jatuh Tempo      : %d\n", data[j].Jatuh_tempo)
						fmt.Printf("   Status Pelunasan : %s\n", data[j].Status_pelunasan)
						fmt.Println("--------------------------------------")
						nomor++
					}
				}
				nomor = 1
			}
			fmt.Println("======================================")
			fmt.Println()
		}else if no == 6 {
			for i := 0; i < jumlah_tagihan; i++ {
       			min := i
       			for j := i + 1; j < jumlah_tagihan; j++ {
           			if data[j].Jatuh_tempo < data[min].Jatuh_tempo {
                		min = j
            		}
       			}
        		data[i], data[min] = data[min], data[i]
    		}
		for i := 0; i < jumlah_tagihan; i++ {
		
			fmt.Printf("%d. Tagihan     : %s\n", i+1, data[i].Nama_tagihan)
			fmt.Println("   Nominal     : Rp", data[i].Nominal)
			fmt.Println("   Jatuh Tempo :", data[i].Jatuh_tempo)
		
		}
		}else if no == 7 {
			var status_lunas int
			for i := 0; i < jumlah_tagihan; i++ {
				if data[i].Status_pelunasan == "Lunas"{
					status_lunas++
				}
			}
			fmt.Printf("Total tagihan yang harus dibayar : Rp%d\n", total)
			fmt.Printf("Presentase tagihan yang lunas    : %.1f%%\n", float64(status_lunas)/float64(jumlah_tagihan) * 100.0)
		}else if no == 8 {
		}else if no == 9 {
			status = false
			fmt.Println("JANGAN LUPA BAYAR HUTANG :)")
		} else {
			fmt.Println("MASUKKAN NOMOR SESUAI KETENTUAN")
		}
		fmt.Println()
		no = 0
	}
}
