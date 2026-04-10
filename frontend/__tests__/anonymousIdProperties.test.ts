// @vitest-environment happy-dom
import * as fc from 'fast-check';
import { describe, it, expect, beforeEach, vi } from 'vitest';
import {
  getOrCreateAnonymousId,
  ANONYMOUS_ID_KEY,
  isValidUUIDv4,
} from '../lib/anonymousId';

const UUID_V4_REGEX = /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i;

beforeEach(() => {
  localStorage.clear();
});

// Feature: anonymous-analytics, Property 1: UUID v4 generation
// Validates: Requirements 1.1, 1.4
describe('Property 1: UUID v4 generation', () => {
  it('TestUuidV4Generation — fresh state always yields a valid UUID v4', () => {
    fc.assert(
      fc.property(fc.constant(null), () => {
        localStorage.clear();
        const id = getOrCreateAnonymousId();
        return UUID_V4_REGEX.test(id);
      }),
      { numRuns: 100 }
    );
  });

  it('stores the generated ID under the correct key', () => {
    localStorage.clear();
    const id = getOrCreateAnonymousId();
    expect(localStorage.getItem(ANONYMOUS_ID_KEY)).toBe(id);
  });
});

// Feature: anonymous-analytics, Property 2: localStorage round-trip
// Validates: Requirements 1.2, 5.1
describe('Property 2: localStorage round-trip', () => {
  it('TestLocalStorageRoundTrip — stored UUID v4 is returned unchanged', () => {
    // Use only valid UUID v4 strings (version 4, variant [89ab])
    const uuidV4Arb = fc.stringMatching(
      /^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/
    );
    fc.assert(
      fc.property(uuidV4Arb, (uuid) => {
        localStorage.clear();
        localStorage.setItem(ANONYMOUS_ID_KEY, uuid);
        const result = getOrCreateAnonymousId();
        return result === uuid;
      }),
      { numRuns: 100 }
    );
  });
});

// Feature: anonymous-analytics, Property 3: anonymousId included in submit request
// Validates: Requirements 2.1
describe('Property 3: anonymousId included in submit request', () => {
  it('TestAnonymousIdIncludedInRequest — submitSolution body always contains anonymousId', async () => {
    await fc.assert(
      fc.asyncProperty(
        fc.record({
          problemId: fc.uuid(),
          code: fc.string({ minLength: 1 }),
          language: fc.constantFrom('javascript', 'python'),
        }),
        async (req) => {
          localStorage.clear();

          let capturedBody: Record<string, unknown> | null = null;
          const mockFetch = vi.fn().mockImplementation((_url: string, options: RequestInit) => {
            capturedBody = JSON.parse(options.body as string);
            return Promise.resolve({
              ok: true,
              json: () =>
                Promise.resolve({
                  data: { status: 'accepted', score: 1, total: 1, results: [] },
                }),
            });
          });

          const originalFetch = global.fetch;
          global.fetch = mockFetch as typeof fetch;

          try {
            const { submitSolution } = await import('../lib/api');
            await submitSolution(req).catch(() => {});
          } finally {
            global.fetch = originalFetch;
          }

          return (
            capturedBody !== null &&
            typeof (capturedBody as Record<string, unknown>).anonymousId === 'string' &&
            UUID_V4_REGEX.test((capturedBody as Record<string, unknown>).anonymousId as string)
          );
        }
      ),
      { numRuns: 20 }
    );
  });
});

// Unit test: localStorage unavailable
// Validates: Requirements 1.3
describe('TestGetOrCreateAnonymousId_LocalStorageUnavailable', () => {
  it('returns a valid UUID even when localStorage throws', () => {
    const getItem = vi.spyOn(Storage.prototype, 'getItem').mockImplementation(() => {
      throw new Error('Storage blocked');
    });
    const setItem = vi.spyOn(Storage.prototype, 'setItem').mockImplementation(() => {
      throw new Error('Storage blocked');
    });

    let result: string | undefined;
    expect(() => {
      result = getOrCreateAnonymousId();
    }).not.toThrow();

    expect(result).toBeDefined();
    expect(UUID_V4_REGEX.test(result!)).toBe(true);

    getItem.mockRestore();
    setItem.mockRestore();
  });
});
