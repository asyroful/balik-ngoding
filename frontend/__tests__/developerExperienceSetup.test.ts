import * as fc from 'fast-check';
import * as fs from 'fs';
import * as path from 'path';
import { describe, it, expect } from 'vitest';

// Resolve root project path (two levels up from frontend/__tests__)
const ROOT = path.resolve(__dirname, '../../');
const STEERING_DIR = path.join(ROOT, '.kiro/steering');

const readSteering = (filename: string): string =>
  fs.readFileSync(path.join(STEERING_DIR, filename), 'utf-8');

// ─── Property 4: Steering files tidak mengandung referensi stack lama ────────
// Feature: developer-experience-setup, Property 4: Steering files tidak mengandung referensi stack lama
// Validates: Requirements 3.5
describe('Property 4: Steering files tidak mengandung referensi stack lama', () => {
  it('tidak ada kata kunci stack lama di semua steering files', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('NestJS', 'TypeORM', 'LogicLab', 'nestjs', 'typeorm'),
        fc.constantFrom('backend.md', 'frontend.md', 'project-overview.md'),
        (keyword, file) => {
          const content = readSteering(file);
          return !content.includes(keyword);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 5: backend.md mendokumentasikan semua teknologi backend ─────────
// Feature: developer-experience-setup, Property 5: backend.md mendokumentasikan semua teknologi backend
// Validates: Requirements 3.2, 3.7
describe('Property 5: backend.md mendokumentasikan semua teknologi backend', () => {
  it('semua kata kunci teknologi backend ada di backend.md', () => {
    fc.assert(
      fc.property(
        fc.constantFrom(
          'Go',
          'Gin',
          'GORM',
          'PostgreSQL',
          'DATABASE_URL',
          'PORT',
          'FRONTEND_ORIGIN',
          'backend/internal'
        ),
        (keyword) => {
          const content = readSteering('backend.md');
          return content.includes(keyword);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 6: frontend.md mendokumentasikan semua teknologi frontend ───────
// Feature: developer-experience-setup, Property 6: frontend.md mendokumentasikan semua teknologi frontend
// Validates: Requirements 3.3, 3.8
describe('Property 6: frontend.md mendokumentasikan semua teknologi frontend', () => {
  it('semua kata kunci teknologi frontend ada di frontend.md', () => {
    fc.assert(
      fc.property(
        fc.constantFrom(
          'Next.js',
          'TypeScript',
          'Tailwind',
          'Zustand',
          'Monaco',
          'NEXT_PUBLIC_API_URL',
          'App Router'
        ),
        (keyword) => {
          const content = readSteering('frontend.md');
          return content.includes(keyword);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap ──────
// Feature: developer-experience-setup, Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap
// Validates: Requirements 4.1, 4.3, 4.5
describe('Property 7: mcp.json valid dengan konfigurasi postgres yang lengkap', () => {
  it('mcp.json adalah JSON valid dan mengandung semua field yang diperlukan', () => {
    const mcpPath = path.join(ROOT, '.kiro/settings/mcp.json');
    const raw = fs.readFileSync(mcpPath, 'utf-8');
    const parsed = JSON.parse(raw); // throws if invalid JSON

    fc.assert(
      fc.property(
        fc.constantFrom('mcpServers', 'postgres', 'command', 'args'),
        (field) => {
          if (field === 'mcpServers') return field in parsed;
          if (field === 'postgres') return field in parsed.mcpServers;
          if (field === 'command') return field in parsed.mcpServers.postgres;
          if (field === 'args') return field in parsed.mcpServers.postgres;
          return false;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 3: frontend/Dockerfile multi-stage build yang benar ─────────────
// Feature: developer-experience-setup, Property 3: frontend/Dockerfile multi-stage build yang benar
// Validates: Requirements 2.1, 2.2, 2.3, 2.4, 2.5
describe('Property 3: frontend/Dockerfile multi-stage build yang benar', () => {
  const dockerfilePath = path.resolve(__dirname, '../Dockerfile');
  const dockerfileContent = fs.readFileSync(dockerfilePath, 'utf-8');

  it('setiap stage ada, menggunakan node:20-alpine, dan memiliki instruksi yang benar', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('deps', 'builder', 'runner'),
        (stage) => {
          // Each stage must exist as a named stage
          const stageExists = dockerfileContent.includes(`AS ${stage}`);
          if (!stageExists) return false;

          // Each stage must use node:20-alpine as base image
          const stagePattern = new RegExp(
            `FROM node:20-alpine AS ${stage}[\\s\\S]*?(?=FROM |$)`
          );
          const stageBlock = dockerfileContent.match(stagePattern)?.[0] ?? '';
          if (!stageBlock) return false;

          // Stage-specific checks
          if (stage === 'runner') {
            // runner must EXPOSE 3000
            return stageBlock.includes('EXPOSE 3000');
          }

          if (stage === 'builder') {
            // builder must have ARG NEXT_PUBLIC_API_URL with default http://localhost:8080
            return stageBlock.includes('ARG NEXT_PUBLIC_API_URL=http://localhost:8080');
          }

          // deps stage: just needs to exist with node:20-alpine (already checked above)
          return true;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan ──
// Feature: developer-experience-setup, Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan
// Validates: Requirements 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 1.12
describe('Property 1: docker-compose.yml berisi semua konfigurasi yang diperlukan', () => {
  it('setiap service terdefinisi di docker-compose.yml', () => {
    const composePath = path.resolve(__dirname, '../../docker-compose.yml');
    const composeContent = fs.readFileSync(composePath, 'utf-8');

    fc.assert(
      fc.property(
        fc.constantFrom('db', 'backend', 'frontend'),
        (service) => {
          return composeContent.includes(service);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 2: Setiap env var memiliki nilai default ────────────────────────
// Feature: developer-experience-setup, Property 2: Setiap env var memiliki nilai default
// Validates: Requirements 1.11
describe('Property 2: Setiap env var memiliki nilai default', () => {
  it('setiap env var menggunakan sintaks ${VAR:-default} di docker-compose.yml', () => {
    const composePath = path.resolve(__dirname, '../../docker-compose.yml');
    const composeContent = fs.readFileSync(composePath, 'utf-8');

    fc.assert(
      fc.property(
        fc.constantFrom(
          'DATABASE_URL',
          'PORT',
          'FRONTEND_ORIGIN',
          'NEXT_PUBLIC_API_URL',
          'POSTGRES_USER',
          'POSTGRES_PASSWORD',
          'POSTGRES_DB'
        ),
        (envVar) => {
          return composeContent.includes(`\${${envVar}:-`);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── Property 8: README berisi semua konten setup yang diperlukan ─────────────
// Feature: developer-experience-setup, Property 8: README berisi semua konten setup yang diperlukan
// Validates: Requirements 5.1, 5.2, 5.3, 5.4, 5.5, 5.6
describe('Property 8: README berisi semua konten setup yang diperlukan', () => {
  it('semua konten setup yang diperlukan ada di README.md', () => {
    const readmePath = path.resolve(__dirname, '../../README.md');
    const readmeContent = fs.readFileSync(readmePath, 'utf-8');

    fc.assert(
      fc.property(
        fc.constantFrom(
          'docker compose up',
          'docker compose down',
          'docker compose logs',
          'npm test',
          'go test',
          'Troubleshooting'
        ),
        (content) => {
          return readmeContent.includes(content);
        }
      ),
      { numRuns: 100 }
    );
  });
});
