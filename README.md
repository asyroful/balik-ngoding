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

## Menjalankan Lokal

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

## Environment Variables

**Backend** (`backend/.env`):

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/balik_ngoding?sslmode=disable
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000
```

**Frontend** (`frontend/.env.local`):

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

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

## Docker (Backend)

```bash
cd backend
docker build -t balik-ngoding-backend .
docker run -p 8080:8080 --env-file .env balik-ngoding-backend
```

## Testing

```bash
# Frontend
cd frontend
npm test

# Backend
cd backend
go test ./...
```

## Lisensi

MIT
