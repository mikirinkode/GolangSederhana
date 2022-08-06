package main

import (
	"fmt"
	"math"
)

/*
	Ditulis oleh Wafa
	6 Agustus 2022
	github.com/mikirinkode
*/

func main() {

	var userChoice int

	for i:= 1; i > 0; i++{
		fmt.Println("\nMenu:")
		fmt.Println("1. Hitung Luas Persegi")
		fmt.Println("2. Hitung Luas Segitiga")
		fmt.Println("3. Hitung Luas Lingkaran")
		fmt.Println("4. Hitung Sudut Trigonometri ")
		fmt.Println("5. Hitung Gerak Lurus Berarturan")
		fmt.Println("6. Hitung Gerak Lurus Berubah Berarturan")
		fmt.Println("7. Hitung Energi Potensial dan Kinetik")
		fmt.Println("8. Hitung frekuensi dan getaran")
		fmt.Println("9. Hitung massa jenis")
		fmt.Println("10. Hitung daya, tekanan, usaha dan gaya")
		fmt.Println("11. Konversi suhu")
		fmt.Println("99. Exit Program")
		fmt.Print("pilih [1-11]: ")
		fmt.Scan(&userChoice)
		switch userChoice{
			case 1:
				LuasPersegi()
			case 2:
				LuasSegitiga()
			case 3:
				GetLuasLingkaran()
			case 4:
				HitungTrigonometry()
			case 5:
				HitungGLB()
			case 6:
				HitungGLBB()
			case 7:
				HitungPotensialKinetik()
			case 8:
				HitungFrekuensiGetaran()
			case 9:
				HitungMassaJenis()
			case 10:
				HitungDayaGaya()
			case 11:
				KonversiSuhu()
			case 99:
				i = -1	// stop looping
			default:
				fmt.Println("Invalid input!")
		}

	}
}

func LuasPersegi() {
	var sisi float64
	fmt.Println("\n#1. Menghitung Luas Persegi")
	fmt.Println("Rumus -> Luas Persegi = sisi * sisi")
	fmt.Print("Panjang sisi (cm) = ")
	fmt.Scan(&sisi)
	fmt.Printf("Luas Persegi = %.1f cm\n", (sisi * sisi))
}

func LuasSegitiga() {
	var alas, tinggi float64
	fmt.Println("\n#2. Menghitung Luas Segitiga")
	fmt.Println("Rumus -> Luas Segitiga = 1/2 * alas * tinggi")
	fmt.Print("Panjang alas (cm) = ")
	fmt.Scan(&alas)
	fmt.Print("Tinggi segitiga (cm) = ")
	fmt.Scan(&tinggi)
	fmt.Printf("Luas segitiga = %.1f cm \n", (1.0/2.0 * alas * tinggi))
}

func GetLuasLingkaran() {
	var jariJari float64
	fmt.Println("\n#3. Menghitung Luas Lingkaran")
	fmt.Println("Rumus -> Luas Lingkaran = phi * r * r")
	fmt.Print("Panjang Jari Jari (cm) = ")
	fmt.Scan(&jariJari)

	fmt.Printf("Luas Lingkaran = %.1f cm \n", (math.Pi * jariJari * jariJari))
}

func HitungTrigonometry(){
	var alas, miring, tinggi float64
	fmt.Println("\n#4. Menghitung Sudut Trigonometri")
	fmt.Print("Panjang Alas = ")
	fmt.Scan(&alas)
	fmt.Print("Panjang Sisi Miring = ")
	fmt.Scan(&miring)
	fmt.Print("Tinggi = ")
	fmt.Scan(&tinggi)
	fmt.Printf("Sin = %.2f \n", tinggi / miring)
	fmt.Printf("Cos = %.2f \n", alas / miring)
	fmt.Printf("Tan = %.2f \n", tinggi / alas)
}

func HitungGLB(){
	var userChoice int
	var kecepatan, jarakTempuh, waktuTempuh float64
	fmt.Println("\n#5. Menghitung Gerak Lurus Beraturan")
	fmt.Println("Menu:")
	fmt.Println("1. Hitung kecepatan")
	fmt.Println("2. Hitung jarak tempuh")
	fmt.Println("3. Hitung waktu tempuh")
	fmt.Print("Pilih [1-3]: ")
	fmt.Scan(&userChoice)

	switch userChoice {
		case 1:
			fmt.Println("\n#Menghitung Kecepatan")
			fmt.Println("Rumus -> Kecepatan = JarakTempuh / WaktuTempuh")
			fmt.Print("Jarak tempuh (km) = ")
			fmt.Scan(&jarakTempuh)
			fmt.Print("Waktu tempuh (jam) = ")
			fmt.Scan(&waktuTempuh)
			fmt.Printf("Kecepatan = %.2f km/jam \n", jarakTempuh/waktuTempuh )
		case 2:
			fmt.Println("\n#Menghitung Jarak Tempuh")
			fmt.Println("Rumus -> JarakTempuh = Kecepatan * WaktuTempuh")
			fmt.Print("Kecepatan (km/jam) = ")
			fmt.Scan(&kecepatan)
			fmt.Print("Waktu tempuh (jam) = ")
			fmt.Scan(&waktuTempuh)
			fmt.Printf("Jarak Tempuh = %.2f km \n", kecepatan*waktuTempuh )
		case 3:
			fmt.Println("\n#Menghitung Waktu Tempuh")
			fmt.Println("Rumus -> WaktuTempuh = jarak tempuh / kecepatan")
			fmt.Print("Jarak tempuh (km) = ")
			fmt.Scan(&jarakTempuh)
			fmt.Print("Kecepatan (km/jam) = ")
			fmt.Scan(&kecepatan)
			fmt.Printf("Waktu Tempuh = %.2f jam \n", jarakTempuh/kecepatan )
		default:
			fmt.Println("Invalid input!")
	}
}

func HitungGLBB(){
	var userChoice int
	var kecepatanAkhir, kecepatanAwal, percepatan, jarak, waktu float64
	fmt.Println("\n#6. Menghitung Gerak Lurus Berubah Beraturan")
	fmt.Println("Menu:")
	fmt.Println("1. Hitung kecepatan akhir")
	fmt.Println("2. Hitung waktu tempuh")
	fmt.Print("Pilih [1 atau 2]: ")
	fmt.Scan(&userChoice)

	switch userChoice {
		case 1:
			fmt.Println("\n#Menghitung Kecepatan Akhir")
			fmt.Println("Rumus -> Kecepatan Akhir = Kecepatan Awal + percepatan * waktu")
			fmt.Print("Kecepatan Awal (m/s)= ")
			fmt.Scan(&kecepatanAwal)
			fmt.Print("Percepatan (m/s^2)= ")
			fmt.Scan(&percepatan)
			fmt.Print("Waktu (s)= ")
			fmt.Scan(&waktu)
			kecepatanAkhir = kecepatanAwal + percepatan*waktu
			fmt.Printf("Kecepatan AKhir = %.2f m/s \n",  kecepatanAkhir)

		case 2:
			fmt.Println("\n#Menghitung Jarak")
			fmt.Println("Rumus -> Jarak  = (Kecepatan Akhir - Kecepatan Awal) / percepatan")
			fmt.Print("Kecepatan Awal (m/s)= ")
			fmt.Scan(&kecepatanAwal)
			fmt.Print("Kecepatan Akhir (m/s)= ")
			fmt.Scan(&kecepatanAkhir)
			fmt.Print("Percepatan (m/s^2)= ")
			fmt.Scan(&waktu)
			jarak = (kecepatanAkhir - kecepatanAwal) / waktu
			fmt.Printf("Waktu = %.2f s \n",  jarak)
		default: 
			fmt.Println("Invalid Input!")
	}
}

func HitungPotensialKinetik(){
	var userChoice int
	var energiPotensial, energiKinetik, massaBenda, gravitasi, ketinggian, kecepatan float64
	fmt.Println("\n#7. Menghitung Energi Potensial dan Kinetik")
	fmt.Println("Menu:")
	fmt.Println("1. Hitung Energi Potensial Gravitasi")
	fmt.Println("2. Hitung Energi Kinetik")
	fmt.Print("Pilih [1 atau 2]: ")
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		fmt.Println("\n#Menghitung Energi Potensial Gravitasi")
		fmt.Println("Rumus -> Ep  = massa benda * gravitasi * ketinggian")
		fmt.Print("Massa benda (kg)= ")
		fmt.Scan(&massaBenda)
		fmt.Print("Gravitasi (m/s2)= ")
		fmt.Scan(&gravitasi)
		fmt.Print("Ketinggian (m)= ")
		fmt.Scan(&ketinggian)
		energiPotensial = massaBenda * gravitasi * ketinggian
		fmt.Printf("Energi Potensial = %.2f Joule \n",  energiPotensial)
	case 2: 		
		fmt.Println("\n#Menghitung Energi Kinetik")
		fmt.Println("Rumus -> Ek  = (massa * kecepatan^2) / 2")
		fmt.Print("Massa (kg)= ")
		fmt.Scan(&massaBenda)
		fmt.Print("Kecepatan (m/s)= ")
		fmt.Scan(&kecepatan)
		energiKinetik = 1.0/2.0 * massaBenda * kecepatan * kecepatan
		fmt.Printf("Energi Kinetik = %.2f Joule \n",  energiKinetik)
	default:
		fmt.Println("Invalid Input!")
	}
}

func HitungFrekuensiGetaran(){
	var jumlahGetaran, waktu float64
	fmt.Println("\n#8. Menghitung Frekuensi dan Getaran")
	fmt.Println("Rumus -> Frekuensi  = jumlah getaran / waktu yang dibutuhkan")
	fmt.Print("Jumlah Getaran = ")
	fmt.Scan(&jumlahGetaran)
	fmt.Print("Waktu (s) = ")
	fmt.Scan(&waktu)
	fmt.Printf("Frekuensi = %.2f hz \n",  jumlahGetaran / waktu)
}

func HitungMassaJenis(){
	var massa, volume float64
	fmt.Println("\n#9. Menghitung Massa Jenis")
	fmt.Println("Rumus -> massa / volume")
	fmt.Print("Massa (kg) = ")
	fmt.Scan(&massa)
	fmt.Print("Volume (m3)) = ")
	fmt.Scan(&volume)
	fmt.Printf("Massa jenis = %.2f kg/m3 \n",  massa / volume)
}

func HitungDayaGaya(){
	var userChoice int
	var usaha, gaya, alas, waktu, jarak float64
	fmt.Println("\n#10. Menghitung Daya, Tekanan, Usaha dan Gaya")
	fmt.Println("Menu:")
	fmt.Println("1. Hitung Daya")
	fmt.Println("2. Hitung Tekanan")
	fmt.Println("3. Hitung Usaha")
	fmt.Println("4. Hitung Gaya")
	fmt.Print("Pilih [1-4]: ")
	fmt.Scan(&userChoice)

	switch userChoice{
		case 1:
			fmt.Println("\n#Menghitung Daya")
			fmt.Println("Rumus -> Daya = Usaha / waktu")
			fmt.Print("Usaha (joule) = ")
			fmt.Scan(&usaha)
			fmt.Print("Waktu (s) = ")
			fmt.Scan(&waktu)		
			fmt.Printf("Gaya = %.2f J/s \n",  usaha / waktu)
		case 2:
			fmt.Println("\n#Menghitung Tekanan")
			fmt.Println("Rumus -> Tekanan = Gaya / Luas alas ")
			fmt.Print("Gaya (N) = ")
			fmt.Scan(&gaya)
			fmt.Print("Luas alas (m2) = ")
			fmt.Scan(&alas)
			fmt.Printf("Tekanan = %.2f N/m2 \n",  gaya / alas)
		case 3:
			fmt.Println("\n#Menghitung Usaha")
			fmt.Println("Rumus -> Usaha = Gaya * Jarak")
			fmt.Print("Gaya (N) = ")
			fmt.Scan(&gaya)
			fmt.Print("Jarak (m) = ")
			fmt.Scan(&jarak)		
			fmt.Printf("Usaha = %.2f J\n",  gaya * jarak)
		case 4:
			fmt.Println("\n#Menghitung Gaya")
			fmt.Println("Rumus -> Gaya = Usaha / Jarak")
			fmt.Print("Usaha (J) = ")
			fmt.Scan(&usaha)
			fmt.Print("Waktu (s) = ")
			fmt.Scan(&jarak)		
			fmt.Printf("Gaya = %.2f N\n",  usaha / jarak)
		default : 
			fmt.Println("Invalid input!")
	}
}

func KonversiSuhu(){
	var userChoice int
	var suhu float64
	fmt.Println("\n#11. Konversi Susu")
	fmt.Println("Pilih Jenis Suhu yang dimiliki")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Reamur")
	fmt.Println("4. Kelvin")
	fmt.Print("Pilihan [1-4] : ")
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		fmt.Print("Masukkan suhu dalam Celcius = ")
		fmt.Scan(&suhu)
		fmt.Printf("%.2f Celcius = %.2f Reamur \n", suhu, suhu * 4.0/5.0)
		fmt.Printf("%.2f Celcius = %.2f Fahrenheit \n", suhu, (suhu / 5.0 * 9.0) + 32)
		fmt.Printf("%.2f Celcius = %.2f Kelvin \n", suhu, (suhu / 5.0 * 5.0) + 273)
	case 2:
		fmt.Print("Masukkan suhu dalam Fahrenheit = ")
		fmt.Scan(&suhu)
		fmt.Printf("%.2f Fahrenheit = %.2f Celcius \n", suhu, (suhu - 32) / 9.0 * 5.0)
		fmt.Printf("%.2f Fahrenheit = %.2f Reamur \n", suhu, (suhu - 32) / 9.0 * 4.0)
		fmt.Printf("%.2f Fahrenheit = %.2f Kelvin \n", suhu, ((suhu - 32) / 9.0 * 5.0) + 273)
	case 3:
		fmt.Print("Masukkan suhu dalam Reamur = ")
		fmt.Scan(&suhu)
		fmt.Printf("%.2f Reamur = %.2f Celcius \n", suhu, suhu / 4.0 * 5.0)
		fmt.Printf("%.2f Reamur = %.2f Fahrenheit \n", suhu, (suhu / 4.0 * 9.0) + 32)
		fmt.Printf("%.2f Reamur = %.2f Kelvin \n", suhu, (suhu / 4.0 * 5.0) + 273)
	case 4:
		fmt.Print("Masukkan suhu dalam Kelvin = ")
		fmt.Scan(&suhu)
		fmt.Printf("%.2f Kelvin = %.2f Celcius \n", suhu, (suhu - 273) / 5.0 * 5.0)
		fmt.Printf("%.2f Kelvin = %.2f Reamur \n", suhu, (suhu - 273) / 5.0 * 4.0)
		fmt.Printf("%.2f Kelvin = %.2f Fahrenheit \n", suhu, ((suhu - 273) / 5.0 * 9.0) + 32)
	default:
		fmt.Println("Invalid input!")
	}
}

/* 
	Referensi
	https://www.zenius.net/blog/rumus-luas-bangun-datar
	https://kumparan.com/kabar-harian/gerak-lurus-beraturan-rumus-grafik-dan-contoh-soalnya-1wbnZ3RGlwN
	https://www.zenius.net/blog/rumus-gerak-lurus-berubah-beraturan-glbb
	https://www.zenius.net/blog/rumus-energi-kinetik-dalam-fisika
	https://www.zenius.net/blog/rumus-energi-potensial
	https://www.zenius.net/blog/rumus-frekuensi
	https://www.zenius.net/blog/rumus-massa-jenis
	https://www.ruangguru.com/blog/daya-usaha
	https://id.wikibooks.org/wiki/Rumus-Rumus_Fisika_Lengkap/Gaya_dan_tekanan
	https://www.zenius.net/blog/mengenal-4-skala-suhu-dan-cara-konversinya
*/