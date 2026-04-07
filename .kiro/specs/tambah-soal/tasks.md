# Implementation Plan: Tambah Soal

## Overview

Implementasi fitur ini mencakup: (1) tambah field `Schema` ke model Problem, (2) buat SQL Evaluator berbasis SQLite in-memory, (3) update routing evaluator dan submissions service, (4) tambah 100 soal baru ke seed data, dan (5) update frontend untuk mendukung language SQL otomatis.

## Tasks

- [x] 1. Tambah field `Schema` ke model Problem
  - Tambah field `Schema string` ke struct `Problem` di `backend/internal/models/models.go` dengan tag `json:"schema,omitempty" gorm:"type:text"`
  - GORM AutoMigrate akan otomatis menambah kolom `schema TEXT` ke tabel `problems` saat startup
  - _Requirements: 4.3_

- [x] 2. Buat SQL Evaluator
  - [x] 2.1 Tambah dependency `modernc.org/sqlite` ke `backend/go.mod`
    - Jalankan `go get modernc.org/sqlite` di direktori `backend/`
    - Import driver di file baru dengan blank import `_ "modernc.org/sqlite"`
    - _Requirements: 4.1_

  - [x] 2.2 Implementasi `SQLEvaluatorService` di `backend/internal/evaluator/sql_evaluator.go`
    - Buat struct `SQLEvaluatorService` dan constructor `NewSQLEvaluatorService()`
    - Implementasi method `Evaluate(schema, query, expected string) EvalResult`
    - Validasi query: tolak jika mengandung keyword `DROP|DELETE|UPDATE|INSERT|CREATE|ALTER|TRUNCATE|ATTACH|DETACH` menggunakan regex, return `EvalResult{Passed: false, Error: "Operasi tidak diizinkan: hanya SELECT yang diperbolehkan"}`
    - Buka koneksi SQLite in-memory (`file::memory:?cache=shared&mode=memory`)
    - Jalankan `schema` (DDL + INSERT) dalam transaksi
    - Jalankan query user dengan `context.WithTimeout` 5 detik
    - Baca semua rows, serialize ke JSON array of objects (key = nama kolom)
    - Bandingkan actual vs expected menggunakan normalisasi JSON (unmarshal → marshal ulang)
    - Handle error: schema error, timeout, serialize error, expected output bukan JSON valid
    - _Requirements: 4.1, 4.4_

  - [x] 2.3 Tulis property test untuk SQL Evaluator — blocked ops (Property 6)
    - **Property 6: SQL Evaluator menolak operasi berbahaya**
    - **Validates: Requirements 4.4**
    - Gunakan `pgregory.net/rapid`, sample dari slice `[]string{"DROP","DELETE","UPDATE","INSERT","CREATE","ALTER","TRUNCATE","ATTACH","DETACH"}`
    - Verifikasi `result.Passed == false` dan `result.Error != ""`
    - Tag: `Feature: tambah-soal, Property 6`

  - [x] 2.4 Tulis property test untuk SQL Evaluator — deterministic (Property 7)
    - **Property 7: SQL Evaluator menghasilkan output deterministik**
    - **Validates: Requirements 4.1**
    - Jalankan evaluasi dua kali dengan schema dan query yang sama, verifikasi `r1.Actual == r2.Actual`
    - Tag: `Feature: tambah-soal, Property 7`

  - [x] 2.5 Tulis unit tests untuk SQL Evaluator
    - Happy path: SELECT * dari tabel sederhana, verifikasi `result.Passed == true`
    - Timeout: recursive CTE tanpa batas, verifikasi error mengandung "Waktu eksekusi habis"
    - Schema error: DDL tidak valid, verifikasi `result.Passed == false` dan error non-kosong
    - Wrong answer: query valid tapi output tidak cocok, verifikasi `result.Passed == false`
    - _Requirements: 4.1, 4.4_

- [x] 3. Update EvaluatorService — tambah routing method
  - Tambah method `EvaluateWithLanguage(language, schema, code, input, expected string) EvalResult` ke `EvaluatorService` di `backend/internal/evaluator/evaluator.go`
  - Jika `language == "sql"`: delegate ke `SQLEvaluatorService.Evaluate(schema, code, expected)` (field `input` diabaikan untuk SQL)
  - Jika `language == "javascript"` atau default: delegate ke method `Evaluate` yang sudah ada
  - Jika language tidak dikenal: log warning, fallback ke JS evaluator
  - Inisialisasi `SQLEvaluatorService` sebagai field di `EvaluatorService` struct
  - _Requirements: 7.1, 7.3_

  - [x] 3.1 Tulis property test untuk routing EvaluatorService (Property 8)
    - **Property 8: EvaluatorService routing berdasarkan language**
    - **Validates: Requirements 7.1, 7.3**
    - Untuk `language = "sql"`: verifikasi hasil konsisten dengan `SQLEvaluatorService.Evaluate` langsung
    - Untuk `language = "javascript"`: verifikasi hasil konsisten dengan `Evaluate` langsung
    - Tag: `Feature: tambah-soal, Property 8`

- [x] 4. Update SubmissionsService — pass language dan schema ke evaluator
  - Modifikasi `Submit` di `backend/internal/submissions/service.go`
  - Fetch `problem.Schema` saat fetch problem dari DB (sudah ada query `database.DB.First(&problem, ...)`)
  - Ganti panggilan `s.evaluator.Evaluate(...)` dengan `s.evaluator.EvaluateWithLanguage(req.Language, problem.Schema, req.Code, tc.Input, tc.ExpectedOutput)`
  - Jika `problem.Category == "sql"` dan `problem.Schema == ""`: return error 500 dengan pesan "Schema soal SQL tidak ditemukan"
  - _Requirements: 7.1, 7.3_

- [x] 5. Checkpoint — pastikan semua tests backend pass
  - Ensure all tests pass, ask the user if questions arise.

- [x] 6. Tambah 100 soal baru ke seed.go
  - [x] 6.1 Tambah helper `GetSeedProblems() []models.Problem` di `backend/internal/database/seed.go`
    - Ekstrak slice soal ke fungsi terpisah yang bisa dipanggil dari tests
    - Fungsi `Seed` tetap memanggil `GetSeedProblems()` secara internal
    - _Requirements: 5.1, 5.4_

  - [x] 6.2 Tambah 25 soal kategori `loop` ke `GetSeedProblems()`
    - Distribusi: 10 easy, 10 medium, 5 hard
    - Setiap soal: title unik, description Bahasa Indonesia, starterCode fungsi JS, isActive true
    - Easy/medium: minimal 3 visible + 2 hidden test case; hard: minimal 4 visible + 3 hidden test case
    - Contoh soal: Hitung Mundur, Deret Fibonacci, Segitiga Bintang, Spiral Matrix, dll.
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 5.1, 5.2, 5.3, 6.1, 6.2_

  - [x] 6.3 Tambah 25 soal kategori `string` ke `GetSeedProblems()`
    - Distribusi: 10 easy, 10 medium, 5 hard
    - Test case mencakup edge case: string kosong, satu karakter, karakter spesial/spasi
    - Contoh soal: Hitung Vokal, Anagram, Kompresi String, Validasi Palindrom, dll.
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 5.1, 5.2, 5.3, 6.1, 6.2_

  - [x] 6.4 Tambah 25 soal kategori `array` ke `GetSeedProblems()`
    - Distribusi: 10 easy, 10 medium, 5 hard
    - Test case mencakup edge case: array kosong, satu elemen, nilai negatif, duplikat
    - Contoh soal: Elemen Terbesar, Rotasi Array, Subarray Terpanjang, Two Sum, dll.
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 5.1, 5.2, 5.3, 6.1, 6.2_

  - [x] 6.5 Tambah 25 soal kategori `sql` ke `GetSeedProblems()`
    - Distribusi: 10 easy, 10 medium, 5 hard
    - Setiap soal: field `Schema` non-kosong dengan `CREATE TABLE` + `INSERT` statements
    - StarterCode berupa template query SQL (bukan fungsi JS)
    - Input test case diisi string kosong `""`; ExpectedOutput berupa JSON array of objects
    - Deskripsi menjelaskan skema tabel, contoh data, dan expected output
    - Contoh soal: Pilih Semua Pengguna, Filter WHERE, JOIN dua tabel, GROUP BY + AVG, Top N, dll.
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5, 5.1, 5.3, 6.1, 6.2_

  - [x] 6.6 Tulis property tests untuk seed data
    - **Property 1: Semua soal JS memiliki field wajib yang valid** — Validates: Requirements 1.3, 2.3, 3.3, 5.2, 7.1
    - **Property 2: Semua soal SQL memiliki field Schema non-kosong** — Validates: Requirements 4.3
    - **Property 3: Setiap soal memiliki jumlah test case minimum sesuai difficulty** — Validates: Requirements 1.4, 2.4, 3.4, 4.4, 6.2
    - **Property 4: Tidak ada dua soal dengan judul yang sama** — Validates: Requirements 5.4
    - **Property 5: Format expected output valid JSON** — Validates: Requirements 5.3, 7.5 
    - **Property 9: StarterCode soal non-SQL mengandung deklarasi fungsi** — Validates: Requirements 5.2, 7.1, 7.3
    - Tambahkan ke `backend/internal/database/seed_test.go` (file baru)
    - Tag setiap test: `Feature: tambah-soal, Property {N}`

- [x] 7. Checkpoint — verifikasi seed data
  - Ensure all tests pass, ask the user if questions arise.

- [x] 8. Update frontend — tambah language support
  - [x] 8.1 Tambah field `language` ke `submissionStore` di `frontend/store/submissionStore.ts`
    - Tambah `language: string` ke interface `SubmissionStore` dan `initialState` (default `'javascript'`)
    - Tambah `setLanguage: (language: string) => void` ke store
    - _Requirements: 7.1_

  - [x] 8.2 Update `ProblemDetailPage` di `frontend/app/problems/[id]/page.tsx`
    - Import `setLanguage` dari `useSubmissionStore`
    - Setelah fetch problem, set language otomatis: `setLanguage(problem.category === 'sql' ? 'sql' : 'javascript')`
    - Pass `language` sebagai prop ke `SubmitButton` dan `CodeEditor`
    - Update editor top bar: ganti hardcoded `"solution.js"` dengan `language === 'sql' ? 'solution.sql' : 'solution.js'`
    - _Requirements: 7.1_

  - [x] 8.3 Update `CodeEditor` di `frontend/components/Editor/CodeEditor.tsx`
    - Tambah prop `language?: string` ke interface `CodeEditorProps`
    - Pass `language === 'sql' ? 'sql' : 'javascript'` ke prop `defaultLanguage` Monaco Editor
    - _Requirements: 7.1_

  - [x] 8.4 Update `SubmitButton` di `frontend/components/Editor/SubmitButton.tsx`
    - Tambah prop `language: string` ke interface `SubmitButtonProps`
    - Ganti hardcoded `language: 'javascript'` dengan nilai prop `language` saat memanggil `submitSolution`
    - _Requirements: 7.1_

  - [x] 8.5 Tulis unit tests untuk language selector otomatis
    - Test: kategori `sql` → language `'sql'`
    - Test property: kategori `loop | string | array` → language `'javascript'` (gunakan `fast-check`)
    - Tambahkan ke `frontend/__tests__/examples.test.ts` atau file baru
    - _Requirements: 7.1_

- [x] 9. Final checkpoint — semua tests pass
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Tasks bertanda `*` bersifat opsional dan bisa dilewati untuk MVP lebih cepat
- Setiap task mereferensikan requirements spesifik untuk traceability
- SQL Evaluator menggunakan `modernc.org/sqlite` (pure Go, tidak butuh CGO)
- Soal SQL: field `Input` di TestCase selalu `""` — query user adalah "input"-nya
- `GetSeedProblems()` memungkinkan property tests mengakses data seed tanpa DB
- Property tests menggunakan `pgregory.net/rapid` (backend) dan `fast-check` (frontend)
