// @vitest-environment happy-dom
import * as fc from 'fast-check';
import { describe, it, expect, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { useProgress } from '../hooks/useProgress';

// Clear cookies between tests
beforeEach(() => {
  document.cookie.split(';').forEach((cookie) => {
    const name = cookie.split('=')[0].trim();
    if (name) {
      document.cookie = `${name}=; max-age=0; path=/`;
    }
  });
});

// Feature: user-experience-improvements, Property 8: markAccepted round-trip
// Validates: Requirements 3.1, 3.2
describe('Property 8: markAccepted round-trip', () => {
  it('after markAccepted(id), bn_progress cookie can be parsed and contains id: "accepted"', () => {
    fc.assert(
      fc.property(
        fc.uuid(),
        (problemId) => {
          // Clear cookies before each run
          document.cookie.split(';').forEach((cookie) => {
            const name = cookie.split('=')[0].trim();
            if (name) {
              document.cookie = `${name}=; max-age=0; path=/`;
            }
          });

          const { result } = renderHook(() => useProgress());

          act(() => {
            result.current.markAccepted(problemId);
          });

          const match = document.cookie
            .split('; ')
            .find((row) => row.startsWith('bn_progress='));

          if (!match) return false;

          const value = decodeURIComponent(match.split('=').slice(1).join('='));
          let parsed: Record<string, unknown>;
          try {
            parsed = JSON.parse(value);
          } catch {
            return false;
          }

          return parsed[problemId] === 'accepted';
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 9: Accepted status is preserved
// Validates: Requirements 3.8
describe('Property 9: Accepted status is preserved', () => {
  it('after markAccepted(id), isAccepted(id) is always true without calling it again', () => {
    fc.assert(
      fc.property(
        fc.uuid(),
        (problemId) => {
          // Clear cookies before each run
          document.cookie.split(';').forEach((cookie) => {
            const name = cookie.split('=')[0].trim();
            if (name) {
              document.cookie = `${name}=; max-age=0; path=/`;
            }
          });

          const { result } = renderHook(() => useProgress());

          act(() => {
            result.current.markAccepted(problemId);
          });

          return result.current.isAccepted(problemId) === true;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 10: Invalid cookie graceful fallback
// Validates: Requirements 3.7
describe('Property 10: Invalid cookie graceful fallback', () => {
  it('any non-JSON string as cookie value does not throw and results in progress === {}', () => {
    fc.assert(
      fc.property(
        fc.oneof(
          fc.string().filter((s) => {
            try {
              JSON.parse(s);
              return false;
            } catch {
              return true;
            }
          }),
          fc.constant(''),
          fc.constant('{invalid'),
          fc.constant('not-json'),
          fc.constant('undefined'),
          fc.constant('null'),
          fc.constant('[1,2,3]')
        ),
        (invalidValue) => {
          // Clear cookies before each run
          document.cookie.split(';').forEach((cookie) => {
            const name = cookie.split('=')[0].trim();
            if (name) {
              document.cookie = `${name}=; max-age=0; path=/`;
            }
          });

          document.cookie = `bn_progress=${encodeURIComponent(invalidValue)}; path=/`;

          let progressResult: Record<string, unknown> = { _error: true };
          let threw = false;

          try {
            const { result } = renderHook(() => useProgress());
            progressResult = result.current.progress;
          } catch {
            threw = true;
          }

          return !threw && JSON.stringify(progressResult) === '{}';
        }
      ),
      { numRuns: 100 }
    );
  });
});
