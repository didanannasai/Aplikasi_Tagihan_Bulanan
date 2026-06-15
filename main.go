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
			var setuju string
			if len(kategori) > 0 {
				fmt.Println("Daftar Kategori-Kategori: ")
				for i := 0; i < len(kategori); i++ {
					fmt.Printf("%d. %s\n", i + 1, kategori[i])
				}
				fmt.Println()
				fmt.Println("Tambah Kategori? (Y/n)")
				fmt.Scanln(&setuju)
				if setuju == "Y" {
					fmt.Print("Masukkan nama Kategori : ")
					fmt.Scanln(&data[jumlah_tagihan].Kategori)
					kategori = append(kategori, data[jumlah_tagihan].Kategori)
				}else if setuju == "n" {
					fmt.Print("Masukkan salah satu nama Kategori di atas: ")
					fmt.Scanln(&data[jumlah_tagihan].Kategori)
				}
			}else {
				fmt.Print("Masukkan nama Kategori : ")
				fmt.Scanln(&data[jumlah_tagihan].Kategori)
				kategori = append(kategori, data[jumlah_tagihan].Kategori)
			}
			fmt.Print("Nama tagihan        : ")
			fmt.Scanln(&data[jumlah_tagihan].Nama_tagihan)
			fmt.Print("Nominal tagihan     : ")
			fmt.Scanln(&data[jumlah_tagihan].Nominal)
			total += data[jumlah_tagihan].Nominal
			fmt.Print("Tanggal jatuh tempo : ")
			fmt.Scanln(&data[jumlah_tagihan].Jatuh_tempo)
			data[jumlah_tagihan].Status_pelunasan = "Belum Lunas"
			jumlah_tagihan++
		}else if no == 2 {
			functions.UbahData(data, jumlah_tagihan, &total)
		}else if no == 3 {

		}else if no == 4 {
			var nama string
			fmt.Print("Nama tagihan yang dicari: ")
			fmt.Scanln(&nama)
			// Sequential Search
			var found bool = false
			var x int = 0
			for x < jumlah_tagihan && !found {
				found = data[x].Nama_tagihan == nama
				x++
			}
			if found == true {
				fmt.Println("Nama tagihan SUDAH terdata")
			}else if found == false {
				fmt.Println("Nama tagihan BELUM terdata")
			}
		}else if no == 5 {
			for i := 0; i < jumlah_tagihan; i++ {
				fmt.Printf("%d. Tagihan          : %s\n", i+1, data[i].Nama_tagihan)
				fmt.Printf("   Nominal          : Rp%d\n", data[i].Nominal)
				fmt.Println("   Jatuh Tempo      :", data[i].Jatuh_tempo)
				fmt.Println("   Status Pelunasan :", data[i].Status_pelunasan)
				fmt.Println("   Kategori         :", data[i].Kategori)
			}
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
			var nomor int = 1
			for i := 0; i < len(kategori); i++ {
				fmt.Println("======================================")
				fmt.Printf("Kategori %s:\n", kategori[i])
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
