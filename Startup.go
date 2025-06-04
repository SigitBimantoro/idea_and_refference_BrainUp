package main

import (
	"fmt"
	"time"
)

func MenuStartup() {
	var pilihan string
	var nama, kategori, startup, status, waktuFormat string
	var dataBaru Riwayat
	var slot, noKategori, vote int

	for {
		fmt.Println("=========================================================")
		fmt.Println("                        STARTUP")
		fmt.Println("=========================================================")
		fmt.Println("             Yuk Buat Startup kamu sendiri !")
		fmt.Println("")
		fmt.Println("                   [1. Tambah ide]")
		fmt.Println("                   [2. Hapus ide ]")
		fmt.Println("                   [3. Rename ide]")
		fmt.Println("                   [4. kembali   ]")
		fmt.Print("Input angka untuk Navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			fmt.Println("")
			fmt.Println("=========================================================")
			fmt.Println("		   KATEGORI STARTUP")
			fmt.Println("=========================================================")
			fmt.Println("")
			fmt.Println("1. Teknologi & Software")
			fmt.Println("2. Aplikasi & Platform Digital")
			fmt.Println("3. Kesehatan & Medis")
			fmt.Println("4. Pendidikan")
			fmt.Println("5. Perdagangan & E-Commerce")
			fmt.Println("6. Keuangan")
			fmt.Println("7. Properti")
			fmt.Println("8. Transportasi")
			fmt.Println("9. Lingkungan & Energi")
			fmt.Println("10. Food & Beverages")
			fmt.Println("11. Sumber Daya Manusia")
			fmt.Println("12. Kreator & Industri Kreatif")
			fmt.Println("")
			fmt.Println("Pilih kategori dengan input angka:")

			fmt.Scan(&noKategori)
			kategori = pilihKategori(noKategori)
			fmt.Print("Nama Owner: ")
			fmt.Scan(&nama)
			fmt.Print("//Format penamaan:Crispy_Bakar\n")
			fmt.Print("Nama Startup (maksimal 25 karakter): ")
			fmt.Scan(&startup)

			waktuFormat = "02-01-2006 15:04:05"
			dataBaru = BuatRiwayat(nama, kategori, startup, status, waktuFormat, vote)

			if jumlahRiwayat < maxRiwayat {
				simpanRiwayat(&daftarRiwayat, dataBaru, jumlahRiwayat)
			} else {
				slot = mintaSlotPengganti(&daftarRiwayat, jumlahRiwayat, 0)
				simpanRiwayat(&daftarRiwayat, dataBaru, slot)
			}
		case "2":
			var posisi int
			if jumlahRiwayat == 0 {
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
			} else {
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
				fmt.Println("Pilih Histori Ide yang mau dihapus")
				fmt.Scan(&posisi)
				hapus(&daftarRiwayat, &jumlahRiwayat, posisi)
			}

		case "3":
			var posisi int
			var namaBaru string

			if jumlahRiwayat == 0 {
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
			} else {
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
				fmt.Println("Pilih Histori Ide yang mau direname:")
				fmt.Scan(&posisi)
				fmt.Print("//Format penamaan:Crispy_Bakar\n")
				fmt.Println("Nama Ide Baru:")
				fmt.Scan(&namaBaru)
				rename(&daftarRiwayat, posisi, namaBaru)
			}

		case "4":
			return
		default:
			fmt.Println("tidak valid")
		}
	}
}

// initeh fungsi ambil waktu realtime
func AmbilWaktuSekarang() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

// initeh fungsi buat riwayat
func BuatRiwayat(nama, kategori, startup, status, formatWaktu string, vote int) Riwayat {
	var riwayat Riwayat
	riwayat.Nama = nama
	riwayat.Kategori = kategori
	riwayat.Startup = startup
	riwayat.Status = "Pending"
	riwayat.Waktu = AmbilWaktuSekarang()
	riwayat.Vote = 0
	return riwayat
}

// simpan riwayat
func simpanRiwayat(t *tabRiwayat, riwayat Riwayat, slot int) {
	if slot >= 0 && slot < maxRiwayat {
		(*t)[slot] = riwayat
		if slot == jumlahRiwayat {
			jumlahRiwayat++
		}
		fmt.Println("Riwayat berhasil disimpan.")
	} else {
		fmt.Println("Slot tidak valid.")
	}
}

// procedure perlihatkan slot penuh
func slotpenuh(r Riwayat, nomor int) {
	fmt.Printf("[%d] %s | %s | %s | %s | %s\n", nomor, r.Nama, r.Kategori, r.Startup, r.Status, r.Waktu)
}

// fungsi meminta slot ke user (waktu penuh))
func mintaSlotPengganti(daftar *tabRiwayat, jumlah int, defaultslot int) int {
	fmt.Println("\nSlot Penuh. Pilih slot yang ingin diganti:")
	var i, slot int
	for i = 0; i < jumlah; i++ {
		slotpenuh((*daftar)[i], i+1)
	}
	fmt.Printf("pilih slot(1-%d): ", jumlah)
	fmt.Scan(&slot)
	if slot >= 1 && slot <= jumlah {
		return slot - 1
	}
	fmt.Println("Slot tidak valid, gunakan slot default.")
	return defaultslot
}

// pilihkategori
func pilihKategori(no int) string {
	switch no {
	case 1:
		return "Teknologi & Software"
	case 2:
		return "Aplikasi & Platform Digital"
	case 3:
		return "Kesehatan & Medis"
	case 4:
		return "Pendidikan"
	case 5:
		return "Perdagangan & E-Commerce"
	case 6:
		return "Keuangan"
	case 7:
		return "Properti"
	case 8:
		return "Transportasi"
	case 9:
		return "Lingkungan & Energi"
	case 10:
		return "Food & Beverages"
	case 11:
		return "Sumber Daya Manusia"
	case 12:
		return "Kreator & Industri Kreatif"
	default:
		return "Lainnya"
	}
}

// procedure hapus
func hapus(h *tabRiwayat, jumlah *int, posisi int) {
	if posisi <= 0 || posisi > *jumlah {
		fmt.Println("Posisi tidak valid")
		return
	}
	for i := posisi - 1; i < *jumlah-1; i++ {
		(*h)[i] = (*h)[i+1]
	}
	*jumlah--
	fmt.Println("Riwayat berhasil dihapus")
}

// procdure rename
func rename(h *tabRiwayat, posisi int, namaBaru string) {
	if posisi <= 0 || posisi > jumlahRiwayat {
		fmt.Println("Posisi tidak valid.")
		return
	}
	(*h)[posisi-1].Startup = namaBaru
	fmt.Println("Nama startup berhasil diubah.")
}
