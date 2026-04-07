# Senior Backend Engineer — LogicLab MVP

## Stack

- **Framework**: NestJS (Node.js)
- **Database**: PostgreSQL
- **ORM**: TypeORM
- **Language**: TypeScript
- **Runtime**: Node.js 20+

---

## Arsitektur Backend (Clean Architecture Sederhana)

```
src/
├── main.ts                         # Entry point
├── app.module.ts                   # Root module
├── problems/
│   ├── problems.module.ts
│   ├── problems.controller.ts      # Handle HTTP request
│   ├── problems.service.ts         # Business logic
│   ├── problems.repository.ts      # Query database
│   └── dto/
│       └── problem.dto.ts
├── submissions/
│   ├── submissions.module.ts
│   ├── submissions.controller.ts
│   ├── submissions.service.ts
│   ├── submissions.repository.ts
│   └── dto/
│       └── submit.dto.ts
├── evaluator/
│   ├── evaluator.module.ts
│   └── evaluator.service.ts        # Logic evaluasi kode (mock di MVP)
├── entities/
│   ├── problem.entity.ts
│   ├── test-case.entity.ts
│   ├── submission.entity.ts
│   └── user.entity.ts
└── common/
    ├── filters/
    │   └── http-exception.filter.ts
    └── interceptors/
        └── response.interceptor.ts
```

---

## Entities

### `problem.entity.ts`

```typescript
import { Entity, PrimaryGeneratedColumn, Column, OneToMany, CreateDateColumn } from 'typeorm';
import { TestCase } from './test-case.entity';

@Entity('problems')
export class Problem {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  title: string;

  @Column('text')
  description: string;

  @Column({ type: 'enum', enum: ['loop', 'string', 'array', 'sql'] })
  category: string;

  @Column({ type: 'enum', enum: ['easy', 'medium', 'hard'], default: 'easy' })
  difficulty: string;

  @Column('text')
  starterCode: string;

  @Column({ default: true })
  isActive: boolean;

  @OneToMany(() => TestCase, (tc) => tc.problem, { cascade: true })
  testCases: TestCase[];

  @CreateDateColumn()
  createdAt: Date;
}
```

### `test-case.entity.ts`

```typescript
import { Entity, PrimaryGeneratedColumn, Column, ManyToOne } from 'typeorm';
import { Problem } from './problem.entity';

@Entity('test_cases')
export class TestCase {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @ManyToOne(() => Problem, (p) => p.testCases, { onDelete: 'CASCADE' })
  problem: Problem;

  @Column('text')
  input: string;

  @Column('text')
  expectedOutput: string;

  @Column({ default: false })
  isHidden: boolean;  // hidden test case tidak ditampilkan ke user
}
```

### `submission.entity.ts`

```typescript
import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn } from 'typeorm';

@Entity('submissions')
export class Submission {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  problemId: string;

  @Column('text')
  code: string;

  @Column({ default: 'javascript' })
  language: string;

  @Column({ type: 'enum', enum: ['accepted', 'wrong_answer', 'error'] })
  status: string;

  @Column({ type: 'int', default: 0 })
  score: number;

  @Column({ type: 'int', default: 0 })
  total: number;

  @Column('jsonb', { nullable: true })
  resultDetail: object;  // simpan detail per test case

  @CreateDateColumn()
  createdAt: Date;
}
```

### `user.entity.ts` (minimal, untuk future use)

```typescript
import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn } from 'typeorm';

@Entity('users')
export class User {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ unique: true })
  email: string;

  @Column()
  name: string;

  @CreateDateColumn()
  createdAt: Date;
}
```

---

## API Endpoints

### `GET /problems`

Mengambil semua soal aktif.

**Query params (opsional):**
- `category` — filter berdasarkan kategori (`loop`, `string`, `array`, `sql`)
- `difficulty` — filter berdasarkan difficulty

**Response:**
```json
{
  "data": [
    {
      "id": "uuid",
      "title": "FizzBuzz",
      "category": "loop",
      "difficulty": "easy"
    }
  ]
}
```

---

### `GET /problems/:id`

Mengambil detail satu soal beserta contoh test case (non-hidden).

**Response:**
```json
{
  "data": {
    "id": "uuid",
    "title": "FizzBuzz",
    "description": "Tulis program yang...",
    "category": "loop",
    "difficulty": "easy",
    "starterCode": "function solution(n) {\n  // kode di sini\n}",
    "examples": [
      {
        "input": "n = 5",
        "expectedOutput": "1 2 Fizz 4 Buzz"
      }
    ],
    "constraints": ["1 ≤ n ≤ 1000"]
  }
}
```

---

### `POST /submit`

Submit jawaban user untuk dievaluasi.

**Request body:**
```json
{
  "problemId": "uuid",
  "code": "function solution(n) { ... }",
  "language": "javascript"
}
```

**Response:**
```json
{
  "data": {
    "status": "wrong_answer",
    "score": 1,
    "total": 3,
    "results": [
      {
        "passed": true,
        "input": "n = 5",
        "expected": "1 2 Fizz 4 Buzz",
        "actual": "1 2 Fizz 4 Buzz"
      },
      {
        "passed": false,
        "input": "n = 15",
        "expected": "1 2 Fizz 4 Buzz 6 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz",
        "actual": "1 2 3 4 5 ..."
      }
    ]
  }
}
```

---

## Logic Evaluasi Submission

### Catatan Penting — MVP

> Di MVP, eksekusi kode user dilakukan dengan **mock evaluator** atau menggunakan library sandbox ringan seperti `vm2` atau `isolated-vm`. Jangan eksekusi kode user langsung di Node.js process utama.

### Alur Evaluasi

```
POST /submit
    ↓
Validasi input (problemId valid, code tidak kosong)
    ↓
Ambil semua test case dari database (termasuk hidden)
    ↓
Untuk setiap test case:
    → Jalankan kode user dengan input test case (via sandbox)
    → Bandingkan output dengan expectedOutput
    → Catat: passed/failed, actual output
    ↓
Hitung score (jumlah passed / total)
    ↓
Tentukan status: accepted / wrong_answer / error
    ↓
Simpan submission ke database
    ↓
Return hasil ke frontend
```

### Implementasi Mock Evaluator (MVP)

```typescript
// evaluator/evaluator.service.ts
import { Injectable } from '@nestjs/common';

@Injectable()
export class EvaluatorService {
  async evaluate(code: string, input: string, language: string): Promise<string> {
    if (language !== 'javascript') {
      throw new Error('Hanya JavaScript yang didukung di MVP');
    }

    try {
      // PERINGATAN: ini hanya untuk MVP/development
      // Ganti dengan vm2 atau isolated-vm untuk keamanan
      const fn = new Function(`
        ${code}
        return solution(${input});
      `);
      const result = fn();
      return String(result).trim();
    } catch (err) {
      throw new Error(`Runtime error: ${err.message}`);
    }
  }
}
```

> **TODO setelah MVP**: Ganti `new Function()` dengan `isolated-vm` atau Docker sandbox untuk keamanan produksi.

---

## Struktur Database

### Tabel & Relasi

```sql
-- Tabel problems
CREATE TABLE problems (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  category VARCHAR(50) NOT NULL CHECK (category IN ('loop', 'string', 'array', 'sql')),
  difficulty VARCHAR(20) NOT NULL DEFAULT 'easy',
  starter_code TEXT NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT NOW()
);

-- Tabel test_cases
CREATE TABLE test_cases (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  problem_id UUID NOT NULL REFERENCES problems(id) ON DELETE CASCADE,
  input TEXT NOT NULL,
  expected_output TEXT NOT NULL,
  is_hidden BOOLEAN DEFAULT FALSE
);

-- Tabel submissions
CREATE TABLE submissions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  problem_id UUID NOT NULL,
  code TEXT NOT NULL,
  language VARCHAR(50) DEFAULT 'javascript',
  status VARCHAR(50) NOT NULL,
  score INT DEFAULT 0,
  total INT DEFAULT 0,
  result_detail JSONB,
  created_at TIMESTAMP DEFAULT NOW()
);

-- Tabel users (untuk future use)
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(255) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
```

### Relasi

```
problems (1) ──── (N) test_cases
problems (1) ──── (N) submissions
users    (1) ──── (N) submissions  [future]
```

---

## Konfigurasi Environment

```env
# .env
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=secret
DATABASE_NAME=logiclab
PORT=3001
```

---

## Validasi Input (DTO)

```typescript
// submissions/dto/submit.dto.ts
import { IsString, IsNotEmpty, IsUUID } from 'class-validator';

export class SubmitDto {
  @IsUUID()
  problemId: string;

  @IsString()
  @IsNotEmpty()
  code: string;

  @IsString()
  language: string = 'javascript';
}
```

---

## Checklist Sebelum Deploy

- [ ] Semua endpoint sudah ada validasi input (class-validator)
- [ ] Error response konsisten menggunakan exception filter
- [ ] Evaluator tidak mengeksekusi kode di main process (gunakan sandbox)
- [ ] Database sudah ada index pada `problem_id` di tabel `test_cases` dan `submissions`
- [ ] CORS sudah dikonfigurasi untuk domain frontend
- [ ] Environment variables tidak di-commit ke git
