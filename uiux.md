# UI/UX Designer — LogicLab MVP

## Design Principles

1. **Simple** — Tidak ada elemen yang tidak perlu. Setiap pixel punya tujuan.
2. **Focus** — User harus bisa langsung mengerjakan soal tanpa distraksi.
3. **Distraction-free** — Tidak ada iklan, pop-up, atau notifikasi yang mengganggu.
4. **Fast feedback** — Hasil evaluasi muncul dalam hitungan detik.
5. **Accessible** — Kontras warna cukup, font readable, layout responsif.

---

## User Flow

```
Landing Page
    ↓
Pilih Kategori Soal (loop / string / array / SQL)
    ↓
List Soal (berdasarkan kategori)
    ↓
Klik Soal → Problem Detail Page
    ↓
Baca deskripsi + contoh input/output
    ↓
Tulis kode di Code Editor
    ↓
Klik tombol "Submit"
    ↓
Lihat Hasil Evaluasi (benar/salah + expected vs actual output)
    ↓
(Opsional) Coba lagi atau pilih soal lain
```

---

## Wireframe (Deskripsi Teks)

### Halaman 1 — Landing Page

```
┌─────────────────────────────────────────────────────┐
│  NAVBAR: [Logo LogicLab]                [GitHub]    │
├─────────────────────────────────────────────────────┤
│                                                     │
│         Latih Logikamu, Bukan Copy-Paste-mu         │
│    Platform latihan logika dasar untuk developer    │
│                                                     │
│              [ Mulai Latihan → ]                    │
│                                                     │
├─────────────────────────────────────────────────────┤
│  Kategori:  [Loop]  [String]  [Array]  [SQL Basic]  │
└─────────────────────────────────────────────────────┘
```

### Halaman 2 — Problem List Page

```
┌─────────────────────────────────────────────────────┐
│  NAVBAR: [Logo]              [Kategori ▼]           │
├─────────────────────────────────────────────────────┤
│  Filter: [Semua] [Loop] [String] [Array] [SQL]      │
├─────────────────────────────────────────────────────┤
│  #  │ Judul Soal              │ Kategori │ Difficulty│
│  1  │ FizzBuzz                │ Loop     │ Easy      │
│  2  │ Reverse String          │ String   │ Easy      │
│  3  │ Sum of Array            │ Array    │ Easy      │
│  4  │ Find Duplicate          │ Array    │ Medium    │
│  5  │ SELECT dengan WHERE     │ SQL      │ Easy      │
└─────────────────────────────────────────────────────┘
```

### Halaman 3 — Problem Detail + Code Editor

```
┌──────────────────────┬──────────────────────────────┐
│  PROBLEM DETAIL      │  CODE EDITOR                 │
│                      │                              │
│  # FizzBuzz          │  function solution(n) {      │
│                      │    // tulis kode di sini     │
│  Deskripsi:          │                              │
│  Tulis program yang  │  }                           │
│  mencetak angka 1    │                              │
│  sampai n...         ├──────────────────────────────┤
│                      │  RESULT PANEL                │
│  Contoh:             │                              │
│  Input: n = 15       │  ● Status: Menunggu...       │
│  Output:             │                              │
│  1, 2, Fizz, 4...    │  [ Submit Jawaban ]          │
│                      │                              │
│  Constraints:        │                              │
│  1 ≤ n ≤ 1000        │                              │
└──────────────────────┴──────────────────────────────┘
```

### Halaman 4 — Result Panel (setelah submit)

```
┌──────────────────────────────────────────────────────┐
│  HASIL EVALUASI                                      │
│                                                      │
│  ✅ Test Case 1: PASSED                              │
│     Input    : n = 5                                 │
│     Expected : 1 2 Fizz 4 Buzz                       │
│     Actual   : 1 2 Fizz 4 Buzz                       │
│                                                      │
│  ❌ Test Case 2: FAILED                              │
│     Input    : n = 15                                │
│     Expected : 1 2 Fizz 4 Buzz ... FizzBuzz          │
│     Actual   : 1 2 3 4 5 ...                         │
│                                                      │
│  Score: 1/2 test cases passed                        │
│                                                      │
│  [ Coba Lagi ]        [ Soal Berikutnya → ]          │
└──────────────────────────────────────────────────────┘
```

---

## Komponen Utama

### Navbar
- Logo di kiri
- Link kategori atau dropdown di kanan
- Tidak ada login button di MVP
- Sticky di atas, height minimal

### Problem List
- Tabel sederhana: nomor, judul, kategori, difficulty badge
- Filter tab di atas (Semua / Loop / String / Array / SQL)
- Klik baris → navigasi ke detail soal
- Tidak ada pagination di MVP (max 20–30 soal)

### Problem Detail
- Panel kiri: judul, deskripsi, contoh input/output, constraints
- Teks readable, font monospace untuk contoh kode
- Tidak ada sidebar tambahan

### Code Editor
- Monaco Editor (sama seperti VS Code)
- Default bahasa: JavaScript
- Minimal toolbar: hanya tombol "Submit"
- Tidak ada run tanpa submit di MVP

### Result Panel
- Muncul di bawah editor setelah submit
- Tampilkan per test case: status (✅/❌), input, expected, actual
- Skor total di bawah
- Tombol "Coba Lagi" dan "Soal Berikutnya"

---

## UX Rules

| Rule | Detail |
|------|--------|
| Minimal klik | Dari landing page ke mengerjakan soal maksimal 2 klik |
| Loading fast | Hasil evaluasi tampil < 3 detik |
| Clean layout | Tidak ada elemen dekoratif yang tidak perlu |
| No modal spam | Hindari pop-up konfirmasi yang tidak perlu |
| Error jelas | Jika submit gagal, tampilkan pesan error yang actionable |
| Mobile-friendly | Layout responsif, tapi prioritas desktop |

---

## Color Palette (Rekomendasi)

| Elemen | Warna |
|--------|-------|
| Background | `#0f172a` (dark) atau `#ffffff` (light) |
| Primary accent | `#6366f1` (indigo) |
| Success | `#22c55e` (green) |
| Error | `#ef4444` (red) |
| Text utama | `#f1f5f9` (dark mode) / `#1e293b` (light mode) |
| Border/divider | `#334155` |

---

## Referensi Inspirasi

| Platform | Yang Diambil |
|----------|-------------|
| LeetCode | Layout split panel (problem + editor), result panel |
| HackerRank | Kategori soal yang jelas, difficulty badge |
| Exercism | Tone yang supportif, tidak intimidatif |
| CodePen | Editor yang ringan dan langsung bisa dipakai |

---

## Catatan untuk Developer

- Gunakan layout CSS Grid atau Flexbox untuk split panel
- Monaco Editor sudah punya dark theme bawaan (`vs-dark`)
- Pastikan editor tidak overflow di layar kecil
- Result panel bisa pakai accordion atau slide-down animation sederhana
