package database

import (
	"log"

	"balik-ngoding-backend/internal/models"

	"gorm.io/gorm"
)

// GetSeedProblems returns all seed problems for use in tests and seeding.
func GetSeedProblems() []models.Problem {
	jsProblems := []models.Problem{
		{
			Title:      "FizzBuzz",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan sebuah bilangan bulat N, cetak angka dari 1 sampai N dengan aturan berikut:
- Jika angka habis dibagi 3, cetak "Fizz"
- Jika angka habis dibagi 5, cetak "Buzz"
- Jika angka habis dibagi keduanya, cetak "FizzBuzz"
- Selain itu, cetak angkanya

Kembalikan hasil sebagai array string.

Contoh: fizzBuzz(5) → ["1","2","Fizz","4","Buzz"]`,
			StarterCode: `function fizzBuzz(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "5", ExpectedOutput: `["1","2","Fizz","4","Buzz"]`, IsHidden: false},
				{Input: "15", ExpectedOutput: `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]`, IsHidden: false},
				{Input: "1", ExpectedOutput: `["1"]`, IsHidden: false},
				{Input: "3", ExpectedOutput: `["1","2","Fizz"]`, IsHidden: true},
				{Input: "10", ExpectedOutput: `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz"]`, IsHidden: true},
			},
		},
		{
			Title:      "Balik String",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, kembalikan string tersebut dalam urutan terbalik.

Contoh: balikString("halo") → "olah"`,
			StarterCode: `function balikString(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo"`, ExpectedOutput: `"olah"`, IsHidden: false},
				{Input: `"dunia"`, ExpectedOutput: `"ainud"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"a"`, ExpectedOutput: `"a"`, IsHidden: true},
				{Input: `"abcde"`, ExpectedOutput: `"edcba"`, IsHidden: true},
			},
		},
		{
			Title:      "Jumlah Array",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan jumlah semua elemennya.

Contoh: jumlahArray([1, 2, 3, 4, 5]) → 15`,
			StarterCode: `function jumlahArray(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "[1,2,3,4,5]", ExpectedOutput: "15", IsHidden: false},
				{Input: "[10,20,30]", ExpectedOutput: "60", IsHidden: false},
				{Input: "[]", ExpectedOutput: "0", IsHidden: false},
				{Input: "[-1,-2,3]", ExpectedOutput: "0", IsHidden: true},
				{Input: "[100]", ExpectedOutput: "100", IsHidden: true},
			},
		},
		{
			Title:      "Cari Duplikat",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan array yang berisi semua elemen yang muncul lebih dari satu kali. Hasil harus diurutkan secara ascending.

Contoh: cariDuplikat([1, 2, 3, 2, 4, 3]) → [2, 3]`,
			StarterCode: `function cariDuplikat(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "[1,2,3,2,4,3]", ExpectedOutput: "[2,3]", IsHidden: false},
				{Input: "[1,1,1,1]", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "[1,2,3]", ExpectedOutput: "[]", IsHidden: false},
				{Input: "[4,3,2,1,2,3,4]", ExpectedOutput: "[2,3,4]", IsHidden: true},
				{Input: "[5,5,5,1,2]", ExpectedOutput: "[5]", IsHidden: true},
			},
		},
		{
			Title:      "Bilangan Prima",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan sebuah bilangan bulat positif N, tentukan apakah N adalah bilangan prima. Kembalikan true jika prima, false jika bukan.

Bilangan prima adalah bilangan yang hanya habis dibagi oleh 1 dan dirinya sendiri.

Contoh: bilPrima(7) → true, bilPrima(4) → false`,
			StarterCode: `function bilPrima(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "7", ExpectedOutput: "true", IsHidden: false},
				{Input: "4", ExpectedOutput: "false", IsHidden: false},
				{Input: "2", ExpectedOutput: "true", IsHidden: false},
				{Input: "1", ExpectedOutput: "false", IsHidden: true},
				{Input: "97", ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Palindrom",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string, tentukan apakah string tersebut merupakan palindrom. Palindrom adalah string yang sama jika dibaca dari depan maupun belakang. Abaikan huruf besar/kecil.

Kembalikan true jika palindrom, false jika bukan.

Contoh: palindrom("katak") → true, palindrom("halo") → false`,
			StarterCode: `function palindrom(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"katak"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"halo"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"Katak"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `""`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"abcba"`, ExpectedOutput: "true", IsHidden: true},
			},
		},
		// ── Loop Easy (10 soal) ──────────────────────────────────────────────
		{
			Title:      "Hitung Mundur",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat N, kembalikan array berisi angka dari N turun ke 1.

Format input: bilangan bulat N (N >= 0)
Format output: array bilangan bulat

Contoh: hitungMundur(5) → [5,4,3,2,1]
Contoh: hitungMundur(3) → [3,2,1]
Contoh: hitungMundur(0) → []`,
			StarterCode: `function hitungMundur(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "5", ExpectedOutput: "[5,4,3,2,1]", IsHidden: false},
				{Input: "3", ExpectedOutput: "[3,2,1]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "0", ExpectedOutput: "[]", IsHidden: true},
				{Input: "10", ExpectedOutput: "[10,9,8,7,6,5,4,3,2,1]", IsHidden: true},
			},
		},
		{
			Title:      "Deret Aritmatika",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan nilai awal (start), selisih (diff), dan jumlah suku (n), kembalikan array berisi n suku deret aritmatika.

Format input: "start diff n" (tiga bilangan bulat dipisah spasi)
Format output: array bilangan bulat

Contoh: deretAritmatika(1, 3, 5) → [1,4,7,10,13]
Contoh: deretAritmatika(0, 2, 4) → [0,2,4,6]`,
			StarterCode: `function deretAritmatika(start, diff, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "1 3 5", ExpectedOutput: "[1,4,7,10,13]", IsHidden: false},
				{Input: "0 2 4", ExpectedOutput: "[0,2,4,6]", IsHidden: false},
				{Input: "5 0 3", ExpectedOutput: "[5,5,5]", IsHidden: false},
				{Input: "10 -2 5", ExpectedOutput: "[10,8,6,4,2]", IsHidden: true},
				{Input: "0 1 1", ExpectedOutput: "[0]", IsHidden: true},
			},
		},
		{
			Title:      "Pangkat Dua",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat N, kembalikan array berisi pangkat dua dari 2^0 sampai 2^N.

Format input: bilangan bulat N (N >= 0)
Format output: array bilangan bulat

Contoh: pangkatDua(4) → [1,2,4,8,16]
Contoh: pangkatDua(0) → [1]`,
			StarterCode: `function pangkatDua(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "4", ExpectedOutput: "[1,2,4,8,16]", IsHidden: false},
				{Input: "0", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "3", ExpectedOutput: "[1,2,4,8]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[1,2]", IsHidden: true},
				{Input: "7", ExpectedOutput: "[1,2,4,8,16,32,64,128]", IsHidden: true},
			},
		},
		{
			Title:      "Kelipatan",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan N dan K, kembalikan array berisi semua kelipatan K yang tidak melebihi N.

Format input: "N K" (dua bilangan bulat positif dipisah spasi)
Format output: array bilangan bulat

Contoh: kelipatan(20, 3) → [3,6,9,12,15,18]
Contoh: kelipatan(10, 5) → [5,10]`,
			StarterCode: `function kelipatan(n, k) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "20 3", ExpectedOutput: "[3,6,9,12,15,18]", IsHidden: false},
				{Input: "10 5", ExpectedOutput: "[5,10]", IsHidden: false},
				{Input: "1 1", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "0 3", ExpectedOutput: "[]", IsHidden: true},
				{Input: "15 4", ExpectedOutput: "[4,8,12]", IsHidden: true},
			},
		},
		{
			Title:      "Jumlah Digit",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat non-negatif N, kembalikan jumlah semua digitnya.

Format input: bilangan bulat non-negatif N
Format output: bilangan bulat

Contoh: jumlahDigit(1234) → 10
Contoh: jumlahDigit(0) → 0`,
			StarterCode: `function jumlahDigit(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "1234", ExpectedOutput: "10", IsHidden: false},
				{Input: "0", ExpectedOutput: "0", IsHidden: false},
				{Input: "99", ExpectedOutput: "18", IsHidden: false},
				{Input: "100", ExpectedOutput: "1", IsHidden: true},
				{Input: "9999", ExpectedOutput: "36", IsHidden: true},
			},
		},
		{
			Title:      "Faktorial",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat non-negatif N, kembalikan nilai N! (faktorial N).

Faktorial N didefinisikan sebagai: N! = N × (N-1) × ... × 2 × 1, dan 0! = 1.

Format input: bilangan bulat non-negatif N (0 <= N <= 12)
Format output: bilangan bulat

Contoh: faktorial(5) → 120
Contoh: faktorial(0) → 1`,
			StarterCode: `function faktorial(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "5", ExpectedOutput: "120", IsHidden: false},
				{Input: "0", ExpectedOutput: "1", IsHidden: false},
				{Input: "3", ExpectedOutput: "6", IsHidden: false},
				{Input: "1", ExpectedOutput: "1", IsHidden: true},
				{Input: "10", ExpectedOutput: "3628800", IsHidden: true},
			},
		},
		{
			Title:      "Angka Genap",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat positif N, kembalikan array berisi semua bilangan genap dari 1 sampai N (inklusif).

Format input: bilangan bulat positif N
Format output: array bilangan bulat

Contoh: angkaGenap(10) → [2,4,6,8,10]
Contoh: angkaGenap(5) → [2,4]`,
			StarterCode: `function angkaGenap(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "10", ExpectedOutput: "[2,4,6,8,10]", IsHidden: false},
				{Input: "5", ExpectedOutput: "[2,4]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[2]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[]", IsHidden: true},
				{Input: "20", ExpectedOutput: "[2,4,6,8,10,12,14,16,18,20]", IsHidden: true},
			},
		},
		{
			Title:      "Angka Ganjil",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat positif N, kembalikan array berisi semua bilangan ganjil dari 1 sampai N (inklusif).

Format input: bilangan bulat positif N
Format output: array bilangan bulat

Contoh: angkaGanjil(9) → [1,3,5,7,9]
Contoh: angkaGanjil(6) → [1,3,5]`,
			StarterCode: `function angkaGanjil(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "9", ExpectedOutput: "[1,3,5,7,9]", IsHidden: false},
				{Input: "6", ExpectedOutput: "[1,3,5]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[1]", IsHidden: true},
				{Input: "15", ExpectedOutput: "[1,3,5,7,9,11,13,15]", IsHidden: true},
			},
		},
		{
			Title:      "Perkalian",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan bilangan bulat positif N, kembalikan array berisi tabel perkalian N dari N×1 sampai N×10.

Format input: bilangan bulat positif N
Format output: array bilangan bulat (10 elemen)

Contoh: perkalian(3) → [3,6,9,12,15,18,21,24,27,30]
Contoh: perkalian(5) → [5,10,15,20,25,30,35,40,45,50]`,
			StarterCode: `function perkalian(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "3", ExpectedOutput: "[3,6,9,12,15,18,21,24,27,30]", IsHidden: false},
				{Input: "5", ExpectedOutput: "[5,10,15,20,25,30,35,40,45,50]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[1,2,3,4,5,6,7,8,9,10]", IsHidden: false},
				{Input: "10", ExpectedOutput: "[10,20,30,40,50,60,70,80,90,100]", IsHidden: true},
				{Input: "7", ExpectedOutput: "[7,14,21,28,35,42,49,56,63,70]", IsHidden: true},
			},
		},
		{
			Title:      "Hitung Huruf",
			Category:   "loop",
			Difficulty: "easy",
			Description: `Diberikan sebuah string s, kembalikan objek yang berisi jumlah kemunculan setiap karakter dalam string tersebut. Kunci objek diurutkan secara alfabetis.

Format input: string s (tanpa tanda kutip di input, tapi output adalah JSON object)
Format output: JSON object dengan key karakter dan value jumlah kemunculan

Contoh: hitungHuruf("aabbc") → {"a":2,"b":2,"c":1}
Contoh: hitungHuruf("hello") → {"e":1,"h":1,"l":2,"o":1}`,
			StarterCode: `function hitungHuruf(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"aabbc"`, ExpectedOutput: `{"a":2,"b":2,"c":1}`, IsHidden: false},
				{Input: `"hello"`, ExpectedOutput: `{"e":1,"h":1,"l":2,"o":1}`, IsHidden: false},
				{Input: `"a"`, ExpectedOutput: `{"a":1}`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `{}`, IsHidden: true},
				{Input: `"abcabc"`, ExpectedOutput: `{"a":2,"b":2,"c":2}`, IsHidden: true},
			},
		},
		// ── Loop Medium (10 soal) ────────────────────────────────────────────
		{
			Title:      "Deret Fibonacci",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, kembalikan array berisi N bilangan pertama dari deret Fibonacci.

Deret Fibonacci: 0, 1, 1, 2, 3, 5, 8, 13, ...
Setiap bilangan adalah jumlah dua bilangan sebelumnya. F(0)=0, F(1)=1.

Format input: bilangan bulat positif N
Format output: array bilangan bulat

Contoh: deretFibonacci(6) → [0,1,1,2,3,5]
Contoh: deretFibonacci(1) → [0]`,
			StarterCode: `function deretFibonacci(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "6", ExpectedOutput: "[0,1,1,2,3,5]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[0]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[0,1]", IsHidden: false},
				{Input: "10", ExpectedOutput: "[0,1,1,2,3,5,8,13,21,34]", IsHidden: true},
				{Input: "3", ExpectedOutput: "[0,1,1]", IsHidden: true},
			},
		},
		{
			Title:      "Bilangan Sempurna",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, tentukan apakah N adalah bilangan sempurna.

Bilangan sempurna adalah bilangan yang sama dengan jumlah semua pembaginya yang lebih kecil dari bilangan itu sendiri.
Contoh: 6 = 1 + 2 + 3 (pembagi 6 selain 6 sendiri adalah 1, 2, 3)

Format input: bilangan bulat positif N
Format output: boolean (true/false)

Contoh: bilanganSempurna(6) → true
Contoh: bilanganSempurna(28) → true
Contoh: bilanganSempurna(12) → false`,
			StarterCode: `function bilanganSempurna(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "6", ExpectedOutput: "true", IsHidden: false},
				{Input: "28", ExpectedOutput: "true", IsHidden: false},
				{Input: "12", ExpectedOutput: "false", IsHidden: false},
				{Input: "1", ExpectedOutput: "false", IsHidden: true},
				{Input: "496", ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Konversi Basis",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan desimal non-negatif N, kembalikan representasi binernya sebagai string.

Format input: bilangan bulat non-negatif N
Format output: string biner

Contoh: konversiBasis(10) → "1010"
Contoh: konversiBasis(0) → "0"
Contoh: konversiBasis(1) → "1"`,
			StarterCode: `function konversiBasis(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "10", ExpectedOutput: `"1010"`, IsHidden: false},
				{Input: "0", ExpectedOutput: `"0"`, IsHidden: false},
				{Input: "1", ExpectedOutput: `"1"`, IsHidden: false},
				{Input: "255", ExpectedOutput: `"11111111"`, IsHidden: true},
				{Input: "8", ExpectedOutput: `"1000"`, IsHidden: true},
			},
		},
		{
			Title:      "Segitiga Bintang",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, kembalikan array berisi N string yang membentuk segitiga siku-siku dari bintang (*).

Baris ke-i (mulai dari 1) berisi i bintang.

Format input: bilangan bulat positif N
Format output: array string

Contoh: segitigaBintang(4) → ["*","**","***","****"]
Contoh: segitigaBintang(1) → ["*"]`,
			StarterCode: `function segitigaBintang(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "4", ExpectedOutput: `["*","**","***","****"]`, IsHidden: false},
				{Input: "1", ExpectedOutput: `["*"]`, IsHidden: false},
				{Input: "3", ExpectedOutput: `["*","**","***"]`, IsHidden: false},
				{Input: "2", ExpectedOutput: `["*","**"]`, IsHidden: true},
				{Input: "5", ExpectedOutput: `["*","**","***","****","*****"]`, IsHidden: true},
			},
		},
		{
			Title:      "Jumlah Digit Berulang",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, jumlahkan digit-digitnya secara berulang sampai hasilnya menjadi satu digit, lalu kembalikan hasilnya.

Format input: bilangan bulat positif N
Format output: bilangan bulat (1 digit)

Contoh: jumlahDigitBerulang(9875) → 2 (9+8+7+5=29 → 2+9=11 → 1+1=2)
Contoh: jumlahDigitBerulang(0) → 0`,
			StarterCode: `function jumlahDigitBerulang(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "9875", ExpectedOutput: "2", IsHidden: false},
				{Input: "0", ExpectedOutput: "0", IsHidden: false},
				{Input: "9", ExpectedOutput: "9", IsHidden: false},
				{Input: "199", ExpectedOutput: "1", IsHidden: true},
				{Input: "12345", ExpectedOutput: "6", IsHidden: true},
			},
		},
		{
			Title:      "Angka Armstrong",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, tentukan apakah N adalah bilangan Armstrong.

Bilangan Armstrong adalah bilangan yang sama dengan jumlah setiap digitnya dipangkatkan dengan jumlah digit bilangan tersebut.
Contoh: 153 = 1³ + 5³ + 3³ = 1 + 125 + 27 = 153

Format input: bilangan bulat positif N
Format output: boolean (true/false)

Contoh: angkaArmstrong(153) → true
Contoh: angkaArmstrong(370) → true
Contoh: angkaArmstrong(100) → false`,
			StarterCode: `function angkaArmstrong(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "153", ExpectedOutput: "true", IsHidden: false},
				{Input: "370", ExpectedOutput: "true", IsHidden: false},
				{Input: "100", ExpectedOutput: "false", IsHidden: false},
				{Input: "1", ExpectedOutput: "true", IsHidden: true},
				{Input: "9474", ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Pola Berlian",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan ganjil positif N, kembalikan array berisi N string yang membentuk pola berlian dari bintang (*).

Baris tengah (ke-(N+1)/2) memiliki N bintang. Baris-baris di atas dan bawahnya simetris.

Format input: bilangan ganjil positif N
Format output: array string

Contoh: polaBerlian(5) → ["  *  ","  ***  "," ***** ","  ***  ","  *  "]

Catatan: gunakan spasi untuk padding agar lebar setiap baris sama dengan N karakter.
Contoh: polaBerlian(3) → [" * ","***"," * "]`,
			StarterCode: `function polaBerlian(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "3", ExpectedOutput: `[" * ","***"," * "]`, IsHidden: false},
				{Input: "1", ExpectedOutput: `["*"]`, IsHidden: false},
				{Input: "5", ExpectedOutput: `["  *  "," *** ","*****"," *** ","  *  "]`, IsHidden: false},
				{Input: "7", ExpectedOutput: `["   *   ","  ***  "," ***** ","*******"," ***** ","  ***  ","   *   "]`, IsHidden: true},
				{Input: "9", ExpectedOutput: `["    *    ","   ***   ","  *****  "," ******* ","*********"," ******* ","  *****  ","   ***   ","    *    "]`, IsHidden: true},
			},
		},
		{
			Title:      "Bilangan Romawi",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat N (1 ≤ N ≤ 3999), kembalikan representasi angka Romawi dari N.

Simbol Romawi: I=1, V=5, X=10, L=50, C=100, D=500, M=1000
Aturan pengurangan: IV=4, IX=9, XL=40, XC=90, CD=400, CM=900

Format input: bilangan bulat N (1 ≤ N ≤ 3999)
Format output: string angka Romawi

Contoh: bilanganRomawi(3) → "III"
Contoh: bilanganRomawi(4) → "IV"
Contoh: bilanganRomawi(1994) → "MCMXCIV"`,
			StarterCode: `function bilanganRomawi(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "3", ExpectedOutput: `"III"`, IsHidden: false},
				{Input: "4", ExpectedOutput: `"IV"`, IsHidden: false},
				{Input: "1994", ExpectedOutput: `"MCMXCIV"`, IsHidden: false},
				{Input: "1", ExpectedOutput: `"I"`, IsHidden: true},
				{Input: "3999", ExpectedOutput: `"MMMCMXCIX"`, IsHidden: true},
			},
		},
		{
			Title:      "Urutan Collatz",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan bulat positif N, kembalikan urutan Collatz mulai dari N hingga mencapai 1 sebagai array.

Aturan Collatz:
- Jika N genap: N berikutnya = N / 2
- Jika N ganjil: N berikutnya = 3 × N + 1

Format input: bilangan bulat positif N
Format output: array bilangan bulat (termasuk N dan 1)

Contoh: urutanCollatz(6) → [6,3,10,5,16,8,4,2,1]
Contoh: urutanCollatz(1) → [1]`,
			StarterCode: `function urutanCollatz(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "6", ExpectedOutput: "[6,3,10,5,16,8,4,2,1]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[1]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[2,1]", IsHidden: false},
				{Input: "5", ExpectedOutput: "[5,16,8,4,2,1]", IsHidden: true},
				{Input: "27", ExpectedOutput: "[27,82,41,124,62,31,94,47,142,71,214,107,322,161,484,242,121,364,182,91,274,137,412,206,103,310,155,466,233,700,350,175,526,263,790,395,1186,593,1780,890,445,1336,668,334,167,502,251,754,377,1132,566,283,850,425,1276,638,319,958,479,1438,719,2158,1079,3238,1619,4858,2429,7288,3644,1822,911,2734,1367,4102,2051,6154,3077,9232,4616,2308,1154,577,1732,866,433,1300,650,325,976,488,244,122,61,184,92,46,23,70,35,106,53,160,80,40,20,10,5,16,8,4,2,1]", IsHidden: true},
			},
		},
		{
			Title:      "Pangkat N",
			Category:   "loop",
			Difficulty: "medium",
			Description: `Diberikan bilangan basis (base) dan eksponen (exp), kembalikan hasil base^exp tanpa menggunakan Math.pow atau operator **.

Format input: "base exp" (dua bilangan bulat dipisah spasi, exp >= 0)
Format output: bilangan bulat

Contoh: pangkatN(2, 10) → 1024
Contoh: pangkatN(3, 0) → 1
Contoh: pangkatN(5, 3) → 125`,
			StarterCode: `function pangkatN(base, exp) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "2 10", ExpectedOutput: "1024", IsHidden: false},
				{Input: "3 0", ExpectedOutput: "1", IsHidden: false},
				{Input: "5 3", ExpectedOutput: "125", IsHidden: false},
				{Input: "1 100", ExpectedOutput: "1", IsHidden: true},
				{Input: "7 4", ExpectedOutput: "2401", IsHidden: true},
			},
		},
		// ── Loop Hard (5 soal) ───────────────────────────────────────────────
		{
			Title:      "Spiral Matrix",
			Category:   "loop",
			Difficulty: "hard",
			Description: `Diberikan bilangan bulat positif N, buat matrix N×N yang diisi angka 1 sampai N² secara spiral searah jarum jam, dimulai dari pojok kiri atas.

Kembalikan matrix sebagai array 2D (array of arrays).

Format input: bilangan bulat positif N
Format output: array 2D bilangan bulat

Contoh: spiralMatrix(3) → [[1,2,3],[8,9,4],[7,6,5]]
Contoh: spiralMatrix(1) → [[1]]
Contoh: spiralMatrix(2) → [[1,2],[4,3]]`,
			StarterCode: `function spiralMatrix(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "3", ExpectedOutput: "[[1,2,3],[8,9,4],[7,6,5]]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[[1]]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[[1,2],[4,3]]", IsHidden: false},
				{Input: "4", ExpectedOutput: "[[1,2,3,4],[12,13,14,5],[11,16,15,6],[10,9,8,7]]", IsHidden: false},
				{Input: "5", ExpectedOutput: "[[1,2,3,4,5],[16,17,18,19,6],[15,24,25,20,7],[14,23,22,21,8],[13,12,11,10,9]]", IsHidden: true},
				{Input: "6", ExpectedOutput: "[[1,2,3,4,5,6],[20,21,22,23,24,7],[19,32,33,34,25,8],[18,31,36,35,26,9],[17,30,29,28,27,10],[16,15,14,13,12,11]]", IsHidden: true},
				{Input: "7", ExpectedOutput: "[[1,2,3,4,5,6,7],[24,25,26,27,28,29,8],[23,40,41,42,43,30,9],[22,39,48,49,44,31,10],[21,38,47,46,45,32,11],[20,37,36,35,34,33,12],[19,18,17,16,15,14,13]]", IsHidden: true},
			},
		},
		{
			Title:      "Segitiga Pascal",
			Category:   "loop",
			Difficulty: "hard",
			Description: `Diberikan bilangan bulat positif N, kembalikan N baris pertama dari segitiga Pascal sebagai array 2D.

Setiap baris ke-i (mulai dari 0) berisi i+1 elemen. Elemen pertama dan terakhir selalu 1. Elemen lainnya adalah jumlah dua elemen di atasnya.

Format input: bilangan bulat positif N
Format output: array 2D bilangan bulat

Contoh: segitigaPascal(5) → [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Contoh: segitigaPascal(1) → [[1]]`,
			StarterCode: `function segitigaPascal(n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "5", ExpectedOutput: "[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]", IsHidden: false},
				{Input: "1", ExpectedOutput: "[[1]]", IsHidden: false},
				{Input: "2", ExpectedOutput: "[[1],[1,1]]", IsHidden: false},
				{Input: "3", ExpectedOutput: "[[1],[1,1],[1,2,1]]", IsHidden: false},
				{Input: "6", ExpectedOutput: "[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1],[1,5,10,10,5,1]]", IsHidden: true},
				{Input: "7", ExpectedOutput: "[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1],[1,5,10,10,5,1],[1,6,15,20,15,6,1]]", IsHidden: true},
				{Input: "4", ExpectedOutput: "[[1],[1,1],[1,2,1],[1,3,3,1]]", IsHidden: true},
			},
		},
		{
			Title:      "Permutasi",
			Category:   "loop",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat yang unik, kembalikan semua permutasi yang mungkin sebagai array 2D, diurutkan secara leksikografis.

Format input: array bilangan bulat (elemen unik)
Format output: array 2D bilangan bulat (diurutkan leksikografis)

Contoh: permutasi([1,2,3]) → [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Contoh: permutasi([1]) → [[1]]
Contoh: permutasi([]) → [[]]`,
			StarterCode: `function permutasi(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "[1,2,3]", ExpectedOutput: "[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]", IsHidden: false},
				{Input: "[1]", ExpectedOutput: "[[1]]", IsHidden: false},
				{Input: "[]", ExpectedOutput: "[[]]", IsHidden: false},
				{Input: "[1,2]", ExpectedOutput: "[[1,2],[2,1]]", IsHidden: false},
				{Input: "[1,2,3,4]", ExpectedOutput: "[[1,2,3,4],[1,2,4,3],[1,3,2,4],[1,3,4,2],[1,4,2,3],[1,4,3,2],[2,1,3,4],[2,1,4,3],[2,3,1,4],[2,3,4,1],[2,4,1,3],[2,4,3,1],[3,1,2,4],[3,1,4,2],[3,2,1,4],[3,2,4,1],[3,4,1,2],[3,4,2,1],[4,1,2,3],[4,1,3,2],[4,2,1,3],[4,2,3,1],[4,3,1,2],[4,3,2,1]]", IsHidden: true},
				{Input: "[2,1,3]", ExpectedOutput: "[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]", IsHidden: true},
				{Input: "[3,1]", ExpectedOutput: "[[1,3],[3,1]]", IsHidden: true},
			},
		},
		{
			Title:      "Bilangan Romawi ke Desimal",
			Category:   "loop",
			Difficulty: "hard",
			Description: `Diberikan sebuah string angka Romawi yang valid, kembalikan nilai desimalnya.

Simbol Romawi: I=1, V=5, X=10, L=50, C=100, D=500, M=1000
Aturan pengurangan: IV=4, IX=9, XL=40, XC=90, CD=400, CM=900

Format input: string angka Romawi (valid, 1 ≤ nilai ≤ 3999)
Format output: bilangan bulat

Contoh: romanKeDesimal("III") → 3
Contoh: romanKeDesimal("IV") → 4
Contoh: romanKeDesimal("MCMXCIV") → 1994`,
			StarterCode: `function romanKeDesimal(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"III"`, ExpectedOutput: "3", IsHidden: false},
				{Input: `"IV"`, ExpectedOutput: "4", IsHidden: false},
				{Input: `"MCMXCIV"`, ExpectedOutput: "1994", IsHidden: false},
				{Input: `"I"`, ExpectedOutput: "1", IsHidden: false},
				{Input: `"MMMCMXCIX"`, ExpectedOutput: "3999", IsHidden: true},
				{Input: `"LVIII"`, ExpectedOutput: "58", IsHidden: true},
				{Input: `"CDXLIV"`, ExpectedOutput: "444", IsHidden: true},
			},
		},
		{
			Title:      "Jumlah Subarray",
			Category:   "loop",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat dan sebuah target, kembalikan jumlah subarray yang memiliki jumlah (sum) sama dengan target.

Subarray adalah bagian berurutan dari array (minimal 1 elemen).

Format input: "arr target" — array JSON diikuti spasi dan bilangan bulat target
Format output: bilangan bulat (jumlah subarray)

Contoh: jumlahSubarray([1,1,1], 2) → 2
Contoh: jumlahSubarray([1,2,3], 3) → 2 (subarray [1,2] dan [3])
Contoh: jumlahSubarray([], 0) → 0`,
			StarterCode: `function jumlahSubarray(arr, target) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "[1,1,1] 2", ExpectedOutput: "2", IsHidden: false},
				{Input: "[1,2,3] 3", ExpectedOutput: "2", IsHidden: false},
				{Input: "[] 0", ExpectedOutput: "0", IsHidden: false},
				{Input: "[1] 1", ExpectedOutput: "1", IsHidden: false},
				{Input: "[3,4,7,2,-3,1,4,2] 7", ExpectedOutput: "4", IsHidden: true},
				{Input: "[1,-1,1,-1] 0", ExpectedOutput: "4", IsHidden: true},
				{Input: "[0,0,0,0] 0", ExpectedOutput: "10", IsHidden: true},
			},
		},
		// ── String Easy (10 soal) ───────────────────────────────────────────────
		{
			Title:      "Hitung Vokal",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, hitung jumlah huruf vokal (a, e, i, o, u) di dalamnya. Huruf besar dan kecil dihitung sama.

Format input: string
Format output: bilangan bulat

Contoh: hitungVokal("halo dunia") → 5
Contoh: hitungVokal("") → 0
Contoh: hitungVokal("xyz") → 0`,
			StarterCode: `function hitungVokal(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia"`, ExpectedOutput: "5", IsHidden: false},
				{Input: `""`, ExpectedOutput: "0", IsHidden: false},
				{Input: `"xyz"`, ExpectedOutput: "0", IsHidden: false},
				{Input: `"a"`, ExpectedOutput: "1", IsHidden: true},
				{Input: `"AEIOU"`, ExpectedOutput: "5", IsHidden: true},
			},
		},
		{
			Title:      "Kapital Pertama",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, ubah huruf pertama setiap kata menjadi huruf kapital. Kata dipisahkan oleh spasi.

Format input: string
Format output: string

Contoh: kapitalPertama("halo dunia") → "Halo Dunia"
Contoh: kapitalPertama("") → ""
Contoh: kapitalPertama("a b c") → "A B C"`,
			StarterCode: `function kapitalPertama(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia"`, ExpectedOutput: `"Halo Dunia"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"a b c"`, ExpectedOutput: `"A B C"`, IsHidden: false},
				{Input: `"satu dua tiga"`, ExpectedOutput: `"Satu Dua Tiga"`, IsHidden: true},
				{Input: `"javascript adalah bahasa pemrograman"`, ExpectedOutput: `"Javascript Adalah Bahasa Pemrograman"`, IsHidden: true},
			},
		},
		{
			Title:      "Hapus Spasi",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, hapus semua karakter spasi di dalamnya.

Format input: string
Format output: string

Contoh: hapusSpasi("halo dunia") → "halodunia"
Contoh: hapusSpasi("  a b  ") → "ab"
Contoh: hapusSpasi("") → ""`,
			StarterCode: `function hapusSpasi(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia"`, ExpectedOutput: `"halodunia"`, IsHidden: false},
				{Input: `"  a b  "`, ExpectedOutput: `"ab"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"tanpa spasi"`, ExpectedOutput: `"tanpaspasi"`, IsHidden: true},
				{Input: `"   "`, ExpectedOutput: `""`, IsHidden: true},
			},
		},
		{
			Title:      "Ulangi String",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string s dan bilangan bulat n, kembalikan string s yang diulang sebanyak n kali.

Format input: "s n" (string dikutip, diikuti spasi dan bilangan bulat)
Format output: string

Contoh: ulangiString("ab", 3) → "ababab"
Contoh: ulangiString("x", 0) → ""
Contoh: ulangiString("", 5) → ""`,
			StarterCode: `function ulangiString(s, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"ab" 3`, ExpectedOutput: `"ababab"`, IsHidden: false},
				{Input: `"x" 0`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"" 5`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"ha" 4`, ExpectedOutput: `"hahahaha"`, IsHidden: true},
				{Input: `"abc" 1`, ExpectedOutput: `"abc"`, IsHidden: true},
			},
		},
		{
			Title:      "Cek Angka",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, kembalikan true jika string tersebut hanya mengandung digit (0-9), dan false jika tidak. String kosong dianggap false.

Format input: string
Format output: boolean (true/false)

Contoh: cekAngka("12345") → true
Contoh: cekAngka("123a5") → false
Contoh: cekAngka("") → false`,
			StarterCode: `function cekAngka(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"12345"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"123a5"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `""`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"0"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"99 99"`, ExpectedOutput: "false", IsHidden: true},
			},
		},
		{
			Title:      "Balik Kata",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah kalimat, balik urutan kata-katanya (bukan karakter). Kata dipisahkan oleh satu spasi.

Format input: string
Format output: string

Contoh: balikKata("halo dunia") → "dunia halo"
Contoh: balikKata("satu dua tiga") → "tiga dua satu"
Contoh: balikKata("satu") → "satu"`,
			StarterCode: `function balikKata(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia"`, ExpectedOutput: `"dunia halo"`, IsHidden: false},
				{Input: `"satu dua tiga"`, ExpectedOutput: `"tiga dua satu"`, IsHidden: false},
				{Input: `"satu"`, ExpectedOutput: `"satu"`, IsHidden: false},
				{Input: `"a b c d"`, ExpectedOutput: `"d c b a"`, IsHidden: true},
				{Input: `"belajar koding itu menyenangkan"`, ExpectedOutput: `"menyenangkan itu koding belajar"`, IsHidden: true},
			},
		},
		{
			Title:      "Hitung Kata",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah kalimat, hitung jumlah kata di dalamnya. Kata dipisahkan oleh satu atau lebih spasi. String kosong menghasilkan 0.

Format input: string
Format output: bilangan bulat

Contoh: hitungKata("halo dunia") → 2
Contoh: hitungKata("satu dua tiga") → 3
Contoh: hitungKata("") → 0`,
			StarterCode: `function hitungKata(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia"`, ExpectedOutput: "2", IsHidden: false},
				{Input: `"satu dua tiga"`, ExpectedOutput: "3", IsHidden: false},
				{Input: `""`, ExpectedOutput: "0", IsHidden: false},
				{Input: `"satu"`, ExpectedOutput: "1", IsHidden: true},
				{Input: `"a b c d e"`, ExpectedOutput: "5", IsHidden: true},
			},
		},
		{
			Title:      "Ganti Karakter",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string s, karakter lama (oldChar), dan karakter baru (newChar), ganti semua kemunculan oldChar dengan newChar.

Format input: "s oldChar newChar" (string dikutip, diikuti dua karakter tunggal)
Format output: string

Contoh: gantiKarakter("halo", "a", "e") → "helo"
Contoh: gantiKarakter("banana", "a", "o") → "bonono"
Contoh: gantiKarakter("", "a", "b") → ""`,
			StarterCode: `function gantiKarakter(s, oldChar, newChar) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo" "a" "e"`, ExpectedOutput: `"helo"`, IsHidden: false},
				{Input: `"banana" "a" "o"`, ExpectedOutput: `"bonono"`, IsHidden: false},
				{Input: `"" "a" "b"`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"mississippi" "s" "z"`, ExpectedOutput: `"mizzizzippi"`, IsHidden: true},
				{Input: `"aaa" "a" "a"`, ExpectedOutput: `"aaa"`, IsHidden: true},
			},
		},
		{
			Title:      "Cek Huruf Besar",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string, kembalikan true jika semua huruf dalam string tersebut adalah huruf kapital (huruf besar). Karakter non-huruf diabaikan. String kosong atau string tanpa huruf dianggap false.

Format input: string
Format output: boolean (true/false)

Contoh: cekHurufBesar("HALO") → true
Contoh: cekHurufBesar("Halo") → false
Contoh: cekHurufBesar("HALO 123") → true`,
			StarterCode: `function cekHurufBesar(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"HALO"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"Halo"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"HALO 123"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `""`, ExpectedOutput: "false", IsHidden: true},
				{Input: `"ABC"`, ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Potong String",
			Category:   "string",
			Difficulty: "easy",
			Description: `Diberikan sebuah string s dan bilangan bulat n, potong string menjadi maksimal n karakter. Jika string lebih panjang dari n, tambahkan "..." di akhir. Jika tidak, kembalikan string aslinya.

Format input: "s n" (string dikutip, diikuti spasi dan bilangan bulat)
Format output: string

Contoh: potongString("halo dunia", 4) → "halo..."
Contoh: potongString("halo", 10) → "halo"
Contoh: potongString("", 5) → ""`,
			StarterCode: `function potongString(s, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"halo dunia" 4`, ExpectedOutput: `"halo..."`, IsHidden: false},
				{Input: `"halo" 10`, ExpectedOutput: `"halo"`, IsHidden: false},
				{Input: `"" 5`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"javascript" 4`, ExpectedOutput: `"java..."`, IsHidden: true},
				{Input: `"abc" 3`, ExpectedOutput: `"abc"`, IsHidden: true},
			},
		},
		// ── String Medium (10 soal) ─────────────────────────────────────────────
		{
			Title:      "Anagram",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan dua string, kembalikan true jika keduanya adalah anagram satu sama lain. Abaikan huruf besar/kecil dan spasi.

Anagram adalah kata yang dibentuk dari huruf-huruf yang sama dengan urutan berbeda.

Format input: "s1 s2" (dua string dikutip dipisah spasi)
Format output: boolean (true/false)

Contoh: anagram("listen", "silent") → true
Contoh: anagram("hello", "world") → false
Contoh: anagram("Astronomer", "Moon starer") → true`,
			StarterCode: `function anagram(s1, s2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"listen" "silent"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"hello" "world"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"Astronomer" "Moon starer"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"" ""`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"abc" "cba"`, ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Kompresi String",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string, lakukan kompresi run-length encoding: ganti setiap urutan karakter yang berulang dengan karakter diikuti jumlah pengulangannya. Jika hasil kompresi lebih panjang atau sama dengan string asli, kembalikan string aslinya.

Format input: string
Format output: string

Contoh: kompresiString("aabcccdddd") → "a2bc3d4"
Contoh: kompresiString("abc") → "abc"
Contoh: kompresiString("aabb") → "aabb"`,
			StarterCode: `function kompresiString(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"aabcccdddd"`, ExpectedOutput: `"a2bc3d4"`, IsHidden: false},
				{Input: `"abc"`, ExpectedOutput: `"abc"`, IsHidden: false},
				{Input: `"aabb"`, ExpectedOutput: `"aabb"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: true},
				{Input: `"aaaaaaaaaa"`, ExpectedOutput: `"a10"`, IsHidden: true},
			},
		},
		{
			Title:      "Validasi Email",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string, kembalikan true jika string tersebut adalah alamat email yang valid. Aturan validasi dasar: harus mengandung tepat satu karakter '@', bagian setelah '@' harus mengandung setidaknya satu titik ('.'), dan tidak boleh ada spasi.

Format input: string
Format output: boolean (true/false)

Contoh: validasiEmail("user@example.com") → true
Contoh: validasiEmail("userexample.com") → false
Contoh: validasiEmail("user@") → false`,
			StarterCode: `function validasiEmail(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"user@example.com"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"userexample.com"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"user@"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"a@b.c"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"user @example.com"`, ExpectedOutput: "false", IsHidden: true},
			},
		},
		{
			Title:      "Hitung Substring",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string teks dan sebuah substring, hitung berapa kali substring tersebut muncul secara non-overlapping di dalam teks.

Format input: "teks sub" (dua string dikutip dipisah spasi)
Format output: bilangan bulat

Contoh: hitungSubstring("banana", "an") → 2
Contoh: hitungSubstring("hello", "ll") → 1
Contoh: hitungSubstring("aaa", "aa") → 1`,
			StarterCode: `function hitungSubstring(teks, sub) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"banana" "an"`, ExpectedOutput: "2", IsHidden: false},
				{Input: `"hello" "ll"`, ExpectedOutput: "1", IsHidden: false},
				{Input: `"aaa" "aa"`, ExpectedOutput: "1", IsHidden: false},
				{Input: `"" "a"`, ExpectedOutput: "0", IsHidden: true},
				{Input: `"abcabcabc" "abc"`, ExpectedOutput: "3", IsHidden: true},
			},
		},
		{
			Title:      "Rotasi String",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan dua string s1 dan s2, kembalikan true jika s2 adalah rotasi dari s1. Rotasi berarti memindahkan beberapa karakter dari awal ke akhir (atau sebaliknya).

Format input: "s1 s2" (dua string dikutip dipisah spasi)
Format output: boolean (true/false)

Contoh: rotasiString("abcde", "cdeab") → true
Contoh: rotasiString("abcde", "abced") → false
Contoh: rotasiString("", "") → true`,
			StarterCode: `function rotasiString(s1, s2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"abcde" "cdeab"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"abcde" "abced"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"" ""`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"waterbottle" "erbottlewat"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"abc" "abcd"`, ExpectedOutput: "false", IsHidden: true},
			},
		},
		{
			Title:      "Hapus Duplikat",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string, hapus karakter yang duplikat dan pertahankan hanya kemunculan pertama setiap karakter. Urutan karakter yang tersisa harus dipertahankan.

Format input: string
Format output: string

Contoh: hapusDuplikat("programming") → "progamin"
Contoh: hapusDuplikat("aabbcc") → "abc"
Contoh: hapusDuplikat("") → ""`,
			StarterCode: `function hapusDuplikat(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"programming"`, ExpectedOutput: `"progamin"`, IsHidden: false},
				{Input: `"aabbcc"`, ExpectedOutput: `"abc"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"abcabc"`, ExpectedOutput: `"abc"`, IsHidden: true},
				{Input: `"a"`, ExpectedOutput: `"a"`, IsHidden: true},
			},
		},
		{
			Title:      "Caesar Cipher",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string s dan bilangan bulat n (shift), geser setiap huruf sebanyak n posisi dalam alfabet. Huruf besar tetap besar, huruf kecil tetap kecil. Karakter non-huruf tidak diubah. Pergeseran melingkar (z+1 = a).

Format input: "s n" (string dikutip, diikuti spasi dan bilangan bulat)
Format output: string

Contoh: caesarCipher("Hello, World!", 3) → "Khoor, Zruog!"
Contoh: caesarCipher("xyz", 3) → "abc"
Contoh: caesarCipher("", 5) → ""`,
			StarterCode: `function caesarCipher(s, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"Hello, World!" 3`, ExpectedOutput: `"Khoor, Zruog!"`, IsHidden: false},
				{Input: `"xyz" 3`, ExpectedOutput: `"abc"`, IsHidden: false},
				{Input: `"" 5`, ExpectedOutput: `""`, IsHidden: false},
				{Input: `"ABC" 1`, ExpectedOutput: `"BCD"`, IsHidden: true},
				{Input: `"Kode 123!" 13`, ExpectedOutput: `"Xbqr 123!"`, IsHidden: true},
			},
		},
		{
			Title:      "Cek Tanda Kurung",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string yang hanya berisi karakter '(' dan ')', kembalikan true jika tanda kurung seimbang (setiap kurung buka memiliki pasangan kurung tutup yang sesuai).

Format input: string
Format output: boolean (true/false)

Contoh: cekTandaKurung("((()))") → true
Contoh: cekTandaKurung("(()") → false
Contoh: cekTandaKurung("") → true`,
			StarterCode: `function cekTandaKurung(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"((()))"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"(()"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `""`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"()()()"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `")("`, ExpectedOutput: "false", IsHidden: true},
			},
		},
		{
			Title:      "Kata Terpanjang",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah kalimat, kembalikan kata terpanjang di dalamnya. Jika ada beberapa kata dengan panjang yang sama, kembalikan kata yang pertama muncul. Kata dipisahkan oleh spasi.

Format input: string
Format output: string

Contoh: kataTerpanjang("saya suka belajar koding") → "belajar"
Contoh: kataTerpanjang("halo") → "halo"
Contoh: kataTerpanjang("a bb ccc") → "ccc"`,
			StarterCode: `function kataTerpanjang(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"saya suka belajar koding"`, ExpectedOutput: `"belajar"`, IsHidden: false},
				{Input: `"halo"`, ExpectedOutput: `"halo"`, IsHidden: false},
				{Input: `"a bb ccc"`, ExpectedOutput: `"ccc"`, IsHidden: false},
				{Input: `"satu dua"`, ExpectedOutput: `"satu"`, IsHidden: true},
				{Input: `"javascript python golang"`, ExpectedOutput: `"javascript"`, IsHidden: true},
			},
		},
		{
			Title:      "Pangkat String",
			Category:   "string",
			Difficulty: "medium",
			Description: `Diberikan sebuah string s dan bilangan bulat n, kembalikan string s yang diulang sebanyak n kali dengan pemisah "-" di antara setiap pengulangan.

Format input: "s n" (string dikutip, diikuti spasi dan bilangan bulat positif)
Format output: string

Contoh: pangkatString("ha", 3) → "ha-ha-ha"
Contoh: pangkatString("go", 1) → "go"
Contoh: pangkatString("ab", 2) → "ab-ab"`,
			StarterCode: `function pangkatString(s, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"ha" 3`, ExpectedOutput: `"ha-ha-ha"`, IsHidden: false},
				{Input: `"go" 1`, ExpectedOutput: `"go"`, IsHidden: false},
				{Input: `"ab" 2`, ExpectedOutput: `"ab-ab"`, IsHidden: false},
				{Input: `"x" 5`, ExpectedOutput: `"x-x-x-x-x"`, IsHidden: true},
				{Input: `"kode" 4`, ExpectedOutput: `"kode-kode-kode-kode"`, IsHidden: true},
			},
		},
		// ── String Hard (5 soal) ────────────────────────────────────────────────
		{
			Title:      "Anagram Substring",
			Category:   "string",
			Difficulty: "hard",
			Description: `Diberikan sebuah string teks dan sebuah pola (pattern), kembalikan array semua indeks awal di mana anagram dari pola ditemukan di dalam teks. Hasil harus diurutkan secara ascending.

Format input: "teks pola" (dua string dikutip dipisah spasi)
Format output: array bilangan bulat (indeks, 0-based)

Contoh: anagramSubstring("cbaebabacd", "abc") → [0,6]
Contoh: anagramSubstring("abab", "ab") → [0,1,2]
Contoh: anagramSubstring("af", "be") → []`,
			StarterCode: `function anagramSubstring(teks, pola) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"cbaebabacd" "abc"`, ExpectedOutput: "[0,6]", IsHidden: false},
				{Input: `"abab" "ab"`, ExpectedOutput: "[0,1,2]", IsHidden: false},
				{Input: `"af" "be"`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `"baa" "aa"`, ExpectedOutput: "[1]", IsHidden: false},
				{Input: `"aaaaaaaaaa" "aaa"`, ExpectedOutput: "[0,1,2,3,4,5,6,7]", IsHidden: true},
				{Input: `"" "a"`, ExpectedOutput: "[]", IsHidden: true},
				{Input: `"abcdefg" "xyz"`, ExpectedOutput: "[]", IsHidden: true},
			},
		},
		{
			Title:      "Validasi Palindrom Lanjutan",
			Category:   "string",
			Difficulty: "hard",
			Description: `Diberikan sebuah string, kembalikan true jika string tersebut adalah palindrom setelah mengabaikan semua karakter non-alfanumerik dan mengabaikan perbedaan huruf besar/kecil.

Format input: string
Format output: boolean (true/false)

Contoh: validasiPalindromLanjutan("A man, a plan, a canal: Panama") → true
Contoh: validasiPalindromLanjutan("race a car") → false
Contoh: validasiPalindromLanjutan("") → true`,
			StarterCode: `function validasiPalindromLanjutan(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"A man, a plan, a canal: Panama"`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"race a car"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `""`, ExpectedOutput: "true", IsHidden: false},
				{Input: `"hello"`, ExpectedOutput: "false", IsHidden: false},
				{Input: `"Was it a car or a cat I saw?"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"No lemon, no melon"`, ExpectedOutput: "true", IsHidden: true},
				{Input: `"a"`, ExpectedOutput: "true", IsHidden: true},
			},
		},
		{
			Title:      "Encode URL",
			Category:   "string",
			Difficulty: "hard",
			Description: `Diberikan sebuah URL string, lakukan encoding dengan mengganti karakter-karakter khusus berikut:
- Spasi → %20
- & → %26
- = → %3D
- ? → %3F

Karakter lain tidak diubah.

Format input: string
Format output: string

Contoh: encodeURL("hello world") → "hello%20world"
Contoh: encodeURL("a=1&b=2") → "a%3D1%26b%3D2"
Contoh: encodeURL("search?q=hello world&lang=id") → "search%3Fq%3Dhello%20world%26lang%3Did"`,
			StarterCode: `function encodeURL(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"hello world"`, ExpectedOutput: `"hello%20world"`, IsHidden: false},
				{Input: `"a=1&b=2"`, ExpectedOutput: `"a%3D1%26b%3D2"`, IsHidden: false},
				{Input: `"search?q=hello world&lang=id"`, ExpectedOutput: `"search%3Fq%3Dhello%20world%26lang%3Did"`, IsHidden: false},
				{Input: `"normal"`, ExpectedOutput: `"normal"`, IsHidden: false},
				{Input: `""`, ExpectedOutput: `""`, IsHidden: true},
				{Input: `"no special chars"`, ExpectedOutput: `"no%20special%20chars"`, IsHidden: true},
				{Input: `"a?b=c&d=e f"`, ExpectedOutput: `"a%3Fb%3Dc%26d%3De%20f"`, IsHidden: true},
			},
		},
		{
			Title:      "Jarak Edit",
			Category:   "string",
			Difficulty: "hard",
			Description: `Diberikan dua string, hitung jarak edit (Levenshtein distance) antara keduanya. Jarak edit adalah jumlah minimum operasi (sisipkan, hapus, atau ganti satu karakter) yang diperlukan untuk mengubah string pertama menjadi string kedua.

Format input: "s1 s2" (dua string dikutip dipisah spasi)
Format output: bilangan bulat

Contoh: jarakEdit("kitten", "sitting") → 3
Contoh: jarakEdit("", "abc") → 3
Contoh: jarakEdit("abc", "abc") → 0`,
			StarterCode: `function jarakEdit(s1, s2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"kitten" "sitting"`, ExpectedOutput: "3", IsHidden: false},
				{Input: `"" "abc"`, ExpectedOutput: "3", IsHidden: false},
				{Input: `"abc" "abc"`, ExpectedOutput: "0", IsHidden: false},
				{Input: `"a" "b"`, ExpectedOutput: "1", IsHidden: false},
				{Input: `"" ""`, ExpectedOutput: "0", IsHidden: true},
				{Input: `"horse" "ros"`, ExpectedOutput: "3", IsHidden: true},
				{Input: `"intention" "execution"`, ExpectedOutput: "5", IsHidden: true},
			},
		},
		{
			Title:      "Ekspresi Matematika",
			Category:   "string",
			Difficulty: "hard",
			Description: `Diberikan sebuah string ekspresi matematika sederhana yang hanya mengandung bilangan bulat non-negatif dan operator '+' serta '-' (tanpa spasi), evaluasi dan kembalikan hasilnya.

Format input: string ekspresi (contoh: "1+2-3+4")
Format output: bilangan bulat

Contoh: ekspresiMatematika("1+2-3+4") → 4
Contoh: ekspresiMatematika("10+5-3") → 12
Contoh: ekspresiMatematika("0") → 0`,
			StarterCode: `function ekspresiMatematika(s) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `"1+2-3+4"`, ExpectedOutput: "4", IsHidden: false},
				{Input: `"10+5-3"`, ExpectedOutput: "12", IsHidden: false},
				{Input: `"0"`, ExpectedOutput: "0", IsHidden: false},
				{Input: `"5"`, ExpectedOutput: "5", IsHidden: false},
				{Input: `"100-50+25-10"`, ExpectedOutput: "65", IsHidden: true},
				{Input: `"1+1+1+1+1"`, ExpectedOutput: "5", IsHidden: true},
				{Input: `"99-99"`, ExpectedOutput: "0", IsHidden: true},
			},
		},
		// ===== ARRAY EASY PROBLEMS =====
		{
			Title:      "Elemen Terbesar",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan elemen terbesar dalam array tersebut.

Format Input: Array bilangan bulat
Format Output: Bilangan bulat terbesar

Contoh:
- elemenTerbesar([3, 1, 4, 1, 5, 9, 2, 6]) → 9
- elemenTerbesar([1]) → 1
- elemenTerbesar([-5, -1, -3]) → -1`,
			StarterCode: `function elemenTerbesar(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[3, 1, 4, 1, 5, 9, 2, 6]`, ExpectedOutput: "9", IsHidden: false},
				{Input: `[1]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[-5, -1, -3]`, ExpectedOutput: "-1", IsHidden: false},
				{Input: `[100, 200, 150]`, ExpectedOutput: "200", IsHidden: true},
				{Input: `[0, 0, 0]`, ExpectedOutput: "0", IsHidden: true},
			},
		},
		{
			Title:      "Elemen Terkecil",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan elemen terkecil dalam array tersebut.

Format Input: Array bilangan bulat
Format Output: Bilangan bulat terkecil

Contoh:
- elemenTerkecil([3, 1, 4, 1, 5, 9, 2, 6]) → 1
- elemenTerkecil([1]) → 1
- elemenTerkecil([-5, -1, -3]) → -5`,
			StarterCode: `function elemenTerkecil(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[3, 1, 4, 1, 5, 9, 2, 6]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[1]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[-5, -1, -3]`, ExpectedOutput: "-5", IsHidden: false},
				{Input: `[100, 200, 150]`, ExpectedOutput: "100", IsHidden: true},
				{Input: `[0, 0, 0]`, ExpectedOutput: "0", IsHidden: true},
			},
		},
		{
			Title:      "Rata-rata Array",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan rata-rata dari semua elemennya. Hasil dibulatkan hingga 2 angka desimal.

Format Input: Array bilangan bulat
Format Output: Bilangan desimal dengan 2 angka di belakang koma

Contoh:
- rataRataArray([1, 2, 3, 4, 5]) → 3.00
- rataRataArray([10, 20]) → 15.00
- rataRataArray([7]) → 7.00`,
			StarterCode: `function rataRataArray(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5]`, ExpectedOutput: "3.00", IsHidden: false},
				{Input: `[10, 20]`, ExpectedOutput: "15.00", IsHidden: false},
				{Input: `[7]`, ExpectedOutput: "7.00", IsHidden: false},
				{Input: `[1, 1, 1, 1]`, ExpectedOutput: "1.00", IsHidden: true},
				{Input: `[-2, 0, 2]`, ExpectedOutput: "0.00", IsHidden: true},
			},
		},
		{
			Title:      "Balik Array",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array, kembalikan array tersebut dalam urutan terbalik.

Format Input: Array bilangan bulat
Format Output: Array dalam urutan terbalik

Contoh:
- balikArray([1, 2, 3, 4, 5]) → [5, 4, 3, 2, 1]
- balikArray([1]) → [1]
- balikArray([]) → []`,
			StarterCode: `function balikArray(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5]`, ExpectedOutput: "[5,4,3,2,1]", IsHidden: false},
				{Input: `[1]`, ExpectedOutput: "[1]", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[10, 20, 30]`, ExpectedOutput: "[30,20,10]", IsHidden: true},
				{Input: `[-1, -2, -3]`, ExpectedOutput: "[-3,-2,-1]", IsHidden: true},
			},
		},
		{
			Title:      "Array Unik",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan array baru yang hanya berisi elemen unik (tanpa duplikat), dengan mempertahankan urutan kemunculan pertama.

Format Input: Array bilangan bulat
Format Output: Array tanpa duplikat

Contoh:
- arrayUnik([1, 2, 2, 3, 3, 3]) → [1, 2, 3]
- arrayUnik([5, 5, 5]) → [5]
- arrayUnik([]) → []`,
			StarterCode: `function arrayUnik(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 2, 3, 3, 3]`, ExpectedOutput: "[1,2,3]", IsHidden: false},
				{Input: `[5, 5, 5]`, ExpectedOutput: "[5]", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[3, 1, 2, 1, 3]`, ExpectedOutput: "[3,1,2]", IsHidden: true},
				{Input: `[1, 2, 3, 4]`, ExpectedOutput: "[1,2,3,4]", IsHidden: true},
			},
		},
		{
			Title:      "Cek Terurut",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan true jika array tersebut terurut secara ascending (dari kecil ke besar), dan false jika tidak.

Format Input: Array bilangan bulat
Format Output: Boolean (true/false)

Contoh:
- cekTerurut([1, 2, 3, 4, 5]) → true
- cekTerurut([1, 3, 2, 4]) → false
- cekTerurut([1]) → true`,
			StarterCode: `function cekTerurut(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5]`, ExpectedOutput: "true", IsHidden: false},
				{Input: `[1, 3, 2, 4]`, ExpectedOutput: "false", IsHidden: false},
				{Input: `[1]`, ExpectedOutput: "true", IsHidden: false},
				{Input: `[5, 5, 5]`, ExpectedOutput: "true", IsHidden: true},
				{Input: `[3, 2, 1]`, ExpectedOutput: "false", IsHidden: true},
			},
		},
		{
			Title:      "Gabung Array",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan dua array bilangan bulat yang sudah terurut, gabungkan keduanya menjadi satu array yang juga terurut secara ascending.

Format Input: Dua array bilangan bulat terurut, dipisahkan koma: arr1, arr2
Format Output: Array gabungan yang terurut

Contoh:
- gabungArray([1, 3, 5], [2, 4, 6]) → [1, 2, 3, 4, 5, 6]
- gabungArray([], [1, 2]) → [1, 2]
- gabungArray([1], [1]) → [1, 1]`,
			StarterCode: `function gabungArray(arr1, arr2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 3, 5], [2, 4, 6]`, ExpectedOutput: "[1,2,3,4,5,6]", IsHidden: false},
				{Input: `[], [1, 2]`, ExpectedOutput: "[1,2]", IsHidden: false},
				{Input: `[1], [1]`, ExpectedOutput: "[1,1]", IsHidden: false},
				{Input: `[1, 2, 3], []`, ExpectedOutput: "[1,2,3]", IsHidden: true},
				{Input: `[-2, 0], [-1, 1]`, ExpectedOutput: "[-2,-1,0,1]", IsHidden: true},
			},
		},
		{
			Title:      "Hitung Kemunculan",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat dan sebuah nilai, kembalikan berapa kali nilai tersebut muncul dalam array.

Format Input: Array bilangan bulat dan nilai yang dicari, dipisahkan koma: arr, nilai
Format Output: Bilangan bulat (jumlah kemunculan)

Contoh:
- hitungKemunculan([1, 2, 2, 3, 2], 2) → 3
- hitungKemunculan([1, 2, 3], 5) → 0
- hitungKemunculan([], 1) → 0`,
			StarterCode: `function hitungKemunculan(arr, nilai) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 2, 3, 2], 2`, ExpectedOutput: "3", IsHidden: false},
				{Input: `[1, 2, 3], 5`, ExpectedOutput: "0", IsHidden: false},
				{Input: `[], 1`, ExpectedOutput: "0", IsHidden: false},
				{Input: `[5, 5, 5, 5], 5`, ExpectedOutput: "4", IsHidden: true},
				{Input: `[-1, -1, 0, 1], -1`, ExpectedOutput: "2", IsHidden: true},
			},
		},
		{
			Title:      "Array Positif",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan array baru yang hanya berisi bilangan positif (lebih besar dari 0).

Format Input: Array bilangan bulat
Format Output: Array yang hanya berisi bilangan positif

Contoh:
- arrayPositif([1, -2, 3, -4, 5]) → [1, 3, 5]
- arrayPositif([-1, -2, -3]) → []
- arrayPositif([]) → []`,
			StarterCode: `function arrayPositif(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, -2, 3, -4, 5]`, ExpectedOutput: "[1,3,5]", IsHidden: false},
				{Input: `[-1, -2, -3]`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[0, 1, 2]`, ExpectedOutput: "[1,2]", IsHidden: true},
				{Input: `[10, -5, 0, 3]`, ExpectedOutput: "[10,3]", IsHidden: true},
			},
		},
		{
			Title:      "Indeks Elemen",
			Category:   "array",
			Difficulty: "easy",
			Description: `Diberikan sebuah array bilangan bulat dan sebuah nilai, kembalikan indeks pertama dari nilai tersebut dalam array. Jika tidak ditemukan, kembalikan -1.

Format Input: Array bilangan bulat dan nilai yang dicari, dipisahkan koma: arr, nilai
Format Output: Bilangan bulat (indeks pertama atau -1)

Contoh:
- indeksElemen([10, 20, 30, 20], 20) → 1
- indeksElemen([1, 2, 3], 5) → -1
- indeksElemen([], 1) → -1`,
			StarterCode: `function indeksElemen(arr, nilai) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[10, 20, 30, 20], 20`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[1, 2, 3], 5`, ExpectedOutput: "-1", IsHidden: false},
				{Input: `[], 1`, ExpectedOutput: "-1", IsHidden: false},
				{Input: `[5, 5, 5], 5`, ExpectedOutput: "0", IsHidden: true},
				{Input: `[1, 2, 3, 4, 5], 5`, ExpectedOutput: "4", IsHidden: true},
			},
		},
		// ===== ARRAY MEDIUM PROBLEMS =====
		{
			Title:      "Rotasi Array",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat dan bilangan K, rotasikan array ke kanan sebanyak K posisi. Elemen yang keluar dari ujung kanan akan masuk kembali dari ujung kiri.

Format Input: Array bilangan bulat dan K, dipisahkan koma: arr, K
Format Output: Array setelah dirotasi

Contoh:
- rotasiArray([1, 2, 3, 4, 5], 2) → [4, 5, 1, 2, 3]
- rotasiArray([1, 2, 3], 1) → [3, 1, 2]
- rotasiArray([1], 5) → [1]`,
			StarterCode: `function rotasiArray(arr, k) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5], 2`, ExpectedOutput: "[4,5,1,2,3]", IsHidden: false},
				{Input: `[1, 2, 3], 1`, ExpectedOutput: "[3,1,2]", IsHidden: false},
				{Input: `[1], 5`, ExpectedOutput: "[1]", IsHidden: false},
				{Input: `[1, 2, 3, 4], 4`, ExpectedOutput: "[1,2,3,4]", IsHidden: true},
				{Input: `[1, 2, 3], 0`, ExpectedOutput: "[1,2,3]", IsHidden: true},
			},
		},
		{
			Title:      "Two Sum",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat dan sebuah target, kembalikan indeks dari dua bilangan yang jika dijumlahkan menghasilkan target. Asumsikan selalu ada tepat satu solusi.

Format Input: Array bilangan bulat dan target, dipisahkan koma: arr, target
Format Output: Array berisi dua indeks (indeks lebih kecil duluan)

Contoh:
- twoSum([2, 7, 11, 15], 9) → [0, 1]
- twoSum([3, 2, 4], 6) → [1, 2]
- twoSum([3, 3], 6) → [0, 1]`,
			StarterCode: `function twoSum(arr, target) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[2, 7, 11, 15], 9`, ExpectedOutput: "[0,1]", IsHidden: false},
				{Input: `[3, 2, 4], 6`, ExpectedOutput: "[1,2]", IsHidden: false},
				{Input: `[3, 3], 6`, ExpectedOutput: "[0,1]", IsHidden: false},
				{Input: `[1, 5, 3, 2], 4`, ExpectedOutput: "[2,3]", IsHidden: true},
				{Input: `[-1, -2, -3, -4], -7`, ExpectedOutput: "[2,3]", IsHidden: true},
			},
		},
		{
			Title:      "Subarray Terpanjang",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat, temukan panjang subarray terpanjang yang tidak mengandung elemen duplikat.

Format Input: Array bilangan bulat
Format Output: Bilangan bulat (panjang subarray terpanjang)

Contoh:
- subarrayTerpanjang([2, 1, 5, 2, 3, 2]) → 4
- subarrayTerpanjang([1, 2, 3, 4]) → 4
- subarrayTerpanjang([1, 1, 1]) → 1`,
			StarterCode: `function subarrayTerpanjang(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[2, 1, 5, 2, 3, 2]`, ExpectedOutput: "4", IsHidden: false},
				{Input: `[1, 2, 3, 4]`, ExpectedOutput: "4", IsHidden: false},
				{Input: `[1, 1, 1]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "0", IsHidden: true},
				{Input: `[1, 2, 1, 3, 2, 4]`, ExpectedOutput: "4", IsHidden: true},
			},
		},
		{
			Title:      "Rata-rata Bergerak",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat dan ukuran jendela K, kembalikan array rata-rata bergerak (moving average) dengan jendela berukuran K. Setiap elemen hasil adalah rata-rata dari K elemen berturut-turut. Hasil dibulatkan hingga 2 angka desimal.

Format Input: Array bilangan bulat dan K, dipisahkan koma: arr, K
Format Output: Array rata-rata bergerak

Contoh:
- rataRataBergerak([1, 2, 3, 4, 5], 3) → [2.00, 3.00, 4.00]
- rataRataBergerak([1, 2], 2) → [1.50]
- rataRataBergerak([5, 5, 5, 5], 2) → [5.00, 5.00, 5.00]`,
			StarterCode: `function rataRataBergerak(arr, k) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5], 3`, ExpectedOutput: "[2.00,3.00,4.00]", IsHidden: false},
				{Input: `[1, 2], 2`, ExpectedOutput: "[1.50]", IsHidden: false},
				{Input: `[5, 5, 5, 5], 2`, ExpectedOutput: "[5.00,5.00,5.00]", IsHidden: false},
				{Input: `[10, 20, 30, 40], 2`, ExpectedOutput: "[15.00,25.00,35.00]", IsHidden: true},
				{Input: `[1, 3, 5, 7, 9], 4`, ExpectedOutput: "[4.00,6.00]", IsHidden: true},
			},
		},
		{
			Title:      "Flatten Array",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array yang mungkin berisi array bersarang satu level, ratakan menjadi satu array tunggal.

Format Input: Array yang mungkin berisi sub-array satu level
Format Output: Array yang sudah diratakan

Contoh:
- flattenArray([1, [2, 3], [4, 5], 6]) → [1, 2, 3, 4, 5, 6]
- flattenArray([[1, 2], [3, 4]]) → [1, 2, 3, 4]
- flattenArray([1, 2, 3]) → [1, 2, 3]`,
			StarterCode: `function flattenArray(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, [2, 3], [4, 5], 6]`, ExpectedOutput: "[1,2,3,4,5,6]", IsHidden: false},
				{Input: `[[1, 2], [3, 4]]`, ExpectedOutput: "[1,2,3,4]", IsHidden: false},
				{Input: `[1, 2, 3]`, ExpectedOutput: "[1,2,3]", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "[]", IsHidden: true},
				{Input: `[[1], [2], [3]]`, ExpectedOutput: "[1,2,3]", IsHidden: true},
			},
		},
		{
			Title:      "Irisan Array",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan dua array bilangan bulat, kembalikan array yang berisi elemen-elemen yang ada di kedua array (irisan/intersection). Hasil hanya berisi elemen unik dan diurutkan secara ascending.

Format Input: Dua array bilangan bulat, dipisahkan koma: arr1, arr2
Format Output: Array irisan yang terurut

Contoh:
- irisanArray([1, 2, 3, 4], [3, 4, 5, 6]) → [3, 4]
- irisanArray([1, 2, 3], [4, 5, 6]) → []
- irisanArray([1, 1, 2], [1, 2, 2]) → [1, 2]`,
			StarterCode: `function irisanArray(arr1, arr2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4], [3, 4, 5, 6]`, ExpectedOutput: "[3,4]", IsHidden: false},
				{Input: `[1, 2, 3], [4, 5, 6]`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[1, 1, 2], [1, 2, 2]`, ExpectedOutput: "[1,2]", IsHidden: false},
				{Input: `[], [1, 2, 3]`, ExpectedOutput: "[]", IsHidden: true},
				{Input: `[5, 3, 1], [1, 3, 5]`, ExpectedOutput: "[1,3,5]", IsHidden: true},
			},
		},
		{
			Title:      "Gabungan Array",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan dua array bilangan bulat, kembalikan array yang berisi semua elemen unik dari kedua array (gabungan/union), diurutkan secara ascending.

Format Input: Dua array bilangan bulat, dipisahkan koma: arr1, arr2
Format Output: Array gabungan yang terurut

Contoh:
- gabunganArray([1, 2, 3], [3, 4, 5]) → [1, 2, 3, 4, 5]
- gabunganArray([1, 1, 2], [2, 3]) → [1, 2, 3]
- gabunganArray([], [1, 2]) → [1, 2]`,
			StarterCode: `function gabunganArray(arr1, arr2) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3], [3, 4, 5]`, ExpectedOutput: "[1,2,3,4,5]", IsHidden: false},
				{Input: `[1, 1, 2], [2, 3]`, ExpectedOutput: "[1,2,3]", IsHidden: false},
				{Input: `[], [1, 2]`, ExpectedOutput: "[1,2]", IsHidden: false},
				{Input: `[5, 3, 1], [2, 4, 6]`, ExpectedOutput: "[1,2,3,4,5,6]", IsHidden: true},
				{Input: `[1, 2], [1, 2]`, ExpectedOutput: "[1,2]", IsHidden: true},
			},
		},
		{
			Title:      "Cari Pasangan",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat dan sebuah target, kembalikan semua pasangan unik yang jika dijumlahkan menghasilkan target. Setiap pasangan diurutkan secara ascending, dan hasil akhir diurutkan berdasarkan elemen pertama pasangan.

Format Input: Array bilangan bulat dan target, dipisahkan koma: arr, target
Format Output: Array of array berisi pasangan-pasangan

Contoh:
- cariPasangan([1, 2, 3, 4, 5], 6) → [[1,5],[2,4]]
- cariPasangan([1, 1, 2, 3], 4) → [[1,3]]
- cariPasangan([1, 2, 3], 10) → []`,
			StarterCode: `function cariPasangan(arr, target) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5], 6`, ExpectedOutput: "[[1,5],[2,4]]", IsHidden: false},
				{Input: `[1, 1, 2, 3], 4`, ExpectedOutput: "[[1,3]]", IsHidden: false},
				{Input: `[1, 2, 3], 10`, ExpectedOutput: "[]", IsHidden: false},
				{Input: `[-1, 0, 1, 2], 1`, ExpectedOutput: "[[-1,2],[0,1]]", IsHidden: true},
				{Input: `[2, 4, 6, 8], 10`, ExpectedOutput: "[[2,8],[4,6]]", IsHidden: true},
			},
		},
		{
			Title:      "Rotasi Kiri",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array bilangan bulat dan bilangan K, rotasikan array ke kiri sebanyak K posisi. Elemen yang keluar dari ujung kiri akan masuk kembali dari ujung kanan.

Format Input: Array bilangan bulat dan K, dipisahkan koma: arr, K
Format Output: Array setelah dirotasi ke kiri

Contoh:
- rotasiKiri([1, 2, 3, 4, 5], 2) → [3, 4, 5, 1, 2]
- rotasiKiri([1, 2, 3], 1) → [2, 3, 1]
- rotasiKiri([1], 5) → [1]`,
			StarterCode: `function rotasiKiri(arr, k) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5], 2`, ExpectedOutput: "[3,4,5,1,2]", IsHidden: false},
				{Input: `[1, 2, 3], 1`, ExpectedOutput: "[2,3,1]", IsHidden: false},
				{Input: `[1], 5`, ExpectedOutput: "[1]", IsHidden: false},
				{Input: `[1, 2, 3, 4], 4`, ExpectedOutput: "[1,2,3,4]", IsHidden: true},
				{Input: `[1, 2, 3], 0`, ExpectedOutput: "[1,2,3]", IsHidden: true},
			},
		},
		{
			Title:      "Chunk Array",
			Category:   "array",
			Difficulty: "medium",
			Description: `Diberikan sebuah array dan ukuran chunk N, bagi array menjadi sub-array dengan ukuran N. Sub-array terakhir boleh memiliki elemen lebih sedikit jika array tidak habis dibagi.

Format Input: Array bilangan bulat dan N, dipisahkan koma: arr, N
Format Output: Array of array (array yang sudah dibagi)

Contoh:
- chunkArray([1, 2, 3, 4, 5], 2) → [[1,2],[3,4],[5]]
- chunkArray([1, 2, 3, 4], 2) → [[1,2],[3,4]]
- chunkArray([1, 2, 3], 5) → [[1,2,3]]`,
			StarterCode: `function chunkArray(arr, n) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4, 5], 2`, ExpectedOutput: "[[1,2],[3,4],[5]]", IsHidden: false},
				{Input: `[1, 2, 3, 4], 2`, ExpectedOutput: "[[1,2],[3,4]]", IsHidden: false},
				{Input: `[1, 2, 3], 5`, ExpectedOutput: "[[1,2,3]]", IsHidden: false},
				{Input: `[], 3`, ExpectedOutput: "[]", IsHidden: true},
				{Input: `[1, 2, 3, 4, 5, 6], 3`, ExpectedOutput: "[[1,2,3],[4,5,6]]", IsHidden: true},
			},
		},
		// ===== ARRAY HARD PROBLEMS =====
		{
			Title:      "Subarray Terpanjang Tanpa Duplikat",
			Category:   "array",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat, temukan panjang subarray terpanjang yang tidak mengandung elemen duplikat. Gunakan pendekatan sliding window.

Format Input: Array bilangan bulat
Format Output: Bilangan bulat (panjang subarray terpanjang)

Contoh:
- subarrayTerpanjang([1, 2, 3, 1, 2, 3]) → 3
- subarrayTerpanjang([1, 2, 3, 4, 5]) → 5
- subarrayTerpanjang([1, 1, 1, 1]) → 1
- subarrayTerpanjang([]) → 0`,
			StarterCode: `function subarrayTerpanjang(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 1, 2, 3]`, ExpectedOutput: "3", IsHidden: false},
				{Input: `[1, 2, 3, 4, 5]`, ExpectedOutput: "5", IsHidden: false},
				{Input: `[1, 1, 1, 1]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[]`, ExpectedOutput: "0", IsHidden: false},
				{Input: `[1, 2, 1, 3, 2, 4]`, ExpectedOutput: "4", IsHidden: true},
				{Input: `[1]`, ExpectedOutput: "1", IsHidden: true},
				{Input: `[1, 2, 3, 2, 1]`, ExpectedOutput: "3", IsHidden: true},
			},
		},
		{
			Title:      "Median Array",
			Category:   "array",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat yang belum terurut, temukan nilai median. Jika jumlah elemen genap, median adalah rata-rata dua elemen tengah. Hasil dibulatkan hingga 1 angka desimal.

Format Input: Array bilangan bulat
Format Output: Bilangan desimal (median)

Contoh:
- medianArray([3, 1, 4, 1, 5]) → 3.0
- medianArray([1, 2, 3, 4]) → 2.5
- medianArray([7]) → 7.0
- medianArray([5, 3, 1, 2, 4]) → 3.0`,
			StarterCode: `function medianArray(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[3, 1, 4, 1, 5]`, ExpectedOutput: "3.0", IsHidden: false},
				{Input: `[1, 2, 3, 4]`, ExpectedOutput: "2.5", IsHidden: false},
				{Input: `[7]`, ExpectedOutput: "7.0", IsHidden: false},
				{Input: `[5, 3, 1, 2, 4]`, ExpectedOutput: "3.0", IsHidden: false},
				{Input: `[1, 2]`, ExpectedOutput: "1.5", IsHidden: true},
				{Input: `[-3, -1, -2]`, ExpectedOutput: "-2.0", IsHidden: true},
				{Input: `[10, 20, 30, 40, 50, 60]`, ExpectedOutput: "35.0", IsHidden: true},
			},
		},
		{
			Title:      "Produk Kecuali Diri",
			Category:   "array",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat, kembalikan array baru di mana setiap elemen adalah hasil perkalian semua elemen lain kecuali dirinya sendiri. Tidak boleh menggunakan operasi pembagian.

Format Input: Array bilangan bulat
Format Output: Array hasil perkalian

Contoh:
- produkKecualiDiri([1, 2, 3, 4]) → [24, 12, 8, 6]
- produkKecualiDiri([2, 3, 4]) → [12, 8, 6]
- produkKecualiDiri([1, 1, 1, 1]) → [1, 1, 1, 1]
- produkKecualiDiri([-1, 1, 0, -3, 3]) → [0, 0, 9, 0, 0]`,
			StarterCode: `function produkKecualiDiri(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[1, 2, 3, 4]`, ExpectedOutput: "[24,12,8,6]", IsHidden: false},
				{Input: `[2, 3, 4]`, ExpectedOutput: "[12,8,6]", IsHidden: false},
				{Input: `[1, 1, 1, 1]`, ExpectedOutput: "[1,1,1,1]", IsHidden: false},
				{Input: `[-1, 1, 0, -3, 3]`, ExpectedOutput: "[0,0,9,0,0]", IsHidden: false},
				{Input: `[2, 2, 2, 2]`, ExpectedOutput: "[8,8,8,8]", IsHidden: true},
				{Input: `[1, 2]`, ExpectedOutput: "[2,1]", IsHidden: true},
				{Input: `[0, 0]`, ExpectedOutput: "[0,0]", IsHidden: true},
			},
		},
		{
			Title:      "Subarray Maksimum",
			Category:   "array",
			Difficulty: "hard",
			Description: `Diberikan sebuah array bilangan bulat, temukan subarray yang memiliki jumlah (sum) terbesar dan kembalikan jumlah tersebut. Gunakan algoritma Kadane.

Format Input: Array bilangan bulat
Format Output: Bilangan bulat (jumlah subarray maksimum)

Contoh:
- subarrayMaksimum([-2, 1, -3, 4, -1, 2, 1, -5, 4]) → 6
- subarrayMaksimum([1]) → 1
- subarrayMaksimum([-1, -2, -3]) → -1
- subarrayMaksimum([5, 4, -1, 7, 8]) → 23`,
			StarterCode: `function subarrayMaksimum(arr) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[-2, 1, -3, 4, -1, 2, 1, -5, 4]`, ExpectedOutput: "6", IsHidden: false},
				{Input: `[1]`, ExpectedOutput: "1", IsHidden: false},
				{Input: `[-1, -2, -3]`, ExpectedOutput: "-1", IsHidden: false},
				{Input: `[5, 4, -1, 7, 8]`, ExpectedOutput: "23", IsHidden: false},
				{Input: `[1, 2, 3, 4, 5]`, ExpectedOutput: "15", IsHidden: true},
				{Input: `[-2, -1]`, ExpectedOutput: "-1", IsHidden: true},
				{Input: `[2, -1, 2, 3, 4, -5]`, ExpectedOutput: "10", IsHidden: true},
			},
		},
		{
			Title:      "Merge Interval",
			Category:   "array",
			Difficulty: "hard",
			Description: `Diberikan sebuah array interval [start, end], gabungkan semua interval yang saling tumpang tindih (overlapping) dan kembalikan array interval yang sudah digabungkan, diurutkan berdasarkan start.

Format Input: Array of interval [start, end]
Format Output: Array of interval setelah digabungkan

Contoh:
- mergeInterval([[1,3],[2,6],[8,10],[15,18]]) → [[1,6],[8,10],[15,18]]
- mergeInterval([[1,4],[4,5]]) → [[1,5]]
- mergeInterval([[1,4],[2,3]]) → [[1,4]]
- mergeInterval([[1,2],[3,4]]) → [[1,2],[3,4]]`,
			StarterCode: `function mergeInterval(intervals) {
  // Tulis kode kamu di sini
}`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: `[[1,3],[2,6],[8,10],[15,18]]`, ExpectedOutput: "[[1,6],[8,10],[15,18]]", IsHidden: false},
				{Input: `[[1,4],[4,5]]`, ExpectedOutput: "[[1,5]]", IsHidden: false},
				{Input: `[[1,4],[2,3]]`, ExpectedOutput: "[[1,4]]", IsHidden: false},
				{Input: `[[1,2],[3,4]]`, ExpectedOutput: "[[1,2],[3,4]]", IsHidden: false},
				{Input: `[[1,10],[2,3],[4,5]]`, ExpectedOutput: "[[1,10]]", IsHidden: true},
				{Input: `[[1,2]]`, ExpectedOutput: "[[1,2]]", IsHidden: true},
				{Input: `[[5,10],[1,3],[2,6]]`, ExpectedOutput: "[[1,10]]", IsHidden: true},
			},
		},
	}
	return append(jsProblems, GetSQLSeedProblems()...)
}

func Seed(db *gorm.DB) {
	problems := GetSeedProblems()

	inserted := 0
	for i := range problems {
		var existing models.Problem
		result := db.Where("title = ?", problems[i].Title).First(&existing)
		if result.Error == nil {
			// soal dengan judul ini sudah ada, skip
			continue
		}
		if err := db.Create(&problems[i]).Error; err != nil {
			log.Printf("Failed to seed problem '%s': %v", problems[i].Title, err)
			return
		}
		inserted++
	}

	if inserted > 0 {
		log.Printf("Seeded %d new problems", inserted)
	} else {
		log.Println("No new problems to seed")
	}
}
