# Requirements Document

## Introduction

Learning Experience Overhaul adalah perbaikan sistematis menyeluruh pada platform Balik Ngoding yang menyentuh tiga layer utama: data layer (seed), API efficiency (backend + frontend), dan UX belajar. Tujuannya bukan menambah fitur baru dari nol, melainkan memperbaiki fondasi yang sudah ada agar platform benar-benar bisa dipakai untuk belajar coding secara progresif dan efisien.

Perbaikan mencakup:
1. Migrasi seed data dari satu file Go monolitik ke JSON files per kategori
2. Penambahan endpoint `/problems/summary` untuk mengurangi over-fetching di frontend
3. Refactor halaman `/problems` agar hanya butuh 2 API calls (bukan 5)
4. Penulisan ulang konten 80 soal (20 per kategori) dengan kualitas pedagogis yang konsisten

## Glossary

- **Seed_Loader**: Komponen Go di `backend/internal/database/seed.go` yang bertanggung jawab memuat data soal ke database
- **Problem_JSON**: File JSON per kategori di `backend/internal/database/seeds/` yang menjadi sumber data soal
- **Problems_Service**: Service Go di `backend/internal/problems/service.go` yang menangani business logic soal
- **Problems_Handler**: Handler Go di `backend/internal/problems/handler.go` yang menangani HTTP request untuk soal
- **Summary_Endpoint**: Endpoint baru `GET /problems/summary` yang mengembalikan jumlah soal per kategori
- **Problems_Page**: Halaman Next.js di `frontend/app/problems/page.tsx` yang menampilkan daftar soal
- **Progress_Tracker**: Komponen UI di Problems_Page yang menampilkan progress belajar user
- **Category**: Kelompok soal berdasarkan topik — `loop`, `string`, `array`, atau `sql`
- **Difficulty**: Tingkat kesulitan soal — `easy`, `medium`, atau `hard`
- **Hint**: Petunjuk progresif yang menuntun user berpikir tanpa langsung memberi jawaban
- **ThinkingGuide**: Panduan berpikir yang menjelaskan konsep di balik soal
- **PrerequisiteID**: UUID soal yang harus diselesaikan sebelum soal ini bisa dikerjakan
- **TestCase**: Input/output pair untuk mengevaluasi jawaban user; bisa visible (3) atau hidden (2)
- **CategorySummary**: Objek `{category: string, total: number}` yang dikembalikan Summary_Endpoint

---

## Requirements

### Requirement 1: Migrasi Seed Data ke JSON Files

**User Story:** Sebagai developer yang maintain platform, saya ingin data soal disimpan di JSON files terpisah per kategori, sehingga saya bisa menambah, mengedit, dan me-review soal tanpa harus membuka file Go 2200+ baris.

#### Acceptance Criteria

1. THE Seed_Loader SHALL membaca data soal dari file JSON di direktori `backend/internal/database/seeds/` menggunakan Go `embed.FS`
2. THE Problem_JSON SHALL menggunakan format yang mencakup field: `id`, `title`, `description`, `category`, `difficulty`, `starterCode`, `hints`, `thinkingGuide`, `prerequisiteId`, dan `testCases`
3. WHEN Seed_Loader memuat Problem_JSON, THE Seed_Loader SHALL memvalidasi bahwa setiap soal memiliki minimal 3 test case visible dan 2 test case hidden
4. IF Problem_JSON tidak dapat di-parse, THEN THE Seed_Loader SHALL mengembalikan error yang menyebutkan nama file dan baris yang bermasalah
5. THE Seed_Loader SHALL memuat soal dari empat file JSON terpisah: `loop.json`, `string.json`, `array.json`, dan `sql.json`
6. WHEN seeding dijalankan, THE Seed_Loader SHALL melakukan upsert berdasarkan `id` sehingga data yang sudah ada tidak diduplikasi

---

### Requirement 2: Konten 80 Soal Berkualitas

**User Story:** Sebagai pemula yang baru belajar programming, saya ingin setiap soal punya hints yang menuntun cara berpikir dan urutan soal yang progresif, sehingga saya bisa belajar secara bertahap tanpa merasa tersesat.

#### Acceptance Criteria

1. THE Problem_JSON SHALL berisi tepat 20 soal untuk setiap Category (loop, string, array, sql) — total 80 soal
2. THE Problem_JSON SHALL mendistribusikan difficulty dalam setiap Category dengan urutan: easy (soal 1–7), medium (soal 8–14), hard (soal 15–20)
3. WHEN soal ditampilkan dalam satu Category, THE Problems_Service SHALL mengurutkan soal berdasarkan difficulty (easy → medium → hard) lalu berdasarkan urutan dalam JSON
4. THE Problem_JSON SHALL memastikan setiap soal memiliki field `hints` berisi tepat 3 string yang bersifat progresif: hint pertama memberi arah umum, hint kedua mempersempit pendekatan, hint ketiga hampir memberi jawaban
5. THE Problem_JSON SHALL memastikan setiap soal memiliki field `thinkingGuide` berisi penjelasan konsep yang relevan dengan soal tersebut
6. THE Problem_JSON SHALL mendefinisikan `prerequisiteId` yang membentuk rantai belajar logis dalam satu Category — soal pertama setiap Category tidak memiliki prerequisite
7. IF sebuah soal memiliki `prerequisiteId`, THEN `prerequisiteId` tersebut SHALL merujuk ke soal lain dalam Category yang sama dengan difficulty lebih rendah atau sama

---

### Requirement 3: Endpoint Summary untuk Progress Tracker

**User Story:** Sebagai user yang membuka halaman daftar soal, saya ingin progress tracker langsung muncul tanpa harus menunggu semua soal dari semua kategori di-load, sehingga halaman terasa lebih cepat.

#### Acceptance Criteria

1. THE Problems_Handler SHALL menyediakan endpoint `GET /problems/summary` yang mengembalikan array of CategorySummary
2. WHEN `GET /problems/summary` dipanggil, THE Problems_Service SHALL menghitung jumlah soal aktif per Category menggunakan satu query database dengan GROUP BY
3. THE Summary_Endpoint SHALL mengembalikan response dengan format: `{"data": [{"category": "loop", "total": 20}, ...]}`
4. THE Summary_Endpoint SHALL selalu mengembalikan entry untuk keempat Category (loop, string, array, sql) meskipun salah satu Category tidak memiliki soal aktif (total: 0)
5. IF terjadi error database saat mengambil summary, THEN THE Problems_Handler SHALL mengembalikan HTTP 500 dengan pesan error dalam Bahasa Indonesia
6. THE Summary_Endpoint SHALL tidak mengembalikan data soal (title, description, starterCode, dll) — hanya category dan total

---

### Requirement 4: Refactor Frontend — Efisiensi API Calls

**User Story:** Sebagai user yang membuka halaman `/problems`, saya ingin halaman load lebih cepat dengan network request yang minimal, sehingga pengalaman browsing tidak terasa lambat terutama di koneksi terbatas.

#### Acceptance Criteria

1. WHEN Problems_Page pertama kali dimuat, THE Problems_Page SHALL melakukan tepat 2 API calls: satu ke `GET /problems/summary` dan satu ke `GET /problems?category={selectedCategory}`
2. THE Progress_Tracker SHALL menghitung total soal berdasarkan data dari `GET /problems/summary`, bukan dari fetching semua soal per kategori
3. WHEN user mengganti Category, THE Problems_Page SHALL hanya melakukan 1 API call tambahan ke `GET /problems?category={newCategory}` — tidak memanggil ulang summary
4. THE Problems_Page SHALL menampilkan Progress_Tracker dengan data yang sudah tersedia dari summary (total soal) dikombinasikan dengan data solved yang sudah ada di local state
5. IF `GET /problems/summary` gagal, THEN THE Problems_Page SHALL tetap menampilkan daftar soal dari category yang dipilih, dengan Progress_Tracker menampilkan state kosong (0/0)
6. THE Problems_Page SHALL menghapus semua pemanggilan `getProblems` paralel untuk keempat Category yang sebelumnya digunakan untuk menghitung total soal

---

### Requirement 5: Ordering Soal yang Konsisten

**User Story:** Sebagai user yang belajar dari awal, saya ingin soal-soal dalam satu kategori muncul dari yang paling mudah ke yang paling sulit, sehingga saya bisa mengikuti alur belajar yang masuk akal.

#### Acceptance Criteria

1. WHEN `GET /problems?category={category}` dipanggil, THE Problems_Service SHALL mengembalikan soal yang diurutkan: difficulty `easy` terlebih dahulu, kemudian `medium`, kemudian `hard`
2. WHEN soal dengan difficulty yang sama ditampilkan, THE Problems_Service SHALL mengurutkan soal tersebut berdasarkan `created_at` ascending untuk menjaga urutan yang konsisten
3. THE Problems_Service SHALL menggunakan SQL ORDER BY dengan CASE expression untuk mengurutkan difficulty: `CASE WHEN difficulty='easy' THEN 1 WHEN difficulty='medium' THEN 2 WHEN difficulty='hard' THEN 3 END`
4. THE ProblemTable SHALL menampilkan badge difficulty dengan warna yang berbeda: hijau untuk easy, kuning untuk medium, merah untuk hard

---

### Requirement 6: Hints dan ThinkingGuide Selalu Tersedia

**User Story:** Sebagai user yang stuck mengerjakan soal, saya ingin bisa melihat hints yang membantu saya berpikir, sehingga saya tidak langsung menyerah dan mencari jawaban di internet.

#### Acceptance Criteria

1. WHEN user membuka halaman detail soal, THE HintPanel SHALL menampilkan hints jika field `hints` pada soal tidak kosong dan tidak null
2. WHEN user membuka halaman detail soal, THE ThinkingGuide SHALL menampilkan thinking guide jika field `thinkingGuide` pada soal tidak null
3. THE Problems_Service SHALL memastikan field `hints` dan `thinkingGuide` selalu ter-populate untuk semua soal yang di-seed dari Problem_JSON
4. IF field `hints` pada soal adalah null atau array kosong, THEN THE HintPanel SHALL tidak ditampilkan (bukan menampilkan panel kosong)
5. IF field `thinkingGuide` pada soal adalah null atau string kosong, THEN THE ThinkingGuide SHALL tidak ditampilkan (bukan menampilkan panel kosong)
