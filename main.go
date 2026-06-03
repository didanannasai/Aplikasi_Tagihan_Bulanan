package main

import "fmt"

type Tagihan struct {
	nama_tagihan string
	nominal, jatuh_tempo int
}

func main() {
	var status bool = true
	var no, jumlah_tagihan, urutan int
	var data = make([]Tagihan, 100)
	fmt.Println("       ===== PROGRAM TAGIHAN (SIMTAB) =====\n")
	fmt.Println(">>> INFORMASI PENGGUNAAN APLIKASI <<<")
	fmt.Println("Ketik 1 untuk menambahkan data")
	fmt.Println("Ketik 2 untuk mengubah data")
	fmt.Println("Ketik 3 untuk delete data")
	fmt.Println("Ketik 4 untuk menampilkan")
	fmt.Println("Ketik 5 untuk menghentikan program\n")

	for status {
		fmt.Print("Ketik : ")
		fmt.Scanln(&no)
		if no == 1 {
			fmt.Print("Nama tagihan: ")
			fmt.Scanln(&data[jumlah_tagihan].nama_tagihan)
			fmt.Print("Nominal tagihan: ")
			fmt.Scanln(&data[jumlah_tagihan].nominal)
			fmt.Print("Tanggal jatuh tempo: ")
			fmt.Scanln(&data[jumlah_tagihan].jatuh_tempo)
			jumlah_tagihan++
		}else if no == 2 {
			fmt.Print("Masukkan urutan Tagihan yang ingin diubah: ")
			fmt.Scanln(&urutan)
			fmt.Print("Nama tagihan: ")
			fmt.Scanln(&data[urutan - 1].nama_tagihan)
			fmt.Print("Nominal tagihan: ")
			fmt.Scanln(&data[urutan - 1].nominal)
			fmt.Print("Tanggal jatuh tempo: ")
			fmt.Scanln(&data[urutan - 1].jatuh_tempo)
		}else if no == 3 {
			var temp [1000]Tagihan
			fmt.Println("Masukkan urutan Tagihan yang ingin didelete: ")
			fmt.Scanln(&urutan)
			urutan 
		}else if no == 4 {
			for i := 0; i < jumlah_tagihan; i++ {
				fmt.Printf("%d. Tagihan     : %s\n", i+1, data[i].nama_tagihan)
				fmt.Println("   Nominal     : Rp", data[i].nominal)
				fmt.Println("   Jatuh Tempo :", data[i].jatuh_tempo)
			}
		}else if no == 5 {
			status = false
		}else {
			fmt.Println("MASUKKAN NOMOR SESUAI KETENTUAN")
		}
		fmt.Println()
	}
}