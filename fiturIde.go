package main

import "fmt"

func main() {
	menuIde()
}

//Berikut merupakan tampilan awal atau beranda pada fitur

func menuIde() {
	var tampilanIde int

	fmt.Println("===============================================================")
	fmt.Println("                   BERANDA IDE DAN REFFERENSI")
	fmt.Println("===============================================================")
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("                    1. Lihat Ide Startup ")
	fmt.Println("                          2. Keluar ")
	fmt.Println("\n")
	fmt.Print("Masukan Pilihan : ")
	fmt.Scan(&tampilanIde)
	switch tampilanIde {
	case 1:
		ideBrainUp1()
	case 2:
		keluar()
	default: //navigasi jika pilihan tidak sesuai
		fmt.Println("Pilihan Tidak valid kocak,  pilih lagi noh!")
		menuIde()
	}

}

func ideBrainUp1() { //NOTE: IDE CUMAN AMPE 6 DOANG BUAT CONTOH AJA, KALO LEBIH KAGOK JG COY!! TTD ADMIN GEET
	//ini buat ide, ntr keluarnya 3
	ide1()
	ide2()
	ide3()

	var navigasiIde int
	fmt.Println("Tertarik sama Ide yang mana nih?")
	fmt.Println("1. Pilih ide ")
	fmt.Println("2. Cari ide lain")
	fmt.Println("3. Kembali")
	fmt.Print("Masukan Pilihan :")
	fmt.Scan(&navigasiIde)

	switch navigasiIde {
	case 1:
		var ide string
		fmt.Println("1. Ide 1")
		fmt.Println("2. Ide 2")
		fmt.Println("3. Ide 3")
		fmt.Print("Masukan ide yang kamu inginkan : ")
		fmt.Scan(&ide)

		switch ide {
		case "1":
			deskripsiIde1()
		case "2":
			deskripsiIde2()
		case "3":
			deskripsiIde3()
		}
	case 2:
		ideBrainUp2()
	case 3:
		menuIde()
	default: //navigasi jika pilihan tidak sesuai
		fmt.Println("Pilihan Tidak valid kocak,  pilih lagi noh!")
		ideBrainUp1()
	}
}

func ideBrainUp2() {
	ide4()
	ide5()
	ide6()

	var navigasiIde int
	fmt.Println("Tertarik sama Ide yang mana nih?")
	fmt.Println("1. Pilih ide ")
	fmt.Println("2. Cari ide lain")
	fmt.Println("3. Kembali")
	fmt.Print("Masukan Pilihan :")
	fmt.Scan(&navigasiIde)

	switch navigasiIde {
	case 1:
		var ide string

		fmt.Println("1. Ide 1")
		fmt.Println("2. Ide 2")
		fmt.Println("3. Ide 3")
		fmt.Print("Masukan ide yang kamu inginkan : ")
		fmt.Scan(&ide)

		switch ide {
		case "1":
			deskripsiIde4()
		case "2":
			deskripsiIde5()
		case "3":
			deskripsiIde6()
		}
	case 2: //ini navigasi karena ide yang diberikan hanya 6 dan akan muncul tampilan ini
		var adios string

		fmt.Println("\n")
		fmt.Println("Yah... Maaf yah idenya baru itu aja")
		fmt.Println("\n")
		fmt.Println("            (╥﹏╥)")
		fmt.Println("\n")
		fmt.Println("    Support kami terus yah!!")
		fmt.Println("")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Print("Masukan pilihan: ")
		fmt.Scan(&adios)
		switch adios {
		case "1":
			menuIde()
		}
	case 3:
		menuIde()
	default: //navigasi jika pilihan tidak sesuai
		fmt.Println("Pilihan Tidak valid kocak,  pilih lagi noh!")
		ideBrainUp2()
	}
}

// nah pada bagian ini merupakan func yang digunakan untuk keluar dari fitur, tapi harusnya bakalan keluar ke menu utama pada main program

func keluar() {
	fmt.Println("yakin nih pengen keluar ?")
	fmt.Println("1. Ya")
	fmt.Println("2. gk jadi deh :)")

	var yakin int
	fmt.Print("Masukan Pilihan :")
	fmt.Scan(&yakin)
	switch yakin {
	case 1:
		fmt.Println("Yowes kalau mau keluar mah :(")
		fmt.Println("Btw Thanks yah udh mau pake jasa kami ;)")
	case 2:
		menuIde()
	default: //navigasi jika pilihan tidak sesuai
		fmt.Println("Pilihan Tidak valid kocak,  pilih lagi noh!")
		keluar()
	}

}

func roadMap1() { // nah ini roadmap ide 1
	fmt.Println("1. RoadMap")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik1 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                          FoodFashion")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan riset dan survei pada pasar yang akan di pakai dalam")
		fmt.Println("startup. Dengan melihat target pasar atau calon pengguna, contohnya gen-Z,")
		fmt.Println("maupun milenial yang tertarik dengan fashion dan kuliner.")
		fmt.Println(" ")
		fmt.Println("|Tahap 2|: Merancang konsep dan desain produk awal, dengan menentukan apakah")
		fmt.Println("berupa online shop atau secara toko berbasis offline. serta melakukan desain")
		fmt.Println("pada produk yang akan pertama kali dilauching.")
		fmt.Println(" ")
		fmt.Println("|Tahap 3|: Produksi dan uji coba pasar, melakukan produksi dengan skala kecil")
		fmt.Println("untuk test pada pasar. Melakukan uji coba pemasaran baik secara online maupun")
		fmt.Println("offline, dengan membuat campaign di medsos maupun marketplace.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan branding dan collaboration pada startup, dengan membangun")
		fmt.Println("identitas brand yang kuat yang disesuaikan dengan target pasar. serta melakukan")
		fmt.Println("kolaborasi baik sesama brand Fashion maupun dengan brand makanan. Pada tahap ini")
		fmt.Println("juga dapat digunakan untuk melakukan perluasan pasar yakni seperti membuat website")
		fmt.Println("sendiri, maupaun menerapkan sistem pre-order maupun bundling pada produk.")
		fmt.Println(" ")
		fmt.Println("|Tahap 5| Pada tahap ini merupakan tahap akhir dengan melakukan variasi produk yang")
		fmt.Println("akan di jual. tahap ini akan dilakukan jika pada tahap sebelumnya terbukti berhasil")
		fmt.Println("dilakukan. Contohnya dengan menambahkan variasi produk berupa aksesoris fashion, dll.")
		fmt.Println("")
		fmt.Println("===============================================================")
		fmt.Println("")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik1)
		switch balik1 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp1()
		}
	case 2:
		menuIde()
	}
}

func roadMap2() { // nah ini roadmap ide 2
	fmt.Println("1. Pengen tau RoadMapnya ?")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik2 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                            RoboCV")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan survei dan riset pada pasar, dengan melihat kebutuhan")
		fmt.Println("pengguna pada aplikasi. serta diberlakukanya survei target pasar baik kepada")
		fmt.Println("mahasiswa,HR,maupun jobseeker. dan menganalasis kompetitor lain pada pasar. ")
		fmt.Println("")
		fmt.Println("|Tahap 2|: Melakukan  Minimum Viable Product dan early user, membangun MVP")
		fmt.Println("sederhana dengan contoh aplikasi dapat melakukan input data dan output CV secara")
		fmt.Println("otomatis, dengan menggunakan tools yang telah dipelajari dalam pengembangan. serta")
		fmt.Println("melakukan uji coba pada early user untuk memperoleh feedback yang dibutuhkan.")
		fmt.Println("")
		fmt.Println("|Tahap 3|: Melakukan pengembangan lanjutan dan Branding, jika tahap 2 telah berhasil")
		fmt.Println("maka diberlakukanya penambahan fitur tambahan seperti AI Assistant, dan tamplate CV.")
		fmt.Println("serta melakukan branding dan pemasaran pada medsos maupun pada website aplikasi.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan kerjasama dan collaboration, dengan melakukan kerjasama maupun")
		fmt.Println("kolaborasi dengan brand tertentu atau dengan universitas maka akan membuka peluang")
		fmt.Println("untuk pemasaran yang lebih lanjut.")
		fmt.Println("")
		fmt.Println("|Tahap 5|: Pada tahap ini sebenarnya hanyalah pengembangan lanjutan dan menjaga agar")
		fmt.Println("startup tetap berjalan dengan stabil. Serta tetap melakukan survei dan feedback terkait")
		fmt.Println("penambahan atau perubahan fitur terhadap pengguna yang dilakukan selama masa tahapan awal")
		fmt.Println(" ")
		fmt.Println("===============================================================")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik2)
		switch balik2 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp1()
		}
	case 2:
		menuIde()
	}
}

func roadMap3() { // nah ini roadmap ide 3
	fmt.Println("1. Pengen tau RoadMapnya ?")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik3 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                            FokusIn!")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan survei dan riset pada pasar, dengan melihat kebutuhan")
		fmt.Println("pengguna pada aplikasi. serta diberlakukanya survei target pasar baik kepada")
		fmt.Println("mahasiswa,pelajar,pekerja. dan menganalasis kompetitor lain pada pasar. ")
		fmt.Println("")
		fmt.Println("|Tahap 2|: Melakukan  Minimum Viable Product dan identifikasi, merancang MVP")
		fmt.Println("dengan membuat prototype awal seperti fitur blokir distraksi, maupun to-do list.")
		fmt.Println("selain itu, dapat diberlakukan juga identifikasi masalah dengan menganalisis data")
		fmt.Println("yang telah diperoleh dari survei yang dilakukan pada tahap sebelumnya.")
		fmt.Println("")
		fmt.Println("|Tahap 3|: Melakukan pengembangan lanjutan dan Branding, jika tahap 2 telah berhasil")
		fmt.Println("maka diberlakukanya penambahan fitur tambahan seperti reminder,maupun statistik harian.")
		fmt.Println("serta melakukan branding dan pemasaran pada medsos maupun pada website aplikasi.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan peluncuran secara publik, dengan melakukan publikasi aplikasi")
		fmt.Println("pada GooglePlaystore maupun Appstore, serta melakukan branding lanjutan dengan skala")
		fmt.Println("yang lebih luas lagi.")
		fmt.Println("")
		fmt.Println("|Tahap 5|: Melakukan kolaborasi dan kerjasama, dengan melakukan kolaborasi dengan aplikasi")
		fmt.Println("yang berbasis edukasi, maupun aplikasi produktifitas lain. serta pada tahap ini juga dilakukan")
		fmt.Println("penambahan atau perubahan fitur seperti musik maupun fitur cross-device.")
		fmt.Println("===============================================================")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik3)
		switch balik3 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp1()
		}
	case 2:
		menuIde()
	}
}

func roadMap4() { // nah ini roadmap ide 4
	fmt.Println("1. Pengen tau RoadMapnya ?")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik4 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                            MedTrack")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan survei dan riset pada pasar, dengan melihat kebutuhan")
		fmt.Println("pengguna pada aplikasi. serta diberlakukanya survei target pasar pada layanan")
		fmt.Println("medis, serta melakukan wawancara dan survei dengan tenaga kesehatan dan pengguna baru.")
		fmt.Println("")
		fmt.Println("|Tahap 2|: Melakukan  Minimum Viable Product dan identifikasi, merancang MVP")
		fmt.Println("dengan membuat prototype awal seperti fitur cek gejala penyakit dan rekomendasi.")
		fmt.Println("selain itu, dapat melakukan integrasi dan kerja sama dengan aplikasi tracking")
		fmt.Println("kesehatan harian yang lain seperti Google Fit atau Apple Health.")
		fmt.Println("")
		fmt.Println("|Tahap 3|: Melakukan pengembangan lanjutan dan Branding, jika tahap 2 telah berhasil")
		fmt.Println("maka diberlakukanya penambahan fitur tambahan seperti konsultasi dokter harian dan")
		fmt.Println("reminder Obat dan pemantauan harian dalam mengkonsumsi obat. selain itu dapat juga")
		fmt.Println("melakukan Branding pada sosmed maupun website layanan aplikasi yang telah di buat.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan peluncuran secara publik, dengan melakukan publikasi aplikasi")
		fmt.Println("pada platform yang mudah dijangkau seperti pada wearable device dan juga aplikasi mobile.")
		fmt.Println("selain itu hal ini juga berbarengan dengan kampanye edukasi digital terkait kesehatan.")
		fmt.Println("")
		fmt.Println("|Tahap 5|: Melakukan kolaborasi dan kerjasama, dengan melakukan kolaborasi dengan rumahsakit,")
		fmt.Println("klinik kesehatan, maupun dengan dokter-dokter dan tenaga medis yang ahli. Sehingga dapat menye-")
		fmt.Println("mpurnakan aplikasi karena langsung di awasi oleh sumber yang terpercaya. selain itu juga dapat mela-")
		fmt.Println("kolaborasi antar aplikasi maupun website layanan kesehatan yang lain sperti Halodoc dan sejenisnya.")
		fmt.Println("===============================================================")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik4)
		switch balik4 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp2()
		}
	case 2:
		menuIde()
	}
}

func roadMap5() { // nah ini roadmap ide 5
	fmt.Println("1. Pengen tau RoadMapnya ?")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik5 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                            Platera")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan riset pada pasar dan kemitraan, dengan melihat kebutuhan")
		fmt.Println("pengguna pada aplikasi. Maka diberlakukanya identifikasi pada UMKM dan petani lokal")
		fmt.Println("yang potensial. Serta melakukan survei preferensi konsumen tentang makanan lokal, agar")
		fmt.Println("dapat melakukan peningkatan dan pengembangan fitur yang akan ditambah pada platform.")
		fmt.Println("")
		fmt.Println("|Tahap 2|: Melakukan  Minimum Viable Product dan identifikasi, merancang MVP")
		fmt.Println("dengan membuat prototype awal atau platform dasar dalam bentuk website atau aplikasi.")
		fmt.Println("contohnya seperti fitur listing produk maupun pemesanan, yang pasti akan menjadi")
		fmt.Println("fitur utama pada platform, dan fitur lain seperti pencarian, dan kategori makanan.")
		fmt.Println("")
		fmt.Println("|Tahap 3|: Melakukan Uji coba pada skala kecil, dengan melakukan uji coba ini maka")
		fmt.Println("dapat mengembangkan sistem logistik yang lebih baik, contohnya dengan membuat sistem")
		fmt.Println("logistik sendiri, maupun menggunakan mitra ekspedisi. Hal inidapat dilakukan dalam skala")
		fmt.Println("kecil, seperti pada kota misalnya jakarta, maupun bandung. dengan adanya uji coba ini")
		fmt.Println("maka dapat menambah preferensi pembukaan platform pada kota lain sehingga akan lebih terarah.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan peningkatan pada fitur dan branding, dengan telah dilakukanya tahap 3")
		fmt.Println("maka dapat melakukan peningkatan fitur lain seperti langganan minggunan pada toko yang diminati.")
		fmt.Println("selain itu dapat juga melakukan branding pada sosmed maupun kampanye dengan mengadakan event")
		fmt.Println("atau kampanye kecil-kecilan antar UMKM maupun petani lokal dengan pengembang platform.")
		fmt.Println("")
		fmt.Println("|Tahap 5|: Melakukan kolaborasi dan kerjasama, dengan melakukan kolaborasi dengan aplikasi atau")
		fmt.Println("paltform makanan lain seperti GrabFood, ShopeeFood, maupun Gojek, diharapkan dapat memperbesar")
		fmt.Println("branding dan ekspansi secara regional, dengan demekian dapat membuka platform ke banyak kota.")
		fmt.Println("===============================================================")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik5)
		switch balik5 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp2()
		}
	case 2:
		menuIde()
	}
}

func roadMap6() { // nah ini roadmap ide 6
	fmt.Println("1. Pengen tau RoadMapnya ?")
	fmt.Println("2. Kyknya cukup udh cukup deh")

	var road int
	var balik6 string
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&road)

	switch road {
	case 1:
		fmt.Println("===============================================================")
		fmt.Println("                            Creativy")
		fmt.Println("===============================================================")
		fmt.Println("Road Map :")
		fmt.Println("|Tahap 1|: Melakukan riset dan survei kreator, dengan melakukan survei dan riset ini,")
		fmt.Println("maka dapat dengan mudah mengidentifikasi kebutuhan utama kreator digital baik dalam")
		fmt.Println("bidang desain, musik maupun yang lainya. dengan survei yang dilakukan ini maka dapat")
		fmt.Println("memperoleh data yang akan digunakan pada peningkatan dan penambahan fitur pada aplikasi.")
		fmt.Println("")
		fmt.Println("|Tahap 2|: Melakukan  Minimum Viable Product dan identifikasi, merancang MVP")
		fmt.Println("dengan membuat prototype awal seperti dengan menambahkan fitur upload karya, profil kreator,")
		fmt.Println("maupun listing jasa. serta menambahkan fitur lain seperti integrasi pada pembayaran")
		fmt.Println("secara digital, contohnya Gopay, OVO, ShopeePay, maupun kartu kredit dan lainya.")
		fmt.Println("")
		fmt.Println("|Tahap 3|: Melakukan peluncuran aplikasi dan membangun komunitas, dengan melakukan tahap")
		fmt.Println("ini, akan dilakukannya uji coba dengan kreator-kretor potensial, sebagai BETA tester pada")
		fmt.Println("aplikasi. selain itu juga dapat membangun komunitas antara sesama kreator dalam tahap uji")
		fmt.Println("coba atau peluncuran aplikasi secara BETA ini.")
		fmt.Println("")
		fmt.Println("|Tahap 4|: Melakukan peningkatan pada fitur dan branding, dengan telah dilakukanya tahap 3")
		fmt.Println("maka dapat melakukan peningkatan fitur lain seperti fitur showcase, tools kolaborasi, serta")
		fmt.Println("fitur voting pada karya yang disukai. selain itu juga dapat melakukan branding dengan menar-")
		fmt.Println("getkan kepada para kreator digital, dengan melakukan branding pada sosial media.")
		fmt.Println("")
		fmt.Println("|Tahap 5|: Melakukan kolaborasi dan kerjasama, dengan melakukan kolaborasi dengan aplikasi atau")
		fmt.Println("paltform karya seni yang lain seperti etsy, eBay, maupun yang lainya, dengan begitu maka akan dapat")
		fmt.Println("memperluas branding dan  juga dapat melakukan ekspansi ke beberapa paltform lain selain aplikasi.")
		fmt.Println("===============================================================")
		fmt.Println("1. Kembali ke menu awal")
		fmt.Println("2. Kembali ke menu Ide")
		fmt.Println("")
		fmt.Print("Masukan Pilihan:")
		fmt.Scan(&balik6)
		switch balik6 {
		case "1":
			menuIde()
		case "2":
			ideBrainUp2()
		}
	case 2:
		menuIde()
	}
}

func ide1() { // ini teh buat tampilan awal ide 1
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                          FoodFashion")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama foodfashion, FoodFashion merupakan startup dengan ide")
	fmt.Println("yang keren lho! dimana ide startup ini dengan menggabungkan tampilan")
	fmt.Println("makanan yang bikin ngiler sama fashion yang kece abis !!")
	fmt.Println("\n")

}
func ide2() { //ini teh buat tampilan awal ide 2
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                            RoboCV")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama RoboCV, RoboCV merupakan startup yang dapat meng-")
	fmt.Println("analisis CV kamu lho! dimana ide startup ini cara kerjanya ialah dengan")
	fmt.Println("membaca apakah CV anda sudah dalam format yang benar!!")
	fmt.Println("\n")
}

func ide3() { // in teh buat tampilan awal ide 3
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                            FokusIn!")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama FokusIn, FokusIn merupakan startup yang dapat membantu")
	fmt.Println("kamu untuk terhindar dari distraksi sosmed maupun game, dengan tetap fokus")
	fmt.Println("untuk belajar. dengan memblokir akses sosmed dan membuat kamu harus produktif")
	fmt.Println("agar dapat membukanya, dengan adanya timer,to-do list, dan juga musik. yang")
	fmt.Println("pastinya dapat membuat kamu terhindar dari distraksi dan menemanimu produktif.")
	fmt.Println("\n")
}

func ide4() { // ini buat tampilan awal ide 4
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                           MedTrack")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama MedTrack, MedTrack merupakan platform digital yang")
	fmt.Println("berbasis AI! yang dimana pengguna dapat memeriksa gejala penyakit")
	fmt.Println("secara mandiri, dengan memberikan rekomendasi awal sebelum ke dokter.")
	fmt.Println("\n")
}

func ide5() { // nah initamwal ide 5
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                            Platera")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama Platera, Platera merupakan marketplace makanan dan")
	fmt.Println("dan minuman lokal, seperti UMKM kuliner maupun bahan makanan langsung")
	fmt.Println("dari petani dan produsen lokal ke tangan konsumen atau pengguna.")
	fmt.Println("\n")
}

func ide6() { // nah ini tamwal ide 6
	fmt.Println("                       Nih ide buat kamu !")
	fmt.Println("===============================================================")
	fmt.Println("                            Creativy")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi: ")
	fmt.Println("yuk kenalan sama Creativy, Creativy adalah platform digital All-in-one")
	fmt.Println("yang berperan sebagai pendukung kreator digital dalam menjual karya,")
	fmt.Println("mengelola portofolio, yang langsung terhubung dengan klien potensial.")
	fmt.Println("\n")
}

func deskripsiIde1() { //nah ini deskripsi lanjutan buat ide 1
	fmt.Println("===============================================================")
	fmt.Println("                          FoodFashion")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("FoodFashion merupakan startup dengan ide menggabungkan tampilan ")
	fmt.Println("makanan yang bikin ngiler sama fashion yang keren.Contohnya dengan")
	fmt.Println("membuat pakaian dan produk fashion dengan tema telur ceplok misalnya.")
	fmt.Println("membuat makanan dengan tema kuliner Indonesia yakni telur ceplok balado,")
	fmt.Println("dengan menggunakan perpaduan warna putih,kuning, dan merah yang menarik!!")
	fmt.Println("selain dari ide dari kuliner indonesia kamu juga bisa kombinasi berbagai")
	fmt.Println("kuliner western, atau asia yang lain lho!, misalnya roti beguette atau yang lain.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap1() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}

func deskripsiIde2() { // nah ini deskripsi lanjutan buat ide 2
	fmt.Println("===============================================================")
	fmt.Println("                            RoboCV")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("RoboCV merupakan startup yang dapat menganalisis CV kamu, dengan")
	fmt.Println("mendeteksi apakah CV kamu sudah sesuai dengan standar CV yang baik atau")
	fmt.Println("tidak. Selain itu juga, RoboCV membantu anda seputar pertanyaan-pertanyaan")
	fmt.Println("terkait dunia kerja. Dengan berorientasi pada AI sehingga membuat RoboCV")
	fmt.Println("menjadi asisten pembantu anda dalam menunjang karir dimasa depan. Platfrom")
	fmt.Println("ini sangat direkomendasikan jika anda ingin melamar kerja baik mahasiswa fresh-")
	fmt.Println("graduate, hingga profesional yang ingin meningkatkan karier. dengan pendekatan")
	fmt.Println("yang relatif cepat, efesien, dan adaptif sehingga dapat membantu dengan optimal.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap2() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}

func deskripsiIde3() { // nah ini deskripsi lanjutan buat ide 3
	fmt.Println("===============================================================")
	fmt.Println("                            FokusIn!")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("FokusIn adalah aplikasi produktifitas, yang membantu pengguna untuk tetap")
	fmt.Println("fokus untuk produktif. Baik saat belajar, bekerja, maupun aktifitas lain")
	fmt.Println("yang membutuhkan tanpa adanya distraksi. Aplikasi ini membantu pengguna agar")
	fmt.Println("tetap fokus dengan memblokir aktifitas medsos, game maupun notifikasi yang dapat")
	fmt.Println("membuat anda terganggu. Selain memblokir, pengguna juga diberikan fitur-fitur yang")
	fmt.Println("menarik tentunya, contohnya pada aplikasi ini, disertai dengan fitur timer yang dapat")
	fmt.Println("anda atur, selain itu juga terdapat fitur musik yang tenang sehingga dapat membantu")
	fmt.Println("anda untuk tetap fokus dengan apa yang sedang dikerjakan dan terhindar dari distraksi.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap3() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}

func deskripsiIde4() {
	fmt.Println("===============================================================")
	fmt.Println("                            MedTrack")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("MedTrack adalah platform kesehatan yang berbasis digital dengan AI, dengan")
	fmt.Println("memungkinkan pengguna untuk memeriksa gejala secara mandiri serta mendapatkan")
	fmt.Println("rekomendasi awal ketika terkena gejala tertentu. dengan mengintegrasikan teknologi")
	fmt.Println("dan machine learning dan data medis secara real-time. MedTrack juga mampu memberikan")
	fmt.Println("rekomendasi kesehatan personal dengan meningkatkan kulitas hidup bagi pengguna.")
	fmt.Println("fitur lain yang diberikan contohnya, disertai dengan fitur jadwal konsumsi obat atau")
	fmt.Println("reminder, serta penambahan fitur penyimpanan riwayat kesehatan pengguna. pada aplikasi")
	fmt.Println("ini juga akan dihadirkan pada wearwble device seperti smartwatch yang dapat memantau")
	fmt.Println("aktifitas jantung, tekanan darah dan pola tidur secara langsung. dengan fokus utama")
	fmt.Println("ialah dengan memberikan akses yang cepat,efesien,dan akurat untuk meningkatkan")
	fmt.Println("kualitas hidup pengguna dan memperluas akses layanan kesehatan secara global.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap4() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}

func deskripsiIde5() {
	fmt.Println("===============================================================")
	fmt.Println("                            Platera")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("Platera adalah marketplace makanan dan minuman lokal yang menghubungkan produsen")
	fmt.Println("berskala kecil contohnya UMKM dan petani dengan pengguna secara langsung. Platform")
	fmt.Println("ini menyediakan makanan-makanan yang siap disantap kapanpun baik kuliner indonesia")
	fmt.Println("maupun makanan diluar kuliner nusantara, seperti korea, jepang, maupun amerika. selain")
	fmt.Println("kuliner yang siap saji, platform ini juga memberikan bahan makanan yang segar yang langsung")
	fmt.Println("dari produsen lokal baik sayuran, buah-buahan, maupun daging, telur dan olahan susu.")
	fmt.Println("dengan menggunakan pengiriman yang cepat dan sistem langganan toko yang dapat diperbarui.")
	fmt.Println("Platera memberikan konsumen atau pengguna makanan yang sehat,enak, dan segar dari")
	fmt.Println("sumber yang terpercaya, sekaligus dapat membantu, mendukung ekonomi lokal seperti UMKM.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap5() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}

func deskripsiIde6() {
	fmt.Println("===============================================================")
	fmt.Println("                            Creativy")
	fmt.Println("===============================================================")
	fmt.Println("Deskripsi :")
	fmt.Println(" ")
	fmt.Println("Creativy adalah platform aplikasi yang All-in-one bagi seniman maupun kreator digital")
	fmt.Println("untuk menjual karya mereka baik desain, musik, lukisan, maupun karya seni yang lain.")
	fmt.Println("Selain menjual, platform ini juga memberikan sistem pengelolaan portofolio, yang dapat")
	fmt.Println("terhubung langsung dengan klien potensial. platform ini juga dilengkapi dengan fitur-fitur")
	fmt.Println("yang menarik yakni seperti fitur showcase yang interaktif, tools kolaborasi yag dapat")
	fmt.Println("dilakukan secara real-time, serta sistem rating dan review yang dapat membantu kreator")
	fmt.Println("dan seniman baru dalam pasar kreatifitas. Creativy berfokus pada pengembangan kreator digital")
	fmt.Println("dan seniman yang baru berkecimpung di dunia industri kreatif di era digital, dengan membantu,")
	fmt.Println("mengelola serta memberikan wadah para talenta-talenta baru dalam pengembangan karir mereka.")
	fmt.Println(" ")
	fmt.Println("===============================================================")
	fmt.Println("              kalau kamu pengen tahu lebih lanjut")
	fmt.Println("               boleh pilih bagian RoadMap yah...")
	fmt.Println(" ")
	fmt.Println("                          (˶ᵔ ᵕ ᵔ˶)")
	fmt.Println("\n")
	roadMap6() // navigasi ini akan menuju pada roadmap sesuai ide yang diberikan
}
