package main

import (
	"fmt"
)

var count1 int = 20
var count2 int = 7
var count3 int = 34
var count4 int = 15
var count5 int = 6
var count6 int = 10

func Menuvote() { //tampilan awal menu vote
	var nVote int
	fmt.Println("===============================================================")
	fmt.Println("                           MENU VOTE")
	fmt.Println("===============================================================")
	fmt.Println("")
	fmt.Println("              Yuk vote ide kesukaan kamu di sini!")
	fmt.Println("")
	fmt.Println("                    [1. Vote Ide Pengembang]")
	fmt.Println("                    [2. Vote Ide Pengguna  ]")
	fmt.Println("                    [3. Hasil Vote         ]")
	fmt.Println("                    [4. Kembali            ]")
	fmt.Print("Masukan pilihan:")
	fmt.Scan(&nVote)
	switch nVote {
	case 1:
		Tampilanvote()
	case 2:
		if jumlahRiwayat == 0 {
			fmt.Println("Slot masih kosong")
		} else {
			TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
			ngulang()
		}
	case 3:
		fmt.Println("Hasil voting Ide pengembang")
		fmt.Print("\n")
		fmt.Printf("Hasil voting FoodFashion : %d vote \n", count1)
		fmt.Printf("Hasil voting RoboCV      : %d vote \n", count2)
		fmt.Printf("Hasil voting FokusIn!    : %d vote \n", count3)
		fmt.Printf("Hasil voting MedTrack    : %d vote \n", count4)
		fmt.Printf("Hasil voting Platera     : %d vote \n", count5)
		fmt.Printf("Hasil voting Creativy    : %d vote \n", count6)
		fmt.Println("\nTekan ENTER untuk kembali...")
		fmt.Scanln()
		fmt.Scanln()
	case 4:
		MenuBeranda()
	}
}

func Tampilanvote() { //Tampilan saat vote
	var voting int
	var balik string

	fmt.Println("")
	ide1()
	ide2()
	ide3()
	ide4()
	ide5()
	ide6()
	fmt.Print("\n")
	fmt.Println("Pilih ide yang akan di vote sesuai urutan angka ide")
	fmt.Print("Pilih ide :")
	fmt.Scan(&voting)
	switch voting {
	case 1:
		count1++
		fmt.Println("Anda telah vote ide FoodFashion")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	case 2:
		count2++
		fmt.Println("Anda telah vote ide RoboCV")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	case 3:
		count3++
		fmt.Println("Anda telah vote ide FokusIn!")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	case 4:
		count4++
		fmt.Println("Anda telah vote ide MedTrack")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	case 5:
		count5++
		fmt.Println("Anda telah vote ide Platera")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	case 6:
		count6++
		fmt.Println("Anda telah vote ide Creativy")
		fmt.Println("1. Kembali ke menu vote")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&balik)
		switch balik {
		case "1":
			Tampilanvote()
		case "2":
			Menuvote()
		}
	}
}

func beriVote(V *tabRiwayat, index int) {
	var pilih string
	V[index].Vote++
	fmt.Println("Anda telah melakukan vote")
	fmt.Print("\n")
	fmt.Println("1. Kembali ke menu vote")
	fmt.Println("2. Kembali ke menu awal")
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&pilih)
	switch pilih {
	case "1":
		TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
		ngulang()
	case "2":
		MenuBeranda()
	}
}

func ngulang() {
	var nomor int
	fmt.Println("Pilih ide yang akan di vote sesuai urutan angka ide")
	fmt.Print("Masukan pilihan: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > 10 {
		fmt.Println("Slot tidak ada")
	} else {
		beriVote(&daftarRiwayat, nomor-1)
		TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
	}
}
