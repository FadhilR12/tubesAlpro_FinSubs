package main

import (
	"fmt"
	"time"
)

const NMAX int = 15

type langganan struct {
	namaAplikasi string
	harga        int
	tenggatBayar time.Time
}

type subcription [NMAX]langganan

func main() {
	menuUtama()
}

func listSubs(subskripsi *subcription) {
	subskripsi[0].namaAplikasi = "Netflix"
	subskripsi[1].namaAplikasi = "Spotify"
	subskripsi[2].namaAplikasi = "Youtube Premium"
	subskripsi[3].namaAplikasi = "Disney+"
	subskripsi[4].namaAplikasi = "Apple Music"
	subskripsi[5].namaAplikasi = "Amazon Prime Video"
	subskripsi[6].namaAplikasi = "HBO Max"
	subskripsi[7].namaAplikasi = "Microsoft 365"
	subskripsi[8].namaAplikasi = "Google One"
	subskripsi[9].namaAplikasi = "Adobe Creative Cloud"
	subskripsi[10].namaAplikasi = "Crunchyroll"
	subskripsi[11].namaAplikasi = "LinkedIn Premium"
	subskripsi[12].namaAplikasi = "Canva Pro"
	subskripsi[13].namaAplikasi = "Duolingo Plus"
	subskripsi[14].namaAplikasi = "NordVPN"

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
}

func menuUtama() {
	var pilih, saldo int

	for pilih != 5 {
		fmt.Println("===========================")
		fmt.Println("          FinSubs")
		fmt.Println("===========================")
		fmt.Println("1) Keuangan")
		fmt.Println("2) Subkripsi")
		fmt.Println("3) Pengeluaran")
		fmt.Println("4) Pembayaran Subkripsi")
		fmt.Println("5) Keluar")
		fmt.Println()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuKeuangan(&saldo)
		} else if pilih == 2 {
			menuSubksripsi()
		} else if pilih == 3 {
			//bisa saja ditambahkan fitur lain didalam 'menuPengeluaran' atau tidak sama sekali menggunakan menu, jadi langsung mengeluarkan riwayat transaksi//
			menuPengeluaran()
		} else if pilih == 4 {
			pembayaran()
		} else if pilih == 5 {
			fmt.Println("Terimakasih telah menggunakan aplikasi kami.")
		} else {
			fmt.Println("Invalid, mohon coba lagi.")
		}
	}
}

func menuKeuangan(saldo *int) {
	var pilih int
	for pilih != 3 {
		fmt.Println("===========================")
		fmt.Println("          Keuangan")
		fmt.Println("===========================")
		fmt.Println("1) Cek Saldo")
		fmt.Println("2) Tambah Saldo")
		fmt.Println("3) Kembali")
		fmt.Println()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			//menampilkan saldo//

		} else if pilih == 2 {

			//memasukan nominal agar saldo terisi//

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println("Invalid, mohon coba lagi.")
		}
	}
}

func menuSubksripsi() {
	var pilih int
	var subskripsi subcription
	listSubs(&subskripsi)

	for pilih != 3 {
		fmt.Println("===========================")
		fmt.Println("          Subskripsi")
		fmt.Println("===========================")
		fmt.Println("1) Cek Subkripsi")
		fmt.Println("2) Tambah Subkripsi")
		fmt.Println("3) Kembali")
		fmt.Println()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			// Menampilkan daftar sub pribadi user
			if len(userSubs) == 0 {
				fmt.Println("Belum ada subkripsi yang ditambahkan.")
			} else {
				fmt.Println("Subkripsi Anda:")
				for i, sub := range userSubs {
					fmt.Printf("%d) %s | Rp%d | Mulai: %s\n", i+1, sub.namaAplikasi, sub.harga, sub.tenggatBayar.Format("2006-01-02"))
				}
			}

		} else if pilih == 2 {

			// Menampilkan listSubs
			fmt.Println("Pilih aplikasi untuk ditambahkan:")
			for i := 0; i < NMAX; i++ {
				fmt.Printf("%d) %s | Rp%d\n", i+1, subskripsi[i].namaAplikasi, subskripsi[i].harga)
			}
			var idx int
			fmt.Print("Masukkan nomor aplikasi: ")
			fmt.Scan(&idx)
			if idx >= 1 && idx <= NMAX {
				// Tambahkan ke list pribadi user
				selected := subskripsi[idx-1]
				selected.tenggatBayar = time.Now() // otomatis catat tanggal sekarang
				userSubs = append(userSubs, selected)
				fmt.Println("Subkripsi berhasil ditambahkan!")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println("Invalid, mohon coba lagi.")
		}
	}
}

func menuPengeluaran() {
	// jaga jaga jika ingin menggunakan menu//
	var pilih int
	for pilih != 3 {
		fmt.Println("===========================")
		fmt.Println("          Pengeluaran")
		fmt.Println("===========================")
		fmt.Println("1) Riwayat Transaksi")
		fmt.Println("2) COMING SOON")
		fmt.Println("3) Kembali")
		fmt.Println()
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			//menampilkan list subkripsi yang telah dibayar//

		} else if pilih == 2 {

			// jika ingin menambahkan fitur lain tambahkan saja//

		} else if pilih == 3 {
			fmt.Println("Kembali ke menu utama...")
		} else {
			fmt.Println("Invalid, mohon coba lagi.")
		}
	}
}

func pembayaran() {

}
