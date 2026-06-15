package main

import (
	"fmt"
	"time"
)

type Barang struct{ ID int; Nama string; Harga float64; Stok int }
type Transaksi struct{ ID int; Barang string; Jumlah int; Total float64; Tanggal string }

var barang []Barang
var transaksi []Transaksi
var idB, idT = 1, 1

func cariBarang(id int) int {
	for i, b := range barang { if b.ID == id { return i } }
	return -1
}

func tambah() {
	var b Barang; b.ID = idB
	fmt.Print("Nama: "); fmt.Scan(&b.Nama)
	fmt.Print("Harga: "); fmt.Scan(&b.Harga)
	fmt.Print("Stok: "); fmt.Scan(&b.Stok)
	barang = append(barang, b); idB++; fmt.Println("Barang ditambahkan.")
}

func ubah() {
	var id int; fmt.Print("ID barang: "); fmt.Scan(&id)
	i := cariBarang(id)
	if i < 0 { fmt.Println("Tidak ditemukan."); return }
	fmt.Print("Nama baru: "); fmt.Scan(&barang[i].Nama)
	fmt.Print("Harga baru: "); fmt.Scan(&barang[i].Harga)
	fmt.Print("Stok baru: "); fmt.Scan(&barang[i].Stok)
	fmt.Println("Barang diubah.")
}

func hapus() {
	var id int; fmt.Print("ID barang: "); fmt.Scan(&id)
	i := cariBarang(id)
	if i < 0 { fmt.Println("Tidak ditemukan."); return }
	barang = append(barang[:i], barang[i+1:]...); fmt.Println("Barang dihapus.")
}

func lihatBarang() {
	if len(barang) == 0 { fmt.Println("Belum ada barang."); return }
	fmt.Printf("\n%-4s %-15s %-12s %-5s\n", "ID", "Nama", "Harga", "Stok")
	fmt.Println("--------------------------------------")
	for _, b := range barang { fmt.Printf("%-4d %-15s Rp%-9.0f %-5d\n", b.ID, b.Nama, b.Harga, b.Stok) }
}

func catat() {
	var id, jml int
	fmt.Print("ID barang: "); fmt.Scan(&id)
	i := cariBarang(id)
	if i < 0 { fmt.Println("Tidak ditemukan."); return }
	fmt.Print("Jumlah: "); fmt.Scan(&jml)
	if jml > barang[i].Stok { fmt.Println("Stok tidak cukup."); return }
	total := float64(jml) * barang[i].Harga; barang[i].Stok -= jml
	transaksi = append(transaksi, Transaksi{idT, barang[i].Nama, jml, total, time.Now().Format("2006-01-02")}); idT++
	fmt.Printf("Transaksi OK. Total: Rp%.0f\n", total)
}

func lihatTransaksi() {
	hari := time.Now().Format("2006-01-02"); var omzet float64
	fmt.Printf("\n== Transaksi %s ==\n", hari)
	fmt.Printf("%-4s %-15s %-7s %-12s\n", "ID", "Barang", "Jumlah", "Total"); fmt.Println("------------------------------------------")
	for _, t := range transaksi {
		if t.Tanggal == hari { fmt.Printf("%-4d %-15s %-7d Rp%.0f\n", t.ID, t.Barang, t.Jumlah, t.Total); omzet += t.Total }
	}
	fmt.Printf("Omzet Harian: Rp%.0f\n", omzet)
}

// Selection Sort - urutkan by Harga ascending
func selectionSort() {
	n := len(barang)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ { if barang[j].Harga < barang[min].Harga { min = j } }
		barang[i], barang[min] = barang[min], barang[i]
	}
	fmt.Println("[Selection Sort] Urut by Harga:"); lihatBarang()
}

// Insertion Sort - urutkan by Nama ascending
func insertionSort() {
	for i := 1; i < len(barang); i++ {
		key := barang[i]; j := i - 1
		for j >= 0 && barang[j].Nama > key.Nama { barang[j+1] = barang[j]; j-- }
		barang[j+1] = key
	}
	fmt.Println("[Insertion Sort] Urut by Nama:"); lihatBarang()
}

// Binary Search - cari by ID (sort sementara dulu)
func binarySearch() {
	var id int; fmt.Print("Cari ID: "); fmt.Scan(&id)
	tmp := make([]Barang, len(barang)); copy(tmp, barang)
	for i := 1; i < len(tmp); i++ {
		key := tmp[i]; j := i - 1
		for j >= 0 && tmp[j].ID > key.ID { tmp[j+1] = tmp[j]; j-- }
		tmp[j+1] = key
	}
	lo, hi := 0, len(tmp)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if tmp[mid].ID == id { fmt.Printf("[Binary Search] Ketemu -> %-15s Rp%.0f Stok:%d\n", tmp[mid].Nama, tmp[mid].Harga, tmp[mid].Stok); return } else if tmp[mid].ID < id { lo = mid + 1 } else { hi = mid - 1 }
	}
	fmt.Println("Tidak ditemukan.")
}

func main() {
	var p int
	for {
		fmt.Println("\n=== KASIR MINIMART ===")
		fmt.Println("1.Tambah  2.Ubah  3.Hapus  4.Lihat Barang  5.Transaksi")
		fmt.Println("6.Omzet   7.Selection Sort  8.Insertion Sort  9.Binary Search  0.Keluar")
		fmt.Print("Pilih: "); fmt.Scan(&p)
		if p == 1 { tambah() } else if p == 2 { ubah() } else if p == 3 { hapus() } else if p == 4 { lihatBarang() } else if p == 5 { catat() } else if p == 6 { lihatTransaksi() } else if p == 7 { selectionSort() } else if p == 8 { insertionSort() } else if p == 9 { binarySearch() } else if p == 0 { fmt.Println("Keluar!"); return } else { fmt.Println("Pilihan tidak valid.") }
	}
}