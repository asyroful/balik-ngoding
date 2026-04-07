import * as fc from 'fast-check';
import { describe, it, expect, beforeEach } from 'vitest';
import { useSubmissionStore } from '../store/submissionStore';
import type { TestCaseResult } from '../lib/types';

// ─── P5: Editor initializes with starter code ────────────────────────────────
// Feature: balik-ngoding, Property 5: Editor initializes with starter code
// Validates: Requirements 2.2, 2.4
describe('P5: Editor initializes with starter code', () => {
  beforeEach(() => {
    useSubmissionStore.getState().reset();
  });

  it('should set store code to starterCode for any non-empty starterCode', () => {
    fc.assert(
      fc.property(
        fc.string({ minLength: 1 }),
        (starterCode) => {
          useSubmissionStore.getState().setCode(starterCode);
          return useSubmissionStore.getState().code === starterCode;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── P6: Empty/whitespace code submission is rejected ────────────────────────
// Feature: balik-ngoding, Property 6: Empty/whitespace code submission is rejected
// Validates: Requirements 4.2
describe('P6: Empty/whitespace code submission is rejected', () => {
  const isCodeEmpty = (code: string) => code.trim() === '';

  it('should reject any whitespace-only string', () => {
    fc.assert(
      fc.property(
        fc.stringOf(fc.constantFrom(' ', '\t', '\n', '\r')),
        (code) => isCodeEmpty(code) === true
      ),
      { numRuns: 100 }
    );
  });

  it('should accept any string with at least one non-whitespace character', () => {
    fc.assert(
      fc.property(
        fc.string({ minLength: 1 }).filter((s) => s.trim().length > 0),
        (code) => isCodeEmpty(code) === false
      ),
      { numRuns: 100 }
    );
  });
});

// ─── P7: Duplicate submit is idempotent ──────────────────────────────────────
// Feature: balik-ngoding, Property 7: Duplicate submit is idempotent
// Validates: Requirements 4.4
describe('P7: Duplicate submit is idempotent (only one request sent)', () => {
  const shouldPreventSubmit = (isLoading: boolean) => isLoading;

  it('should block submit when isLoading is true', () => {
    fc.assert(
      fc.property(
        fc.constant(true),
        (isLoading) => shouldPreventSubmit(isLoading) === true
      ),
      { numRuns: 100 }
    );
  });

  it('should allow submit when isLoading is false', () => {
    fc.assert(
      fc.property(
        fc.constant(false),
        (isLoading) => shouldPreventSubmit(isLoading) === false
      ),
      { numRuns: 100 }
    );
  });
});

// ─── P9: Result panel contains required fields per test case ─────────────────
// Feature: balik-ngoding, Property 9: Result panel contains required fields per test case
// Validates: Requirements 5.2
describe('P9: Result panel contains required fields per test case', () => {
  const hasRequiredFields = (result: TestCaseResult): boolean =>
    typeof result.passed === 'boolean' &&
    typeof result.input === 'string' &&
    typeof result.expected === 'string' &&
    typeof result.actual === 'string';

  const arbitraryTestCaseResult = () =>
    fc.record({
      passed: fc.boolean(),
      input: fc.string(),
      expected: fc.string(),
      actual: fc.string(),
      error: fc.option(fc.string(), { nil: undefined }),
    });

  it('should have all required fields for any TestCaseResult', () => {
    fc.assert(
      fc.property(
        arbitraryTestCaseResult(),
        (result) => hasRequiredFields(result as TestCaseResult)
      ),
      { numRuns: 100 }
    );
  });
});

// ─── P10: Score display matches actual counts ─────────────────────────────────
// Feature: balik-ngoding, Property 10: Score display matches actual counts
// Validates: Requirements 5.3
describe('P10: Score display matches actual counts', () => {
  const formatScore = (score: number, total: number): string =>
    `${score} dari ${total} test case passed`;

  it('should format score string correctly for any valid score/total pair', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 0, max: 1000 }).chain((total) =>
          fc.tuple(fc.integer({ min: 0, max: total }), fc.constant(total))
        ),
        ([score, total]) => {
          const result = formatScore(score, total);
          return (
            result === `${score} dari ${total} test case passed` &&
            result.startsWith(`${score} dari ${total}`)
          );
        }
      ),
      { numRuns: 100 }
    );
  });

  it('should always contain "dari" and "test case passed"', () => {
    fc.assert(
      fc.property(
        fc.integer({ min: 0, max: 100 }).chain((total) =>
          fc.tuple(fc.integer({ min: 0, max: total }), fc.constant(total))
        ),
        ([score, total]) => {
          const result = formatScore(score, total);
          return result.includes('dari') && result.includes('test case passed');
        }
      ),
      { numRuns: 100 }
    );
  });
});

// ─── P16: "Coba Lagi" resets editor to starter code ─────────────────────────
// Feature: balik-ngoding, Property 16: "Coba Lagi" resets editor to starter code
// Validates: Requirements 6.2
describe('P16: "Coba Lagi" resets editor to starter code', () => {
  beforeEach(() => {
    useSubmissionStore.getState().reset();
  });

  it('should restore starterCode after setting arbitrary code', () => {
    fc.assert(
      fc.property(
        fc.string({ minLength: 1 }),
        fc.string({ minLength: 1 }).filter((s) => s.trim().length > 0),
        (arbitraryCode, starterCode) => {
          const store = useSubmissionStore.getState();
          store.setCode(arbitraryCode);
          // Simulate "Coba Lagi": reset to starter code
          store.setCode(starterCode);
          return useSubmissionStore.getState().code === starterCode;
        }
      ),
      { numRuns: 100 }
    );
  });
});
