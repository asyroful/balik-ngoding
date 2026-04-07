# Senior Frontend Engineer — LogicLab MVP

## Stack

- **Framework**: Next.js 14 (App Router)
- **Styling**: TailwindCSS
- **Code Editor**: Monaco Editor (`@monaco-editor/react`)
- **State Management**: Zustand (ringan, tidak perlu Redux)
- **HTTP Client**: fetch bawaan Next.js / axios
- **Language**: TypeScript

---

## Arsitektur Folder

```
src/
├── app/
│   ├── layout.tsx              # Root layout (navbar, font, metadata)
│   ├── page.tsx                # Landing page
│   ├── problems/
│   │   ├── page.tsx            # Problem list page
│   │   └── [id]/
│   │       └── page.tsx        # Problem detail + editor page
├── components/
│   ├── Navbar.tsx
│   ├── ProblemList/
│   │   ├── ProblemTable.tsx
│   │   └── CategoryFilter.tsx
│   ├── ProblemDetail/
│   │   ├── ProblemDescription.tsx
│   │   └── ExampleBlock.tsx
│   ├── Editor/
│   │   ├── CodeEditor.tsx      # Wrapper Monaco Editor
│   │   └── SubmitButton.tsx
│   └── Result/
│       ├── ResultPanel.tsx
│       └── TestCaseResult.tsx
├── store/
│   └── submissionStore.ts      # Zustand store untuk submission state
├── lib/
│   ├── api.ts                  # Fungsi-fungsi API call
│   └── types.ts                # TypeScript types/interfaces
└── styles/
    └── globals.css
```

---

## Routing Structure

| Route | Halaman | Deskripsi |
|-------|---------|-----------|
| `/` | Landing Page | Hero section + CTA + kategori |
| `/problems` | Problem List | Daftar semua soal dengan filter |
| `/problems/[id]` | Problem Detail | Deskripsi soal + editor + result |

---

## TypeScript Types

```typescript
// lib/types.ts

export interface Problem {
  id: string;
  title: string;
  category: 'loop' | 'string' | 'array' | 'sql';
  difficulty: 'easy' | 'medium' | 'hard';
  description: string;
  examples: Example[];
  constraints: string[];
  starterCode: string;
}

export interface Example {
  input: string;
  output: string;
  explanation?: string;
}

export interface TestCaseResult {
  passed: boolean;
  input: string;
  expected: string;
  actual: string;
}

export interface SubmissionResult {
  status: 'accepted' | 'wrong_answer' | 'error';
  score: number;
  total: number;
  results: TestCaseResult[];
  errorMessage?: string;
}
```

---

## Komponen Utama

### `CodeEditor.tsx`

```tsx
'use client';
import Editor from '@monaco-editor/react';

interface CodeEditorProps {
  value: string;
  onChange: (value: string) => void;
  language?: string;
}

export default function CodeEditor({ value, onChange, language = 'javascript' }: CodeEditorProps) {
  return (
    <Editor
      height="400px"
      language={language}
      value={value}
      theme="vs-dark"
      onChange={(val) => onChange(val ?? '')}
      options={{
        minimap: { enabled: false },
        fontSize: 14,
        lineNumbers: 'on',
        scrollBeyondLastLine: false,
      }}
    />
  );
}
```

### `ResultPanel.tsx`

```tsx
import { SubmissionResult } from '@/lib/types';
import TestCaseResult from './TestCaseResult';

interface ResultPanelProps {
  result: SubmissionResult | null;
  isLoading: boolean;
}

export default function ResultPanel({ result, isLoading }: ResultPanelProps) {
  if (isLoading) return <p className="text-gray-400">Mengevaluasi jawaban...</p>;
  if (!result) return null;

  return (
    <div className="mt-4 space-y-3">
      <p className="font-semibold">
        Score: {result.score}/{result.total} test cases passed
      </p>
      {result.results.map((tc, i) => (
        <TestCaseResult key={i} index={i + 1} result={tc} />
      ))}
    </div>
  );
}
```

---

## State Management (Zustand)

```typescript
// store/submissionStore.ts
import { create } from 'zustand';
import { SubmissionResult } from '@/lib/types';

interface SubmissionStore {
  code: string;
  isLoading: boolean;
  result: SubmissionResult | null;
  setCode: (code: string) => void;
  setLoading: (loading: boolean) => void;
  setResult: (result: SubmissionResult | null) => void;
  reset: () => void;
}

export const useSubmissionStore = create<SubmissionStore>((set) => ({
  code: '',
  isLoading: false,
  result: null,
  setCode: (code) => set({ code }),
  setLoading: (isLoading) => set({ isLoading }),
  setResult: (result) => set({ result }),
  reset: () => set({ code: '', isLoading: false, result: null }),
}));
```

---

## Integrasi API

```typescript
// lib/api.ts
const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:3001';

export async function getProblems() {
  const res = await fetch(`${BASE_URL}/problems`);
  if (!res.ok) throw new Error('Gagal mengambil daftar soal');
  return res.json();
}

export async function getProblemById(id: string) {
  const res = await fetch(`${BASE_URL}/problems/${id}`);
  if (!res.ok) throw new Error('Soal tidak ditemukan');
  return res.json();
}

export async function submitSolution(problemId: string, code: string, language: string) {
  const res = await fetch(`${BASE_URL}/submit`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ problemId, code, language }),
  });
  if (!res.ok) throw new Error('Gagal submit jawaban');
  return res.json();
}
```

---

## Handling Submission

Flow submit di halaman `/problems/[id]`:

```tsx
// app/problems/[id]/page.tsx (simplified)
'use client';
import { useSubmissionStore } from '@/store/submissionStore';
import { submitSolution } from '@/lib/api';

export default function ProblemPage({ params }: { params: { id: string } }) {
  const { code, isLoading, result, setCode, setLoading, setResult } = useSubmissionStore();

  const handleSubmit = async () => {
    setLoading(true);
    setResult(null);
    try {
      const data = await submitSolution(params.id, code, 'javascript');
      setResult(data);
    } catch (err) {
      // tampilkan error ke user
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="grid grid-cols-2 gap-4 h-screen">
      {/* Panel kiri: deskripsi soal */}
      {/* Panel kanan: editor + result */}
    </div>
  );
}
```

---

## Error Handling

| Skenario | Handling |
|----------|----------|
| API tidak bisa diakses | Tampilkan toast/banner "Server tidak tersedia, coba lagi" |
| Soal tidak ditemukan | Redirect ke `/problems` dengan pesan error |
| Submit gagal (network) | Tampilkan pesan error di result panel, tombol retry |
| Kode kosong saat submit | Validasi di frontend, disable tombol submit jika editor kosong |
| Response timeout | Tampilkan "Evaluasi memakan waktu terlalu lama, coba lagi" |

---

## Coding Standards

### Umum
- Gunakan TypeScript strict mode
- Semua komponen pakai `'use client'` hanya jika butuh interaktivitas
- Gunakan Server Components untuk halaman yang hanya fetch data

### Penamaan
- Komponen: `PascalCase` (contoh: `ResultPanel.tsx`)
- Fungsi/variabel: `camelCase`
- File non-komponen: `camelCase` (contoh: `api.ts`, `types.ts`)
- CSS class: gunakan Tailwind utility, hindari custom CSS kecuali terpaksa

### Struktur Komponen
```tsx
// Urutan yang disarankan dalam file komponen:
// 1. Import
// 2. Types/Interface
// 3. Komponen utama
// 4. Export default
```

### Environment Variables
```env
# .env.local
NEXT_PUBLIC_API_URL=http://localhost:3001
```

---

## Checklist Sebelum Deploy

- [ ] Semua API call sudah ada error handling
- [ ] Loading state sudah ditampilkan
- [ ] Tidak ada `console.log` yang tertinggal
- [ ] Semua tipe sudah didefinisikan (tidak ada `any`)
- [ ] Responsif di layar 1280px ke atas (prioritas desktop)
- [ ] Monaco Editor tidak crash di SSR (gunakan dynamic import)

```tsx
// Cara import Monaco yang aman untuk SSR
import dynamic from 'next/dynamic';
const CodeEditor = dynamic(() => import('@/components/Editor/CodeEditor'), { ssr: false });
```
