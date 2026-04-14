package evaluator

import (
	"testing"
)

// Comprehensive test covering all problem types and input formats
func TestAllProblemTypesComprehensive(t *testing.T) {
	svc := NewEvaluatorService()

	tests := []struct {
		name     string
		code     string
		input    string
		expected string
		wantPass bool
	}{
		// Loop problems - single number input
		{
			name:     "jumlahSampaiN(5)",
			code:     "function jumlahSampaiN(n) { let total = 0; for (let i = 1; i <= n; i++) { total += i; } return total; }",
			input:    "5",
			expected: "15",
			wantPass: true,
		},
		{
			name:     "jumlahSampaiN(0)",
			code:     "function jumlahSampaiN(n) { let total = 0; for (let i = 1; i <= n; i++) { total += i; } return total; }",
			input:    "0",
			expected: "0",
			wantPass: true,
		},
		// Loop problems - array input
		{
			name:     "hitungMundur(5)",
			code:     "function hitungMundur(n) { let hasil = []; for (let i = n; i >= 1; i--) { hasil.push(i); } return hasil; }",
			input:    "5",
			expected: "[5,4,3,2,1]",
			wantPass: true,
		},
		// Loop problems - comma-separated input
		{
			name:     "kaliTanpaBintang(3, 4)",
			code:     "function kaliTanpaBintang(a, b) { let hasil = 0; for (let i = 0; i < b; i++) { hasil += a; } return hasil; }",
			input:    "3, 4",
			expected: "12",
			wantPass: true,
		},
		{
			name:     "kaliTanpaBintang(5, 2)",
			code:     "function kaliTanpaBintang(a, b) { let hasil = 0; for (let i = 0; i < b; i++) { hasil += a; } return hasil; }",
			input:    "5, 2",
			expected: "10",
			wantPass: true,
		},
		{
			name:     "kaliTanpaBintang(0, 7)",
			code:     "function kaliTanpaBintang(a, b) { let hasil = 0; for (let i = 0; i < b; i++) { hasil += a; } return hasil; }",
			input:    "0, 7",
			expected: "0",
			wantPass: true,
		},
		// Array problems
		{
			name:     "nilaiMaksimum([3,1,4,1,5,9,2])",
			code:     "function nilaiMaksimum(arr) { let maks = arr[0]; for (let i = 1; i < arr.length; i++) { if (arr[i] > maks) { maks = arr[i]; } } return maks; }",
			input:    "[3,1,4,1,5,9,2]",
			expected: "9",
			wantPass: true,
		},
		{
			name:     "nilaiMaksimum([10,5,8])",
			code:     "function nilaiMaksimum(arr) { let maks = arr[0]; for (let i = 1; i < arr.length; i++) { if (arr[i] > maks) { maks = arr[i]; } } return maks; }",
			input:    "[10,5,8]",
			expected: "10",
			wantPass: true,
		},
		{
			name:     "nilaiMaksimum([7])",
			code:     "function nilaiMaksimum(arr) { let maks = arr[0]; for (let i = 1; i < arr.length; i++) { if (arr[i] > maks) { maks = arr[i]; } } return maks; }",
			input:    "[7]",
			expected: "7",
			wantPass: true,
		},
		{
			name:     "hitungGenap([1,2,3,4,5,6])",
			code:     "function hitungGenap(arr) { let count = 0; for (let i = 0; i < arr.length; i++) { if (arr[i] % 2 === 0) { count++; } } return count; }",
			input:    "[1,2,3,4,5,6]",
			expected: "3",
			wantPass: true,
		},
		{
			name:     "hitungGenap([1,3,5])",
			code:     "function hitungGenap(arr) { let count = 0; for (let i = 0; i < arr.length; i++) { if (arr[i] % 2 === 0) { count++; } } return count; }",
			input:    "[1,3,5]",
			expected: "0",
			wantPass: true,
		},
		// FizzBuzz
		{
			name:     "fizzBuzz(5)",
			code:     "function fizzBuzz(n) { let hasil = []; for (let i = 1; i <= n; i++) { if (i % 15 === 0) { hasil.push('FizzBuzz'); } else if (i % 3 === 0) { hasil.push('Fizz'); } else if (i % 5 === 0) { hasil.push('Buzz'); } else { hasil.push(String(i)); } } return hasil; }",
			input:    "5",
			expected: `["1","2","Fizz","4","Buzz"]`,
			wantPass: true,
		},
		// While loop
		{
			name:     "hitungBagi2(8)",
			code:     "function hitungBagi2(n) { let langkah = 0; while (n > 0) { n = Math.floor(n / 2); langkah++; } return langkah; }",
			input:    "8",
			expected: "4",
			wantPass: true,
		},
		// Prime check
		{
			name:     "isPrima(7)",
			code:     "function isPrima(n) { if (n < 2) return false; for (let i = 2; i * i <= n; i++) { if (n % i === 0) return false; } return true; }",
			input:    "7",
			expected: "true",
			wantPass: true,
		},
		{
			name:     "isPrima(4)",
			code:     "function isPrima(n) { if (n < 2) return false; for (let i = 2; i * i <= n; i++) { if (n % i === 0) return false; } return true; }",
			input:    "4",
			expected: "false",
			wantPass: true,
		},
		// Fibonacci
		{
			name:     "fibonacci(5)",
			code:     "function fibonacci(n) { if (n <= 0) return []; let hasil = [0]; if (n >= 2) hasil.push(1); for (let i = 2; i < n; i++) { hasil.push(hasil[i-1] + hasil[i-2]); } return hasil; }",
			input:    "5",
			expected: "[0,1,1,2,3]",
			wantPass: true,
		},
		// Multiplication table
		{
			name:     "tabelPerkalian(2)",
			code:     "function tabelPerkalian(n) { let tabel = []; for (let i = 0; i < n; i++) { let baris = []; for (let j = 0; j < n; j++) { baris.push((i+1)*(j+1)); } tabel.push(baris); } return tabel; }",
			input:    "2",
			expected: "[[1,2],[2,4]]",
			wantPass: true,
		},
		// Power
		{
			name:     "pangkat(2, 10)",
			code:     "function pangkat(basis, eksponen) { let hasil = 1; for (let i = 0; i < eksponen; i++) { hasil *= basis; } return hasil; }",
			input:    "2, 10",
			expected: "1024",
			wantPass: true,
		},
		// Factors
		{
			name:     "cariFaktor(12)",
			code:     "function cariFaktor(n) { let faktor = []; for (let i = 1; i <= n; i++) { if (n % i === 0) { faktor.push(i); } } return faktor; }",
			input:    "12",
			expected: "[1,2,3,4,6,12]",
			wantPass: true,
		},
		// Triangle pattern
		{
			name:     "segitigaBintang(3)",
			code:     "function segitigaBintang(n) { let pola = []; for (let i = 1; i <= n; i++) { pola.push('*'.repeat(i)); } return pola; }",
			input:    "3",
			expected: `["*","**","***"]`,
			wantPass: true,
		},
		// Conditional sum
		{
			name:     "jumlahDiAtas([1,5,3,8,2,9], 4)",
			code:     "function jumlahDiAtas(arr, threshold) { let total = 0; for (let i = 0; i < arr.length; i++) { if (arr[i] > threshold) { total += arr[i]; } } return total; }",
			input:    "[1,5,3,8,2,9], 4",
			expected: "22",
			wantPass: true,
		},
		// Bubble sort
		{
			name:     "bubbleSort([5,1,4,2,8])",
			code:     "function bubbleSort(arr) { let hasil = [...arr]; for (let i = 0; i < hasil.length - 1; i++) { for (let j = 0; j < hasil.length - 1 - i; j++) { if (hasil[j] > hasil[j+1]) { let temp = hasil[j]; hasil[j] = hasil[j+1]; hasil[j+1] = temp; } } } return hasil; }",
			input:    "[5,1,4,2,8]",
			expected: "[1,2,4,5,8]",
			wantPass: true,
		},
		// Sieve of Eratosthenes
		{
			name:     "sieveEratosthenes(10)",
			code:     "function sieveEratosthenes(n) { let isPrima = new Array(n + 1).fill(true); isPrima[0] = false; isPrima[1] = false; for (let i = 2; i * i <= n; i++) { if (isPrima[i]) { for (let j = i * i; j <= n; j += i) { isPrima[j] = false; } } } let hasil = []; for (let i = 2; i <= n; i++) { if (isPrima[i]) { hasil.push(i); } } return hasil; }",
			input:    "10",
			expected: "[2,3,5,7]",
			wantPass: true,
		},
		// Base conversion
		{
			name:     "konversiBasis(10, 2)",
			code:     "function konversiBasis(n, basis) { if (n === 0) return '0'; const digit = '0123456789abcdef'; let hasil = ''; while (n > 0) { hasil = digit[n % basis] + hasil; n = Math.floor(n / basis); } return hasil; }",
			input:    "10, 2",
			expected: `"1010"`,
			wantPass: true,
		},
		// Digit sum
		{
			name:     "jumlahDigitBerulang(9875)",
			code:     "function jumlahDigitBerulang(n) { let langkah = 0; while (n >= 10) { n = String(n).split('').reduce((sum, d) => sum + Number(d), 0); langkah++; } return { langkah, nilai: n }; }",
			input:    "9875",
			expected: `{"langkah":3,"nilai":2}`,
			wantPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.Evaluate(tt.code, tt.input, tt.expected)
			if result.Passed != tt.wantPass {
				t.Errorf("expected Passed=%v, got %v\n  actual=%q\n  error=%q", tt.wantPass, result.Passed, result.Actual, result.Error)
			}
		})
	}
}
