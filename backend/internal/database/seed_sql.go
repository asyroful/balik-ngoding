package database

import "balik-ngoding-backend/internal/models"

// GetSQLSeedProblems returns 25 SQL seed problems (10 easy, 10 medium, 5 hard).
func GetSQLSeedProblems() []models.Problem {
	return []models.Problem{
		// ===== EASY (10 soal) =====
		{
			Title:      "Pilih Semua Pengguna",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel users dengan kolom id, name, dan age, tulis query untuk mengambil semua data pengguna.

Skema tabel:
  users(id INTEGER, name TEXT, age INTEGER)

Contoh data:
  (1, 'Alice', 30), (2, 'Bob', 25), (3, 'Charlie', 35)

Expected output: semua baris dari tabel users.`,
			StarterCode: "",
			Schema: `CREATE TABLE users (id INTEGER, name TEXT, age INTEGER);
INSERT INTO users VALUES (1, 'Alice', 30);
INSERT INTO users VALUES (2, 'Bob', 25);
INSERT INTO users VALUES (3, 'Charlie', 35);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":2,"name":"Bob","age":25},{"id":3,"name":"Charlie","age":35}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":2,"name":"Bob","age":25},{"id":3,"name":"Charlie","age":35}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":2,"name":"Bob","age":25},{"id":3,"name":"Charlie","age":35}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":2,"name":"Bob","age":25},{"id":3,"name":"Charlie","age":35}]`, IsHidden: true},
			},
		},
		{
			Title:      "Filter Usia",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel users, tulis query untuk mengambil semua pengguna yang berusia lebih dari 25 tahun.

Skema tabel:
  users(id INTEGER, name TEXT, age INTEGER)

Contoh data:
  (1, 'Alice', 30), (2, 'Bob', 25), (3, 'Charlie', 35), (4, 'Diana', 22)

Expected output: pengguna dengan age > 25.`,
			StarterCode: "",
			Schema: `CREATE TABLE users (id INTEGER, name TEXT, age INTEGER);
INSERT INTO users VALUES (1, 'Alice', 30);
INSERT INTO users VALUES (2, 'Bob', 25);
INSERT INTO users VALUES (3, 'Charlie', 35);
INSERT INTO users VALUES (4, 'Diana', 22);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":3,"name":"Charlie","age":35}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":3,"name":"Charlie","age":35}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":3,"name":"Charlie","age":35}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","age":30},{"id":3,"name":"Charlie","age":35}]`, IsHidden: true},
			},
		},
		{
			Title:      "Urutkan Nama",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel users, tulis query untuk mengambil semua pengguna dan mengurutkannya berdasarkan nama secara ascending (A-Z).

Skema tabel:
  users(id INTEGER, name TEXT, age INTEGER)

Contoh data:
  (1, 'Charlie', 35), (2, 'Alice', 30), (3, 'Bob', 25), (4, 'Diana', 22)

Expected output: semua pengguna diurutkan berdasarkan name ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE users (id INTEGER, name TEXT, age INTEGER);
INSERT INTO users VALUES (1, 'Charlie', 35);
INSERT INTO users VALUES (2, 'Alice', 30);
INSERT INTO users VALUES (3, 'Bob', 25);
INSERT INTO users VALUES (4, 'Diana', 22);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Alice","age":30},{"id":3,"name":"Bob","age":25},{"id":1,"name":"Charlie","age":35},{"id":4,"name":"Diana","age":22}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Alice","age":30},{"id":3,"name":"Bob","age":25},{"id":1,"name":"Charlie","age":35},{"id":4,"name":"Diana","age":22}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Alice","age":30},{"id":3,"name":"Bob","age":25},{"id":1,"name":"Charlie","age":35},{"id":4,"name":"Diana","age":22}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Alice","age":30},{"id":3,"name":"Bob","age":25},{"id":1,"name":"Charlie","age":35},{"id":4,"name":"Diana","age":22}]`, IsHidden: true},
			},
		},
		{
			Title:      "Hitung Baris",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel products, tulis query untuk menghitung jumlah total produk yang ada.

Skema tabel:
  products(id INTEGER, name TEXT, price REAL)

Contoh data:
  5 produk dengan berbagai harga.

Expected output: [{"count":5}]`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, price REAL);
INSERT INTO products VALUES (1, 'Apel', 5000.0);
INSERT INTO products VALUES (2, 'Jeruk', 8000.0);
INSERT INTO products VALUES (3, 'Mangga', 12000.0);
INSERT INTO products VALUES (4, 'Pisang', 3000.0);
INSERT INTO products VALUES (5, 'Anggur', 25000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"count":5}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"count":5}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"count":5}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"count":5}]`, IsHidden: true},
			},
		},
		{
			Title:      "Pilih Kolom Tertentu",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel products, tulis query untuk mengambil hanya kolom name dan price dari semua produk.

Skema tabel:
  products(id INTEGER, name TEXT, price REAL)

Contoh data:
  (1, 'Apel', 5000.0), (2, 'Jeruk', 8000.0), (3, 'Mangga', 12000.0), (4, 'Pisang', 3000.0)

Expected output: hanya kolom name dan price.`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, price REAL);
INSERT INTO products VALUES (1, 'Apel', 5000.0);
INSERT INTO products VALUES (2, 'Jeruk', 8000.0);
INSERT INTO products VALUES (3, 'Mangga', 12000.0);
INSERT INTO products VALUES (4, 'Pisang', 3000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"name":"Apel","price":5000},{"name":"Jeruk","price":8000},{"name":"Mangga","price":12000},{"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Apel","price":5000},{"name":"Jeruk","price":8000},{"name":"Mangga","price":12000},{"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Apel","price":5000},{"name":"Jeruk","price":8000},{"name":"Mangga","price":12000},{"name":"Pisang","price":3000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"name":"Apel","price":5000},{"name":"Jeruk","price":8000},{"name":"Mangga","price":12000},{"name":"Pisang","price":3000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Filter Harga",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel products, tulis query untuk mengambil semua produk dengan harga kurang dari 10000.

Skema tabel:
  products(id INTEGER, name TEXT, price REAL)

Contoh data:
  (1, 'Apel', 5000.0), (2, 'Jeruk', 8000.0), (3, 'Mangga', 12000.0), (4, 'Pisang', 3000.0), (5, 'Anggur', 25000.0)

Expected output: produk dengan price < 10000.`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, price REAL);
INSERT INTO products VALUES (1, 'Apel', 5000.0);
INSERT INTO products VALUES (2, 'Jeruk', 8000.0);
INSERT INTO products VALUES (3, 'Mangga', 12000.0);
INSERT INTO products VALUES (4, 'Pisang', 3000.0);
INSERT INTO products VALUES (5, 'Anggur', 25000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Apel","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Apel","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Apel","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Apel","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Nilai Unik",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel products dengan kolom category, tulis query untuk mendapatkan daftar kategori yang unik (tidak duplikat).

Skema tabel:
  products(id INTEGER, name TEXT, category TEXT)

Contoh data:
  6 produk dengan 3 kategori berbeda: 'Buah', 'Sayur', 'Minuman'

Expected output: daftar kategori unik diurutkan secara ascending.`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, category TEXT);
INSERT INTO products VALUES (1, 'Apel', 'Buah');
INSERT INTO products VALUES (2, 'Jeruk', 'Buah');
INSERT INTO products VALUES (3, 'Bayam', 'Sayur');
INSERT INTO products VALUES (4, 'Kangkung', 'Sayur');
INSERT INTO products VALUES (5, 'Teh', 'Minuman');
INSERT INTO products VALUES (6, 'Kopi', 'Minuman');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"category":"Buah"},{"category":"Minuman"},{"category":"Sayur"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"category":"Buah"},{"category":"Minuman"},{"category":"Sayur"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"category":"Buah"},{"category":"Minuman"},{"category":"Sayur"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"category":"Buah"},{"category":"Minuman"},{"category":"Sayur"}]`, IsHidden: true},
			},
		},
		{
			Title:      "Batas Hasil",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel employees, tulis query untuk mengambil 3 karyawan dengan gaji tertinggi.

Skema tabel:
  employees(id INTEGER, name TEXT, salary REAL)

Contoh data:
  6 karyawan dengan gaji berbeda-beda.

Expected output: 3 karyawan dengan salary tertinggi, diurutkan dari tertinggi.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, salary REAL);
INSERT INTO employees VALUES (1, 'Alice', 8000.0);
INSERT INTO employees VALUES (2, 'Bob', 5000.0);
INSERT INTO employees VALUES (3, 'Charlie', 9500.0);
INSERT INTO employees VALUES (4, 'Diana', 7000.0);
INSERT INTO employees VALUES (5, 'Eve', 11000.0);
INSERT INTO employees VALUES (6, 'Frank', 6500.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9500},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9500},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9500},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9500},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Urutkan Harga Descending",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel products, tulis query untuk mengambil semua produk dan mengurutkannya berdasarkan harga dari yang termahal ke termurah.

Skema tabel:
  products(id INTEGER, name TEXT, price REAL)

Contoh data:
  (1, 'Apel', 5000.0), (2, 'Jeruk', 8000.0), (3, 'Mangga', 12000.0), (4, 'Pisang', 3000.0)

Expected output: semua produk diurutkan berdasarkan price DESC.`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, price REAL);
INSERT INTO products VALUES (1, 'Apel', 5000.0);
INSERT INTO products VALUES (2, 'Jeruk', 8000.0);
INSERT INTO products VALUES (3, 'Mangga', 12000.0);
INSERT INTO products VALUES (4, 'Pisang', 3000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":3,"name":"Mangga","price":12000},{"id":2,"name":"Jeruk","price":8000},{"id":1,"name":"Apel","price":5000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":3,"name":"Mangga","price":12000},{"id":2,"name":"Jeruk","price":8000},{"id":1,"name":"Apel","price":5000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":3,"name":"Mangga","price":12000},{"id":2,"name":"Jeruk","price":8000},{"id":1,"name":"Apel","price":5000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":3,"name":"Mangga","price":12000},{"id":2,"name":"Jeruk","price":8000},{"id":1,"name":"Apel","price":5000},{"id":4,"name":"Pisang","price":3000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Cari Nama",
			Category:   "sql",
			Difficulty: "easy",
			Description: `Diberikan tabel users, tulis query untuk mencari semua pengguna yang namanya diawali dengan huruf 'A'.

Skema tabel:
  users(id INTEGER, name TEXT, city TEXT)

Contoh data:
  5 pengguna dari berbagai kota, beberapa namanya diawali 'A'.

Expected output: pengguna dengan name LIKE 'A%'.`,
			StarterCode: "",
			Schema: `CREATE TABLE users (id INTEGER, name TEXT, city TEXT);
INSERT INTO users VALUES (1, 'Alice', 'Jakarta');
INSERT INTO users VALUES (2, 'Bob', 'Bandung');
INSERT INTO users VALUES (3, 'Andi', 'Surabaya');
INSERT INTO users VALUES (4, 'Charlie', 'Medan');
INSERT INTO users VALUES (5, 'Ayu', 'Yogyakarta');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Andi","city":"Surabaya"},{"id":5,"name":"Ayu","city":"Yogyakarta"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Andi","city":"Surabaya"},{"id":5,"name":"Ayu","city":"Yogyakarta"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Andi","city":"Surabaya"},{"id":5,"name":"Ayu","city":"Yogyakarta"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Andi","city":"Surabaya"},{"id":5,"name":"Ayu","city":"Yogyakarta"}]`, IsHidden: true},
			},
		},

		// ===== MEDIUM (10 soal) =====
		{
			Title:      "Rata-rata Gaji",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel employees, tulis query untuk menghitung rata-rata gaji semua karyawan.

Skema tabel:
  employees(id INTEGER, name TEXT, salary REAL, department TEXT)

Expected output: [{"avg_salary": <nilai rata-rata>}]`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, salary REAL, department TEXT);
INSERT INTO employees VALUES (1, 'Alice', 8000.0, 'Engineering');
INSERT INTO employees VALUES (2, 'Bob', 6000.0, 'Marketing');
INSERT INTO employees VALUES (3, 'Charlie', 9000.0, 'Engineering');
INSERT INTO employees VALUES (4, 'Diana', 7000.0, 'HR');
INSERT INTO employees VALUES (5, 'Eve', 5000.0, 'Marketing');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"avg_salary":7000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"avg_salary":7000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"avg_salary":7000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"avg_salary":7000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Group By Departemen",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel employees, tulis query untuk menghitung jumlah karyawan per departemen.

Skema tabel:
  employees(id INTEGER, name TEXT, salary REAL, department TEXT)

Expected output: jumlah karyawan per department, diurutkan berdasarkan department ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, salary REAL, department TEXT);
INSERT INTO employees VALUES (1, 'Alice', 8000.0, 'Engineering');
INSERT INTO employees VALUES (2, 'Bob', 6000.0, 'Marketing');
INSERT INTO employees VALUES (3, 'Charlie', 9000.0, 'Engineering');
INSERT INTO employees VALUES (4, 'Diana', 7000.0, 'HR');
INSERT INTO employees VALUES (5, 'Eve', 5000.0, 'Marketing');
INSERT INTO employees VALUES (6, 'Frank', 11000.0, 'Engineering');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"department":"Engineering","total":3},{"department":"HR","total":1},{"department":"Marketing","total":2}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","total":3},{"department":"HR","total":1},{"department":"Marketing","total":2}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","total":3},{"department":"HR","total":1},{"department":"Marketing","total":2}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","total":3},{"department":"HR","total":1},{"department":"Marketing","total":2}]`, IsHidden: true},
			},
		},
		{
			Title:      "INNER JOIN Pesanan",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel orders dan customers, tulis query untuk menampilkan nama customer beserta total pesanan mereka.

Skema tabel:
  customers(id INTEGER, name TEXT)
  orders(id INTEGER, customer_id INTEGER, amount REAL)

Expected output: name dan total_amount per customer, diurutkan berdasarkan name ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE customers (id INTEGER, name TEXT);
INSERT INTO customers VALUES (1, 'Alice');
INSERT INTO customers VALUES (2, 'Bob');
INSERT INTO customers VALUES (3, 'Charlie');
CREATE TABLE orders (id INTEGER, customer_id INTEGER, amount REAL);
INSERT INTO orders VALUES (1, 1, 150000.0);
INSERT INTO orders VALUES (2, 1, 75000.0);
INSERT INTO orders VALUES (3, 2, 200000.0);
INSERT INTO orders VALUES (4, 3, 50000.0);
INSERT INTO orders VALUES (5, 2, 100000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"name":"Alice","total_amount":225000},{"name":"Bob","total_amount":300000},{"name":"Charlie","total_amount":50000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Alice","total_amount":225000},{"name":"Bob","total_amount":300000},{"name":"Charlie","total_amount":50000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Alice","total_amount":225000},{"name":"Bob","total_amount":300000},{"name":"Charlie","total_amount":50000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"name":"Alice","total_amount":225000},{"name":"Bob","total_amount":300000},{"name":"Charlie","total_amount":50000}]`, IsHidden: true},
			},
		},
		{
			Title:      "HAVING Filter",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel orders, tulis query untuk menampilkan customer_id yang memiliki total pesanan lebih dari 200000.

Skema tabel:
  orders(id INTEGER, customer_id INTEGER, amount REAL)

Expected output: customer_id dengan total_amount > 200000, diurutkan berdasarkan customer_id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE orders (id INTEGER, customer_id INTEGER, amount REAL);
INSERT INTO orders VALUES (1, 1, 150000.0);
INSERT INTO orders VALUES (2, 1, 75000.0);
INSERT INTO orders VALUES (3, 2, 200000.0);
INSERT INTO orders VALUES (4, 3, 50000.0);
INSERT INTO orders VALUES (5, 2, 100000.0);
INSERT INTO orders VALUES (6, 4, 300000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"customer_id":1,"total_amount":225000},{"customer_id":2,"total_amount":300000},{"customer_id":4,"total_amount":300000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"customer_id":1,"total_amount":225000},{"customer_id":2,"total_amount":300000},{"customer_id":4,"total_amount":300000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"customer_id":1,"total_amount":225000},{"customer_id":2,"total_amount":300000},{"customer_id":4,"total_amount":300000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"customer_id":1,"total_amount":225000},{"customer_id":2,"total_amount":300000},{"customer_id":4,"total_amount":300000}]`, IsHidden: true},
			},
		},
		{
			Title:      "Subquery Gaji Tertinggi",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel employees, tulis query untuk menampilkan karyawan yang gajinya di atas rata-rata gaji semua karyawan.

Skema tabel:
  employees(id INTEGER, name TEXT, salary REAL)

Expected output: karyawan dengan salary > AVG(salary), diurutkan berdasarkan salary DESC.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, salary REAL);
INSERT INTO employees VALUES (1, 'Alice', 8000.0);
INSERT INTO employees VALUES (2, 'Bob', 5000.0);
INSERT INTO employees VALUES (3, 'Charlie', 9000.0);
INSERT INTO employees VALUES (4, 'Diana', 6000.0);
INSERT INTO employees VALUES (5, 'Eve', 11000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9000},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9000},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9000},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":5,"name":"Eve","salary":11000},{"id":3,"name":"Charlie","salary":9000},{"id":1,"name":"Alice","salary":8000}]`, IsHidden: true},
			},
		},
		{
			Title:      "LEFT JOIN Produk",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel products dan categories, tulis query untuk menampilkan semua produk beserta nama kategorinya (termasuk produk tanpa kategori).

Skema tabel:
  products(id INTEGER, name TEXT, category_id INTEGER)
  categories(id INTEGER, name TEXT)

Expected output: product name dan category name (NULL jika tidak ada kategori), diurutkan berdasarkan product id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE categories (id INTEGER, name TEXT);
INSERT INTO categories VALUES (1, 'Elektronik');
INSERT INTO categories VALUES (2, 'Pakaian');
CREATE TABLE products (id INTEGER, name TEXT, category_id INTEGER);
INSERT INTO products VALUES (1, 'Laptop', 1);
INSERT INTO products VALUES (2, 'Kaos', 2);
INSERT INTO products VALUES (3, 'Meja', NULL);
INSERT INTO products VALUES (4, 'Headphone', 1);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"product":"Laptop","category":"Elektronik"},{"product":"Kaos","category":"Pakaian"},{"product":"Meja","category":null},{"product":"Headphone","category":"Elektronik"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"product":"Laptop","category":"Elektronik"},{"product":"Kaos","category":"Pakaian"},{"product":"Meja","category":null},{"product":"Headphone","category":"Elektronik"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"product":"Laptop","category":"Elektronik"},{"product":"Kaos","category":"Pakaian"},{"product":"Meja","category":null},{"product":"Headphone","category":"Elektronik"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"product":"Laptop","category":"Elektronik"},{"product":"Kaos","category":"Pakaian"},{"product":"Meja","category":null},{"product":"Headphone","category":"Elektronik"}]`, IsHidden: true},
			},
		},
		{
			Title:      "Nilai Maksimum Per Grup",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel scores, tulis query untuk menampilkan nilai tertinggi per mata pelajaran.

Skema tabel:
  scores(id INTEGER, student TEXT, subject TEXT, score INTEGER)

Expected output: subject dan max_score per mata pelajaran, diurutkan berdasarkan subject ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE scores (id INTEGER, student TEXT, subject TEXT, score INTEGER);
INSERT INTO scores VALUES (1, 'Alice', 'Math', 90);
INSERT INTO scores VALUES (2, 'Bob', 'Math', 75);
INSERT INTO scores VALUES (3, 'Alice', 'Science', 85);
INSERT INTO scores VALUES (4, 'Bob', 'Science', 92);
INSERT INTO scores VALUES (5, 'Charlie', 'Math', 88);
INSERT INTO scores VALUES (6, 'Charlie', 'Science', 78);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"subject":"Math","max_score":90},{"subject":"Science","max_score":92}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"subject":"Math","max_score":90},{"subject":"Science","max_score":92}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"subject":"Math","max_score":90},{"subject":"Science","max_score":92}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"subject":"Math","max_score":90},{"subject":"Science","max_score":92}]`, IsHidden: true},
			},
		},
		{
			Title:      "Filter NULL",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel employees, tulis query untuk menampilkan karyawan yang belum memiliki manager (manager_id adalah NULL).

Skema tabel:
  employees(id INTEGER, name TEXT, manager_id INTEGER)

Expected output: karyawan dengan manager_id IS NULL, diurutkan berdasarkan id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, manager_id INTEGER);
INSERT INTO employees VALUES (1, 'Alice', NULL);
INSERT INTO employees VALUES (2, 'Bob', 1);
INSERT INTO employees VALUES (3, 'Charlie', 1);
INSERT INTO employees VALUES (4, 'Diana', NULL);
INSERT INTO employees VALUES (5, 'Eve', 2);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice"},{"id":4,"name":"Diana"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice"},{"id":4,"name":"Diana"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice"},{"id":4,"name":"Diana"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice"},{"id":4,"name":"Diana"}]`, IsHidden: true},
			},
		},
		{
			Title:      "BETWEEN Range",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel products, tulis query untuk menampilkan produk dengan harga antara 5000 dan 15000 (inklusif).

Skema tabel:
  products(id INTEGER, name TEXT, price REAL)

Expected output: produk dengan 5000 <= price <= 15000, diurutkan berdasarkan price ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE products (id INTEGER, name TEXT, price REAL);
INSERT INTO products VALUES (1, 'Apel', 3000.0);
INSERT INTO products VALUES (2, 'Jeruk', 8000.0);
INSERT INTO products VALUES (3, 'Mangga', 12000.0);
INSERT INTO products VALUES (4, 'Pisang', 5000.0);
INSERT INTO products VALUES (5, 'Anggur', 25000.0);
INSERT INTO products VALUES (6, 'Semangka', 15000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":4,"name":"Pisang","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":3,"name":"Mangga","price":12000},{"id":6,"name":"Semangka","price":15000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":4,"name":"Pisang","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":3,"name":"Mangga","price":12000},{"id":6,"name":"Semangka","price":15000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":4,"name":"Pisang","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":3,"name":"Mangga","price":12000},{"id":6,"name":"Semangka","price":15000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":4,"name":"Pisang","price":5000},{"id":2,"name":"Jeruk","price":8000},{"id":3,"name":"Mangga","price":12000},{"id":6,"name":"Semangka","price":15000}]`, IsHidden: true},
			},
		},
		{
			Title:      "IN Operator",
			Category:   "sql",
			Difficulty: "medium",
			Description: `Diberikan tabel users, tulis query untuk menampilkan pengguna yang tinggal di Jakarta, Bandung, atau Surabaya.

Skema tabel:
  users(id INTEGER, name TEXT, city TEXT)

Expected output: pengguna dari kota yang ditentukan, diurutkan berdasarkan id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE users (id INTEGER, name TEXT, city TEXT);
INSERT INTO users VALUES (1, 'Alice', 'Jakarta');
INSERT INTO users VALUES (2, 'Bob', 'Medan');
INSERT INTO users VALUES (3, 'Charlie', 'Bandung');
INSERT INTO users VALUES (4, 'Diana', 'Surabaya');
INSERT INTO users VALUES (5, 'Eve', 'Bali');
INSERT INTO users VALUES (6, 'Frank', 'Jakarta');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Charlie","city":"Bandung"},{"id":4,"name":"Diana","city":"Surabaya"},{"id":6,"name":"Frank","city":"Jakarta"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Charlie","city":"Bandung"},{"id":4,"name":"Diana","city":"Surabaya"},{"id":6,"name":"Frank","city":"Jakarta"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Charlie","city":"Bandung"},{"id":4,"name":"Diana","city":"Surabaya"},{"id":6,"name":"Frank","city":"Jakarta"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":1,"name":"Alice","city":"Jakarta"},{"id":3,"name":"Charlie","city":"Bandung"},{"id":4,"name":"Diana","city":"Surabaya"},{"id":6,"name":"Frank","city":"Jakarta"}]`, IsHidden: true},
			},
		},
		// ===== HARD (5 soal) =====
		{
			Title:      "Top 3 Departemen Bergaji Tertinggi",
			Category:   "sql",
			Difficulty: "hard",
			Description: `Diberikan tabel employees, tulis query untuk mendapatkan 3 departemen dengan rata-rata gaji tertinggi. Tampilkan department dan avg_salary (dibulatkan 2 desimal), diurutkan dari tertinggi.

Skema tabel:
  employees(id INTEGER, name TEXT, salary REAL, department TEXT)

Expected output: 3 departemen dengan avg_salary tertinggi.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, salary REAL, department TEXT);
INSERT INTO employees VALUES (1, 'Alice', 8000.0, 'Engineering');
INSERT INTO employees VALUES (2, 'Bob', 6000.0, 'Marketing');
INSERT INTO employees VALUES (3, 'Charlie', 9000.0, 'Engineering');
INSERT INTO employees VALUES (4, 'Diana', 7000.0, 'HR');
INSERT INTO employees VALUES (5, 'Eve', 5000.0, 'Marketing');
INSERT INTO employees VALUES (6, 'Frank', 11000.0, 'Engineering');
INSERT INTO employees VALUES (7, 'Grace', 8500.0, 'HR');
INSERT INTO employees VALUES (8, 'Hank', 4500.0, 'Support');`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"department":"Engineering","avg_salary":9333.33},{"department":"HR","avg_salary":7750},{"department":"Marketing","avg_salary":5500}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","avg_salary":9333.33},{"department":"HR","avg_salary":7750},{"department":"Marketing","avg_salary":5500}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","avg_salary":9333.33},{"department":"HR","avg_salary":7750},{"department":"Marketing","avg_salary":5500}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"department":"Engineering","avg_salary":9333.33},{"department":"HR","avg_salary":7750},{"department":"Marketing","avg_salary":5500}]`, IsHidden: true},
			},
		},
		{
			Title:      "Karyawan Tanpa Pesanan",
			Category:   "sql",
			Difficulty: "hard",
			Description: `Diberikan tabel customers dan orders, tulis query untuk menampilkan customer yang belum pernah melakukan pesanan.

Skema tabel:
  customers(id INTEGER, name TEXT)
  orders(id INTEGER, customer_id INTEGER, amount REAL)

Expected output: customer yang tidak ada di tabel orders, diurutkan berdasarkan id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE customers (id INTEGER, name TEXT);
INSERT INTO customers VALUES (1, 'Alice');
INSERT INTO customers VALUES (2, 'Bob');
INSERT INTO customers VALUES (3, 'Charlie');
INSERT INTO customers VALUES (4, 'Diana');
CREATE TABLE orders (id INTEGER, customer_id INTEGER, amount REAL);
INSERT INTO orders VALUES (1, 1, 150000.0);
INSERT INTO orders VALUES (2, 3, 75000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Bob"},{"id":4,"name":"Diana"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Bob"},{"id":4,"name":"Diana"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Bob"},{"id":4,"name":"Diana"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"id":2,"name":"Bob"},{"id":4,"name":"Diana"}]`, IsHidden: true},
			},
		},
		{
			Title:      "Ranking Nilai Siswa",
			Category:   "sql",
			Difficulty: "hard",
			Description: `Diberikan tabel students dan grades, tulis query untuk menampilkan nama siswa beserta total nilai mereka, diurutkan dari nilai tertinggi. Tampilkan hanya siswa yang memiliki total nilai di atas 200.

Skema tabel:
  students(id INTEGER, name TEXT)
  grades(id INTEGER, student_id INTEGER, subject TEXT, score INTEGER)

Expected output: name dan total_score, diurutkan berdasarkan total_score DESC.`,
			StarterCode: "",
			Schema: `CREATE TABLE students (id INTEGER, name TEXT);
INSERT INTO students VALUES (1, 'Alice');
INSERT INTO students VALUES (2, 'Bob');
INSERT INTO students VALUES (3, 'Charlie');
CREATE TABLE grades (id INTEGER, student_id INTEGER, subject TEXT, score INTEGER);
INSERT INTO grades VALUES (1, 1, 'Math', 90);
INSERT INTO grades VALUES (2, 1, 'Science', 85);
INSERT INTO grades VALUES (3, 1, 'English', 88);
INSERT INTO grades VALUES (4, 2, 'Math', 70);
INSERT INTO grades VALUES (5, 2, 'Science', 65);
INSERT INTO grades VALUES (6, 2, 'English', 72);
INSERT INTO grades VALUES (7, 3, 'Math', 95);
INSERT INTO grades VALUES (8, 3, 'Science', 92);
INSERT INTO grades VALUES (9, 3, 'English', 89);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"name":"Charlie","total_score":276},{"name":"Alice","total_score":263}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Charlie","total_score":276},{"name":"Alice","total_score":263}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"name":"Charlie","total_score":276},{"name":"Alice","total_score":263}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"name":"Charlie","total_score":276},{"name":"Alice","total_score":263}]`, IsHidden: true},
			},
		},
		{
			Title:      "Self Join Hierarki",
			Category:   "sql",
			Difficulty: "hard",
			Description: `Diberikan tabel employees dengan kolom manager_id yang merujuk ke id karyawan lain, tulis query untuk menampilkan nama karyawan beserta nama manager mereka.

Skema tabel:
  employees(id INTEGER, name TEXT, manager_id INTEGER)

Expected output: employee_name dan manager_name (NULL jika tidak ada manager), diurutkan berdasarkan employee id ASC.`,
			StarterCode: "",
			Schema: `CREATE TABLE employees (id INTEGER, name TEXT, manager_id INTEGER);
INSERT INTO employees VALUES (1, 'Alice', NULL);
INSERT INTO employees VALUES (2, 'Bob', 1);
INSERT INTO employees VALUES (3, 'Charlie', 1);
INSERT INTO employees VALUES (4, 'Diana', 2);
INSERT INTO employees VALUES (5, 'Eve', 2);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"employee_name":"Alice","manager_name":null},{"employee_name":"Bob","manager_name":"Alice"},{"employee_name":"Charlie","manager_name":"Alice"},{"employee_name":"Diana","manager_name":"Bob"},{"employee_name":"Eve","manager_name":"Bob"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"employee_name":"Alice","manager_name":null},{"employee_name":"Bob","manager_name":"Alice"},{"employee_name":"Charlie","manager_name":"Alice"},{"employee_name":"Diana","manager_name":"Bob"},{"employee_name":"Eve","manager_name":"Bob"}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"employee_name":"Alice","manager_name":null},{"employee_name":"Bob","manager_name":"Alice"},{"employee_name":"Charlie","manager_name":"Alice"},{"employee_name":"Diana","manager_name":"Bob"},{"employee_name":"Eve","manager_name":"Bob"}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"employee_name":"Alice","manager_name":null},{"employee_name":"Bob","manager_name":"Alice"},{"employee_name":"Charlie","manager_name":"Alice"},{"employee_name":"Diana","manager_name":"Bob"},{"employee_name":"Eve","manager_name":"Bob"}]`, IsHidden: true},
			},
		},
		{
			Title:      "Pivot Sederhana",
			Category:   "sql",
			Difficulty: "hard",
			Description: `Diberikan tabel sales dengan kolom month dan amount, tulis query untuk menghitung total penjualan per bulan dan menampilkan hanya bulan dengan total penjualan di atas 1000000. Urutkan berdasarkan total DESC.

Skema tabel:
  sales(id INTEGER, month TEXT, amount REAL)

Expected output: month dan total, diurutkan berdasarkan total DESC.`,
			StarterCode: "",
			Schema: `CREATE TABLE sales (id INTEGER, month TEXT, amount REAL);
INSERT INTO sales VALUES (1, 'Jan', 500000.0);
INSERT INTO sales VALUES (2, 'Jan', 700000.0);
INSERT INTO sales VALUES (3, 'Feb', 300000.0);
INSERT INTO sales VALUES (4, 'Feb', 400000.0);
INSERT INTO sales VALUES (5, 'Mar', 800000.0);
INSERT INTO sales VALUES (6, 'Mar', 600000.0);
INSERT INTO sales VALUES (7, 'Apr', 200000.0);`,
			IsActive: true,
			TestCases: []models.TestCase{
				{Input: "", ExpectedOutput: `[{"month":"Jan","total":1200000},{"month":"Mar","total":1400000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"month":"Mar","total":1400000},{"month":"Jan","total":1200000}]`, IsHidden: false},
				{Input: "", ExpectedOutput: `[{"month":"Mar","total":1400000},{"month":"Jan","total":1200000}]`, IsHidden: true},
				{Input: "", ExpectedOutput: `[{"month":"Mar","total":1400000},{"month":"Jan","total":1200000}]`, IsHidden: true},
			},
		},
	}
}

