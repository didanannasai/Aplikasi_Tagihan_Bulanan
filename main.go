package main

import (
	"fmt"
)

type Tagihan struct {
	nama_tagihan, status_pelunasan, kategori string
	nominal, jatuh_tempo int
}

func main() {
	var status bool = true
	var no, jumlah_tagihan, urutan, total int
	var data = make([]Tagihan, 100)
	var kategori []string
	fmt.Println("            ===== PROGRAM TAGIHAN (SIMTAB) =====\n")
	fmt.Println(">>> INFORMASI PENGGUNAAN APLIKASI <<<")
	fmt.Println("Ketik Angka:")
	fmt.Println("1. untuk menambahkan data        5. untuk menampilkan tagihan")
	fmt.Println("2. untuk mengubah data           6. untuk mengurutkan jatuh tempo terdekat")
	fmt.Println("3. untuk delete data             7. untuk menampilkan statistik")
	fmt.Println("4. untuk mencari nama tagihan    8. untuk menghentikan program\n")

	for status {
		fmt.Print("Ketik : ")
		fmt.Scanln(&no)
		if no == 1 {
			var setuju string
			if len(kategori) > 0 {
				fmt.Println("Daftar kategori-kategori: ")
				for i := 0; i < len(kategori); i++ {
					fmt.Printf("%d. %s\n", i + 1, kategori[i])
				}
				fmt.Println()
				fmt.Println("Tambah kategori? (Y/n)")
				fmt.Scanln(&setuju)
				if setuju == "Y" {
					fmt.Print("Masukkan nama kategori : ")
					fmt.Scanln(&data[jumlah_tagihan].kategori)
					kategori = append(kategori, data[jumlah_tagihan].kategori)
				}else if setuju == "n" {
					fmt.Print("Masukkan salah satu nama kategori di atas: ")
					fmt.Scanln(&data[jumlah_tagihan].kategori)
				}
			}else {
				fmt.Print("Masukkan nama kategori : ")
				fmt.Scanln(&data[jumlah_tagihan].kategori)
				kategori = append(kategori, data[jumlah_tagihan].kategori)
			}
			fmt.Print("Nama tagihan        : ")
			fmt.Scanln(&data[jumlah_tagihan].nama_tagihan)
			fmt.Print("Nominal tagihan     : ")
			fmt.Scanln(&data[jumlah_tagihan].nominal)
			total += data[jumlah_tagihan].nominal
			fmt.Print("Tanggal jatuh tempo : ")
			fmt.Scanln(&data[jumlah_tagihan].jatuh_tempo)
			data[jumlah_tagihan].status_pelunasan = "Belum Lunas"
			jumlah_tagihan++
		}else if no == 2 {
			var n string
			fmt.Print("Masukkan urutan Tagihan yang ingin diubah: ")
			fmt.Scanln(&urutan)
			fmt.Println("Data apa yang ingin diubah? (Nama/Nominal/Tanggal/Status)")
			fmt.Print("Data : ")
			fmt.Scanln(&n)
			if n == "Nama" {
				fmt.Print("Nama tagihan : ")
				fmt.Scanln(&data[urutan - 1].nama_tagihan)
			}else if n == "Nominal" {
				total -= data[urutan - 1].nominal
				fmt.Print("Nominal tagihan : ")
				fmt.Scanln(&data[urutan - 1].nominal)
				total += data[urutan - 1].nominal
			}else if n == "Tanggal" {
				fmt.Print("Tanggal jatuh tempo : ")
				fmt.Scanln(&data[urutan - 1].jatuh_tempo)
			}else if n == "Status" {
				fmt.Print("Status pembayaran : ")
				fmt.Scanln(&data[urutan - 1].status_pelunasan)
			}
		}else if no == 3 {

		}else if no == 4 {
			var nama string
			fmt.Print("Nama tagihan yang dicari: ")
			fmt.Scanln(&nama)
			// Sequential Search
			var found bool = false
			var x int = 0
			for x < jumlah_tagihan && !found {
				found = data[x].nama_tagihan == nama
				x++
			}
			if found == true {
				fmt.Println("Nama tagihan SUDAH terdata")
			}else if found == false {
				fmt.Println("Nama tagihan BELUM terdata")
			}
		}else if no == 5 {
			for i := 0; i < jumlah_tagihan; i++ {
				fmt.Printf("%d. Tagihan          : %s\n", i+1, data[i].nama_tagihan)
				fmt.Printf("   Nominal          : Rp%d\n", data[i].nominal)
				fmt.Println("   Jatuh Tempo      :", data[i].jatuh_tempo)
				fmt.Println("   Status Pelunasan :", data[i].status_pelunasan)
				fmt.Println("   Kategori         :", data[i].kategori)
			}
		}else if no == 6 {
			 for i := 0; i < jumlah_tagihan; i++ {
       		 min := i
       		 for j := i + 1; j < jumlah_tagihan; j++ {
           		 if data[j].jatuh_tempo < data[min].jatuh_tempo {
                	min = j
            	}
       		 }
        	data[i], data[min] = data[min], data[i]
    	}
		for i := 0; i < jumlah_tagihan; i++ {
		
		fmt.Printf("%d. Tagihan     : %s\n", i+1, data[i].nama_tagihan)
		fmt.Println("   Nominal     : Rp", data[i].nominal)
		fmt.Println("   Jatuh Tempo :", data[i].jatuh_tempo)
		
		}
		}else if no == 7 {
			var status_lunas int
			for i := 0; i < jumlah_tagihan; i++ {
				if data[i].status_pelunasan == "Lunas"{
					status_lunas++
					fmt.Println(data[i].nominal)
					total -= data[i].nominal
				}
			}
			fmt.Printf("Total tagihan yang harus dibayar : Rp%d\n", total)
			fmt.Printf("Presentase tagihan yang lunas    : %.1f%%\n", float64(status_lunas)/float64(jumlah_tagihan) * 100.0)
		}else if no == 8 {
			status = false
		}else {
			fmt.Println("MASUKKAN NOMOR SESUAI KETENTUAN")
		}
		fmt.Println()
	}
}