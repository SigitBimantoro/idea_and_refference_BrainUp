package main

import "fmt"

func main() {
	menuIde()
}

func menuIde() {
	var tampilanIde int

	fmt.Println("=============================================================")
	fmt.Println("                 BERANDA IDE DAN REFFERENSI")
	fmt.Println("=============================================================")
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("                  1. Lihat Ide Startup ")
	fmt.Println("                        2. Keluar ")
	fmt.Println("\n")
	fmt.Println("Masukan Pilihan : ")
	fmt.Scan(&tampilanIde)
	switch tampilanIde {
	case 1:
		ideBrainUp()
	case 2:
		keluar()
	}

}

func ideBrainUp() {
	fmt.Println("                     Nih ide buat kamu !")
	fmt.Println("=============================================================")
	fmt.Println("                        FoodFashion")
	fmt.Println("=============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama foodfashion, FoodFashion merupakan startup dengan ide")
	fmt.Println("yang keren lho! dimana ide startup ini dengan menggabungkan tampilan")
	fmt.Println("makanan yang bikin ngiler sama fashion yang kece abis !!")
	fmt.Println("\n")
}

func keluar() {
	fmt.Println("yakin nih pengen keluar ?")
	fmt.Println("1. Ya")
	fmt.Println("2. gk jadi deh :)")

	var yakin int
	fmt.Scan(&yakin)
	switch yakin {
	case 1:
		fmt.Println("Yowes kalau mau keluar mah :(")
		fmt.Println("Btw Thanks yah udh mau pake jasa kami ;)")
	case 2:
		ideBrainUp()
	}

}
