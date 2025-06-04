package main

import "fmt"

func MenuSorting() {
	var pilihan string
	for {
		fmt.Println("=========================================================")
		fmt.Println("                     LEADERBOARD")
		fmt.Println("=========================================================")
		fmt.Println("")
		fmt.Println("     [1. Berdasarkan Vote Paling Populer (upvote)]")
		fmt.Println("     [2. Berdasarkan Vote paling baru (waktu)    ]")
		fmt.Println("                     [3. Kembali]")
		fmt.Print("\nInput angka untuk navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			if jumlahRiwayat == 0 {
				fmt.Println("Tidak ada data.")
			} else {
				UrutkanVoteSelectionSort(&daftarRiwayat, jumlahRiwayat)
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
			}
		case "2":
			if jumlahRiwayat == 0 {
				fmt.Println("Tidak ada data.")
			} else {
				UrutkanWaktuInsertionSort(&daftarRiwayat, jumlahRiwayat)
				TampilkanSemuaRiwayat(&daftarRiwayat, jumlahRiwayat)
			}
		case "3":
			return
		default:
			fmt.Println("tidak valid")
		}
	}
}

// sorting vote (selection) (descending)
func UrutkanVoteSelectionSort(A *tabRiwayat, N int) {
	var i, idx, pass int
	var temp Riwayat
	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i].Vote > A[idx].Vote {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

// berdasarkan waktu (insertion) (descending)
func UrutkanWaktuInsertionSort(A *tabRiwayat, N int) {
	var i, pass int
	var temp Riwayat

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.Waktu < A[i-1].Waktu {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
