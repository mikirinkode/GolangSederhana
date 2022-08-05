package main

import (
	"fmt"
	"math"
)

func main() {

	var userChoice int

	for i:= 1; i > 0; i++{
		fmt.Println("\nMenu:")
		fmt.Println("1. Hitung Luas Persegi")
		fmt.Println("2. Hitung Luas Segitiga")
		fmt.Println("3. Hitung Luas Lingkaran")
		fmt.Print("pilih [1-3]: ")
		fmt.Scan(&userChoice)
		switch userChoice{
			case 1:
				LuasPersegi()
			case 2:
				LuasSegitiga()
			case 3:
				GetLuasLingkaran()
			case 99:
				i = -1
			default:
				fmt.Println("Invalid input!")
		}

	}
}

func LuasPersegi() {
	var sisi float64
	fmt.Print("Panjang sisi = ")
	fmt.Scan(&sisi)
	fmt.Println("Luas Persegi = ", (sisi * sisi))
}

func LuasSegitiga() {
	var alas, tinggi float64
	fmt.Print("Panjang alas = ")
	fmt.Scan(&alas)
	fmt.Print("Tinggi segitiga = ")
	fmt.Scan(&tinggi)

	fmt.Println("Luas segitiga = ", (1.0/2.0 * alas * tinggi))
}

func GetLuasLingkaran() {
	var jariJari float64
	fmt.Print("Panjang Jari Jari = ")
	fmt.Scan(&jariJari)

	fmt.Println("Luas Lingkaran = ", (math.Pi * jariJari * jariJari))
}

	// // B. test menghitung luas segitiga
	

	// // C. test menghitung luas lingkaran
	// fmt.Printf("Luas lingkaran (panjang jari jari = 14) adalah %.2f", formula.GetLuasLingkaran(14.0))
	// fmt.Println()

	// /*
	// 	D. test menghitung sudut trigonometri
	// */
	// fmt.Print("Sudut trigonometri sin(Pi/2 atau 90 derajat) adalah ")
	// displayTrigonometryResult(formula.GetSudutTrigonometri(math.Pi/2, "sin"))
	
	// fmt.Print("Sudut trigonometri cos(Pi/6 atau 30 derajat) adalah ")
	// displayTrigonometryResult(formula.GetSudutTrigonometri(math.Pi/6, "cos"))

	// fmt.Print("Sudut trigonometri tan(Pi/3 atau 60 derajat) adalah ")
	// displayTrigonometryResult(formula.GetSudutTrigonometri(math.Pi/3, "tan"))
	
	// // E. test menghitung gerak lurus beraturan
	// fmt.Print("Kecepatan yang diperlukan untuk mencapai 36km dalam 15 menit adalah ")
	// fmt.Printf("%.2f km/jam", formula.GetKecepatan(36, 15.0/60.0))

	// // F. test menghitung gerak lurus berubah beraturan

	// // G. test menghitung energi potensial kinetik

	// // H. test menghitung frekuensi / getaran

	// // I. test menghitung massa jenis

	// // J. test menghitung daya, tekanan, usaha dan gaya!

	// // K. test konversi untuk semua satuan suhu

