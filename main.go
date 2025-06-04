package main

import (
	"fmt"
	"os"
)

type Riwayat struct { // ieu teh struck dan var global ye
	Nama     string
	Kategori string
	Startup  string
	Status   string
	Waktu    string
	Vote     int
}

const maxRiwayat int = 10

type tabRiwayat [maxRiwayat]Riwayat

var daftarRiwayat tabRiwayat
var jumlahRiwayat int = 0

func main() {
	MenuUtama()
}

func MenuUtama() {
	var pilihan string
	for {
		fmt.Println("")
		fmt.Println("               ______________  ________________ ")
		fmt.Println("             /```            \\/     __|        \\ ")
		fmt.Println("  	    /___  |__''	     /\\      '''        \\ ")
		fmt.Println(" 	   /    |_|         /  \\     __'   ___/''\\ ")
		fmt.Println("	  /  _             /    \\      |__|       \\ ")
		fmt.Println("	  | / \\__``       /      \\       |         |")
		fmt.Println("	  |/	         /        \\            /`` /")
		fmt.Println("	  \\	        /          \\          |___/")
		fmt.Println("	   \\  __/``    /___      ___\\   ``\\___/")
		fmt.Println(" 	    \\/__________  |      | _________/")
		fmt.Println("            	       \\__/     /___/         ")
		fmt.Println("              	    	 /     /            ")
		fmt.Println("	     	  	/     /              ")
		fmt.Println("")
		fmt.Println("		BrainUP Find your StarUp!")
		fmt.Println("=========================================================")
		fmt.Println("")
		fmt.Println("			[1. START]")
		fmt.Println("			[2. EXIT ]")
		fmt.Print("\nInput angka untuk navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			MenuBeranda()
		case "2":
			fmt.Println("")
			fmt.Println("               ______________  ________________ ")
			fmt.Println("             /```            \\/     __|        \\ ")
			fmt.Println("  	    /___  |__''	     /\\      '''        \\ ")
			fmt.Println(" 	   /    |_|         /  \\     __'   ___/''\\ ")
			fmt.Println("	  /  _             /    \\      |__|       \\ ")
			fmt.Println("	  | / \\__``       /      \\       |         |")
			fmt.Println("	  |/	         /        \\            /`` /")
			fmt.Println("	  \\	        /          \\          |___/")
			fmt.Println("	   \\  __/``    /___      ___\\   ``\\___/")
			fmt.Println(" 	    \\/__________  |      | _________/")
			fmt.Println("            	       \\__/     /___/         ")
			fmt.Println("              	    	 /     /            ")
			fmt.Println("	     	  	/     /              ")
			fmt.Println("")
			fmt.Println("		BrainUP Find your StarUp!")
			fmt.Println("=========================================================")
			fmt.Println("")
			fmt.Println("      [ Terima kasih sudah menggunakan BrainUP! ^_^ ] ")
			fmt.Println("")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// beranda
func MenuBeranda() {
	var pilihan string
	for {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("               ______________  ________________ ")
		fmt.Println("             /```            \\/     __|        \\ ")
		fmt.Println("  	    /___  |__''	     /\\      '''        \\ ")
		fmt.Println(" 	   /    |_|         /  \\     __'   ___/''\\ ")
		fmt.Println("	  /  _             /    \\      |__|       \\ ")
		fmt.Println("	  | / \\__``       /      \\       |         |")
		fmt.Println("	  |/	         /        \\            /`` /")
		fmt.Println("	  \\	        /          \\          |___/")
		fmt.Println("	   \\  __/``    /___      ___\\   ``\\___/")
		fmt.Println(" 	    \\/__________  |      | _________/")
		fmt.Println("            	       \\__/     /___/         ")
		fmt.Println("              	    	 /     /            ")
		fmt.Println("	     	  	/     /              ")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("		BrainUP Find your StarUp!")
		fmt.Println("=========================================================")
		fmt.Println("                [1. STARTUP BARU      ]")
		fmt.Println("                [2. TEMUKAN INSPIRASI ]")
		fmt.Println("                [3. INFO              ]")
		fmt.Println("                [4. HISTORI           ]")
		fmt.Println("                [5. SEARCH            ]")
		fmt.Println("                [6. VOTE              ]")
		fmt.Println("                [7. LEADERBOARD       ]")
		fmt.Println("                [8. KEMBALI           ]")
		fmt.Print("\nInput angka untuk navigasi: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			MenuStartup()
		case "2":
			menuIde()
		case "3":
			MenuInfo()
		case "4":
			MenuHistori()
		case "5":
			MenuSearch()
		case "6":
			Menuvote()
		case "7":
			MenuSorting()
		case "8":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// info
func MenuInfo() {
	fmt.Println("")
	fmt.Println("=========================================================")
	fmt.Println("                       ABOUT US")
	fmt.Println("=========================================================")
	fmt.Println("                  Tentang BrainUp:")
	fmt.Println("BrainUP merupakan Aplikasi berbasis CLI (Command Line Interface)")
	fmt.Println("yang membantu pengguna untuk melakukan brainstorming terhadap banyak")
	fmt.Println("                ide pengembangan Startup")
	fmt.Println("")
	fmt.Println("                   Tim Pengembang:")
	fmt.Println("Sigit Bimantoro	              Mochamad Raffi Sidiq Awaludin")
	fmt.Println("sigitbimantoro71@gmail.com    mochamadraffi885@gmail.com")
	fmt.Println("")
	fmt.Println("   Mahasiswa S1 Informatika Telkom University Indonesia")
	fmt.Println("\nTekan ENTER untuk kembali...")
	fmt.Scanln()
	fmt.Scanln()
}
