package functions

import (
	"fmt"
	"code/model"
)

func Statistik(data []model.Tagihan, jumlah_tagihan, total int) {
	var status_lunas int
		for i := 0; i < jumlah_tagihan; i++ {
			if data[i].Status_pelunasan == "Lunas"{
				status_lunas++
			}
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("Total tagihan yang harus dibayar : Rp%d\n", total)
		fmt.Printf("Presentase tagihan yang lunas    : %.1f%%\n", float64(status_lunas)/float64(jumlah_tagihan) * 100.0)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func Sorting(data []model.Tagihan, jumlah_tagihan int) {
	var metode int
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Ketik: ")
	fmt.Scanln(&metode)
	if metode == 1 {
		for i := 0; i < jumlah_tagihan; i++ {
    	 	min := i
    		for j := i + 1; j < jumlah_tagihan; j++ {
       			if data[j].Jatuh_tempo < data[min].Jatuh_tempo {
            		min = j
        		}
       		}
        	data[i], data[min] = data[min], data[i]
		}
    } else if metode == 2 {
	    for i := 1; i < jumlah_tagihan; i++ {
			min := data[i]
			j := i - 1
			for j >= 0 && data[j].Jatuh_tempo > min.Jatuh_tempo {
				data[j+1] = data[j]
				j--
			}
		    data[j+1] = min
		}
		} else {
			fmt.Println("MASUKKAN NOMOR SESUAI KETENTUAN")
		}

		for i := 0; i < jumlah_tagihan; i++ {

		fmt.Printf("%d. Tagihan     : %s\n", i+1, data[i].Nama_tagihan)
		fmt.Println("   Nominal     : Rp", data[i].Nominal)
		fmt.Println("   Jatuh Tempo :", data[i].Jatuh_tempo)
		
		}
}

func Menampilkan(data []model.Tagihan, kategori []string) {
	var nomor int = 1
	for i := 0; i < len(kategori); i++ {
		fmt.Println("======================================")
		fmt.Printf("                %s\n", kategori[i])
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
}

func SearchNamaTagihan(data []model.Tagihan, jumlah_tagihan int) {
	var nama, metode string
	fmt.Print("Masukkan nama tagihan : ")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan metode pencarian! (Sequential/Binary)")
	fmt.Scanln(&metode)
	if metode == "Sequential" {
		var found bool = false
		var x int = 0
		for x < jumlah_tagihan && !found {
			found = data[x].Nama_tagihan == nama
			x++
		}
		if found == true {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Printf("   Tagihan          : %s\n", data[x-1].Nama_tagihan)
			fmt.Printf("   Nominal          : Rp%d\n", data[x-1].Nominal)
			fmt.Println("   Jatuh Tempo      :", data[x-1].Jatuh_tempo)
			fmt.Println("   Status Pelunasan :", data[x-1].Status_pelunasan)
			fmt.Println("   Kategori         :", data[x-1].Kategori)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}else if found == false {
			fmt.Println("Nama tagihan BELUM terdata")
		}
		}else if metode == "Binary" {
			for i := 0; i < jumlah_tagihan; i++ {
       			min := i
       			for j := i + 1; j < jumlah_tagihan; j++ {
        			if data[j].Nama_tagihan < data[min].Nama_tagihan {
        				min = j
         			}
       			}
        		data[i], data[min] = data[min], data[i]
    		}
		kr := 0
		kn := jumlah_tagihan - 1
		found := false
		var x int
		for kr <= kn && !found {
			med := (kr + kn) / 2
			if data[med].Nama_tagihan < nama {
				kr = med + 1
			}else if data[med].Nama_tagihan > nama {
				kn = med - 1
			}else {
				found = true
				x = med
			}
		}
		if found == true {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Printf("   Tagihan          : %s\n", data[x].Nama_tagihan)
			fmt.Printf("   Nominal          : Rp%d\n", data[x].Nominal)
			fmt.Println("   Jatuh Tempo      :", data[x].Jatuh_tempo)
			fmt.Println("   Status Pelunasan :", data[x].Status_pelunasan)
			fmt.Println("   Kategori         :", data[x].Kategori)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}else if found == false {
			fmt.Println("Nama tagihan BELUM terdata")
		}
	}
}

func SearchKategori(data []model.Tagihan, kategori []string, jumlah_tagihan int) {
	var nama, metode string
	fmt.Print("Masukkan nama kategori : ")
	fmt.Scanln(&nama)
	fmt.Println("Masukkan metode pencarian! (Sequential/Binary)")
	fmt.Scanln(&metode)
	if metode == "Sequential" {
		var found bool = false
		var x int
		for x < jumlah_tagihan && !found {
			found = kategori[x] == nama
			x++
		}
		x = 0
		if found == true {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			for i := 0; i < jumlah_tagihan; i++ {
				if data[i].Kategori == nama {
					fmt.Printf("   Tagihan          : %s\n", data[i].Nama_tagihan)
					fmt.Printf("   Nominal          : Rp%d\n", data[i].Nominal)
					fmt.Println("   Jatuh Tempo      :", data[i].Jatuh_tempo)
					fmt.Println("   Status Pelunasan :", data[i].Status_pelunasan)
					fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
				}
			}
		}else if found == false {
			fmt.Println("Nama tagihan BELUM terdata")
		}
		found = false
	}else if metode == "Binary" {
		for i := 0; i < len(kategori)-1; i++ {
     		min := i
      		for j := i + 1; j < len(kategori); j++ {
       			if kategori[j] < kategori[min] {
        	    	min = j
         		}
       		}
    		kategori[i], kategori[min] = kategori[min], kategori[i]
    	}

    	kr := 0
    	kn := len(kategori) - 1
    	found := false
    	for kr <= kn && !found {
        	med := (kr + kn) / 2

        	if kategori[med] < nama {
        		kr = med + 1
    	    } else if kategori[med] > nama {
    	        kn = med - 1
    	    } else {
     	       found = true
     	   	}
   		}

    	if found == true{
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
    	    for i := 0; i < jumlah_tagihan; i++ {
        		if data[i].Kategori == nama {
        		    fmt.Printf("   Tagihan          : %s\n", data[i].Nama_tagihan)
        		    fmt.Printf("   Nominal          : Rp%d\n", data[i].Nominal)
        		    fmt.Println("   Jatuh Tempo      :", data[i].Jatuh_tempo)
        			fmt.Println("   Status Pelunasan :", data[i].Status_pelunasan)
        		    fmt.Println("   Kategori         :", data[i].Kategori)
        		    fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
        		}
        	}
    	} else {
    	    fmt.Println("Kategori BELUM terdata")
		}
	}
}

func TambahData(kategori *[]string, data []model.Tagihan, jumlah_tagihan, total *int) {
	var setuju string
	if len(*kategori) > 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("| Daftar Kategori-Kategori: |")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		for i := 0; i < len(*kategori); i++ {
			fmt.Printf("  %d. %s\n", i + 1, (*kategori)[i])
			fmt.Println("-----------------------------")
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
