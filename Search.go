package main

import "fmt"

func MenuSearch() {
	var pilihan string
	for {
		fmt.Println("=========================================================")
		fmt.Println("                         SEARCH")
		fmt.Println("=========================================================")
		fmt.Println("")
		fmt.Println("                       [1. Cari]")
		fmt.Println("                      [2. Kembali]")
		fmt.Print("\nInput angka untuk navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			var cariNama string
			var posisi int

			fmt.Print("//Format penamaan:Crispy_Bakar\n")
			fmt.Print("Nama Startup yang dicari:")
			fmt.Scan(&cariNama)
			posisi = cari(&daftarRiwayat, jumlahRiwayat, cariNama)

			if posisi == -1 {
				fmt.Println("Ide Startup tidak ditemukan")
			} else {
				fmt.Println("Ditemukan di riwayat:", posisi+1)
				fmt.Println("Data:", daftarRiwayat[posisi])
			}
		case "2":
			return
		default:
			fmt.Println("tidak valid")
		}
	}
}

// fitur cari
func cari(h *tabRiwayat, n int, namaStartup string) int {
	var i, found int
	found = -1
	i = 0

	for i < n && found == -1 {
		if h[i].Startup == namaStartup {
			found = i
		}
		i++
	}
	return found
}
