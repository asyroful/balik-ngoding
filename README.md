# Balik Ngoding

Platform latihan coding berbasis web — gratis, tanpa daftar, langsung di browser.

![Next.js](https://img.shields.io/badge/Next.js-14-black?logo=next.js)
![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-blue?logo=postgresql)
![Tailwind CSS](https://img.shields.io/badge/Tailwind-3-38bdf8?logo=tailwindcss)

## Fitur

- Kerjakan soal langsung di browser dengan Monaco Editor (VS Code-like)
- Kategori soal: **Loop**, **String**, **Array**, **SQL**
- Tingkat kesulitan: Easy, Medium, Hard
- Evaluasi JavaScript di sandbox aman (timeout 5 detik, API terbatas)
- Evaluasi SQL dengan SQLite in-memory
- Feedback per test case: input, expected output, actual output
- Progress tracker keseluruhan tersimpan di cookie (tanpa login)
- URL sync per kategori — konteks tetap terjaga saat navigasi

## Tech Stack

| Layer | Teknologi |
|---|---|
| Frontend | Next.js 14, TypeScript, Tailwind CSS, Zustand, Monaco Editor |
| Backend | Go (Gin), GORM |
| Database | PostgreSQL |
| Evaluator | goja (JS sandbox), SQLite in-memory |
| Testing | Vitest, fast-check (property-based) |
| Deploy | Netlify (frontend), Docker (backend) |

## Struktur Project

```
balik-ngoding/
├── frontend/          # Next.js app
│   ├── app/           # Pages (App Router)
│   ├── components/    # UI components
│   ├── hooks/         # useProgress (cookie-based)
│   ├── lib/           # API client, types
│   └── store/         # Zustand submission store
└── backend/           # Go API
    └── internal/
        ├── database/  # Init, seed
        ├── evaluator/ # JS & SQL evaluator
        ├── models/    # GORM models
        ├── problems/  # Handler & service
        └── submissions/
```

---

## Quick Start (Docker Compose)

Cara tercepat untuk menjalankan seluruh stack (database + backend + frontend) dengan satu perintah.

**Prasyarat:** [Docker Desktop](https://www.docker.com/products/docker-desktop/) terinstall dan berjalan.

```bash
# 1. Clone repository
git clone <repo-url>
cd balik-ngoding

# 2. (Opsional) Salin dan sesuaikan environment variables
cp .env.example .env

# 3. Jalankan seluruh stack
docker compose up

# 4. Buka aplikasi di browser
# http://localhost:3000
```

Selesai. Aplikasi berjalan di `http://localhost:3000`.

### Perintah Umum Docker Compose

```bash
# Jalankan stack di background (detached mode)
docker compose up -d

# Hentikan semua service
docker compose down

# Lihat log semua service secara real-time
docker compose logs -f

# Lihat log service tertentu (db / backend / frontend)
docker compose logs -f backend
docker compose logs -f frontend
docker compose logs -f db

# Rebuild image setelah ada perubahan Dockerfile
docker compose up --build
```

---

## Menjalankan Lokal (Manual — Tanpa Docker)

Alternatif jika kamu tidak menggunakan Docker. Pastikan PostgreSQL sudah berjalan di `localhost:5432`.

### Prasyarat

- Node.js 18+
- Go 1.21+
- PostgreSQL

### Backend

```bash
cd backend
cp .env.example .env
# Edit .env sesuai koneksi PostgreSQL kamu

go run .
```

Server berjalan di `http://localhost:8080`.

### Frontend

```bash
cd frontend
cp .env.example .env.local
# NEXT_PUBLIC_API_URL=http://localhost:8080

npm install
npm run dev
```

App berjalan di `http://localhost:3000`.

---

## Environment Variables

**Root** (`.env`, digunakan Docker Compose):

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=balik_ngoding
DATABASE_URL=postgres://postgres:postgres@db:5432/balik_ngoding?sslmode=disable
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000
NEXT_PUBLIC_API_URL=http://localhost:8080
```

**Backend** (`backend/.env`, untuk setup manual):

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/balik_ngoding?sslmode=disable
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000
```

**Frontend** (`frontend/.env.local`, untuk setup manual):

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

---

## API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/health` | Health check |
| GET | `/problems?category=loop` | Daftar soal (filter opsional) |
| GET | `/problems/:id` | Detail soal |
| POST | `/submit` | Submit solusi |

### Contoh Submit

```json
POST /submit
{
  "problemId": "uuid",
  "code": "function solve(n) { return n * 2; }",
  "language": "javascript"
}
```

---

## Testing

```bash
# Frontend
cd frontend
npm test

# Backend
cd backend
go test ./...
```

---

## Troubleshooting

### Port sudah digunakan (5432, 8080, atau 3000)

Jika muncul error `address already in use`, ada proses lain yang menggunakan port tersebut.

**Cari dan hentikan proses yang menggunakan port (contoh port 8080):**

```bash
# Linux / macOS
lsof -ti :8080 | xargs kill -9

# Windows (PowerShell)
netstat -ano | findstr :8080
# Catat PID-nya, lalu:
taskkill /PID <PID> /F
```

**Atau ganti port di `.env`:**

```env
PORT=8081               # ganti port backend
NEXT_PUBLIC_API_URL=http://localhost:8081
```

Lalu jalankan ulang: `docker compose up`.

---

### Database connection error

Jika backend gagal konek ke database:

- **Docker Compose**: pastikan service `db` sudah healthy sebelum backend start. Cek dengan `docker compose logs db`.
- **Manual**: pastikan PostgreSQL berjalan di `localhost:5432` dan `DATABASE_URL` di `backend/.env` sudah benar.

```bash
# Cek status PostgreSQL (Linux/macOS)
pg_isready -h localhost -p 5432
```

---

### Frontend tidak bisa reach backend

Jika frontend menampilkan error network atau data tidak muncul:

- Pastikan `NEXT_PUBLIC_API_URL` sudah di-set dengan benar.
- Untuk Docker Compose: nilai default `http://localhost:8080` sudah benar — pastikan backend container berjalan (`docker compose ps`).
- Untuk setup manual: pastikan backend berjalan di port yang sama dengan nilai `NEXT_PUBLIC_API_URL` di `frontend/.env.local`.

```bash
# Cek apakah backend merespons
curl http://localhost:8080/health
```

---

## Lisensi

MIT
