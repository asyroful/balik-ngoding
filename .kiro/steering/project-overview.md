---
inclusion: always
---

# Balik Ngoding — Project Overview

## Nama Project

**Balik Ngoding** — Platform latihan coding berbasis web untuk belajar pemrograman dengan soal-soal interaktif.

## Tujuan

Balik Ngoding memungkinkan developer (terutama pemula) untuk berlatih coding melalui soal-soal yang dikategorikan berdasarkan topik (loop, string, array, SQL). User dapat menulis kode di browser, submit, dan langsung mendapatkan feedback hasil evaluasi.

## Tech Stack Keseluruhan

| Layer | Teknologi |
|---|---|
| Backend | Go 1.25, Gin framework, GORM, PostgreSQL |
| Frontend | Next.js 14 (App Router), TypeScript, Tailwind CSS, Zustand, Monaco Editor |
| Database | PostgreSQL 16 |
| Containerization | Docker, Docker Compose |

## Struktur Repository

```
balik-ngoding/
├── backend/                    # Go/Gin backend service
│   ├── internal/               # Internal packages
│   │   ├── database/           # Database connection & seeding
│   │   ├── evaluator/          # SQL/code evaluator logic
│   │   ├── models/             # GORM models
│   │   ├── problems/           # Problems handler & service
│   │   └── submissions/        # Submissions handler & service
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── frontend/                   # Next.js 14 frontend
│   ├── app/                    # App Router pages
│   ├── components/             # React components
│   ├── store/                  # Zustand stores
│   ├── lib/                    # API client & types
│   └── Dockerfile
└── docker-compose.yml          # Full stack orchestration
```

## Cara Menjalankan Lokal

### Dengan Docker Compose (Direkomendasikan)

```bash
# Clone repository
git clone <repo-url>
cd balik-ngoding

# Jalankan seluruh stack
docker compose up

# Akses aplikasi di http://localhost:3000
```

### Tanpa Docker (Manual)

**Backend:**
```bash
cd backend
# Pastikan PostgreSQL berjalan di localhost:5432
cp .env.example .env  # sesuaikan DATABASE_URL
go run main.go
```

**Frontend:**
```bash
cd frontend
cp .env.example .env.local  # sesuaikan NEXT_PUBLIC_API_URL
npm install
npm run dev
```

## Port Default

| Service | Port |
|---|---|
| Frontend (Next.js) | 3000 |
| Backend (Go/Gin) | 8080 |
| Database (PostgreSQL) | 5432 |
