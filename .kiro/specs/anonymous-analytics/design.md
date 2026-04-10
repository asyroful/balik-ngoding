# Design Document — Anonymous Analytics

## Overview

Fitur Anonymous Analytics menambahkan kemampuan tracking aktivitas pengguna tanpa autentikasi ke platform Balik Ngoding. Setiap perangkat diidentifikasi oleh UUID v4 anonim (`Anonymous_ID`) yang di-generate di sisi klien dan disimpan di `localStorage`. ID ini dikirim bersama setiap submission sehingga backend dapat menghitung statistik agregat: jumlah perangkat unik, total submission, acceptance rate, dan soal paling populer.

Statistik tersebut ditampilkan di halaman `/dashboard` — sebuah Server Component Next.js yang fetch data dari endpoint `GET /analytics/stats`.

Desain ini sengaja dibuat minimal: tidak ada autentikasi, tidak ada PII, dan tidak ada tracking lintas-sesi yang invasif. Tujuannya hanya memberikan gambaran kasar penggunaan platform kepada pemilik.

---

## Architecture

```mermaid
flowchart TD
    subgraph Frontend
        A[ID_Generator\ngetOrCreateAnonymousId] -->|localStorage| B[anonymousId]
        B --> C[submitSolution API call\nPOST /submit + anonymousId]
        D[/dashboard page\nServer Component] -->|fetch| E[GET /analytics/stats]
    end

    subgraph Backend
        F[submissions.Handler\nPOST /submit] --> G[submissions.Service\nSubmit]
        G -->|persist anonymousId| H[(PostgreSQL\nsubmissions table)]
        I[analytics.Handler\nGET /analytics/stats] --> J[analytics.Service\nGetStats]
        J -->|aggregate queries| H
    end

    C --> F
    E --> I
```

Alur utama:
1. Saat halaman pertama kali dimuat, `getOrCreateAnonymousId()` membaca atau membuat UUID di `localStorage`.
2. Setiap kali user submit, `anonymousId` disertakan di request body ke `POST /submit`.
3. Backend menyimpan `anonymous_id` di tabel `submissions`.
4. Halaman `/dashboard` memanggil `GET /analytics/stats` dan menampilkan hasilnya.

---

## Components and Interfaces

### Backend

#### `backend/internal/analytics/`

Package baru mengikuti konvensi yang sudah ada (`problems/`, `submissions/`).

**`handler.go`**
```go
type Handler struct { service *AnalyticsService }
func NewHandler() *Handler
func (h *Handler) GetStats(c *gin.Context)  // GET /analytics/stats
```

**`service.go`**
```go
type AnalyticsService struct{}
func NewAnalyticsService() *AnalyticsService
func (s *AnalyticsService) GetStats() (*StatsResult, error)
```

**`service.go` — DTOs**
```go
type TopProblem struct {
    ProblemID       string `json:"problemId"`
    Title           string `json:"title"`
    SubmissionCount int    `json:"submissionCount"`
}

type StatsResult struct {
    UniqueDevices      int          `json:"uniqueDevices"`
    TotalSubmissions   int          `json:"totalSubmissions"`
    TotalAccepted      int          `json:"totalAccepted"`
    DevicesWithAccepted int         `json:"devicesWithAccepted"`
    TopProblems        []TopProblem `json:"topProblems"`
}
```

#### Perubahan di `submissions/`

`SubmitRequest` dan model `Submission` diperluas dengan field `AnonymousID`:

```go
// submissions/service.go
type SubmitRequest struct {
    ProblemID   string  `json:"problemId"`
    Code        string  `json:"code"`
    Language    string  `json:"language"`
    AnonymousID *string `json:"anonymousId"` // nullable
}
```

```go
// models/models.go — tambah field di Submission
AnonymousID *string `json:"anonymousId" gorm:"type:varchar(36);index"`
```

#### Route baru di `main.go`

```go
analyticsHandler := analytics.NewHandler()
r.GET("/analytics/stats", analyticsHandler.GetStats)
```

---

### Frontend

#### `lib/anonymousId.ts` — ID_Generator

```typescript
export const ANONYMOUS_ID_KEY = 'balik-ngoding-anonymous-id';

export function getOrCreateAnonymousId(): string
// Baca dari localStorage; jika tidak ada, generate UUID v4 dan simpan.
// Jika localStorage tidak tersedia, kembalikan UUID in-memory (session-scoped).
```

#### Perubahan di `lib/api.ts`

`submitSolution` diperluas untuk menyertakan `anonymousId`:

```typescript
export async function submitSolution(req: SubmitRequest): Promise<SubmissionResult>
// req.anonymousId diisi dari getOrCreateAnonymousId() sebelum dikirim

export async function getAnalyticsStats(): Promise<AnalyticsStats>
// GET /analytics/stats
```

#### Perubahan di `lib/types.ts`

```typescript
export interface SubmitRequest {
  problemId: string;
  code: string;
  language: string;
  anonymousId?: string;
}

export interface TopProblem {
  problemId: string;
  title: string;
  submissionCount: number;
}

export interface AnalyticsStats {
  uniqueDevices: number;
  totalSubmissions: number;
  totalAccepted: number;
  devicesWithAccepted: number;
  topProblems: TopProblem[];
}
```

#### `app/dashboard/page.tsx` — Analytics Dashboard

Server Component (tidak perlu `'use client'`) karena hanya fetch dan render data.

```typescript
// fetch di server, tampilkan StatsCard per metrik + TopProblemsTable
export default async function DashboardPage()
```

#### `components/Analytics/StatsCard.tsx`

```typescript
// Pure presentational component
interface StatsCardProps { label: string; value: number | string }
export function StatsCard({ label, value }: StatsCardProps)
```

#### `components/Analytics/TopProblemsTable.tsx`

```typescript
interface TopProblemsTableProps { problems: TopProblem[] }
export function TopProblemsTable({ problems }: TopProblemsTableProps)
```

---

## Data Models

### Perubahan Database: tabel `submissions`

Tambah kolom `anonymous_id` (nullable, tidak ada foreign key — by design karena anonim):

```sql
ALTER TABLE submissions ADD COLUMN anonymous_id VARCHAR(36) NULL;
CREATE INDEX idx_submissions_anonymous_id ON submissions(anonymous_id);
```

GORM akan auto-migrate kolom ini saat `AutoMigrate` dijalankan.

### Model `Submission` (updated)

```go
type Submission struct {
    ID           string          `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    ProblemID    string          `json:"problemId" gorm:"type:uuid;not null;index:idx_submissions_problem_id"`
    AnonymousID  *string         `json:"anonymousId" gorm:"type:varchar(36);index"`
    Code         string          `json:"code"`
    Language     string          `json:"language" gorm:"default:'javascript'"`
    Status       string          `json:"status"`
    Score        int             `json:"score"`
    Total        int             `json:"total"`
    ResultDetail json.RawMessage `json:"resultDetail" gorm:"type:jsonb"`
    CreatedAt    time.Time       `json:"createdAt"`
}
```

### Query Agregasi (`analytics.Service.GetStats`)

```sql
-- uniqueDevices
SELECT COUNT(DISTINCT anonymous_id) FROM submissions WHERE anonymous_id IS NOT NULL;

-- totalSubmissions
SELECT COUNT(*) FROM submissions;

-- totalAccepted
SELECT COUNT(*) FROM submissions WHERE status = 'accepted';

-- devicesWithAccepted
SELECT COUNT(DISTINCT anonymous_id) FROM submissions
WHERE status = 'accepted' AND anonymous_id IS NOT NULL;

-- topProblems (top 5)
SELECT s.problem_id, p.title, COUNT(*) AS submission_count
FROM submissions s
JOIN problems p ON p.id = s.problem_id
GROUP BY s.problem_id, p.title
ORDER BY submission_count DESC
LIMIT 5;
```

Semua query dijalankan via GORM (`Raw` atau method chaining) di dalam `AnalyticsService`.

---

## Correctness Properties

*A property is a characteristic or behavior that should hold true across all valid executions of a system — essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.*

### Property 1: UUID v4 generation

*For any* fresh localStorage state (no existing key), calling `getOrCreateAnonymousId()` should return a string that matches the UUID v4 format `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx` and store that same value under `balik-ngoding-anonymous-id`.

**Validates: Requirements 1.1, 1.4**

---

### Property 2: localStorage round-trip

*For any* valid UUID v4 string written to localStorage under `balik-ngoding-anonymous-id`, calling `getOrCreateAnonymousId()` should return a string strictly equal to the stored value — no mutation, no regeneration.

**Validates: Requirements 1.2, 5.1**

---

### Property 3: anonymousId included in submit request

*For any* valid submission (non-empty code, valid problemId), the request body serialized by `submitSolution()` should contain an `anonymousId` field whose value equals the result of `getOrCreateAnonymousId()` at call time.

**Validates: Requirements 2.1**

---

### Property 4: anonymousId persisted without modification

*For any* valid UUID v4 sent as `anonymousId` in a submission request, the value stored in the `submissions` table should be byte-for-byte equal to the value that was sent — no trimming, casing change, or transformation.

**Validates: Requirements 2.2, 5.2**

---

### Property 5: Stats aggregation correctness

*For any* set of submission records in the database, `GetStats()` should return:
- `uniqueDevices` equal to `COUNT(DISTINCT anonymous_id)` where `anonymous_id IS NOT NULL`
- `totalSubmissions` equal to `COUNT(*)`
- `totalAccepted` equal to `COUNT(*) WHERE status = 'accepted'`
- `devicesWithAccepted` equal to `COUNT(DISTINCT anonymous_id) WHERE status = 'accepted' AND anonymous_id IS NOT NULL`

Each value must match the corresponding direct SQL query result exactly.

**Validates: Requirements 3.1, 3.2, 3.3, 3.5, 5.3**

---

### Property 6: topProblems ordering and size

*For any* set of submissions, the `topProblems` list returned by `GetStats()` should contain at most 5 entries, and for every adjacent pair `(topProblems[i], topProblems[i+1])`, `topProblems[i].submissionCount >= topProblems[i+1].submissionCount` (non-increasing order).

**Validates: Requirements 3.4**

---

### Property 7: Dashboard renders all required fields

*For any* valid `AnalyticsStats` response, the rendered `/dashboard` page should contain visible representations of `uniqueDevices`, `totalSubmissions`, `totalAccepted`, `devicesWithAccepted`, and all entries in `topProblems`.

**Validates: Requirements 4.4**

---

## Error Handling

### Backend

| Kondisi | Respons |
|---|---|
| `GET /analytics/stats` — DB error | `500 Internal Server Error` + `{"error": "Gagal mengambil statistik"}` |
| `POST /submit` — `anonymousId` hadir tapi bukan string | `400 Bad Request` + pesan validasi |
| `POST /submit` — `anonymousId` tidak hadir | Diterima, disimpan sebagai `NULL` |
| Tabel submissions kosong | `200 OK` dengan semua angka `0` dan `topProblems: []` |

### Frontend

| Kondisi | Perilaku |
|---|---|
| `getOrCreateAnonymousId()` — localStorage tidak tersedia | Generate UUID in-memory, tidak crash |
| `GET /analytics/stats` gagal | Tampilkan pesan error, sembunyikan skeleton |
| `POST /submit` gagal | Perilaku error yang sudah ada tidak berubah |
| `anonymousId` gagal di-generate | Submit tetap dikirim tanpa field `anonymousId` (graceful degradation) |

---

## Testing Strategy

### Dual Testing Approach

Fitur ini menggunakan dua lapisan pengujian yang saling melengkapi:

- **Unit tests**: Verifikasi contoh spesifik, edge case, dan kondisi error
- **Property tests**: Verifikasi properti universal di atas input yang di-generate secara acak

### Backend — Go

**Library**: [`pgregory.net/rapid`](https://github.com/flyingmutant/rapid) (sudah digunakan di codebase)

**Property tests** (minimum 100 iterasi per properti):

| Test | Property | Tag |
|---|---|---|
| `TestStatsAggregationCorrectness` | P5 | `Feature: anonymous-analytics, Property 5: stats aggregation correctness` |
| `TestTopProblemsOrdering` | P6 | `Feature: anonymous-analytics, Property 6: topProblems ordering and size` |
| `TestAnonymousIdPersistedWithoutModification` | P4 | `Feature: anonymous-analytics, Property 4: anonymousId persisted without modification` |

**Unit tests / edge cases**:
- `TestGetStatsEmptyTable` — tabel kosong mengembalikan zero values (Req 3.6)
- `TestSubmitWithoutAnonymousId` — submission tanpa `anonymousId` diterima, disimpan NULL (Req 2.3)
- `TestGetStatsEndpoint` — HTTP handler mengembalikan JSON yang benar

### Frontend — TypeScript/Vitest + fast-check

**Library**: [`fast-check`](https://fast-check.dev/) (sudah digunakan di codebase)

**Property tests** (minimum 100 iterasi per properti):

| Test | Property | Tag |
|---|---|---|
| `TestUuidV4Generation` | P1 | `Feature: anonymous-analytics, Property 1: UUID v4 generation` |
| `TestLocalStorageRoundTrip` | P2 | `Feature: anonymous-analytics, Property 2: localStorage round-trip` |
| `TestAnonymousIdIncludedInRequest` | P3 | `Feature: anonymous-analytics, Property 3: anonymousId included in submit request` |
| `TestDashboardRendersAllFields` | P7 | `Feature: anonymous-analytics, Property 7: dashboard renders all required fields` |

**Unit tests / edge cases**:
- `TestGetOrCreateAnonymousId_LocalStorageUnavailable` — localStorage tidak tersedia, tetap return UUID (Req 1.3)
- `TestDashboardLoadingState` — skeleton ditampilkan saat loading (Req 4.2)
- `TestDashboardErrorState` — pesan error ditampilkan saat fetch gagal (Req 4.3)
- `TestDashboardRendersStats` — halaman menampilkan data dari API (Req 4.1)

### Catatan Implementasi

- Setiap property test harus di-tag dengan komentar: `// Feature: anonymous-analytics, Property N: <teks properti>`
- Property tests backend menggunakan `rapid.MakeRun` dengan minimal 100 iterasi
- Property tests frontend menggunakan `fc.assert(fc.property(...))` dengan `{ numRuns: 100 }`
- Unit tests fokus pada contoh konkret dan edge case yang tidak mudah di-generate secara acak
