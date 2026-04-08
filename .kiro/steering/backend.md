---
inclusion: always
---

# Backend — Go/Gin Service

## Stack

- **Language**: Go 1.25
- **Framework**: Gin (HTTP router & middleware)
- **ORM**: GORM
- **Database**: PostgreSQL 16
- **Runtime**: Go binary (compiled)

## Struktur Direktori

```
backend/
├── main.go                         # Entry point, setup Gin router
├── go.mod
├── go.sum
├── Dockerfile
├── .env
├── .env.example
└── internal/
    ├── database/
    │   ├── database.go             # Koneksi PostgreSQL via GORM
    │   ├── seed.go                 # Seeding data soal
    │   └── seed_sql.go             # Seeding soal SQL
    ├── evaluator/
    │   ├── evaluator.go            # Evaluator utama
    │   ├── sql_evaluator.go        # Evaluator khusus SQL (SQLite sandbox)
    │   └── sqlite_driver.go        # SQLite driver untuk sandbox
    ├── models/
    │   └── models.go               # GORM models: Problem, TestCase, Submission
    ├── problems/
    │   ├── handler.go              # Gin handler untuk /problems endpoints
    │   └── service.go              # Business logic problems
    └── submissions/
        ├── handler.go              # Gin handler untuk /submit endpoint
        └── service.go              # Business logic submissions & evaluasi
```

## API Endpoints

| Method | Path | Deskripsi |
|---|---|---|
| GET | `/problems` | Ambil semua soal aktif (support query param `category`) |
| GET | `/problems/:id` | Ambil detail satu soal |
| POST | `/submit` | Submit jawaban untuk dievaluasi |
| GET | `/health` | Health check endpoint |

## Environment Variables

| Variable | Deskripsi | Contoh Nilai |
|---|---|---|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5432/balik_ngoding?sslmode=disable` |
| `PORT` | Port yang digunakan backend | `8080` |
| `FRONTEND_ORIGIN` | Origin frontend untuk CORS | `http://localhost:3000` |

## Konvensi Kode

- Semua package ada di `backend/internal/` — tidak ada kode bisnis di luar `internal/`
- Handler hanya menangani HTTP request/response, logic ada di service
- GORM digunakan untuk semua operasi database — tidak ada raw SQL kecuali di evaluator
- Error dikembalikan sebagai JSON: `{"error": "pesan error"}`
- Gunakan Go idioms: error handling eksplisit, tidak ada panic di production code

## Menjalankan Backend

```bash
cd backend
go run main.go

# Atau build binary
go build -o balik-ngoding-backend .
./balik-ngoding-backend
```

## Menjalankan Tests

```bash
cd backend
go test ./...

# Dengan verbose output
go test -v ./...
```
