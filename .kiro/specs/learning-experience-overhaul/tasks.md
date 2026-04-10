# Implementation Plan: Learning Experience Overhaul

## Overview

Refactor sistematis tiga layer: data layer (migrasi seed ke JSON), API layer (endpoint summary + ordering fix), dan frontend layer (kurangi API calls dari 5 ke 2). Urutan implementasi: data layer → backend → frontend → tests.

## Tasks

- [x] 1. Buat JSON seed files — loop, string, array, sql (20 soal per file)
  - Buat direktori `backend/internal/database/seeds/`
  - Buat `seeds/loop.json` dengan 20 soal: soal 1–7 easy, 8–14 medium, 15–20 hard
  - Buat `seeds/string.json` dengan 20 soal: soal 1–7 easy, 8–14 medium, 15–20 hard
  - Buat `seeds/array.json` dengan 20 soal: soal 1–7 easy, 8–14 medium, 15–20 hard
  - Buat `seeds/sql.json` dengan 20 soal: soal 1–7 easy, 8–14 medium, 15–20 hard (migrasi dari `seed_sql.go`)
  - Setiap soal wajib punya: `id`, `title`, `description`, `category`, `difficulty`, `starterCode`, `thinkingGuide`, `hints` (tepat 3 string), `prerequisiteId`, dan `testCases` (3 visible + 2 hidden)
  - Prerequisite chain: soal ke-N punya `prerequisiteId` ke soal ke-(N-1) dalam kategori yang sama; soal pertama tiap kategori `prerequisiteId: null`
  - Konten soal harus berkualitas pedagogis — bukan placeholder
  - _Requirements: 2.1, 2.2, 2.4, 2.5, 2.6, 2.7_

- [x] 2. Refactor `seed.go` — ganti ke JSON loader
  - [x] 2.1 Tambah `SeedProblem` dan `SeedTestCase` struct untuk deserialisasi JSON
    - Definisikan struct sesuai design: `SeedProblem` dengan field `Hints []string` (bukan `pq.StringArray`)
    - _Requirements: 1.2_
  - [x] 2.2 Implementasi `loadProblemsFromFile(fs embed.FS, filename string) ([]SeedProblem, error)`
    - Parse JSON dari embed.FS
    - Validasi setiap soal: minimal 3 visible test cases, minimal 2 hidden test cases
    - Validasi field wajib tidak kosong: `id`, `title`, `category`, `difficulty`, `description`, `starterCode`
    - Validasi `hints` tepat 3 elemen, `thinkingGuide` tidak null/kosong
    - Return error yang menyebut nama file dan ID soal jika validasi gagal
    - _Requirements: 1.3, 1.4_
  - [ ]* 2.3 Write property test untuk seed validation (Property 1)
    - **Property 1: Seed validation rejects problems with insufficient test cases**
    - **Validates: Requirements 1.3**
    - Generate `SeedProblem` dengan visible count < 3 atau hidden count < 2
    - Assert `loadProblemsFromFile` return non-nil error yang mengandung problem ID
  - [x] 2.4 Implementasi `SeedFromJSON(db *gorm.DB) error`
    - Setup `//go:embed seeds/*.json` dan `var seedFiles embed.FS`
    - Load keempat file: `seeds/loop.json`, `seeds/string.json`, `seeds/array.json`, `seeds/sql.json`
    - Konversi `SeedProblem` ke `models.Problem` (convert `[]string` hints ke `pq.StringArray`)
    - Upsert problem dengan `clause.OnConflict` berdasarkan `id`
    - Delete-then-insert test cases per problem untuk simplisitas
    - _Requirements: 1.1, 1.5, 1.6_
  - [ ]* 2.5 Write property test untuk upsert idempotency (Property 2)
    - **Property 2: Upsert idempotency**
    - **Validates: Requirements 1.6**
    - Jalankan `SeedFromJSON` dua kali dengan data yang sama
    - Assert row count dan semua field values identik setelah kedua run
  - [x] 2.6 Update `main.go` — ganti `database.Seed(database.DB)` ke `database.SeedFromJSON(database.DB)`
    - Hapus panggilan ke `database.SeedPrerequisites(database.DB)` jika prerequisite sudah di-handle di JSON
    - _Requirements: 1.1_

- [x] 3. Checkpoint — pastikan seed berjalan tanpa error
  - Ensure all tests pass, ask the user if questions arise.

- [x] 4. Tambah `CategorySummary` model dan `FindSummary` service
  - [x] 4.1 Tambah `CategorySummary` struct di `backend/internal/models/models.go`
    - `type CategorySummary struct { Category string \`json:"category"\`; Total int \`json:"total"\` }`
    - Tidak di-persist ke database — hanya DTO
    - _Requirements: 3.1_
  - [x] 4.2 Implementasi `FindSummary() ([]models.CategorySummary, error)` di `service.go`
    - Query: `SELECT category, COUNT(*) as total FROM problems WHERE is_active = true GROUP BY category`
    - Merge hasil dengan daftar tetap `["loop", "string", "array", "sql"]` untuk pastikan 4 entries selalu ada
    - Kategori tanpa soal aktif dikembalikan dengan `total: 0`
    - _Requirements: 3.2, 3.4_
  - [ ]* 4.3 Write property test untuk FindSummary (Property 3)
    - **Property 3: Summary always returns all four categories with correct counts**
    - **Validates: Requirements 3.2, 3.4**
    - Insert random problems ke test DB dengan berbagai kombinasi kategori
    - Assert `FindSummary` selalu return tepat 4 entries dengan total yang akurat
  - [x] 4.4 Fix ordering di `FindAll` — tambah ORDER BY CASE expression
    - Tambah `.Order("CASE difficulty WHEN 'easy' THEN 1 WHEN 'medium' THEN 2 WHEN 'hard' THEN 3 ELSE 4 END, created_at ASC")` ke query
    - _Requirements: 5.1, 5.2, 5.3_
  - [ ]* 4.5 Write property test untuk difficulty ordering (Property 4)
    - **Property 4: Difficulty ordering is monotonically non-decreasing**
    - **Validates: Requirements 2.3, 5.1, 5.2, 5.3**
    - Insert random problems dengan berbagai difficulty untuk satu kategori
    - Assert urutan hasil `FindAll` selalu easy → medium → hard, lalu `created_at` ASC dalam difficulty yang sama

- [x] 5. Tambah `GetSummary` handler dan register route
  - [x] 5.1 Implementasi `GetSummary(c *gin.Context)` di `handler.go`
    - Panggil `h.service.FindSummary()`
    - Return `{"data": [...]}` dengan HTTP 200 jika sukses
    - Return HTTP 500 dengan `{"message": "Gagal mengambil ringkasan soal"}` jika error
    - _Requirements: 3.1, 3.3, 3.5_
  - [x] 5.2 Register route di `main.go`
    - Tambah `r.GET("/problems/summary", problemsHandler.GetSummary)` **sebelum** `r.GET("/problems/:id", ...)`
    - _Requirements: 3.1_
  - [ ]* 5.3 Write unit test untuk GetSummary handler
    - Test HTTP 500 ketika service return error
    - Test response format: `{"data": [...]}` dengan 4 entries
    - _Requirements: 3.3, 3.5_

- [x] 6. Checkpoint — pastikan semua backend tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [x] 7. Update frontend types dan API client
  - [x] 7.1 Tambah `CategorySummary` interface di `frontend/lib/types.ts`
    - `export interface CategorySummary { category: 'loop' | 'string' | 'array' | 'sql'; total: number; }`
    - _Requirements: 4.1_
  - [x] 7.2 Tambah `getProblemsSummary()` di `frontend/lib/api.ts`
    - Fetch ke `GET /problems/summary`
    - Return `CategorySummary[]` via `handleResponse`
    - _Requirements: 4.1_
  - [ ]* 7.3 Write unit test untuk getProblemsSummary
    - Mock fetch response dengan data valid
    - Assert return value sesuai `CategorySummary[]`
    - Assert response tidak mengandung field `title`, `description`, `starterCode`
    - _Requirements: 3.6_

- [x] 8. Refactor `frontend/app/problems/page.tsx`
  - [x] 8.1 Ganti state `allProblems` dengan `summary: CategorySummary[]`
    - Hapus `const [allProblems, setAllProblems] = useState<Problem[]>([])`
    - Tambah `const [summary, setSummary] = useState<CategorySummary[]>([])`
    - _Requirements: 4.2, 4.6_
  - [x] 8.2 Ganti `useEffect` yang fetch 4 kategori paralel dengan single call ke `getProblemsSummary()`
    - Hapus `Promise.all([getProblems('loop'), getProblems('string'), getProblems('array'), getProblems('sql')])`
    - Tambah fetch ke `getProblemsSummary()` → set ke `summary` state
    - Jika `getProblemsSummary()` gagal: biarkan `summary` tetap `[]` (Progress Tracker tampil 0/0)
    - _Requirements: 4.1, 4.5, 4.6_
  - [x] 8.3 Update kalkulasi Progress Tracker untuk pakai `summary`
    - Ganti `const totalCount = allProblems.length` dengan `const totalCount = summary.reduce((acc, s) => acc + s.total, 0)`
    - `completedCount` tetap dari `useProgress().isAccepted`
    - _Requirements: 4.2, 4.4_
  - [ ]* 8.4 Write property test untuk API calls on initial load (Property 5)
    - **Property 5: Problems page makes exactly 2 API calls on initial load**
    - **Validates: Requirements 4.1, 4.6**
    - Render `ProblemsPageInner` dengan mock fetch untuk setiap initial category
    - Assert fetch dipanggil tepat 2 kali: ke `/problems/summary` dan `/problems?category={initialCategory}`
  - [ ]* 8.5 Write property test untuk category switch (Property 6)
    - **Property 6: Category switch makes exactly 1 additional API call**
    - **Validates: Requirements 4.3**
    - Render, tunggu initial load, ganti kategori
    - Assert hanya 1 call tambahan ke `/problems?category={newCategory}`, tidak ada call ke `/summary`
  - [ ]* 8.6 Write property test untuk Progress Tracker total (Property 12)
    - **Property 12: Progress Tracker total is derived from summary data**
    - **Validates: Requirements 4.2, 4.4**
    - Render dengan berbagai kombinasi `summary` state
    - Assert `totalCount` === `summary.reduce((acc, s) => acc + s.total, 0)`

- [x] 9. Verifikasi HintPanel dan ThinkingGuide visibility
  - [x] 9.1 Pastikan `HintPanel` tidak dirender ketika `hints` null atau array kosong
    - Cek kondisi render di `frontend/components/ProblemDetail/HintPanel.tsx`
    - Tambah guard: `if (!hints || hints.length === 0) return null`
    - _Requirements: 6.1, 6.4_
  - [x] 9.2 Pastikan `ThinkingGuide` tidak dirender ketika `thinkingGuide` null atau string kosong
    - Cek kondisi render di `frontend/components/ProblemDetail/ThinkingGuide.tsx`
    - Tambah guard: `if (!thinkingGuide) return null`
    - _Requirements: 6.2, 6.5_
  - [ ]* 9.3 Write property test untuk HintPanel visibility (Property 7)
    - **Property 7: HintPanel visibility matches hints presence**
    - **Validates: Requirements 6.1, 6.4**
    - Render dengan berbagai nilai `hints` (null, [], array berisi string)
    - Assert HintPanel rendered iff `hints !== null && hints.length > 0`
  - [ ]* 9.4 Write property test untuk ThinkingGuide visibility (Property 8)
    - **Property 8: ThinkingGuide visibility matches thinkingGuide presence**
    - **Validates: Requirements 6.2, 6.5**
    - Render dengan berbagai nilai `thinkingGuide` (null, "", string non-empty)
    - Assert ThinkingGuide rendered iff `thinkingGuide !== null && thinkingGuide !== ""`

- [ ] 10. Write property tests untuk seed content validation (Go)
  - [ ]* 10.1 Write property test untuk seeded problems hints dan thinkingGuide (Property 10)
    - **Property 10: Seeded problems have exactly 3 hints and non-empty thinkingGuide**
    - **Validates: Requirements 2.4, 2.5, 6.3**
    - Load semua problems dari seed JSON files
    - Assert setiap problem: `len(hints) == 3`, setiap hint non-empty, `thinkingGuide` non-nil dan non-empty
  - [ ]* 10.2 Write property test untuk prerequisite chain (Property 11)
    - **Property 11: Prerequisite chain stays within the same category**
    - **Validates: Requirements 2.6, 2.7**
    - Load semua problems dari seed JSON files, buat map `id → problem`
    - Assert setiap problem dengan `prerequisiteId != nil`: prerequisite ada di map, kategori sama, difficulty ≤ problem

- [x] 11. Final checkpoint — pastikan semua tests pass
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Tasks bertanda `*` bersifat opsional dan bisa di-skip untuk MVP yang lebih cepat
- Setiap task mereferensikan requirements spesifik untuk traceability
- Route `/problems/summary` HARUS didaftarkan sebelum `/problems/:id` di Gin agar tidak di-match sebagai `:id`
- JSON seed files adalah konten utama platform — kualitas soal menentukan nilai platform
- Property tests menggunakan **rapid** (Go) dan **fast-check** (TypeScript/Vitest)
