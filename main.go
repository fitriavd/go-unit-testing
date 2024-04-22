package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TaxRate   = 0.1  // Persentase pajak sebesar 10%
	ChargeFee = 5000 // Biaya tambahan (charge) sebesar 5000
)

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	// cek hargaTotal
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	// cek metode pembayaran
	validMetode := map[string]bool{
		"cod":      true,
		"transfer": true,
		"debit":    true,
		"credit":   true,
		"gerai":    true,
	}
	if !validMetode[metodePembayaran] {
		return errors.New("metode tidak dikenali")
	}

	// cek apakah dicicil atau tidak
	if dicicil {
		if metodePembayaran != "credit" {
			return errors.New("cicilan harus menggunakan metode pembayaran credit")
		}
		if hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		if metodePembayaran == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	// Hitung total harga
	totalHarga := HitungTotalHarga(hargaTotal)

	fmt.Printf("Total harga: %.2f\n", totalHarga)

	return nil
}

func HitungTotalHarga(hargaTotal float64) float64 {
	// Hitung pajak
	tax := hargaTotal * TaxRate

	// Hitung total harga dengan menambahkan pajak dan biaya tambahan (charge)
	return hargaTotal + tax + ChargeFee
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Input harga total
	fmt.Print("Masukkan harga total: ")
	hargaTotalStr, _ := reader.ReadString('\n')
	hargaTotalStr = strings.TrimSpace(hargaTotalStr)
	hargaTotal, _ := strconv.ParseFloat(hargaTotalStr, 64)

	// Input metode pembayaran
	fmt.Print("Masukkan metode pembayaran: ")
	metodePembayaran, _ := reader.ReadString('\n')
	metodePembayaran = strings.TrimSpace(metodePembayaran)

	// Input apakah pembayaran dicicil atau tidak
	fmt.Print("Apakah pembayaran dicicil (y/n): ")
	cicilStr, _ := reader.ReadString('\n')
	cicilStr = strings.TrimSpace(cicilStr)
	dicicil := false
	if cicilStr == "y" {
		dicicil = true
	}

	// Panggil fungsi PembayaranBarang dengan input dari pengguna
	err := PembayaranBarang(hargaTotal, metodePembayaran, dicicil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil")
	}
}
