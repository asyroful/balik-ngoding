# Requirements Document

## Introduction

Fitur **User Experience Improvements** pada aplikasi coding challenge **Balik Ngoding** mencakup lima peningkatan UX yang saling independen:

1. **Language Coming Soon** — Menampilkan label "coming soon" pada bahasa yang belum tersedia (Python, Java, PHP, C) di language selector, sehingga user memahami bahwa hanya JavaScript dan SQL yang aktif saat ini.
2. **Ctrl+Enter Submit** — Shortcut keyboard Ctrl+Enter di dalam code editor untuk men-submit jawaban tanpa harus berpindah ke mouse.
3. **Cookie-based Progress & Scoring** — Menyimpan progress pengerjaan soal (soal yang sudah dikerjakan beserta statusnya) menggunakan cookie browser tanpa memerlukan autentikasi, dan menampilkan status "accepted" di daftar soal.
4. **Category-based Problem Fetching** — Mengambil daftar soal per kategori dari server saat user memilih tab kategori, menggantikan pendekatan saat ini yang memuat semua soal sekaligus lalu memfilter di sisi client.
5. **Mobile-Friendly Problem Detail Layout** — Mengubah layout halaman detail soal menjadi tab-based di viewport mobile sehingga deskripsi soal dan code editor masing-masing mendapat ruang penuh, menggantikan layout side-by-side yang membuat kedua panel terlalu sempit di layar kecil.

---

## Glossary

- **Language_Selector**: Komponen UI yang menampilkan daftar bahasa pemrograman yang tersedia dan yang belum tersedia di halaman detail soal.
- **Code_Editor**: Komponen Monaco Editor (`CodeEditor.tsx`) tempat user menulis kode jawaban.
- **Submit_Handler**: Fungsi yang dipanggil saat user men-submit jawaban, baik via tombol maupun shortcut keyboard.
- **Progress_Store**: Mekanisme penyimpanan berbasis cookie yang mencatat status pengerjaan soal per `problemId`.
- **Problem_Table**: Komponen `ProblemTable.tsx` yang menampilkan daftar soal dalam bentuk tabel.
- **Accepted_Status**: Status submission dengan nilai `"accepted"` yang menandakan semua test case lulus.
- **Cookie**: Penyimpanan data di browser menggunakan `document.cookie` atau library `js-cookie`, tanpa memerlukan login.
- **Problems_Page**: Halaman `problems/page.tsx` yang menampilkan daftar soal beserta filter kategori.
- **Category_Filter**: Komponen `CategoryFilter.tsx` yang menampilkan tab-tab kategori soal (loop, string, array, sql).
- **API_Client**: Fungsi `getProblems(category?)` di `lib/api.ts` yang memanggil endpoint `GET /problems?category=` pada backend.
- **Problem_Detail_Page**: Halaman `problems/[id]/page.tsx` yang menampilkan deskripsi soal dan code editor secara berdampingan.
- **Mobile_Viewport**: Viewport dengan lebar kurang dari 768px (breakpoint `md` pada Tailwind CSS).
- **Desktop_Viewport**: Viewport dengan lebar 768px atau lebih.
- **Tab_Navigation**: Komponen UI yang menampilkan dua tab ("Soal" dan "Editor") di Mobile_Viewport untuk berpindah antara panel deskripsi dan panel editor.
- **Active_Tab**: Tab yang sedang aktif dan kontennya ditampilkan, ditandai secara visual berbeda dari tab yang tidak aktif.

---

## Requirements

### Requirement 1: Language Coming Soon Label

**User Story:** Sebagai user, saya ingin melihat bahasa pemrograman apa saja yang akan tersedia di masa depan, sehingga saya tahu bahwa hanya JavaScript dan SQL yang bisa digunakan sekarang.

#### Acceptance Criteria

1. THE Language_Selector SHALL menampilkan daftar bahasa: JavaScript, SQL, Python, Java, PHP, dan C.
2. WHEN Language_Selector ditampilkan, THE Language_Selector SHALL menandai JavaScript dan SQL sebagai bahasa yang aktif dan dapat dipilih.
3. WHEN Language_Selector ditampilkan, THE Language_Selector SHALL menampilkan label "Coming Soon" pada bahasa Python, Java, PHP, dan C.
4. WHILE bahasa Python, Java, PHP, atau C ditampilkan dengan label "Coming Soon", THE Language_Selector SHALL menonaktifkan interaksi klik pada bahasa tersebut.
5. THE Language_Selector SHALL menentukan bahasa aktif secara otomatis berdasarkan kategori soal: kategori `sql` menggunakan bahasa SQL, kategori lainnya menggunakan JavaScript.
6. IF user mencoba mengklik bahasa dengan label "Coming Soon", THEN THE Language_Selector SHALL tidak mengubah bahasa yang sedang aktif.

---

### Requirement 2: Ctrl+Enter Keyboard Shortcut untuk Submit

**User Story:** Sebagai user, saya ingin bisa men-submit jawaban menggunakan shortcut keyboard Ctrl+Enter, sehingga saya tidak perlu berpindah ke mouse saat coding.

#### Acceptance Criteria

1. WHEN user menekan Ctrl+Enter di dalam Code_Editor, THE Submit_Handler SHALL memproses submission dengan perilaku yang identik dengan menekan tombol "Submit Jawaban".
2. WHILE submission sedang diproses (`isLoading` bernilai `true`), THE Submit_Handler SHALL mengabaikan input Ctrl+Enter berikutnya.
3. WHILE kode editor kosong (hanya whitespace), THE Submit_Handler SHALL mengabaikan input Ctrl+Enter dan tidak memproses submission.
4. THE Code_Editor SHALL mendaftarkan keyboard shortcut Ctrl+Enter menggunakan Monaco Editor action API saat editor selesai dimuat (`onMount`).
5. IF Code_Editor gagal dimuat dan fallback textarea aktif, THEN THE Submit_Handler SHALL tetap dapat dipanggil melalui tombol "Submit Jawaban".
6. WHEN komponen Code_Editor di-unmount, THE Code_Editor SHALL membersihkan (dispose) action Ctrl+Enter yang telah didaftarkan.

---

### Requirement 3: Cookie-based Progress & Scoring

**User Story:** Sebagai user, saya ingin progress pengerjaan soal saya tersimpan secara otomatis di browser, sehingga saya bisa melihat soal mana yang sudah berhasil saya selesaikan tanpa perlu login.

#### Acceptance Criteria

1. WHEN submission menghasilkan status `"accepted"`, THE Progress_Store SHALL menyimpan `problemId` beserta status `"accepted"` ke dalam cookie browser.
2. THE Progress_Store SHALL menyimpan data progress dalam format JSON yang berisi mapping `problemId` ke status submission terakhir.
3. THE Progress_Store SHALL menggunakan cookie dengan nama `bn_progress` dan masa berlaku (expiry) 365 hari.
4. WHEN Problem_Table dirender, THE Problem_Table SHALL membaca data progress dari cookie `bn_progress` dan menampilkan indikator visual pada baris soal yang memiliki status `"accepted"`.
5. THE Problem_Table SHALL menampilkan ikon centang (✓) atau badge "Selesai" berwarna hijau pada kolom status untuk soal dengan status `"accepted"`.
6. WHEN user membuka halaman detail soal yang sudah pernah di-accepted, THE Problem_Table SHALL tetap menampilkan status "accepted" dari cookie tanpa memerlukan submission ulang.
7. IF cookie `bn_progress` tidak ditemukan atau formatnya tidak valid (corrupt/non-JSON), THEN THE Progress_Store SHALL menginisialisasi progress dengan objek kosong tanpa melempar error.
8. WHEN submission baru menghasilkan status selain `"accepted"` untuk soal yang sebelumnya sudah `"accepted"`, THE Progress_Store SHALL mempertahankan status `"accepted"` yang sudah tersimpan dan tidak menimpanya.
9. THE Progress_Store SHALL dapat diakses oleh komponen manapun tanpa memerlukan prop drilling, menggunakan Zustand store atau custom hook.

---

### Requirement 4: Category-based Problem Fetching

**User Story:** Sebagai user, saya ingin daftar soal dimuat per kategori saat saya memilih tab kategori, sehingga halaman tidak perlu memuat ratusan soal sekaligus dan lebih ringan.

#### Acceptance Criteria

1. WHEN Problems_Page pertama kali dimuat, THE API_Client SHALL memanggil `GET /problems?category=loop` untuk mengambil hanya soal dengan kategori default `loop`.
2. WHEN user mengklik tab kategori pada Category_Filter, THE Problems_Page SHALL memanggil `GET /problems?category={kategori}` untuk kategori yang dipilih, bukan memfilter data yang sudah ada di client.
3. WHILE request fetch soal sedang berlangsung, THE Problems_Page SHALL menampilkan loading state (skeleton) dan menonaktifkan interaksi pada tabel soal.
4. IF request fetch soal gagal, THEN THE Problems_Page SHALL menampilkan pesan error dan tombol "Coba lagi" yang memicu ulang fetch untuk kategori yang sedang aktif.
5. WHEN fetch soal berhasil, THE Problems_Page SHALL mengganti daftar soal yang ditampilkan dengan hasil fetch terbaru tanpa mempertahankan data kategori sebelumnya.

---

### Requirement 5: Mobile-Friendly Problem Detail Layout

**User Story:** Sebagai user yang mengakses dari perangkat mobile, saya ingin bisa membaca soal dan menulis kode dengan nyaman, sehingga saya tidak perlu scroll horizontal atau bekerja di area yang terlalu sempit.

#### Acceptance Criteria

1. WHEN Problem_Detail_Page dirender pada Mobile_Viewport, THE Problem_Detail_Page SHALL menampilkan Tab_Navigation dengan dua tab: "Soal" dan "Editor", menggantikan layout side-by-side.
2. WHEN tab "Soal" aktif pada Mobile_Viewport, THE Problem_Detail_Page SHALL menampilkan deskripsi soal dan contoh test case secara full-width tanpa batasan lebar kolom.
3. WHEN tab "Editor" aktif pada Mobile_Viewport, THE Problem_Detail_Page SHALL menampilkan Code_Editor secara full-width tanpa batasan lebar kolom.
4. WHILE tab "Editor" aktif pada Mobile_Viewport, THE Problem_Detail_Page SHALL menampilkan Submit_Handler dan Result_Panel sehingga dapat diakses tanpa berpindah tab.
5. THE Tab_Navigation SHALL menandai Active_Tab secara visual berbeda dari tab yang tidak aktif (misalnya dengan warna latar, border bawah, atau perubahan warna teks).
6. WHEN Problem_Detail_Page dirender pada Desktop_Viewport, THE Problem_Detail_Page SHALL menampilkan layout side-by-side dengan panel deskripsi di kiri (45%) dan panel editor di kanan (55%), identik dengan tampilan saat ini.
7. WHILE tab "Editor" aktif pada Mobile_Viewport dan virtual keyboard muncul, THE Submit_Handler SHALL tetap dapat dijangkau oleh user tanpa tertutup oleh virtual keyboard (tombol submit sticky atau dapat di-scroll ke posisinya).

