# Requirements Document

## Introduction

Fitur ini menambahkan 100 soal baru ke platform Balik Ngoding untuk memperkaya konten sebelum campaign LinkedIn. Soal-soal baru harus bervariasi dalam difficulty (easy/medium/hard) dan kategori (loop, string, array, sql), berkualitas tinggi, dan memiliki test case yang memadai termasuk edge case tersembunyi. Target minimal 100 soal aktif setelah penambahan.

## Glossary

- **Platform**: Aplikasi web Balik Ngoding (frontend Next.js + backend Go/Gin)
- **Problem**: Satu soal pemrograman dengan judul, deskripsi, kategori, difficulty, starter code, dan test case
- **Test_Case**: Pasangan input dan expected output untuk mengevaluasi submission; bisa visible atau hidden
- **Seed_File**: File `backend/internal/database/seed.go` yang berisi data soal awal untuk database
- **Evaluator**: Layanan backend yang mengeksekusi kode JavaScript user menggunakan goja
- **Category**: Pengelompokan soal: `loop`, `string`, `array`, `sql`
- **Difficulty**: Tingkat kesulitan: `easy`, `medium`, `hard`
- **Starter_Code**: Template fungsi JavaScript yang ditampilkan di editor saat user membuka soal
- **Hidden_Test_Case**: Test_Case dengan `isHidden: true` yang tidak ditampilkan ke user tapi digunakan saat evaluasi

---

## Requirements

### Requirement 1: Penambahan Soal Kategori Loop

**User Story:** Sebagai user, saya ingin mengerjakan soal-soal bertema perulangan yang bervariasi, sehingga saya dapat melatih kemampuan loop dan iterasi dengan berbagai tingkat kesulitan.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat minimal 6 soal baru berkategori `loop` yang belum ada sebelumnya.
2. THE Seed_File SHALL memuat soal loop dengan distribusi difficulty: minimal 2 soal `easy`, 2 soal `medium`, dan 1 soal `hard`.
3. WHEN soal loop di-seed ke database, THE Platform SHALL menyimpan setiap soal dengan field: title, description (Bahasa Indonesia), category `loop`, difficulty, starterCode (template fungsi JavaScript), dan isActive `true`.
4. THE Seed_File SHALL memuat setiap soal loop dengan minimal 3 visible Test_Case dan minimal 2 hidden Test_Case.
5. WHEN Evaluator menjalankan Test_Case soal loop, THE Evaluator SHALL membandingkan output fungsi JavaScript user dengan expected output secara exact match setelah normalisasi whitespace.

---

### Requirement 2: Penambahan Soal Kategori String

**User Story:** Sebagai user, saya ingin mengerjakan soal-soal manipulasi string yang bervariasi, sehingga saya dapat melatih kemampuan string processing dengan berbagai tingkat kesulitan.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat minimal 6 soal baru berkategori `string` yang belum ada sebelumnya.
2. THE Seed_File SHALL memuat soal string dengan distribusi difficulty: minimal 2 soal `easy`, 2 soal `medium`, dan 1 soal `hard`.
3. WHEN soal string di-seed ke database, THE Platform SHALL menyimpan setiap soal dengan field: title, description (Bahasa Indonesia), category `string`, difficulty, starterCode (template fungsi JavaScript), dan isActive `true`.
4. THE Seed_File SHALL memuat setiap soal string dengan minimal 3 visible Test_Case dan minimal 2 hidden Test_Case.
5. THE Seed_File SHALL memuat Test_Case soal string yang mencakup edge case: string kosong, string satu karakter, dan string dengan karakter spesial atau spasi.

---

### Requirement 3: Penambahan Soal Kategori Array

**User Story:** Sebagai user, saya ingin mengerjakan soal-soal manipulasi array yang bervariasi, sehingga saya dapat melatih kemampuan array processing dengan berbagai tingkat kesulitan.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat minimal 6 soal baru berkategori `array` yang belum ada sebelumnya.
2. THE Seed_File SHALL memuat soal array dengan distribusi difficulty: minimal 2 soal `easy`, 2 soal `medium`, dan 1 soal `hard`.
3. WHEN soal array di-seed ke database, THE Platform SHALL menyimpan setiap soal dengan field: title, description (Bahasa Indonesia), category `array`, difficulty, starterCode (template fungsi JavaScript), dan isActive `true`.
4. THE Seed_File SHALL memuat setiap soal array dengan minimal 3 visible Test_Case dan minimal 2 hidden Test_Case.
5. THE Seed_File SHALL memuat Test_Case soal array yang mencakup edge case: array kosong, array satu elemen, dan array dengan nilai negatif atau duplikat.

---

### Requirement 4: Penambahan Soal Kategori SQL

**User Story:** Sebagai user, saya ingin mengerjakan soal-soal SQL dasar, sehingga saya dapat melatih kemampuan query database yang sering dibutuhkan di dunia kerja.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat minimal 5 soal baru berkategori `sql`.
2. THE Seed_File SHALL memuat soal sql dengan distribusi difficulty: minimal 2 soal `easy`, 2 soal `medium`, dan 1 soal `hard`.
3. WHEN soal sql di-seed ke database, THE Platform SHALL menyimpan setiap soal dengan field: title, description (Bahasa Indonesia), category `sql`, difficulty, starterCode (template query SQL), dan isActive `true`.
4. THE Seed_File SHALL memuat setiap soal sql dengan minimal 2 visible Test_Case dan minimal 2 hidden Test_Case.
5. THE Seed_File SHALL memuat deskripsi soal sql yang menjelaskan skema tabel yang relevan beserta contoh data dan expected output query.

---

### Requirement 5: Kualitas dan Konsistensi Soal

**User Story:** Sebagai user, saya ingin soal-soal yang jelas dan konsisten, sehingga saya tidak bingung dengan instruksi yang ambigu atau test case yang tidak konsisten.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat setiap soal dengan deskripsi yang mencakup: penjelasan masalah, format input, format output, dan minimal satu contoh input/output.
2. THE Seed_File SHALL memuat StarterCode setiap soal dalam format template fungsi JavaScript dengan nama fungsi yang sesuai judul soal (camelCase).
3. THE Seed_File SHALL memuat expected output setiap Test_Case dalam format yang konsisten dengan output fungsi JavaScript (JSON-serializable: string dengan tanda kutip, array dengan bracket, boolean lowercase, number tanpa kutip).
4. IF dua soal memiliki judul yang sama, THEN THE Seed_File SHALL dianggap tidak valid dan harus diperbaiki sebelum di-deploy.
5. THE Seed_File SHALL memuat total minimal 25 soal aktif (termasuk 6 soal yang sudah ada) setelah penambahan.

---

### Requirement 6: Distribusi Difficulty Keseluruhan

**User Story:** Sebagai user, saya ingin platform memiliki soal dari berbagai tingkat kesulitan, sehingga saya dapat belajar secara bertahap dari mudah ke sulit.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat total soal dengan distribusi difficulty: minimal 8 soal `easy`, minimal 10 soal `medium`, dan minimal 5 soal `hard` dari keseluruhan soal aktif.
2. THE Seed_File SHALL memuat soal `hard` dengan minimal 4 visible Test_Case dan minimal 3 hidden Test_Case untuk memastikan coverage yang memadai.
3. WHEN Problem_List menampilkan soal, THE Platform SHALL menampilkan badge difficulty dengan warna berbeda: hijau untuk `easy`, kuning untuk `medium`, merah untuk `hard`.

---

### Requirement 7: Kompatibilitas dengan Evaluator

**User Story:** Sebagai developer, saya ingin semua soal baru kompatibel dengan evaluator yang sudah ada, sehingga tidak perlu mengubah infrastruktur backend.

#### Acceptance Criteria

1. THE Seed_File SHALL memuat StarterCode setiap soal non-SQL sebagai fungsi JavaScript yang dapat dieksekusi oleh Evaluator berbasis goja.
2. THE Seed_File SHALL memuat input setiap Test_Case dalam format yang dapat di-parse sebagai argumen fungsi JavaScript (JSON-valid atau literal JavaScript).
3. WHEN Evaluator mengeksekusi soal baru, THE Evaluator SHALL memanggil fungsi dengan nama yang sama seperti yang didefinisikan di StarterCode.
4. IF soal memiliki multiple parameter, THEN THE Seed_File SHALL memuat input Test_Case sebagai argumen terpisah dengan koma (contoh: `"[1,2,3], 2"` untuk dua parameter).
5. THE Seed_File SHALL memuat expected output setiap Test_Case dalam format string yang merepresentasikan nilai return fungsi setelah `JSON.stringify`.
