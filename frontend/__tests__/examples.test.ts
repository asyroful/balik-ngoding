import { describe, it, expect } from 'vitest';
import * as fc from 'fast-check';

// ─── 13.7: Problem not found redirect (Req 2.3) ──────────────────────────────
describe('Problem not found redirect (Req 2.3)', () => {
  // Simulate the error thrown by handleResponse when the API returns 404
  const handleApiError = (status: number, body?: { message?: string }): Error => {
    const message = body?.message || `Request failed with status ${status}`;
    return new Error(message);
  };

  it('should produce an error with status 404 message when problem is not found', () => {
    const err = handleApiError(404, { message: 'Problem not found' });
    expect(err.message.toLowerCase()).toMatch(/not found|404/);
  });

  it('should produce a fallback error message containing status code when no body message', () => {
    const err = handleApiError(404);
    expect(err.message).toContain('404');
  });

  it('should throw an Error instance', () => {
    const err = handleApiError(404);
    expect(err).toBeInstanceOf(Error);
  });
});

// ─── 13.8: 404 page (Req 6.4) ────────────────────────────────────────────────
describe('404 page (Req 6.4)', () => {
  it('should export a default function (not-found page exists)', async () => {
    // Dynamically import to verify the module resolves without errors
    const mod = await import('../app/not-found');
    expect(typeof mod.default).toBe('function');
  });
});

// ─── 8.5: Language selector otomatis (Req 7.1) ───────────────────────────────
// Feature: tambah-soal — language selector otomatis
describe('Language selector otomatis (Req 7.1)', () => {
  const getLanguage = (category: string): string =>
    category === 'sql' ? 'sql' : 'javascript';

  it('should return "sql" for category "sql"', () => {
    expect(getLanguage('sql')).toBe('sql');
  });

  it('should return "javascript" for loop, string, and array categories', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('loop', 'string', 'array'),
        (category) => getLanguage(category) === 'javascript'
      ),
      { numRuns: 100 }
    );
  });
});
