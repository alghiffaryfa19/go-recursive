package main

import (
	"fmt"
	"unicode"
)

var penyebut float64
var pembilang float64

func row(i float64, j float64) float64 {

	if i <= 1 {
		fmt.Printf("%.1f", j)
		penyebut += j
		return j
	}

	k := row(i-1, j)
	fmt.Printf(" + %.1f", j/i)
	penyebut += j/i
	return k
}

func d(i float64) float64 {
	if i <= 1 {
		penyebut += 1
		pembilang += 2
		fmt.Print("1/2x")
		return i
	}
	y := d(i - 1)
	penyebut += i
	x := i * i
	pembilang += x
	fmt.Printf(" + %.0f/%.0fx", i, x)

	return y
}

func e(i float64) float64 {
	if i <= 1 {
		penyebut += 2
		pembilang += 1
		fmt.Print("2x/1")
		return i
	}
	y := e(i - 1)
	x := i * i
	penyebut += x
	pembilang += i
	fmt.Printf(" + %.0fx/%.0f", x, i)

	return y
}

func isNtCap(i int, word string) int {
	if i < 0 {
		return i
	}
	y := rune(word[i])
	if unicode.IsLower(y) {
		pembilang = pembilang + 1
	}

	return isNtCap(i-1, word)
}

func isNumb(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsNumber(y) {
		pembilang = pembilang + 1
	}
	return isNumb(i-1, x)
}

func menu() {
	var i int
	var max float64
	var num float64
	var word string

	fmt.Println("\n--------------------")
	fmt.Println("1. Deret a b c")
	fmt.Println("2. Deret d")
	fmt.Println("3. Deret e")
	fmt.Println("4. Jumlah huruf non kapital")
	fmt.Println("5. Jumlah angka pada kata")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&i)

	switch i {
	case 1:
		fmt.Print("Bilangan Pertama: ")
		fmt.Scan(&num)
		fmt.Print("Batas : ")
		fmt.Scan(&max)
		row(max, num)
		fmt.Printf(" = %.2f", penyebut)
		menu()
	case 2:
		fmt.Print("Batas : ")
		fmt.Scan(&max)

		d(max)
		fmt.Printf(" = %.0f/%.0fx", penyebut, pembilang)
		menu()
	case 3:
		fmt.Print("Batas : ")
		fmt.Scan(&max)
		e(max)
		fmt.Printf(" = %.0fx/%.0f", penyebut, pembilang)
		menu()
	case 4:
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&word)

		isNtCap(len(word)-1, word)
		fmt.Printf("Jumlah : %.0f", pembilang)
	case 5:
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&word)
		isNumb(len(word)-1, word)
		fmt.Printf("Jumlah : %.0f", pembilang)
		menu()
	case 6:
		fmt.Printf("Bye")
	default:
		fmt.Print("Menu tidak ada")
		menu()
	}
}

func main() {
	menu()
}