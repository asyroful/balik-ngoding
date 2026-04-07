# Implementation Plan: User Experience Improvements

## Overview

Lima fitur UX yang independen diimplementasikan secara bertahap: LanguageSelector baru, Ctrl+Enter shortcut, cookie-based progress, category-based fetching, dan mobile tab navigation. `fast-check` sudah tersedia di devDependencies.

## Tasks

- [x] 1. Buat komponen LanguageSelector
  - [x] 1.1 Buat file `frontend/components/Editor/LanguageSelector.tsx`
    - Definisikan array `LANGUAGES` dengan 6 bahasa (JavaScript, SQL, Python, Java, PHP, C)
    - JavaScript dan SQL: `active: true`; Python, Java, PHP, C: `active: false`
    - Render tombol per bahasa; bahasa `active: false` diberi badge "Coming Soon" dan atribut `disabled`
    - Ekspor fungsi `getDefaultLanguage(category: string): string` untuk digunakan di luar komponen
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5, 1.6_

  - [x] 1.2 Tulis unit tests untuk LanguageSelector
    - Render menampilkan 6 bahasa
    - JS dan SQL tidak disabled, tidak ada teks "Coming Soon"
    - Python/Java/PHP/C memiliki teks "Coming Soon" dan atribut disabled
    - Klik JS/SQL memanggil `onChange`; klik Python tidak memanggil `onChange`
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.6_

  - [x] 1.3 Tulis property test untuk LanguageSelector
    - **Property 1: Language list completeness** — render selalu menghasilkan tepat 6 tombol bahasa
    - **Validates: Requirements 1.1**
    - **Property 2: Active languages are selectable** — JS dan SQL tidak pernah disabled
    - **Validates: Requirements 1.2**
    - **Property 3: Coming soon languages are disabled and labeled** — Python/Java/PHP/C selalu disabled dan berlabel "Coming Soon"
    - **Validates: Requirements 1.3, 1.4, 1.6**
    - **Property 4: Default language matches problem category** — `getDefaultLanguage` mengembalikan `'sql'` hanya jika `category === 'sql'`, selain itu `'javascript'`
    - **Validates: Requirements 1.5**

- [x] 2. Tambahkan Ctrl+Enter shortcut ke CodeEditor
  - [x] 2.1 Modifikasi `frontend/components/Editor/CodeEditor.tsx`
    - Tambahkan prop `onSubmit?: () => void` ke interface `CodeEditorProps`
    - Di callback `onMount`, daftarkan Monaco action dengan `editor.addAction({ id: 'submit-solution', keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter], run: () => onSubmit?.() })`
    - Guard di dalam `run`: cek `isLoading` dari `useSubmissionStore` dan `code.trim() === ''` sebelum memanggil `onSubmit`
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.6_

  - [x] 2.2 Sambungkan `onSubmit` di `frontend/app/problems/[id]/page.tsx`
    - Buat fungsi `handleSubmit` yang memanggil logika submit (sama dengan `SubmitButton.handleSubmit`)
    - Pass `onSubmit={handleSubmit}` ke komponen `<CodeEditor>`
    - _Requirements: 2.1, 2.5_

  - [x] 2.3 Tulis unit tests untuk Ctrl+Enter
    - Render `CodeEditor` dengan prop `onSubmit`; verifikasi prop diterima tanpa error
    - Verifikasi fallback textarea tetap muncul saat Monaco error (Requirement 2.5)
    - _Requirements: 2.4, 2.5_

  - [x] 2.4 Tulis property tests untuk Ctrl+Enter
    - **Property 6: Ctrl+Enter ignored when loading** — saat `isLoading === true`, `onSubmit` tidak dipanggil
    - **Validates: Requirements 2.2**
    - **Property 7: Ctrl+Enter ignored for whitespace-only code** — untuk sembarang string whitespace, `onSubmit` tidak dipanggil
    - **Validates: Requirements 2.3**

- [x] 3. Checkpoint — Pastikan semua tests pass
  - Pastikan semua tests pass, tanyakan ke user jika ada pertanyaan.

- [x] 4. Buat `useProgress` hook dan integrasikan ke ResultPanel
  - [x] 4.1 Buat file `frontend/hooks/useProgress.ts`
    - Definisikan tipe `ProgressMap = Record<string, 'accepted'>`
    - Baca cookie `bn_progress` saat inisialisasi; parse JSON dengan `try/catch`, fallback ke `{}`
    - Implementasikan `markAccepted(problemId)`: hanya tulis jika belum `'accepted'`; tulis ke cookie dengan `max-age=31536000; path=/; SameSite=Lax`
    - Implementasikan `isAccepted(problemId)`: kembalikan `progress[problemId] === 'accepted'`
    - Simpan state di `useState` agar reaktif
    - _Requirements: 3.1, 3.2, 3.3, 3.7, 3.8, 3.9_

  - [x] 4.2 Modifikasi `frontend/components/Result/ResultPanel.tsx`
    - Import dan panggil `useProgress` hook
    - Saat `result?.status === 'accepted'`, panggil `markAccepted(currentProblemId)`
    - _Requirements: 3.1_

  - [x] 4.3 Tulis unit tests untuk useProgress
    - `markAccepted` menulis ke `document.cookie`
    - Cookie corrupt → `progress` adalah `{}`
    - Cookie tidak ada → `progress` adalah `{}`
    - `isAccepted` mengembalikan `true` untuk problemId yang sudah accepted
    - `markAccepted` tidak menimpa status `'accepted'` yang sudah ada
    - _Requirements: 3.1, 3.2, 3.3, 3.7, 3.8_

  - [x] 4.4 Tulis property tests untuk useProgress
    - **Property 8: markAccepted round-trip** — setelah `markAccepted(id)`, cookie `bn_progress` dapat di-parse dan mengandung `id: 'accepted'`
    - **Validates: Requirements 3.1, 3.2**
    - **Property 9: Accepted status is preserved** — setelah `markAccepted(id)`, `isAccepted(id)` selalu `true` tanpa perlu dipanggil ulang
    - **Validates: Requirements 3.8**
    - **Property 10: Invalid cookie graceful fallback** — sembarang string non-JSON sebagai nilai cookie tidak melempar error dan menghasilkan `progress === {}`
    - **Validates: Requirements 3.7**

- [x] 5. Tampilkan indikator progress di ProblemTable
  - [x] 5.1 Modifikasi `frontend/components/ProblemList/ProblemTable.tsx`
    - Import dan panggil `useProgress` hook
    - Tambahkan kolom "Status" di header tabel
    - Untuk setiap baris, jika `isAccepted(problem.id)` tampilkan badge hijau "Selesai" dengan ikon ✓; jika tidak, tampilkan sel kosong
    - _Requirements: 3.4, 3.5, 3.6_

  - [x] 5.2 Tulis property test untuk ProblemTable
    - **Property 11: Accepted problems show visual indicator in table** — setiap baris dengan `id` di progress store menampilkan indikator; baris lain tidak
    - **Validates: Requirements 3.4, 3.5, 3.6**

- [x] 6. Implementasikan category-based fetching di problems/page.tsx
  - [x] 6.1 Modifikasi `frontend/app/problems/page.tsx`
    - Ubah `fetchProblems` agar menerima parameter `category: string` dan memanggil `getProblems(category)`
    - Ubah `useEffect` dependency menjadi `[selectedCategory]` sehingga fetch dipicu saat kategori berubah
    - Hapus variabel `filtered` dan client-side filter; gunakan langsung `problems` dari state
    - Pastikan loading state (skeleton) aktif selama fetch berlangsung
    - Pastikan error state menampilkan tombol "Coba lagi" yang memanggil `fetchProblems(selectedCategory)`
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5_

  - [x] 6.2 Tulis unit tests untuk category-based fetching
    - Initial load memanggil `getProblems('loop')`
    - Klik tab kategori memanggil `getProblems` dengan kategori yang dipilih
    - Error state menampilkan tombol "Coba lagi"
    - _Requirements: 4.1, 4.2, 4.4_

  - [x] 6.3 Tulis property test untuk category-based fetching
    - **Property 12: Category change triggers server fetch** — untuk sembarang kategori yang dipilih, `getProblems` dipanggil dengan kategori tersebut
    - **Validates: Requirements 4.2, 4.5**

- [x] 7. Checkpoint — Pastikan semua tests pass
  - Pastikan semua tests pass, tanyakan ke user jika ada pertanyaan.

- [x] 8. Integrasikan LanguageSelector ke halaman detail soal
  - [x] 8.1 Modifikasi `frontend/app/problems/[id]/page.tsx`
    - Import `LanguageSelector` dan `getDefaultLanguage`
    - Ganti baris `setLanguage(data.category === 'sql' ? 'sql' : 'javascript')` dengan `setLanguage(getDefaultLanguage(data.category))`
    - Render `<LanguageSelector selected={language} onChange={setLanguage} problemCategory={problem.category} />` di editor top bar, menggantikan teks statis `solution.js` / `solution.sql`
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5, 1.6_

- [x] 9. Implementasikan mobile tab navigation di halaman detail soal
  - [x] 9.1 Modifikasi `frontend/app/problems/[id]/page.tsx`
    - Tambahkan state `const [activeTab, setActiveTab] = useState<'soal' | 'editor'>('soal')`
    - Buat komponen tab navigation inline: dua tombol "Soal" dan "Editor" dengan class berbeda untuk tab aktif vs tidak aktif
    - Tampilkan tab navigation hanya di mobile menggunakan `md:hidden`
    - Panel kiri (deskripsi): tampilkan full-width di mobile saat `activeTab === 'soal'`, gunakan `hidden md:block` untuk desktop
    - Panel kanan (editor): tampilkan full-width di mobile saat `activeTab === 'editor'`, gunakan `hidden md:flex` untuk desktop
    - Di mobile, panel editor tetap menyertakan `SubmitButton` dan `ResultPanel` (Requirement 5.4)
    - Pastikan layout desktop side-by-side (45%/55%) tidak berubah
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7_

  - [x] 9.2 Tulis unit tests untuk mobile layout
    - Tab "Soal" menampilkan deskripsi soal
    - Tab "Editor" menampilkan editor dan submit button
    - Desktop layout menampilkan kedua panel sekaligus
    - _Requirements: 5.1, 5.2, 5.3, 5.4, 5.6_

  - [x] 9.3 Tulis property test untuk tab navigation
    - **Property 13: Active tab has distinct visual style** — untuk sembarang `activeTab`, tombol aktif memiliki className berbeda dari tombol tidak aktif
    - **Validates: Requirements 5.5**

- [x] 10. Final checkpoint — Pastikan semua tests pass
  - Pastikan semua tests pass, tanyakan ke user jika ada pertanyaan.

## Notes

- Tasks bertanda `*` bersifat opsional dan dapat dilewati untuk MVP yang lebih cepat
- `fast-check` sudah tersedia di devDependencies, tidak perlu install ulang
- Setiap task mereferensikan requirements spesifik untuk traceability
- Property tests menggunakan minimum 100 iterasi (`numRuns: 100`)
- Kelima fitur bersifat independen dan dapat dikerjakan secara paralel, namun task 8 bergantung pada task 1
