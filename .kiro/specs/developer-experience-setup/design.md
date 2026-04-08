# Dokumen Desain: Developer Experience Setup

## Overview

Fitur ini memperbaiki pengalaman developer di project Balik Ngoding melalui tiga area utama:

1. **Docker Compose** — menjalankan seluruh stack (PostgreSQL + Go backend + Next.js frontend) dengan satu perintah `docker compose up`
2. **Steering Files** — file konteks Kiro yang akurat di `.kiro/steering/` menggantikan template lama yang berisi referensi NestJS/LogicLab
3. **MCP Configuration** — konfigurasi `postgres` MCP server di `.kiro/settings/mcp.json` agar Kiro dapat query langsung ke database lokal

Tidak ada perubahan pada logika bisnis aplikasi. Semua perubahan bersifat infrastruktur dan konfigurasi.

---

## Architecture

### Gambaran Sistem

```mermaid
graph TD
    subgraph "Host Machine"
        DEV[Developer]
        BROWSER[Browser :3000]
    end

    subgraph "Docker Network: balik-ngoding-net"
        FE[frontend\nNext.js :3000]
        BE[backend\nGo/Gin :8080]
        DB[(db\nPostgreSQL :5432)]
    end

    DEV -->|docker compose up| FE
    DEV -->|docker compose up| BE
    DEV -->|docker compose up| DB
    BROWSER --> FE
    FE -->|NEXT_PUBLIC_API_URL=http://localhost:8080| BE
    BE -->|DATABASE_URL| DB

    DB -->|health check OK| BE
    BE -->|health check OK| FE
```

### Urutan Startup

```mermaid
sequenceDiagram
    participant DC as Docker Compose
    participant DB as db (PostgreSQL)
    participant BE as backend (Go)
    participant FE as frontend (Next.js)

    DC->>DB: start container
    loop Health Check
        DC->>DB: pg_isready -U postgres
        DB-->>DC: not ready / ready
    end
    DC->>BE: start container (setelah db healthy)
    loop Health Check
        DC->>BE: GET /health
        BE-->>DC: not ready / {"status":"ok"}
    end
    DC->>FE: start container (setelah backend healthy)
```

### Struktur File yang Dibuat

```
balik-ngoding/
├── docker-compose.yml          # Konfigurasi utama (production-like)
├── docker-compose.override.yml # Override untuk development (hot reload)
├── .env.example                # Template env vars untuk Docker Compose
├── frontend/
│   └── Dockerfile              # Multi-stage build (deps → builder → runner)
└── .kiro/
    ├── settings/
    │   └── mcp.json            # Konfigurasi MCP postgres server
    └── steering/
        ├── project-overview.md # Gambaran umum project
        ├── backend.md          # Konteks Go/Gin backend
        └── frontend.md         # Konteks Next.js frontend
```

---

## Components and Interfaces

### 1. docker-compose.yml

File utama yang mendefinisikan tiga service dengan dependency ordering via `depends_on` + `condition: service_healthy`.

**Service `db`:**
- Image: `postgres:16-alpine`
- Port: `5432:5432`
- Named volume: `postgres_data`
- Health check: `pg_isready -U postgres`
- Env: `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`

**Service `backend`:**
- Build context: `./backend`
- Port: `8080:8080`
- Depends on: `db` (condition: `service_healthy`)
- Health check: `wget -qO- http://localhost:8080/health`
- Env: `DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN`

**Service `frontend`:**
- Build context: `./frontend`
- Port: `3000:3000`
- Depends on: `backend` (condition: `service_healthy`)
- Env: `NEXT_PUBLIC_API_URL`

**Network:** `balik-ngoding-net` (bridge driver) — semua service terhubung ke network ini sehingga dapat saling resolve via nama service sebagai hostname.

### 2. docker-compose.override.yml

Override untuk mode development. Secara otomatis di-merge oleh Docker Compose saat `docker compose up` tanpa flag tambahan.

- Mount `./backend:/app` untuk hot reload backend (dengan `air` atau `go run .`)
- Mount `./frontend:/app` dan `./frontend/node_modules:/app/node_modules` untuk hot reload frontend
- Override command ke `npm run dev` untuk frontend

### 3. frontend/Dockerfile

Multi-stage build dengan tiga stage:

```
Stage 1: deps
  - Base: node:20-alpine
  - Copy package.json + package-lock.json
  - RUN npm ci

Stage 2: builder
  - Base: node:20-alpine
  - Copy dari deps: node_modules
  - Copy source code
  - ARG NEXT_PUBLIC_API_URL (default: http://localhost:8080)
  - RUN npm run build

Stage 3: runner
  - Base: node:20-alpine
  - Copy dari builder: .next/standalone, .next/static, public
  - EXPOSE 3000
  - CMD ["node", "server.js"]
```

> Catatan: Next.js standalone output (`output: 'standalone'` di `next.config.js`) diperlukan agar stage runner dapat menggunakan `server.js` yang minimal.

### 4. .kiro/settings/mcp.json

Konfigurasi MCP dengan satu server: `postgres`.

```json
{
  "mcpServers": {
    "postgres": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-postgres", "${DATABASE_URL}"],
      "env": {
        "DATABASE_URL": "postgres://postgres:postgres@localhost:5432/balik_ngoding?sslmode=disable"
      }
    }
  }
}
```

Server `postgres` memberikan Kiro kemampuan untuk:
- Inspect schema database (tabel, kolom, tipe data)
- Menjalankan query SELECT untuk eksplorasi data dan debugging
- Memahami relasi antar tabel tanpa perlu membuka psql manual

> Keputusan desain: `filesystem` MCP tidak dikonfigurasi karena Kiro sudah memiliki tools bawaan untuk baca/tulis file yang lebih terintegrasi.

### 5. .kiro/steering/ Files

Tiga file steering dengan format markdown:

| File | Konten |
|---|---|
| `project-overview.md` | Nama project, tujuan, tech stack, cara run lokal |
| `backend.md` | Go 1.25, Gin, GORM, PostgreSQL, struktur `backend/internal/`, env vars |
| `frontend.md` | Next.js 14 App Router, TypeScript, Tailwind, Zustand, Monaco, env vars |

Setiap file menggunakan frontmatter `inclusion: always` agar selalu dibaca Kiro.

---

## Data Models

Fitur ini tidak mengubah data model aplikasi. Berikut adalah model konfigurasi yang relevan:

### Environment Variables

**Root `.env.example` (untuk Docker Compose):**

```env
# PostgreSQL
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=balik_ngoding

# Backend
DATABASE_URL=postgres://postgres:postgres@db:5432/balik_ngoding?sslmode=disable
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000

# Frontend
NEXT_PUBLIC_API_URL=http://localhost:8080
```

> Perhatikan: `DATABASE_URL` di Docker Compose menggunakan `db` sebagai hostname (nama service), bukan `localhost`.

**Perbedaan hostname per environment:**

| Environment | DATABASE_URL host |
|---|---|
| Lokal (tanpa Docker) | `localhost` |
| Docker Compose | `db` (nama service) |
| Production (Railway) | URL dari Railway env var |

### Docker Compose Service Dependencies

```
db (PostgreSQL)
  └── backend (Go/Gin)  [depends_on: db, condition: service_healthy]
        └── frontend (Next.js)  [depends_on: backend, condition: service_healthy]
```

### Health Check Configuration

| Service | Command | Interval | Timeout | Retries |
|---|---|---|---|---|
| db | `pg_isready -U postgres` | 5s | 5s | 5 |
| backend | `wget -qO- http://localhost:8080/health` | 10s | 5s | 5 |


---

## Correctness Properties

*A property is a characteristic or behavior that should hold true across all valid executions of a system — essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.*

### Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan

*Untuk setiap* service yang terdefinisi (`db`, `backend`, `frontend`), file `docker-compose.yml` harus mendefinisikan: nama service, port mapping yang sesuai, environment variables yang diperlukan, dependency ordering yang benar, dan named volume untuk PostgreSQL serta internal network.

**Validates: Requirements 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.12**

### Property 2: Setiap environment variable di docker-compose.yml memiliki nilai default

*Untuk setiap* environment variable yang didefinisikan di `docker-compose.yml`, variabel tersebut harus memiliki nilai default (menggunakan sintaks `${VAR:-default}`) sehingga stack dapat berjalan tanpa file `.env`.

**Validates: Requirements 1.11**

### Property 3: frontend/Dockerfile menggunakan multi-stage build yang benar

*Untuk setiap* stage di `frontend/Dockerfile` (`deps`, `builder`, `runner`), stage tersebut harus menggunakan base image `node:20-alpine`, berisi instruksi yang sesuai dengan tujuan stage, dan stage `runner` harus mengekspos port `3000` dengan nilai default `NEXT_PUBLIC_API_URL=http://localhost:8080`.

**Validates: Requirements 2.1, 2.2, 2.3, 2.4, 2.5**

### Property 4: Steering files tidak mengandung referensi ke stack lama

*Untuk setiap* file di `.kiro/steering/` dan *untuk setiap* kata kunci stack lama (`NestJS`, `TypeORM`, `LogicLab`, `nestjs`, `typeorm`), kata kunci tersebut tidak boleh ditemukan di dalam file tersebut.

**Validates: Requirements 3.5**

### Property 5: backend.md mendokumentasikan semua teknologi dan konfigurasi backend

*Untuk setiap* teknologi dan konfigurasi backend yang wajib didokumentasikan (`Go`, `Gin`, `GORM`, `PostgreSQL`, `DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN`, `backend/internal`), kata kunci tersebut harus ditemukan di file `.kiro/steering/backend.md`.

**Validates: Requirements 3.2, 3.7**

### Property 6: frontend.md mendokumentasikan semua teknologi dan konfigurasi frontend

*Untuk setiap* teknologi dan konfigurasi frontend yang wajib didokumentasikan (`Next.js`, `TypeScript`, `Tailwind`, `Zustand`, `Monaco`, `NEXT_PUBLIC_API_URL`, `App Router`), kata kunci tersebut harus ditemukan di file `.kiro/steering/frontend.md`.

**Validates: Requirements 3.3, 3.8**

### Property 7: mcp.json adalah JSON valid dengan konfigurasi postgres server yang lengkap

*Untuk setiap* field yang diperlukan di konfigurasi MCP (`mcpServers`, `postgres`, `command`, `args`), field tersebut harus ada di file `.kiro/settings/mcp.json` yang merupakan JSON valid (dapat di-parse tanpa error).

**Validates: Requirements 4.1, 4.3, 4.5**

### Property 8: README berisi semua konten setup yang diperlukan

*Untuk setiap* konten yang wajib ada di README (`docker compose up`, `docker compose down`, `docker compose logs`, `npm test`, `go test`, `Troubleshooting`, instruksi manual setup), konten tersebut harus ditemukan di file `README.md`.

**Validates: Requirements 5.1, 5.2, 5.3, 5.4, 5.5, 5.6**

---

## Error Handling

### Docker Compose

| Skenario Error | Penanganan |
|---|---|
| Port sudah digunakan (5432/8080/3000) | Docker Compose gagal dengan pesan error jelas. README menyediakan solusi di seksi Troubleshooting. |
| Database belum siap saat backend start | Health check dengan retry (5x, interval 5s) mencegah backend start prematur. |
| Backend belum siap saat frontend start | Health check dengan retry (5x, interval 10s) mencegah frontend start prematur. |
| File `.env` tidak ada | Nilai default di `docker-compose.yml` digunakan, tidak ada error. |
| Image build gagal | Docker Compose menampilkan build log. Developer perlu memeriksa Dockerfile. |

### Frontend Dockerfile

| Skenario Error | Penanganan |
|---|---|
| `npm ci` gagal (dependency conflict) | Build gagal di stage `deps` dengan pesan error dari npm. |
| `npm run build` gagal (TypeScript error) | Build gagal di stage `builder`. Developer perlu fix kode sebelum build ulang. |
| `NEXT_PUBLIC_API_URL` tidak di-set | Nilai default `http://localhost:8080` digunakan via `ARG` di Dockerfile. |
| `next.config.js` tidak memiliki `output: 'standalone'` | Stage `runner` tidak menemukan `server.js`. Perlu tambah konfigurasi di `next.config.js`. |

### MCP Configuration

| Skenario Error | Penanganan |
|---|---|
| PostgreSQL tidak berjalan saat MCP digunakan | MCP server gagal connect. Kiro menampilkan error koneksi. Developer perlu jalankan database. |
| `DATABASE_URL` salah di mcp.json | MCP server gagal connect dengan pesan authentication error. |

---

## Testing Strategy

Fitur ini adalah konfigurasi infrastruktur, bukan logika bisnis. Strategi testing berfokus pada verifikasi struktural file konfigurasi.

### Unit Tests (Verifikasi Struktural)

Unit tests memverifikasi bahwa file konfigurasi yang dibuat memiliki konten yang benar. Tests ini bersifat deterministik dan tidak memerlukan Docker daemon atau database yang berjalan.

**Scope:**
- Parse dan validasi `docker-compose.yml` (YAML valid, service terdefinisi, port mapping, env vars, health checks)
- Parse dan validasi `frontend/Dockerfile` (stage names, base images, EXPOSE, ARG defaults)
- Parse dan validasi `.kiro/settings/mcp.json` (JSON valid, postgres server terkonfigurasi)
- Verifikasi konten steering files (kata kunci yang harus ada, kata kunci yang tidak boleh ada)
- Verifikasi konten `README.md` (perintah yang harus ada, seksi yang harus ada)

### Property-Based Tests

Property-based testing digunakan untuk memverifikasi properties yang berlaku secara universal. Library yang digunakan: **fast-check** (sudah ada di `frontend/package.json` sebagai devDependency).

> Catatan: Karena fitur ini adalah konfigurasi file statis (bukan fungsi dengan input dinamis), sebagian besar properties diimplementasikan sebagai parameterized tests yang iterasi atas daftar kata kunci/field yang diperlukan — bukan random input generation. Ini tetap menggunakan fast-check untuk konsistensi dengan testing strategy project.

**Konfigurasi:** Minimum 100 iterasi per property test.

**Tag format:** `Feature: developer-experience-setup, Property {N}: {property_text}`

#### Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan
```
// Feature: developer-experience-setup, Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan
// Untuk setiap service yang terdefinisi, semua konfigurasi wajib harus ada
fc.assert(fc.property(
  fc.constantFrom('db', 'backend', 'frontend'),
  (service) => { /* verifikasi service ada di parsed YAML */ }
))
```

#### Property 2: Setiap env var memiliki nilai default
```
// Feature: developer-experience-setup, Property 2: Setiap env var memiliki nilai default
// Untuk setiap env var di docker-compose.yml, harus ada nilai default
fc.assert(fc.property(
  fc.constantFrom('DATABASE_URL', 'PORT', 'FRONTEND_ORIGIN', 'NEXT_PUBLIC_API_URL'),
  (envVar) => { /* verifikasi format ${VAR:-default} ada */ }
))
```

#### Property 3: frontend/Dockerfile multi-stage build yang benar
```
// Feature: developer-experience-setup, Property 3: frontend/Dockerfile multi-stage build yang benar
// Untuk setiap stage, instruksi yang diperlukan harus ada
fc.assert(fc.property(
  fc.constantFrom('deps', 'builder', 'runner'),
  (stage) => { /* verifikasi stage ada dengan instruksi yang benar */ }
))
```

#### Property 4: Steering files tidak mengandung referensi stack lama
```
// Feature: developer-experience-setup, Property 4: Steering files tidak mengandung referensi stack lama
// Untuk setiap file steering dan setiap kata kunci terlarang, kata kunci tidak boleh ada
fc.assert(fc.property(
  fc.constantFrom('NestJS', 'TypeORM', 'LogicLab', 'nestjs', 'typeorm'),
  fc.constantFrom('backend.md', 'frontend.md', 'project-overview.md'),
  (keyword, file) => { /* verifikasi keyword tidak ada di file */ }
))
```

#### Property 5: backend.md mendokumentasikan semua teknologi backend
```
// Feature: developer-experience-setup, Property 5: backend.md mendokumentasikan semua teknologi backend
// Untuk setiap kata kunci wajib, harus ada di backend.md
fc.assert(fc.property(
  fc.constantFrom('Go', 'Gin', 'GORM', 'PostgreSQL', 'DATABASE_URL', 'PORT', 'FRONTEND_ORIGIN'),
  (keyword) => { /* verifikasi keyword ada di backend.md */ }
))
```

#### Property 6: frontend.md mendokumentasikan semua teknologi frontend
```
// Feature: developer-experience-setup, Property 6: frontend.md mendokumentasikan semua teknologi frontend
fc.assert(fc.property(
  fc.constantFrom('Next.js', 'TypeScript', 'Tailwind', 'Zustand', 'Monaco', 'NEXT_PUBLIC_API_URL'),
  (keyword) => { /* verifikasi keyword ada di frontend.md */ }
))
```

#### Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap
```
// Feature: developer-experience-setup, Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap
// Untuk setiap field wajib, harus ada di mcp.json
fc.assert(fc.property(
  fc.constantFrom('mcpServers', 'postgres', 'command', 'args'),
  (field) => { /* verifikasi field ada di parsed JSON */ }
))
```

#### Property 8: README berisi semua konten setup yang diperlukan
```
// Feature: developer-experience-setup, Property 8: README berisi semua konten setup yang diperlukan
fc.assert(fc.property(
  fc.constantFrom('docker compose up', 'docker compose down', 'docker compose logs', 'npm test', 'go test', 'Troubleshooting'),
  (content) => { /* verifikasi content ada di README.md */ }
))
```

### Integration Tests

Integration tests tidak diperlukan untuk fitur ini karena tidak ada logika bisnis yang diubah. Verifikasi end-to-end (apakah `docker compose up` benar-benar berhasil) dilakukan secara manual oleh developer saat implementasi.

### Test File Location

```
frontend/__tests__/
└── developerExperienceSetup.test.ts   # Unit + property tests untuk semua konfigurasi
```
