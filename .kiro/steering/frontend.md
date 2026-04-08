---
inclusion: always
---

# Frontend — Next.js 14 App Router

## Stack

- **Framework**: Next.js 14 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: Zustand
- **Code Editor**: Monaco Editor (`@monaco-editor/react`)
- **HTTP Client**: fetch (bawaan Next.js)

## Struktur Direktori

```
frontend/
├── app/                            # Next.js App Router
│   ├── layout.tsx                  # Root layout (Navbar, font, metadata)
│   ├── page.tsx                    # Landing page
│   ├── globals.css
│   └── problems/
│       ├── page.tsx                # Daftar soal
│       └── [id]/
│           └── page.tsx            # Detail soal + editor + result
├── components/
│   ├── Navbar.tsx
│   ├── NavLink.tsx
│   ├── Skeleton.tsx
│   ├── Editor/
│   │   ├── CodeEditor.tsx          # Wrapper Monaco Editor
│   │   └── SubmitButton.tsx
│   ├── ProblemList/
│   │   ├── ProblemTable.tsx
│   │   └── CategoryFilter.tsx
│   ├── ProblemDetail/
│   │   ├── ProblemDescription.tsx
│   │   └── ExampleBlock.tsx
│   └── Result/
│       └── ResultPanel.tsx
├── store/
│   └── submissionStore.ts          # Zustand store untuk submission state
├── lib/
│   ├── api.ts                      # Fungsi API call ke backend
│   └── types.ts                    # TypeScript types/interfaces
├── hooks/
│   └── useProgress.ts              # Custom hooks
├── __tests__/                      # Vitest + fast-check tests
├── next.config.js
├── tailwind.config.ts
├── tsconfig.json
└── vitest.config.ts
```

## Environment Variables

| Variable | Deskripsi | Contoh Nilai |
|---|---|---|
| `NEXT_PUBLIC_API_URL` | URL backend API | `http://localhost:8080` |

## Routing

| Route | Halaman |
|---|---|
| `/` | Landing page |
| `/problems` | Daftar semua soal |
| `/problems/[id]` | Detail soal + Monaco Editor + result panel |

## Konvensi Kode

- Gunakan Server Components untuk halaman yang hanya fetch data
- Tambahkan `'use client'` hanya jika komponen butuh interaktivitas (event handler, hooks)
- Monaco Editor harus di-import dengan `dynamic` dari `next/dynamic` dengan `ssr: false` untuk menghindari SSR crash
- Semua tipe didefinisikan di `lib/types.ts` — hindari penggunaan `any`
- Komponen: `PascalCase`, fungsi/variabel: `camelCase`

## Menjalankan Frontend

```bash
cd frontend
npm install
npm run dev       # Development server di http://localhost:3000
npm run build     # Production build
npm start         # Production server
```

## Menjalankan Tests

```bash
cd frontend
npx vitest run    # Run semua tests sekali
npx vitest        # Watch mode
```
