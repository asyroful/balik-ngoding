# Requirements Document

## Introduction

Balik Ngoding adalah web platform latihan logika pemrograman dasar yang ditujukan untuk fresh graduate dan developer yang ingin melatih kembali kemampuan problem solving secara mandiri. MVP mencakup lima fitur utama: daftar soal berdasarkan kategori, detail soal, code editor berbasis browser, submit jawaban, dan evaluasi hasil (benar/salah + expected vs actual output). Platform ini tidak memerlukan login di tahap MVP.

## Glossary

- **Platform**: Aplikasi web Balik Ngoding secara keseluruhan (frontend Next.js + backend NestJS)
- **Problem_List**: Halaman yang menampilkan daftar soal yang tersedia
- **Problem_Detail**: Halaman yang menampilkan deskripsi lengkap satu soal beserta code editor
- **Code_Editor**: Komponen Monaco Editor di browser tempat user menulis kode
- **Evaluator**: Layanan backend yang mengeksekusi kode user dan membandingkan output dengan expected output
- **Submission**: Satu pengiriman kode oleh user untuk sebuah soal
- **Test_Case**: Pasangan input dan expected output yang digunakan untuk mengevaluasi submission
- **Result_Panel**: Komponen UI yang menampilkan hasil evaluasi setelah submit
- **Category**: Pengelompokan soal berdasarkan topik: `loop`, `string`, `array`, `sql`
- **Difficulty**: Tingkat kesulitan soal: `easy`, `medium`, `hard`
- **Starter_Code**: Template kode awal yang ditampilkan di editor saat user membuka soal

---

## Requirements

### Requirement 1: Menampilkan Daftar Soal

**User Story:** Sebagai user, saya ingin melihat daftar soal yang tersedia, sehingga saya dapat memilih soal yang ingin dikerjakan.

#### Acceptance Criteria

1. WHEN user mengakses halaman `/problems`, THE Problem_List SHALL menampilkan semua soal aktif beserta judul, kategori, dan difficulty.
2. WHEN user memilih filter kategori tertentu, THE Problem_List SHALL menampilkan hanya soal dengan kategori yang dipilih.
3. IF backend tidak dapat diakses saat halaman dimuat, THEN THE Problem_List SHALL menampilkan pesan error yang informatif dan actionable.
4. THE Problem_List SHALL menampilkan maksimal 30 soal tanpa pagination di MVP.
5. WHEN user mengklik baris soal, THE Problem_List SHALL mengarahkan user ke halaman detail soal yang bersangkutan.

---

### Requirement 2: Menampilkan Detail Soal

**User Story:** Sebagai user, saya ingin membaca deskripsi lengkap sebuah soal, sehingga saya memahami apa yang harus dikerjakan sebelum mulai menulis kode.

#### Acceptance Criteria

1. WHEN user mengakses halaman `/problems/:id`, THE Problem_Detail SHALL menampilkan judul, deskripsi, contoh input/output, dan constraints soal.
2. THE Problem_Detail SHALL menampilkan Starter_Code di Code_Editor sebagai template awal.
3. IF soal dengan id yang diminta tidak ditemukan, THEN THE Problem_Detail SHALL mengarahkan user ke halaman `/problems` dengan pesan error yang jelas.
4. WHEN halaman Problem_Detail dimuat, THE Code_Editor SHALL menampilkan Starter_Code sesuai soal yang dipilih.

---

### Requirement 3: Code Editor di Browser

**User Story:** Sebagai user, saya ingin menulis kode langsung di browser, sehingga saya tidak perlu berpindah ke IDE eksternal.

#### Acceptance Criteria

1. THE Code_Editor SHALL mendukung syntax highlighting untuk bahasa JavaScript.
2. THE Code_Editor SHALL menampilkan nomor baris pada setiap baris kode.
3. WHILE user mengetik di Code_Editor, THE Code_Editor SHALL memperbarui konten kode secara real-time tanpa delay yang terasa.
4. WHERE browser tidak mendukung Monaco Editor, THE Platform SHALL menampilkan fallback textarea sederhana agar user tetap dapat menulis kode.
5. THE Code_Editor SHALL memiliki tinggi minimal 400px agar area penulisan kode cukup nyaman.

---

### Requirement 4: Submit Jawaban

**User Story:** Sebagai user, saya ingin mengirimkan kode yang saya tulis untuk dievaluasi, sehingga saya dapat mengetahui apakah jawaban saya benar.

#### Acceptance Criteria

1. WHEN user mengklik tombol "Submit Jawaban", THE Platform SHALL mengirimkan kode beserta `problemId` dan bahasa pemrograman ke endpoint `POST /submit`.
2. IF Code_Editor kosong saat user mengklik tombol submit, THEN THE Platform SHALL menonaktifkan tombol submit dan menampilkan pesan validasi "Kode tidak boleh kosong".
3. WHILE Submission sedang diproses, THE Platform SHALL menonaktifkan tombol submit dan menampilkan indikator loading.
4. IF user mengklik tombol submit lebih dari satu kali secara berurutan sebelum respons diterima, THEN THE Platform SHALL mengabaikan klik berikutnya dan hanya memproses satu Submission.
5. WHEN Submission berhasil dikirim, THE Platform SHALL menampilkan Result_Panel dengan hasil evaluasi.
6. IF terjadi error jaringan saat submit, THEN THE Platform SHALL menampilkan pesan error yang actionable di Result_Panel dan mengaktifkan kembali tombol submit.

---

### Requirement 5: Evaluasi Jawaban

**User Story:** Sebagai user, saya ingin melihat hasil evaluasi kode saya secara detail, sehingga saya dapat memahami di mana letak kesalahan saya.

#### Acceptance Criteria

1. WHEN Evaluator menerima Submission, THE Evaluator SHALL menjalankan kode user terhadap semua Test_Case yang terkait dengan soal tersebut.
2. WHEN Evaluator selesai mengevaluasi, THE Result_Panel SHALL menampilkan status per Test_Case: passed atau failed, beserta input, expected output, dan actual output.
3. THE Result_Panel SHALL menampilkan skor total dalam format "X dari Y test case passed".
4. IF semua Test_Case passed, THEN THE Evaluator SHALL menetapkan status Submission sebagai `accepted`.
5. IF setidaknya satu Test_Case failed dan tidak ada runtime error, THEN THE Evaluator SHALL menetapkan status Submission sebagai `wrong_answer`.
6. IF kode user menghasilkan runtime error atau syntax error, THEN THE Evaluator SHALL menetapkan status Submission sebagai `error` dan menampilkan pesan error yang deskriptif.
7. IF eksekusi kode user melebihi batas waktu yang ditentukan, THEN THE Evaluator SHALL menghentikan eksekusi dan mengembalikan status `error` dengan pesan "Waktu eksekusi habis".
8. THE Evaluator SHALL menormalisasi output (trim whitespace) sebelum membandingkan actual output dengan expected output.
9. WHEN evaluasi selesai, THE Platform SHALL menampilkan hasil dalam waktu kurang dari 5 detik sejak submit dikirim.

---

### Requirement 6: Navigasi dan Alur Pengguna

**User Story:** Sebagai user, saya ingin dapat berpindah antar halaman dengan mudah, sehingga pengalaman menggunakan platform terasa lancar dan tidak membingungkan.

#### Acceptance Criteria

1. THE Platform SHALL menyediakan navbar yang konsisten di semua halaman dengan logo dan navigasi ke halaman Problem_List.
2. WHEN user berada di Result_Panel setelah submit, THE Platform SHALL menampilkan tombol "Coba Lagi" yang mereset Code_Editor ke Starter_Code dan tombol "Soal Berikutnya" yang mengarahkan ke soal berikutnya.
3. THE Platform SHALL memungkinkan user mencapai halaman pengerjaan soal dari landing page dalam maksimal 2 klik.
4. WHEN user mengakses URL yang tidak valid, THE Platform SHALL menampilkan halaman 404 yang jelas dengan tautan kembali ke Problem_List.

---

### Requirement 7: Penyimpanan dan Pengambilan Data Soal

**User Story:** Sebagai developer, saya ingin soal dan test case tersimpan di database, sehingga konten dapat dikelola dan diperbarui tanpa mengubah kode.

#### Acceptance Criteria

1. THE Platform SHALL menyimpan setiap soal dengan atribut: judul, deskripsi, kategori, difficulty, starter code, dan status aktif.
2. THE Platform SHALL menyimpan setiap Test_Case dengan atribut: input, expected output, dan flag `isHidden`.
3. WHEN endpoint `GET /problems` dipanggil, THE Platform SHALL mengembalikan hanya soal dengan status aktif.
4. WHEN endpoint `GET /problems/:id` dipanggil, THE Platform SHALL mengembalikan detail soal beserta Test_Case yang tidak hidden (`isHidden = false`) sebagai contoh.
5. THE Evaluator SHALL menggunakan semua Test_Case (termasuk yang hidden) saat mengevaluasi Submission.
6. THE Platform SHALL menyimpan setiap Submission ke database beserta kode, bahasa, status, skor, dan detail hasil per Test_Case.

---

### Requirement 8: Keamanan Eksekusi Kode

**User Story:** Sebagai developer, saya ingin kode yang dikirim user dieksekusi dalam lingkungan yang terisolasi, sehingga server tidak rentan terhadap kode berbahaya.

#### Acceptance Criteria

1. THE Evaluator SHALL mengeksekusi kode user dalam lingkungan sandbox yang terisolasi dari proses utama server.
2. IF kode user mencoba mengakses sistem file, jaringan, atau environment variables server, THEN THE Evaluator SHALL memblokir akses tersebut dan mengembalikan status `error`.
3. THE Evaluator SHALL membatasi waktu eksekusi setiap Submission dengan timeout yang ditentukan (maksimal 5 detik per Test_Case).
4. THE Evaluator SHALL membatasi penggunaan memori eksekusi kode user untuk mencegah resource exhaustion.
