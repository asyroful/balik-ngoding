// @vitest-environment happy-dom
import * as fc from 'fast-check';
import { describe, it, expect, beforeEach } from 'vitest';
import {
  applyMarkAccepted,
  applyUnlockNextHint,
  readFromStorage,
  writeToStorage,
  STORAGE_KEY,
  type ProgressStore,
} from '../hooks/useProgress';

beforeEach(() => {
  localStorage.clear();
});

// Arbitrary for a valid ProgressStore
const progressStoreArb = fc.dictionary(
  fc.uuid(),
  fc.record({
    solved: fc.boolean(),
    hintsUnlocked: fc.nat({ max: 3 }),
  })
);

// Feature: guided-learning, Property 10: ProgressStore round-trip persistence
// Validates: Requirements 6.2, 6.3, 6.5
describe('Property 10: ProgressStore round-trip persistence', () => {
  it('writeToStorage then readFromStorage yields identical state', () => {
    fc.assert(
      fc.property(progressStoreArb, (store) => {
        localStorage.clear();
        writeToStorage(store);
        const restored = readFromStorage();
        return JSON.stringify(restored) === JSON.stringify(store);
      }),
      { numRuns: 100 }
    );
  });
});

// Feature: guided-learning, Property 11: markAccepted idempotent
// Validates: Requirements 6.2
describe('Property 11: markAccepted idempotent', () => {
  it('calling applyMarkAccepted twice yields same state as once', () => {
    fc.assert(
      fc.property(progressStoreArb, fc.uuid(), (store, problemId) => {
        const once = applyMarkAccepted(store, problemId);
        const twice = applyMarkAccepted(once, problemId);
        return JSON.stringify(once) === JSON.stringify(twice);
      }),
      { numRuns: 100 }
    );
  });

  it('markAccepted always results in solved: true', () => {
    fc.assert(
      fc.property(progressStoreArb, fc.uuid(), (store, problemId) => {
        const updated = applyMarkAccepted(store, problemId);
        return updated[problemId]?.solved === true;
      }),
      { numRuns: 100 }
    );
  });
});

// Feature: guided-learning, Property 12: hintsUnlocked hanya bertambah
// Validates: Requirements 2.4, 6.5
describe('Property 12: hintsUnlocked only increases, never decreases', () => {
  it('sequence of applyUnlockNextHint calls is monotonically non-decreasing', () => {
    fc.assert(
      fc.property(
        fc.uuid(),
        fc.integer({ min: 1, max: 3 }),
        fc.integer({ min: 1, max: 6 }),
        (problemId, totalHints, callCount) => {
          let store: ProgressStore = {};
          const counts: number[] = [];

          for (let i = 0; i < callCount; i++) {
            store = applyUnlockNextHint(store, problemId, totalHints);
            counts.push(store[problemId]?.hintsUnlocked ?? 0);
          }

          // Monotonically non-decreasing
          for (let i = 1; i < counts.length; i++) {
            if (counts[i] < counts[i - 1]) return false;
          }

          // Never exceeds totalHints
          return counts.every((c) => c <= totalHints);
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: guided-learning, Property 13: Isolasi antar soal
// Validates: Requirements 6.4
describe('Property 13: Isolation between problems', () => {
  it('updating problem A does not change problem B', () => {
    fc.assert(
      fc.property(
        progressStoreArb,
        fc.tuple(fc.uuid(), fc.uuid()).filter(([a, b]) => a !== b),
        fc.boolean(),
        (store, [problemA, problemB], useMarkAccepted) => {
          const bBefore = store[problemB];

          const updated = useMarkAccepted
            ? applyMarkAccepted(store, problemA)
            : applyUnlockNextHint(store, problemA, 3);

          const bAfter = updated[problemB];

          return JSON.stringify(bBefore) === JSON.stringify(bAfter);
        }
      ),
      { numRuns: 100 }
    );
  });
});
