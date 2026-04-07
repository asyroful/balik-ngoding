# Product Manager — LogicLab MVP

## Product Vision

Membangun platform latihan logika pemrograman dasar yang sederhana, fokus, dan efektif — membantu developer muda atau yang terlalu bergantung pada AI untuk kembali melatih kemampuan berpikir logis secara mandiri.

---

## Target User Persona

### Persona 1 — Fresh Graduate
- Usia: 21–25 tahun
- Baru lulus dari kampus atau bootcamp
- Punya pengetahuan teori tapi kurang jam terbang problem solving
- Ingin mempersiapkan diri untuk technical interview

### Persona 2 — AI-Dependent Developer
- Usia: 23–30 tahun
- Sudah bekerja 1–3 tahun
- Terbiasa copy-paste dari ChatGPT / GitHub Copilot
- Sadar kemampuan logika dasarnya melemah dan ingin melatih ulang

---

## Problem Statement

1. Banyak fresh graduate tidak terbiasa menulis kode dari nol tanpa bantuan AI
2. Developer yang terlalu bergantung AI kehilangan kemampuan debug dan problem solving mandiri
3. Platform seperti LeetCode terlalu kompleks dan intimidatif untuk pemula
4. Tidak ada platform yang fokus khusus pada logika dasar (loop, string, array, SQL basic) dengan feedback langsung

---

## Value Proposition

> "Latih logikamu, bukan copy-paste-mu."

- Soal yang relevan dan tidak overwhelming
- Feedback langsung: benar/salah + expected output
- Tidak perlu login untuk mulai latihan (MVP)
- Antarmuka bersih, tidak distraktif

---

## MVP Scope (Fitur Wajib)

| # | Fitur | Deskripsi |
|---|-------|-----------|
| 1 | List Soal | Tampilkan daftar soal berdasarkan kategori (loop, string, array, SQL) |
| 2 | Detail Soal | Tampilkan deskripsi soal, contoh input/output, dan constraints |
| 3 | Code Editor | Editor kode di browser (Monaco Editor) dengan syntax highlighting |
| 4 | Submit Jawaban | User submit kode, sistem evaluasi terhadap test case |
| 5 | Hasil Evaluasi | Tampilkan hasil: benar/salah, expected output vs actual output |

---

## Future Scope (Fitur Lanjutan)

- Sistem autentikasi & profil user
- Leaderboard & progress tracking
- Hint system (tampilkan petunjuk bertahap)
- Diskusi per soal (komentar)
- Soal dengan tingkat kesulitan adaptif
- Mode timer (simulasi interview)
- Dukungan multi-bahasa pemrograman (Python, Java, dll)
- Integrasi notifikasi harian (streak system)

---

## Success Metrics (MVP)

| Metrik | Target (bulan pertama) |
|--------|------------------------|
| Daily Active User (DAU) | 50 user/hari |
| Submission Rate | >60% user yang buka soal melakukan submit |
| Completion Rate | >40% submission mendapat status "Benar" |
| Bounce Rate (landing page) | <50% |
| Waktu rata-rata per sesi | >5 menit |

---

## Asumsi & Risiko

**Asumsi:**
- User tidak perlu login untuk mengerjakan soal di MVP
- Eksekusi kode menggunakan mock/sandbox sederhana di awal

**Risiko:**
- Eksekusi kode user di server bisa berbahaya → gunakan sandbox atau mock dulu
- Soal terlalu mudah/sulit → perlu kurasi konten yang baik sejak awal
