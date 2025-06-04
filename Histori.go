package main

import (
	"fmt"
)

func MenuHistori() {
	var pilihan string

	for {
		fmt.Print("\n")
		fmt.Println("=========================================================")
		fmt.Println("                        HISTORY")
		fmt.Println("=========================================================")
		fmt.Println("")
		fmt.Println("                [1. Tampilkan Riwayat]")
		fmt.Println("                [2. Ubah Status      ]")
		fmt.Println("                [3. Kembali          ]")
		fmt.Print("\nInput angka untuk navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
		case "2":
			UbahStatusRiwayat(&daftarRiwayat, jumlahRiwayat)
		case "3":
			return
		default:
			fmt.Println("tidak valid")
		}
	}
}

// sistem histori
// procedure menampilkan riwayat
func TampilkanSemuaRiwayat(daftar *tabRiwayat, jumlah int) {
	var i int
	fmt.Println("\n=== Daftar Riwayat Startup ===")
	if jumlah == 0 {
		fmt.Println("Belum ada riwayat disimpan.")
	} else {
		fmt.Println("|====|=====================|===============================|====================|============|=====================|======|")
		fmt.Println("| No | Nama                | Kategori                      | Startup            | Status     | Waktu               | Vote |")
		fmt.Println("|====|=====================|===============================|====================|============|=====================|======|")
		for i = 0; i < jumlah; i++ {
			CetakRiwayat((*daftar)[i], i+1)
		}
	}
}

// procedure cetak satu riwayat
func CetakRiwayat(r Riwayat, nomor int) {
	fmt.Printf("| %-2d | %-19s | %-29s | %-18s | %-10s | %-19s | %4d |\n", nomor, r.Nama, r.Kategori, r.Startup, r.Status, r.Waktu, r.Vote)
	fmt.Println("|====|=====================|===============================|====================|============|=====================|======|")
}

// ubah status
func UbahStatusRiwayat(A *tabRiwayat, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada riwayat.")
		return
	}

	TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)

	var pilihan int
	fmt.Print("Pilih nomor riwayat yang ingin diubah statusnya: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > jumlah {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var statusBaru string
	fmt.Println("Status yang tersedia:")
	fmt.Println("1. Pending")
	fmt.Println("2. Diterima")
	fmt.Println("3. Ditolak")
	fmt.Print("Masukkan nomor status baru: ")

	var kode int
	fmt.Scan(&kode)

	switch kode {
	case 1:
		statusBaru = "Pending"
	case 2:
		statusBaru = "Diterima"
	case 3:
		statusBaru = "Ditolak"
	default:
		fmt.Println("Pilihan status tidak valid.")
		return
	}

	A[pilihan-1].Status = statusBaru
	fmt.Println("Status berhasil diubah menjadi:", statusBaru)
}
