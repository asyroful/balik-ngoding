# Dokumen Requirements

## Pendahuluan

Fitur ini bertujuan memperbaiki pengalaman developer saat setup dan bekerja di project Balik Ngoding. Saat ini developer harus menjalankan PostgreSQL, backend Go, dan frontend Next.js secara manual dan terpisah. Selain itu, Kiro tidak memiliki konteks project yang akurat karena steering files masih berisi template lama (NestJS/LogicLab), dan MCP belum dikonfigurasi.

Fitur ini mencakup tiga area perbaikan:
1. Docker Compose untuk menjalankan seluruh stack (PostgreSQL + backend Go + frontend Next.js) dengan satu perintah
2. Steering files yang akurat di `.kiro/steering/` yang mencerminkan stack aktual (Go/Gin, Next.js 14, TypeScript)
3. Konfigurasi MCP dasar di `.kiro/settings/mcp.json`

## Glosarium

- **Docker_Compose**: Alat orkestrasi container yang mendefinisikan dan menjalankan multi-container Docker application via file `docker-compose.yml`
- **Stack**: Keseluruhan layanan yang dibutuhkan untuk menjalankan Balik Ngoding secara lokal — PostgreSQL, backend Go, dan frontend Next.js
- **Steering_File**: File markdown di `.kiro/steering/` yang memberikan konteks project kepada Kiro agar menghasilkan kode yang sesuai dengan stack aktual
- **MCP**: Model Context Protocol — mekanisme konfigurasi di `.kiro/settings/mcp.json` untuk mengintegrasikan tools eksternal ke dalam Kiro
- **Backend**: Layanan Go/Gin yang berjalan di port 8080, terhubung ke PostgreSQL
- **Frontend**: Aplikasi Next.js 14 yang berjalan di port 3000, berkomunikasi dengan Backend
- **Database**: Instance PostgreSQL yang menyimpan data soal, test case, dan submission
- **Health_Check**: Mekanisme Docker untuk memverifikasi bahwa sebuah service sudah siap menerima koneksi sebelum service dependen dijalankan
- **Hot_Reload**: Kemampuan service untuk mendeteksi perubahan kode dan me-restart otomatis tanpa perlu restart manual

---

## Requirements

### Requirement 1: Docker Compose Full Stack

**User Story:** Sebagai developer, saya ingin menjalankan seluruh stack (database, backend, frontend) dengan satu perintah, agar saya tidak perlu setup manual setiap kali mulai bekerja.

#### Acceptance Criteria

1. THE Docker_Compose SHALL mendefinisikan tiga service: `db` (PostgreSQL), `backend` (Go/Gin), dan `frontend` (Next.js) dalam satu file `docker-compose.yml` di root project.
2. WHEN developer menjalankan `docker compose up`, THE Docker_Compose SHALL menjalankan ketiga service secara berurutan sesuai dependency (db → backend → frontend).
3. WHEN service `db` belum siap menerima koneksi, THE Docker_Compose SHALL menahan service `backend` agar tidak dimulai sebelum Health_Check Database berhasil.
4. WHEN service `backend` belum siap menerima koneksi, THE Docker_Compose SHALL menahan service `frontend` agar tidak dimulai sebelum Health_Check Backend berhasil.
5. THE Docker_Compose SHALL memetakan port `5432` untuk Database, port `8080` untuk Backend, dan port `3000` untuk Frontend ke host machine.
6. THE Docker_Compose SHALL menyediakan environment variable `DATABASE_URL`, `PORT`, dan `FRONTEND_ORIGIN` ke service `backend` dengan nilai default yang sesuai untuk environment lokal.
7. THE Docker_Compose SHALL menyediakan environment variable `NEXT_PUBLIC_API_URL` ke service `frontend` dengan nilai `http://localhost:8080`.
8. THE Docker_Compose SHALL mendefinisikan named volume untuk data PostgreSQL agar data tidak hilang saat container di-restart.
9. THE Docker_Compose SHALL mendefinisikan Docker network internal agar service dapat berkomunikasi satu sama lain menggunakan nama service sebagai hostname.
10. WHERE developer ingin mode development dengan Hot_Reload, THE Docker_Compose SHALL menyediakan konfigurasi `docker-compose.override.yml` yang me-mount source code sebagai volume.
11. IF file `.env` tidak ditemukan di root project, THEN THE Docker_Compose SHALL menggunakan nilai default yang sudah didefinisikan di `docker-compose.yml` tanpa error.
12. THE Docker_Compose SHALL menyertakan file `.env.example` di root project yang mendokumentasikan semua environment variable yang dapat di-override.

### Requirement 2: Dockerfile Frontend

**User Story:** Sebagai developer, saya ingin Frontend dapat di-containerize, agar dapat dijalankan via Docker Compose bersama service lainnya.

#### Acceptance Criteria

1. THE Frontend SHALL memiliki `Dockerfile` di direktori `frontend/` yang menggunakan multi-stage build (stage `deps`, `builder`, `runner`).
2. WHEN Docker membangun image Frontend, THE Dockerfile SHALL menginstall dependencies di stage `deps`, melakukan build Next.js di stage `builder`, dan menjalankan production server di stage `runner`.
3. THE Dockerfile SHALL menggunakan base image `node:20-alpine` untuk meminimalkan ukuran image.
4. THE Dockerfile SHALL mengekspos port `3000` sebagai port default Next.js.
5. IF environment variable `NEXT_PUBLIC_API_URL` tidak di-set saat build, THEN THE Dockerfile SHALL menggunakan nilai default `http://localhost:8080`.

### Requirement 3: Steering Files Akurat

**User Story:** Sebagai developer yang menggunakan Kiro, saya ingin Kiro memiliki konteks project yang akurat, agar saran dan kode yang dihasilkan sesuai dengan stack aktual (Go/Gin, bukan NestJS).

#### Acceptance Criteria

1. THE Steering_File SHALL dibuat di direktori `.kiro/steering/` dengan file terpisah per domain: `backend.md`, `frontend.md`, `project-overview.md`.
2. THE Steering_File `backend.md` SHALL mendokumentasikan stack aktual: Go 1.25, Gin framework, GORM, PostgreSQL, dengan struktur direktori `backend/internal/` yang benar.
3. THE Steering_File `frontend.md` SHALL mendokumentasikan stack aktual: Next.js 14 App Router, TypeScript, Tailwind CSS, Zustand, Monaco Editor, dengan struktur direktori `frontend/` yang benar.
4. THE Steering_File `project-overview.md` SHALL mendokumentasikan gambaran umum project: nama project (Balik Ngoding), tujuan, tech stack keseluruhan, dan cara menjalankan lokal.
5. THE Steering_File SHALL tidak mengandung referensi ke stack lama: NestJS, TypeORM, LogicLab, atau `src/` directory structure yang tidak ada di project ini.
6. WHEN Kiro membaca Steering_File, THE Steering_File SHALL memberikan informasi yang cukup agar Kiro menghasilkan kode Go (bukan TypeScript/NestJS) untuk perubahan di direktori `backend/`.
7. THE Steering_File `backend.md` SHALL mendokumentasikan environment variables yang digunakan backend: `DATABASE_URL`, `PORT`, `FRONTEND_ORIGIN`.
8. THE Steering_File `frontend.md` SHALL mendokumentasikan environment variables yang digunakan frontend: `NEXT_PUBLIC_API_URL`.

### Requirement 4: Konfigurasi MCP Dasar

**User Story:** Sebagai developer, saya ingin MCP dikonfigurasi di project, agar Kiro dapat menggunakan tools eksternal yang relevan untuk membantu development.

#### Acceptance Criteria

1. THE MCP SHALL memiliki file konfigurasi di `.kiro/settings/mcp.json` dengan struktur JSON yang valid.
2. THE MCP SHALL mengkonfigurasi server `filesystem` untuk memberikan Kiro akses baca/tulis ke direktori project.
3. THE MCP SHALL mengkonfigurasi server `postgres` untuk memberikan Kiro kemampuan query langsung ke Database lokal (read-only) guna membantu debugging dan eksplorasi schema.
4. IF file `.kiro/settings/mcp.json` sudah ada, THEN THE MCP SHALL mempertahankan konfigurasi yang sudah ada dan hanya menambahkan server yang belum terdefinisi.
5. THE MCP SHALL menyertakan komentar atau dokumentasi inline (via field `description` jika didukung) yang menjelaskan tujuan setiap server MCP yang dikonfigurasi.

### Requirement 5: Dokumentasi Setup Developer

**User Story:** Sebagai developer baru yang bergabung ke project, saya ingin ada dokumentasi setup yang jelas dan akurat, agar saya bisa mulai berkontribusi dalam waktu singkat.

#### Acceptance Criteria

1. THE README SHALL diperbarui dengan instruksi setup menggunakan Docker Compose sebagai metode utama yang direkomendasikan.
2. WHEN developer mengikuti instruksi di README, THE README SHALL memandu developer dari clone repository hingga aplikasi berjalan di browser dalam kurang dari 5 langkah.
3. THE README SHALL tetap mempertahankan instruksi setup manual (tanpa Docker) sebagai alternatif untuk developer yang tidak menggunakan Docker.
4. THE README SHALL mendokumentasikan perintah-perintah umum: `docker compose up`, `docker compose down`, `docker compose logs -f [service]`.
5. THE README SHALL mendokumentasikan cara menjalankan test: `npm test` untuk frontend dan `go test ./...` untuk backend.
6. IF terjadi masalah umum saat setup (misalnya port sudah digunakan), THEN THE README SHALL menyediakan seksi "Troubleshooting" dengan solusi untuk masalah tersebut.
