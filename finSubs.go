package main

import (
	"fmt"
	"time"
)

const NMAX int = 16

type langganan struct {
	namaAplikasi string
	harga        int
	tenggatBayar time.Time
}

type subcription [NMAX]langganan

var userSubs []langganan

// colors.go
var (
	reset               = "\033[0m"
	red                 = "\033[31m"
	green               = "\033[32m"
	yellow              = "\033[33m"
	blue                = "\033[34m"
	magenta             = "\033[35m"
	cyan                = "\033[36m"
	white               = "\033[37m"
	gray                = "\033[90m"
	orange              = "\033[38;5;208m"
	textLofi            = "\033[38;2;220;220;180m" // Teks putih kekuningan
	gold                = "\033[38;2;255;180;100m" // Oranye keemasan
	softRed             = "\033[38;2;255;105;180m" // Soft Pink Red
	pastelRed           = "\033[38;2;255;179;184m" // Pastel Red
	lightRed            = "\033[38;2;255;120;120m" // Light Warm Red
	bgDarkPastelUngu    = "\033[48;2;102;67;102m"
	bgDarkestPastelUngu = "\033[48;2;75;40;85m"
)

func main() {
	var saldo int = 0
	menuUtama(&saldo)
}

func listSubs(subskripsi *subcription) {
	subskripsi[0].namaAplikasi = "Netflix              "
	subskripsi[1].namaAplikasi = "Spotify              "
	subskripsi[2].namaAplikasi = "Youtube Premium      "
	subskripsi[3].namaAplikasi = "Disney+              "
	subskripsi[4].namaAplikasi = "Apple Music          "
	subskripsi[5].namaAplikasi = "Amazon Prime Video   "
	subskripsi[6].namaAplikasi = "HBO Max              "
	subskripsi[7].namaAplikasi = "Microsoft 365        "
	subskripsi[8].namaAplikasi = "Google One           "
	subskripsi[9].namaAplikasi = "Adobe Creative Cloud"
	subskripsi[10].namaAplikasi = "Crunchyroll         "
	subskripsi[11].namaAplikasi = "LinkedIn Premium    "
	subskripsi[12].namaAplikasi = "Canva Pro           "
	subskripsi[13].namaAplikasi = "Duolingo Plus       "
	subskripsi[14].namaAplikasi = "NordVPN             "
	subskripsi[15].namaAplikasi = "Roblox Premium      "

	subskripsi[0].harga = 153000
	subskripsi[1].harga = 71490
	subskripsi[2].harga = 69000
	subskripsi[3].harga = 39000
	subskripsi[4].harga = 55000
	subskripsi[5].harga = 89000
	subskripsi[6].harga = 60000
	subskripsi[7].harga = 95999
	subskripsi[8].harga = 26900
	subskripsi[9].harga = 139000
	subskripsi[10].harga = 65000
	subskripsi[11].harga = 299000
	subskripsi[12].harga = 95000
	subskripsi[13].harga = 89000
	subskripsi[14].harga = 64000
	subskripsi[15].harga = 180000
}

func menuUtama(saldo *int) {
	var pilih int

	for pilih != 5 {
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "       🍒 " + lightRed + "FinSubs" + reset + bgDarkestPastelUngu + " 🍒        " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❶  Keuangan                " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❷  Subkripsi               " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❸  Pengeluaran             " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❹  Pembayaran Subkripsi    " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❺  Keluar                  " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println()
		fmt.Print(gold + "  Pilih: ")
		fmt.Scan(&pilih)
		fmt.Print(reset)
		fmt.Println()

		if pilih == 1 {
			menuKeuangan(saldo)
		} else if pilih == 2 {
			menuSubksripsi()
		} else if pilih == 3 {
			//bisa saja ditambahkan fitur lain didalam 'menuPengeluaran' atau tidak sama sekali menggunakan menu, jadi langsung mengeluarkan riwayat transaksi//
			menuPengeluaran()
		} else if pilih == 4 {
			pembayaran()
		} else if pilih == 5 {
			fmt.Println(green + "Terimakasih telah menggunakan aplikasi kami. Have a cozy day! 🌙" + reset)
		} else {
			fmt.Println(red + "Invalid, mohon coba lagi." + reset)
		}
	}
}

func menuKeuangan(saldo *int) {
	var pilih int
	for pilih != 3 {
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "       🍒 " + lightRed + "Keuangan" + reset + bgDarkestPastelUngu + " 🍒       " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❶  Cek Saldo               " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❷  Tambah Saldo            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❸  Kembali                 " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println()
		fmt.Print(gold + "  Pilih: ")
		fmt.Scan(&pilih)
		fmt.Print(reset)

		if pilih == 1 {

			//menampilkan saldo//
			fmt.Print(gold)
			fmt.Printf("Saldo Anda saat ini: Rp%d\n", *saldo)
			fmt.Print(reset)

		} else if pilih == 2 {

			var nominal int
			//memasukan nominal agar saldo terisi//
			fmt.Print(gold)
			fmt.Print("Masukkan nominal yang ingin ditambahkan:")
			fmt.Scan(&nominal)
			fmt.Print(reset)
			if nominal > 0 {
				*saldo += nominal
				fmt.Println(green + "Saldo berhasil ditambahkan!" + reset)
				fmt.Print(gold)
				fmt.Printf("Saldo baru: Rp%d\n", *saldo)
				fmt.Print(reset)
			} else {
				fmt.Println(red + "Nominal tidak valid." + reset)
			}

		} else if pilih == 3 {
			fmt.Println(gold + "Kembali ke menu utama..." + reset)
		} else {
			fmt.Println(red + "Invalid, mohon coba lagi." + reset)
		}
	}
}

func menuSubksripsi() {
	var pilih int
	var subskripsi subcription
	listSubs(&subskripsi)

	for pilih != 3 {
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "      🍒 " + lightRed + "Subskripsi" + reset + bgDarkestPastelUngu + " 🍒      " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❶  Cek Subkripsi           " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❷  Tambah Subkripsi        " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❸  Kembali                 " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println()
		fmt.Print(gold + "  Pilih: ")
		fmt.Scan(&pilih)
		fmt.Print(reset)

		if pilih == 1 {

			// Menampilkan daftar sub pribadi user
			if len(userSubs) == 0 {
				fmt.Println("Belum ada subkripsi yang ditambahkan.")
			} else {
				fmt.Println(gold + "Subkripsi Anda:" + reset)
				for i, sub := range userSubs {
					fmt.Printf("%d) %s | Rp%d | Mulai: %s\n", i+1, sub.namaAplikasi, sub.harga, sub.tenggatBayar.Format("2006-01-02"))
				}
			}

		} else if pilih == 2 {

			// Menampilkan listSubs
			fmt.Print(gold)
			fmt.Println(bgDarkPastelUngu + "                                        " + reset)
			fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "Pilih aplikasi untuk ditambahkan:     " + bgDarkPastelUngu + " " + reset)
			for i := 0; i < NMAX; i++ {
				fmt.Print(bgDarkPastelUngu + " " + reset)
				fmt.Printf("%s%s%-2d) %-22s | Rp%-7d%s %s\n", gold, bgDarkestPastelUngu, i+1, subskripsi[i].namaAplikasi, subskripsi[i].harga, bgDarkPastelUngu, reset)
			}
			fmt.Println(bgDarkPastelUngu + "                                        " + reset)

			var idx int
			fmt.Print(gold)
			fmt.Print("Masukkan nomor aplikasi: ")
			fmt.Scan(&idx)
			fmt.Print(reset)
			if idx >= 1 && idx <= NMAX {
				// Tambahkan ke list pribadi user
				selected := subskripsi[idx-1]
				selected.tenggatBayar = time.Now() // otomatis catat tanggal sekarang
				userSubs = append(userSubs, selected)
				fmt.Print(reset)
				fmt.Println("Subkripsi berhasil ditambahkan!")
			} else {
				fmt.Println(red + "Pilihan tidak valid." + reset)
			}

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println(red + "Invalid, mohon coba lagi." + reset)
		}
	}
}

func menuPengeluaran() {
	// jaga jaga jika ingin menggunakan menu//
	var pilih int
	for pilih != 3 {
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "      🍒 " + lightRed + "Pengeluaran" + reset + bgDarkestPastelUngu + " 🍒     " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❶  Riwayat Pengeluaran     " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❷  -                       " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❸  Kembali                 " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + "                            " + bgDarkPastelUngu + " " + reset)
		fmt.Println(bgDarkPastelUngu + "                              " + reset)
		fmt.Println()
		fmt.Print(gold + "  Pilih: ")
		fmt.Scan(&pilih)
		fmt.Print(reset)

		if pilih == 1 {

			//menampilkan list subkripsi yang telah dibayar//

		} else if pilih == 2 {

			// jika ingin menambahkan fitur lain tambahkan saja//

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println(red + "Invalid, mohon coba lagi." + reset)
		}
	}
}

func pembayaran() {

}
