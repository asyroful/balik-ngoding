# Tasks — Anonymous Analytics

## Task List

- [x] 1. Backend: Data Model & Migration
  - [x] 1.1 Tambah field `AnonymousID *string` ke struct `Submission` di `backend/internal/models/models.go` dengan GORM tag `type:varchar(36);index`
  - [x] 1.2 Update `AutoMigrate` di `backend/internal/database/database.go` agar kolom `anonymous_id` ter-migrate ke tabel `submissions`

- [x] 2. Backend: Perubahan Submissions Package
  - [x] 2.1 Tambah field `AnonymousID *string` ke `SubmitRequest` di `backend/internal/submissions/service.go`
  - [x] 2.2 Update handler `Submit` di `backend/internal/submissions/handler.go` untuk membaca field `anonymousId` dari request body (nullable — tidak wajib)
  - [x] 2.3 Update `submissions.Service.Submit` untuk menyimpan `AnonymousID` ke record `Submission` saat persist ke DB

- [x] 3. Backend: Analytics Package
  - [x] 3.1 Buat file `backend/internal/analytics/service.go` dengan struct `AnalyticsService`, DTOs `TopProblem` dan `StatsResult`, dan method `GetStats() (*StatsResult, error)` yang menjalankan 5 query agregasi
  - [x] 3.2 Buat file `backend/internal/analytics/handler.go` dengan `Handler`, `NewHandler()`, dan `GetStats` gin handler untuk `GET /analytics/stats`
  - [x] 3.3 Daftarkan route `GET /analytics/stats` di `backend/main.go`

- [x] 4. Backend: Tests
  - [x] 4.1 Buat `backend/internal/analytics/service_test.go` dengan unit test `TestGetStatsEmptyTable` (edge case: tabel kosong → semua zero)
  - [x] 4.2 Buat property test `TestStatsAggregationCorrectness` menggunakan `pgregory.net/rapid` — untuk sembarang set submissions, verifikasi `uniqueDevices`, `totalSubmissions`, `totalAccepted`, `devicesWithAccepted` cocok dengan query langsung
    - Tag: `// Feature: anonymous-analytics, Property 5: stats aggregation correctness`
  - [x] 4.3 Buat property test `TestTopProblemsOrdering` — untuk sembarang set submissions, verifikasi `topProblems` max 5 entri dan urutan non-increasing
    - Tag: `// Feature: anonymous-analytics, Property 6: topProblems ordering and size`
  - [x] 4.4 Buat property test `TestAnonymousIdPersistedWithoutModification` di `backend/internal/submissions/service_test.go` — untuk sembarang UUID v4, submit lalu baca kembali dari DB, nilai harus identik
    - Tag: `// Feature: anonymous-analytics, Property 4: anonymousId persisted without modification`
  - [x] 4.5 Buat unit test `TestSubmitWithoutAnonymousId` — submission tanpa `anonymousId` diterima dan disimpan dengan `NULL`

- [x] 5. Frontend: ID Generator
  - [x] 5.1 Buat file `frontend/lib/anonymousId.ts` dengan fungsi `getOrCreateAnonymousId()`: baca dari localStorage, jika tidak ada generate UUID v4 dan simpan; jika localStorage tidak tersedia, return UUID in-memory

- [x] 6. Frontend: Types & API
  - [x] 6.1 Tambah field `anonymousId?: string` ke interface `SubmitRequest` di `frontend/lib/types.ts`
  - [x] 6.2 Tambah interface `TopProblem` dan `AnalyticsStats` ke `frontend/lib/types.ts`
  - [x] 6.3 Update fungsi `submitSolution` di `frontend/lib/api.ts` untuk memanggil `getOrCreateAnonymousId()` dan menyertakan hasilnya sebagai `anonymousId` di request body
  - [x] 6.4 Tambah fungsi `getAnalyticsStats(): Promise<AnalyticsStats>` ke `frontend/lib/api.ts`

- [x] 7. Frontend: Dashboard Page & Components
  - [x] 7.1 Buat komponen `frontend/components/Analytics/StatsCard.tsx` — presentational, props: `label: string`, `value: number | string`
  - [x] 7.2 Buat komponen `frontend/components/Analytics/TopProblemsTable.tsx` — presentational, props: `problems: TopProblem[]`
  - [x] 7.3 Buat halaman `frontend/app/dashboard/page.tsx` sebagai Server Component yang fetch `getAnalyticsStats()`, tampilkan loading skeleton saat fetch, error message jika gagal, dan render `StatsCard` + `TopProblemsTable` dengan data

- [x] 8. Frontend: Tests
  - [x] 8.1 Buat `frontend/__tests__/anonymousIdProperties.test.ts` dengan property test `TestUuidV4Generation` — untuk sembarang fresh state, `getOrCreateAnonymousId()` return string yang match UUID v4 regex
    - Tag: `// Feature: anonymous-analytics, Property 1: UUID v4 generation`
  - [x] 8.2 Buat property test `TestLocalStorageRoundTrip` di file yang sama — untuk sembarang UUID v4 yang disimpan ke localStorage, `getOrCreateAnonymousId()` return nilai yang sama
    - Tag: `// Feature: anonymous-analytics, Property 2: localStorage round-trip`
  - [x] 8.3 Buat property test `TestAnonymousIdIncludedInRequest` — untuk sembarang valid submission, request body yang dikirim `submitSolution()` mengandung field `anonymousId`
    - Tag: `// Feature: anonymous-analytics, Property 3: anonymousId included in submit request`
  - [x] 8.4 Buat `frontend/__tests__/DashboardProperties.test.tsx` dengan property test `TestDashboardRendersAllFields` — untuk sembarang `AnalyticsStats`, semua field tampil di render output
    - Tag: `// Feature: anonymous-analytics, Property 7: dashboard renders all required fields`
  - [x] 8.5 Buat unit test `TestGetOrCreateAnonymousId_LocalStorageUnavailable` — localStorage diblokir, fungsi tetap return UUID valid tanpa throw
  - [x] 8.6 Buat unit test `TestDashboardErrorState` — fetch gagal, halaman menampilkan pesan error
