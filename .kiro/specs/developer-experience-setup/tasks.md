# Rencana Implementasi: Developer Experience Setup

## Overview

Implementasi infrastruktur dan konfigurasi untuk memperbaiki pengalaman developer di project Balik Ngoding: Docker Compose full stack, steering files Kiro yang akurat, konfigurasi MCP postgres, dan pembaruan README.

## Tasks

- [x] 1. Buat steering files di .kiro/steering/
  - Buat direktori `.kiro/steering/` jika belum ada
  - Buat file `.kiro/steering/project-overview.md` dengan frontmatter `inclusion: always`, berisi nama project (Balik Ngoding), tujuan, tech stack keseluruhan, dan cara menjalankan lokal
  - Buat file `.kiro/steering/backend.md` dengan frontmatter `inclusion: always`, berisi stack Go 1.25, Gin, GORM, PostgreSQL, struktur direktori `backend/internal/`, dan env vars (`DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN`)
  - Buat file `.kiro/steering/frontend.md` dengan frontmatter `inclusion: always`, berisi stack Next.js 14 App Router, TypeScript, Tailwind CSS, Zustand, Monaco Editor, struktur direktori `frontend/`, dan env var `NEXT_PUBLIC_API_URL`
  - Pastikan tidak ada referensi ke NestJS, TypeORM, LogicLab, atau struktur direktori `src/` yang tidak ada di project ini
  - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8_

  - [x] 1.1 Tulis property test untuk steering files (Property 4, 5, 6)
    - **Property 4: Steering files tidak mengandung referensi stack lama**
    - **Validates: Requirements 3.5**
    - **Property 5: backend.md mendokumentasikan semua teknologi backend**
    - **Validates: Requirements 3.2, 3.7**
    - **Property 6: frontend.md mendokumentasikan semua teknologi frontend**
    - **Validates: Requirements 3.3, 3.8**
    - Buat file `frontend/__tests__/developerExperienceSetup.test.ts`
    - Gunakan `fast-check` dengan `fc.constantFrom` untuk iterasi kata kunci
    - Tag: `Feature: developer-experience-setup, Property 4/5/6`

- [x] 2. Hapus steering files lama di root project
  - Hapus file `backend.md`, `frontend.md`, `pm.md`, `qa.md`, `uiux.md` dari root project setelah steering files baru di `.kiro/steering/` sudah dibuat di task 1
  - _Requirements: 3.5_

- [x] 3. Buat konfigurasi MCP di .kiro/settings/mcp.json
  - Buat direktori `.kiro/settings/` jika belum ada
  - Buat file `.kiro/settings/mcp.json` dengan konfigurasi `postgres` MCP server menggunakan `npx @modelcontextprotocol/server-postgres`
  - Sertakan `DATABASE_URL` default `postgres://postgres:postgres@localhost:5432/balik_ngoding?sslmode=disable` di env
  - Hanya konfigurasi server `postgres` (tidak perlu `filesystem`)
  - _Requirements: 4.1, 4.3, 4.5_

  - [x] 3.1 Tulis property test untuk mcp.json (Property 7)
    - **Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap**
    - **Validates: Requirements 4.1, 4.3, 4.5**
    - Tambahkan ke `frontend/__tests__/developerExperienceSetup.test.ts`
    - Verifikasi field `mcpServers`, `postgres`, `command`, `args` ada di parsed JSON
    - Tag: `Feature: developer-experience-setup, Property 7`

- [x] 4. Tambahkan output standalone ke next.config.js
  - Edit `frontend/next.config.js`, tambahkan `output: 'standalone'` ke konfigurasi `nextConfig`
  - Ini diperlukan agar stage `runner` di Dockerfile frontend dapat menggunakan `server.js` yang minimal
  - _Requirements: 2.1, 2.2_

- [x] 5. Buat Dockerfile untuk frontend
  - Buat file `frontend/Dockerfile` dengan multi-stage build tiga stage: `deps`, `builder`, `runner`
  - Stage `deps`: base `node:20-alpine`, copy `package.json` + `package-lock.json`, jalankan `npm ci`
  - Stage `builder`: base `node:20-alpine`, copy `node_modules` dari `deps`, copy source code, definisikan `ARG NEXT_PUBLIC_API_URL` dengan default `http://localhost:8080`, jalankan `npm run build`
  - Stage `runner`: base `node:20-alpine`, copy `.next/standalone`, `.next/static`, dan `public` dari `builder`, ekspos port `3000`, jalankan `node server.js`
  - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5_

  - [x] 5.1 Tulis property test untuk frontend/Dockerfile (Property 3)
    - **Property 3: frontend/Dockerfile multi-stage build yang benar**
    - **Validates: Requirements 2.1, 2.2, 2.3, 2.4, 2.5**
    - Tambahkan ke `frontend/__tests__/developerExperienceSetup.test.ts`
    - Gunakan `fc.constantFrom('deps', 'builder', 'runner')` untuk iterasi stage
    - Tag: `Feature: developer-experience-setup, Property 3`

- [x] 6. Buat docker-compose.yml di root project
  - Buat file `docker-compose.yml` di root project dengan tiga service: `db`, `backend`, `frontend`
  - Service `db`: image `postgres:16-alpine`, port `5432:5432`, named volume `postgres_data`, health check `pg_isready -U postgres` (interval 5s, timeout 5s, retries 5), env vars dengan nilai default via sintaks `${VAR:-default}`
  - Service `backend`: build context `./backend`, port `8080:8080`, depends on `db` (condition: `service_healthy`), health check `wget -qO- http://localhost:8080/health` (interval 10s, timeout 5s, retries 5), env vars `DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN` dengan nilai default
  - Service `frontend`: build context `./frontend`, port `3000:3000`, depends on `backend` (condition: `service_healthy`), env var `NEXT_PUBLIC_API_URL` dengan nilai default
  - Definisikan network `balik-ngoding-net` (bridge) dan named volume `postgres_data`
  - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9_

  - [x] 6.1 Tulis property test untuk docker-compose.yml (Property 1, 2)
    - **Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan**
    - **Validates: Requirements 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.12**
    - **Property 2: Setiap env var memiliki nilai default**
    - **Validates: Requirements 1.11**
    - Tambahkan ke `frontend/__tests__/developerExperienceSetup.test.ts`
    - Parse YAML menggunakan library `js-yaml` atau baca sebagai string dan verifikasi pola `${VAR:-default}`
    - Tag: `Feature: developer-experience-setup, Property 1/2`

- [x] 7. Checkpoint — Pastikan semua tests pass
  - Pastikan semua tests pass, tanyakan ke user jika ada pertanyaan.

- [x] 8. Buat docker-compose.override.yml untuk mode development
  - Buat file `docker-compose.override.yml` di root project
  - Mount `./backend:/app` untuk hot reload backend
  - Mount `./frontend:/app` dan `./frontend/node_modules:/app/node_modules` untuk hot reload frontend
  - Override command frontend ke `npm run dev`
  - _Requirements: 1.10_

- [x] 9. Buat .env.example di root project
  - Buat file `.env.example` di root project yang mendokumentasikan semua environment variable yang dapat di-override untuk Docker Compose
  - Sertakan variabel: `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN`, `NEXT_PUBLIC_API_URL`
  - Gunakan nilai default yang sesuai untuk environment lokal (hostname `db` untuk `DATABASE_URL` di Docker Compose)
  - _Requirements: 1.12_

- [x] 10. Perbarui README.md
  - Tambahkan seksi "Quick Start (Docker Compose)" sebagai metode setup utama yang direkomendasikan, dengan langkah-langkah dari clone hingga aplikasi berjalan (maksimal 5 langkah)
  - Dokumentasikan perintah umum: `docker compose up`, `docker compose down`, `docker compose logs -f [service]`
  - Pertahankan instruksi setup manual (tanpa Docker) sebagai alternatif
  - Dokumentasikan cara menjalankan test: `npm test` untuk frontend dan `go test ./...` untuk backend
  - Tambahkan seksi "Troubleshooting" dengan solusi untuk masalah umum (port sudah digunakan, dll.)
  - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5, 5.6_

  - [x] 10.1 Tulis property test untuk README.md (Property 8)
    - **Property 8: README berisi semua konten setup yang diperlukan**
    - **Validates: Requirements 5.1, 5.2, 5.3, 5.4, 5.5, 5.6**
    - Tambahkan ke `frontend/__tests__/developerExperienceSetup.test.ts`
    - Gunakan `fc.constantFrom('docker compose up', 'docker compose down', 'docker compose logs', 'npm test', 'go test', 'Troubleshooting')` untuk iterasi konten wajib
    - Tag: `Feature: developer-experience-setup, Property 8`

- [x] 11. Checkpoint akhir — Pastikan semua tests pass
  - Pastikan semua tests pass, tanyakan ke user jika ada pertanyaan.

## Catatan

- Task bertanda `*` bersifat opsional dan dapat dilewati untuk MVP yang lebih cepat
- Setiap task mereferensikan requirements spesifik untuk traceability
- Hapus file lama di root (`backend.md`, `frontend.md`, `pm.md`, `qa.md`, `uiux.md`) di task 2, setelah steering files baru di `.kiro/steering/` selesai dibuat di task 1
- `output: 'standalone'` di `next.config.js` (task 4) harus dilakukan sebelum membuat Dockerfile frontend (task 5)
- Property tests menggunakan `fast-check` yang sudah tersedia sebagai devDependency di `frontend/package.json`
- Semua property tests ditempatkan di satu file: `frontend/__tests__/developerExperienceSetup.test.ts`
