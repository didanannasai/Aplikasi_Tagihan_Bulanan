package functions

import (
	"fmt"
	"code/model"
)

func TambahData(kategori *[]string, data []model.Tagihan, jumlah_tagihan, total *int) {
	var setuju string
	if len(*kategori) > 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("| Daftar Kategori-Kategori: |")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		for i := 0; i < len(*kategori); i++ {
			fmt.Printf("  %d. %s\n", i + 1, (*kategori)[i])
			fmt.Println("+---------------------------+")
		}
		fmt.Println()
		fmt.Println("Tambah Kategori? (Y/n)")
		fmt.Scanln(&setuju)
		if setuju == "Y" {
			fmt.Print("Masukkan nama Kategori : ")
			fmt.Scanln(&data[*jumlah_tagihan].Kategori)
			*kategori = append(*kategori, data[*jumlah_tagihan].Kategori)
		}else if setuju == "n" {
			fmt.Print("Masukkan salah satu nama Kategori di atas: ")
			fmt.Scanln(&data[*jumlah_tagihan].Kategori)
		}
	}else {
		fmt.Print("Masukkan nama Kategori : ")
		fmt.Scanln(&data[*jumlah_tagihan].Kategori)
		*kategori = append(*kategori, data[*jumlah_tagihan].Kategori)
	}
	fmt.Print("Nama tagihan           : ")
	fmt.Scanln(&data[*jumlah_tagihan].Nama_tagihan)
	fmt.Print("Nominal tagihan        : ")
	fmt.Scanln(&data[*jumlah_tagihan].Nominal)
	*total += data[*jumlah_tagihan].Nominal
	fmt.Print("Tanggal jatuh tempo    : ")
	fmt.Scanln(&data[*jumlah_tagihan].Jatuh_tempo)
	data[*jumlah_tagihan].Status_pelunasan = "Belum Lunas"
	*jumlah_tagihan++
}

func UbahData(data []model.Tagihan, jumlah_tagihan int, total *int) {
	var n, cariNama string
	var found bool = false
	var i int = 0
	fmt.Print("Masukan Nama Tagihan : ")
	fmt.Scanln(&cariNama)
	for i < jumlah_tagihan && !found {
		found = data[i].Nama_tagihan == cariNama
		i++
	}

	if found == true {
		fmt.Println("Data yang ingin diubah? (Nama/Nominal/Tanggal/Status)")
		fmt.Print("Data                 : ")
		fmt.Scanln(&n)
		if n == "Nama" {
			fmt.Print("Nama tagihan         : ")
			fmt.Scanln(&data[i - 1].Nama_tagihan)
		}else if n == "Nominal" {
			*total -= data[i - 1].Nominal
			fmt.Print("Nominal tagihan      : ")
			fmt.Scanln(&data[i - 1].Nominal)
			*total += data[i - 1].Nominal
		}else if n == "Tanggal" {
			fmt.Print("Tanggal jatuh tempo  : ")
			fmt.Scanln(&data[i - 1].Jatuh_tempo)
		}else if n == "Status" {
			if data[i - 1].Status_pelunasan == "Lunas" {
				fmt.Println("Tidak bisa merubah status lagi")
			}else if data[i - 1].Status_pelunasan == "Belum Lunas" {
				fmt.Print("Status pembayaran    : ")
				fmt.Scanln(&data[i - 1].Status_pelunasan)
				if data[i - 1].Status_pelunasan == "Lunas" {
					*total -= data[i - 1].Nominal
				}
			}
		}
	}else {
		fmt.Printf("Nama tagihan %q tidak ditemukan\n", cariNama)
	}
}
