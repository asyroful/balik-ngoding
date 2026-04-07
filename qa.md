# QA Tester — LogicLab MVP

## Test Strategy

Untuk MVP, fokus pada **manual testing** terlebih dahulu. Tidak perlu automated test end-to-end di tahap ini — cukup pastikan semua happy path dan edge case utama sudah diverifikasi secara manual sebelum release.

### Prioritas Testing

1. **Functional testing** — fitur berjalan sesuai spesifikasi
2. **Edge case testing** — input tidak normal / ekstrem
3. **Error handling testing** — sistem gagal dengan baik (graceful)
4. **UI/UX sanity check** — tampilan tidak rusak, flow tidak membingungkan

---

## Test Case Utama

### TC-01: Load Daftar Soal

| Field | Detail |
|-------|--------|
| ID | TC-01 |
| Nama | Load Problem List |
| Precondition | Aplikasi berjalan, database berisi minimal 5 soal |
| Steps | 1. Buka `/problems` |
| Expected | Daftar soal tampil dengan judul, kategori, dan difficulty |
| Status | [ ] Pass / [ ] Fail |

---

### TC-02: Filter Soal Berdasarkan Kategori

| Field | Detail |
|-------|--------|
| ID | TC-02 |
| Nama | Filter by Category |
| Precondition | Halaman `/problems` sudah terbuka |
| Steps | 1. Klik tab "Loop" |
| Expected | Hanya soal dengan kategori "loop" yang tampil |
| Status | [ ] Pass / [ ] Fail |

---

### TC-03: Load Detail Soal

| Field | Detail |
|-------|--------|
| ID | TC-03 |
| Nama | Load Problem Detail |
| Precondition | Halaman list soal sudah terbuka |
| Steps | 1. Klik salah satu soal |
| Expected | Halaman detail tampil: judul, deskripsi, contoh input/output, code editor |
| Status | [ ] Pass / [ ] Fail |

---

### TC-04: Submit Jawaban Benar

| Field | Detail |
|-------|--------|
| ID | TC-04 |
| Nama | Submit Correct Answer |
| Precondition | Berada di halaman detail soal "FizzBuzz" |
| Steps | 1. Tulis kode yang benar di editor<br>2. Klik "Submit Jawaban" |
| Expected | Semua test case passed, status "Accepted", score = total |
| Status | [ ] Pass / [ ] Fail |

**Contoh kode benar untuk FizzBuzz:**
```javascript
function solution(n) {
  const result = [];
  for (let i = 1; i <= n; i++) {
    if (i % 15 === 0) result.push('FizzBuzz');
    else if (i % 3 === 0) result.push('Fizz');
    else if (i % 5 === 0) result.push('Buzz');
    else result.push(String(i));
  }
  return result.join(' ');
}
```

---

### TC-05: Submit Jawaban Salah

| Field | Detail |
|-------|--------|
| ID | TC-05 |
| Nama | Submit Wrong Answer |
| Precondition | Berada di halaman detail soal |
| Steps | 1. Tulis kode yang salah (misal: return nilai hardcoded)<br>2. Klik "Submit Jawaban" |
| Expected | Beberapa/semua test case failed, status "Wrong Answer", tampil expected vs actual output |
| Status | [ ] Pass / [ ] Fail |

---

### TC-06: Submit Kode Kosong

| Field | Detail |
|-------|--------|
| ID | TC-06 |
| Nama | Submit Empty Code |
| Precondition | Berada di halaman detail soal, editor kosong |
| Steps | 1. Hapus semua kode di editor<br>2. Klik "Submit Jawaban" |
| Expected | Tombol submit disabled ATAU muncul pesan validasi "Kode tidak boleh kosong" |
| Status | [ ] Pass / [ ] Fail |

---

### TC-07: Submit Kode dengan Syntax Error

| Field | Detail |
|-------|--------|
| ID | TC-07 |
| Nama | Submit Syntax Error Code |
| Precondition | Berada di halaman detail soal |
| Steps | 1. Tulis kode dengan syntax error (misal: `function solution( {`)<br>2. Klik "Submit Jawaban" |
| Expected | Status "Error", pesan error ditampilkan, tidak crash |
| Status | [ ] Pass / [ ] Fail |

---

### TC-08: Submit Kode dengan Infinite Loop

| Field | Detail |
|-------|--------|
| ID | TC-08 |
| Nama | Submit Infinite Loop |
| Precondition | Berada di halaman detail soal |
| Steps | 1. Tulis kode dengan infinite loop (`while(true){}`)<br>2. Klik "Submit Jawaban" |
| Expected | Evaluasi timeout setelah beberapa detik, tampil pesan "Waktu eksekusi habis" |
| Status | [ ] Pass / [ ] Fail |

---

### TC-09: API Backend Tidak Tersedia

| Field | Detail |
|-------|--------|
| ID | TC-09 |
| Nama | Backend Unavailable |
| Precondition | Backend server dimatikan |
| Steps | 1. Buka halaman `/problems` |
| Expected | Tampil pesan error yang informatif, bukan halaman blank atau crash |
| Status | [ ] Pass / [ ] Fail |

---

### TC-10: Soal Tidak Ditemukan

| Field | Detail |
|-------|--------|
| ID | TC-10 |
| Nama | Problem Not Found |
| Precondition | Aplikasi berjalan |
| Steps | 1. Akses URL `/problems/id-yang-tidak-ada` |
| Expected | Redirect ke `/problems` atau tampil halaman 404 yang jelas |
| Status | [ ] Pass / [ ] Fail |

---

## Edge Cases

| # | Skenario | Expected Behavior |
|---|----------|-------------------|
| E-01 | Kode sangat panjang (>10.000 karakter) | Sistem tetap bisa submit, tidak crash |
| E-02 | Output mengandung karakter spesial (`\n`, `\t`, spasi ekstra) | Perbandingan output di-trim/normalize |
| E-03 | Soal dengan 0 test case | Tidak bisa submit, atau tampil pesan "Soal belum memiliki test case" |
| E-04 | User submit berkali-kali dengan cepat (spam) | Tombol submit di-disable saat loading |
| E-05 | Koneksi internet putus saat submit | Tampil pesan error jaringan, bukan spinner selamanya |
| E-06 | Browser tidak support Monaco Editor | Tampil fallback textarea sederhana |

---

## Bug Reporting Format

Gunakan format berikut saat melaporkan bug (bisa di GitHub Issues atau dokumen internal):

```
## [BUG] Judul singkat bug

**ID Bug**: BUG-001
**Tanggal**: YYYY-MM-DD
**Reporter**: [nama]
**Severity**: Critical / High / Medium / Low

### Deskripsi
Jelaskan bug secara singkat dan jelas.

### Steps to Reproduce
1. Buka halaman ...
2. Lakukan ...
3. Klik ...

### Expected Behavior
Apa yang seharusnya terjadi.

### Actual Behavior
Apa yang sebenarnya terjadi.

### Screenshot / Video
(lampirkan jika ada)

### Environment
- Browser: Chrome 120 / Firefox 121 / dll
- OS: Windows 11 / macOS 14 / dll
- URL: https://...

### Catatan Tambahan
(opsional)
```

---

## Severity Level

| Level | Deskripsi | Contoh |
|-------|-----------|--------|
| Critical | Fitur utama tidak bisa digunakan sama sekali | Submit tidak bisa dilakukan |
| High | Fitur utama terganggu tapi ada workaround | Hasil evaluasi salah untuk beberapa soal |
| Medium | Fitur minor tidak berjalan | Filter kategori tidak bekerja |
| Low | Masalah tampilan / UX kecil | Warna badge salah, typo |

---

## Definition of Done (DoD)

Sebuah fitur dianggap **selesai dan siap release** jika memenuhi semua kriteria berikut:

### Functional
- [ ] Semua test case utama (TC-01 s/d TC-10) passed
- [ ] Tidak ada bug dengan severity Critical atau High yang belum diselesaikan
- [ ] Edge case utama sudah diverifikasi

### UI/UX
- [ ] Tampilan sesuai wireframe yang disepakati
- [ ] Tidak ada elemen UI yang rusak di resolusi 1280px ke atas
- [ ] Loading state ditampilkan saat menunggu response API
- [ ] Pesan error ditampilkan dengan jelas dan actionable

### Performance
- [ ] Halaman list soal load < 2 detik
- [ ] Hasil evaluasi muncul < 5 detik setelah submit

### Code Quality
- [ ] Tidak ada `console.log` yang tertinggal di production build
- [ ] Tidak ada error di browser console saat normal usage

---

## Checklist Pre-Release

- [ ] Semua test case manual sudah dijalankan
- [ ] Bug Critical dan High sudah di-fix
- [ ] Flow utama (landing → pilih soal → submit → lihat hasil) berjalan mulus
- [ ] Tampilan tidak rusak di Chrome, Firefox, dan Edge
- [ ] Environment variables production sudah dikonfigurasi dengan benar
