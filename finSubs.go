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
type user [NMAX]langganan

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
	var subskripsi subcription
	var list user

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
		fmt.Println(bgDarkPastelUngu + " " + bgDarkestPastelUngu + gold + " ❹  Renewal Subkripsi    " + bgDarkPastelUngu + " " + reset)
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
			menuSubksripsi(&subskripsi, &list)
		} else if pilih == 3 {
			//bisa saja ditambahkan fitur lain didalam 'menuPengeluaran' atau tidak sama sekali menggunakan menu, jadi langsung mengeluarkan riwayat transaksi//
			menuPengeluaran()
		} else if pilih == 4 {
			pembayaran(saldo, &list)
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

func menuSubksripsi(subskripsi *subcription, list *user) {
	var pilih, pilihSort, pilihTambah int
	listSubs(subskripsi)

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
			fmt.Println("Daftar subskripsi Anda:")
			cek := false
			for i := 0; i < NMAX && list[i].namaAplikasi != ""; i++ {
				fmt.Printf("%d. %s - Rp%d, tenggat: %s\n", i+1, list[i].namaAplikasi, list[i].harga, list[i].tenggatBayar.Format("02 Jan 2006"))
				cek = true
			}
			if !cek {
				fmt.Println("Belum ada subskripsi.")
			}

		} else if pilih == 2 {

			// Menampilkan listSubs
			for pilih != 4 {
				fmt.Println("Aplikasi yang bisa ditambahkan:")
				for i := 0; i < NMAX; i++ {
					fmt.Printf("%d. %s - Rp%d\n", i+1, subskripsi[i].namaAplikasi, subskripsi[i].harga)
				}
				fmt.Println("====================================")
				fmt.Println("1. Cari Subskripsi")
				fmt.Println("2. Urutkan list berdasarkan harga")
				fmt.Println("3. Tambahkan subskripsi")
				fmt.Println("4. Kembali")
				fmt.Print("Pilih: ")
				fmt.Scan(&pilih)
				if pilih == 1 {
					// fungsi search
				} else if pilih == 2 {
					// fungsi sort
					fmt.Println("1. Termurah")
					fmt.Println("2. Termahal")
					fmt.Print("Pilih: ")
					fmt.Scan(&pilihSort)
					if pilihSort == 1 {
						sortASC(subskripsi)
					} else if pilih == 2 {
						sortDSC(subskripsi)
					} else {
						fmt.Println(red + "Invalid, mohon coba lagi." + reset)
					}
				} else if pilih == 3 {
					// fungsi tambahkan subskripsi ke user
					fmt.Print("Masukkan nomor aplikasi yang ingin ditambahkan: ")
					fmt.Scan(&pilihTambah)
					idx := pilihTambah - 1

					var sudahAda bool = false
					var slotKosong int = -1

					if idx >= 0 && idx < NMAX && subskripsi[idx].namaAplikasi != "" {
						for i := 0; i < NMAX; i++ {
							if list[i].namaAplikasi == subskripsi[idx].namaAplikasi {
								sudahAda = true
							}
							if list[i].namaAplikasi == "" && slotKosong == -1 {
								slotKosong = i
							}
						}

						if sudahAda {
							fmt.Println("Aplikasi sudah ditambahkan ke daftar user.")
						} else if slotKosong != -1 {
							subskripsi[idx].tenggatBayar = time.Now().AddDate(0, 1, 0)
							list[slotKosong] = subskripsi[idx]
							fmt.Println("Aplikasi berhasil ditambahkan ke daftar user.")
						}
					} else {
						fmt.Println("Pilihan tidak valid.")
					}

				} else if pilih == 4 {
					fmt.Println("kembali ke menu utama...")
				} else {
					fmt.Println(red + "Invalid, mohon coba lagi." + reset)
				}
			}

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println(red + "Invalid, mohon coba lagi." + reset)
		}
	}
}

func sortASC(a *subcription) {
	var i, j, idx_min int
	var temp langganan
	const n = NMAX

	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if a[idx_min].harga > a[j].harga {
				idx_min = j
			}
			j = j + 1
		}
		temp = a[idx_min]
		a[idx_min] = a[i-1]
		a[i-1] = temp
		i = i + 1
	}
}

func sortDSC(a *subcription) {
	var i, j int
	var temp langganan

	for i = 1; i < NMAX; i++ {
		j = i
		for j > 0 && a[j].harga > a[j-1].harga {
			temp = a[j]
			a[j] = a[j-1]
			a[j-1] = temp
			j--
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

func pembayaran(saldo *int, list *user) {
	var total int
	var pilih string
	var cek bool = false

	fmt.Println("Rincian tagihan: ")

	for i := 0; i < NMAX && list[i].namaAplikasi != ""; i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, list[i].namaAplikasi, list[i].harga)
		cek = true
	}
	if !cek {
		fmt.Println("Belum ada subskripsi.")
	}
	fmt.Println("=============================")
	for i := 0; i < NMAX; i++ {
		if list[i].namaAplikasi != "" {
			total += list[i].harga
		}
	}
	fmt.Println("Total tagihan anda: Rp.", total)
	fmt.Println("Total saldo anda: Rp.", *saldo)

	fmt.Print("Lanjutkan Pembayaran? (Y/N)")
	fmt.Scan(&pilih)
	if pilih == "Y" || pilih == "y" {
		if *saldo < total {
			fmt.Println("Pembayaran gagal, saldo anda tidak mencukupi")
		} else {
			*saldo = *saldo - total
			for i := 0; i < NMAX; i++ {
				if list[i].namaAplikasi != "" {
					list[i].tenggatBayar = list[i].tenggatBayar.AddDate(0, 1, 0)
				}
			}
			fmt.Println("Pembayaran berhasil. Saldo tersisa: Rp.", *saldo)
		}
	} else {
		fmt.Println("Kembali ke menu utama...")
	}
}
