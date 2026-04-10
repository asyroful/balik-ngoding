// @vitest-environment happy-dom
// Feature: user-experience-improvements, Property 11: Accepted problems show visual indicator in table

import React from 'react';
import * as fc from 'fast-check';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render } from '@testing-library/react';

// Mock useProgress before importing ProblemTable
vi.mock('../hooks/useProgress', () => ({
  useProgress: vi.fn(),
}));

import ProblemTable from '../components/ProblemList/ProblemTable';
import { useProgress } from '../hooks/useProgress';

const mockUseProgress = useProgress as ReturnType<typeof vi.fn>;

beforeEach(() => {
  vi.clearAllMocks();
});

// Arbitraries
const arbCategory = fc.constantFrom('loop', 'string', 'array', 'sql') as fc.Arbitrary<'loop' | 'string' | 'array' | 'sql'>;
const arbDifficulty = fc.constantFrom('easy', 'medium', 'hard') as fc.Arbitrary<'easy' | 'medium' | 'hard'>;

const arbProblem = fc.record({
  id: fc.uuid(),
  title: fc.string({ minLength: 1, maxLength: 40 }),
  category: arbCategory,
  difficulty: arbDifficulty,
  description: fc.constant(''),
  starterCode: fc.constant(''),
  isActive: fc.constant(true),
  createdAt: fc.constant('2024-01-01T00:00:00Z'),
  thinkingGuide: fc.constant(null),
  hints: fc.constant(null),
  prerequisiteId: fc.constant(null),
  prerequisiteTitle: fc.constant(null),
});

// Validates: Requirements 3.4, 3.5, 3.6
describe('Property 11: Accepted problems show visual indicator in table', () => {
  it('each accepted row shows green checkmark SVG; non-accepted rows show gray circle SVG', () => {
    fc.assert(
      fc.property(
        fc.array(arbProblem, { minLength: 1, maxLength: 10 }),
        (problems) => {
          // Accept every other problem (deterministic subset)
          const acceptedIds = new Set(problems.filter((_, i) => i % 2 === 0).map((p) => p.id));

          mockUseProgress.mockReturnValue({
            markAccepted: vi.fn(),
            isAccepted: (id: string) => acceptedIds.has(id),
            getHintsUnlocked: vi.fn().mockReturnValue(0),
            unlockNextHint: vi.fn(),
          });

          const { container, unmount } = render(
            <ProblemTable problems={problems} onRowClick={() => {}} />
          );

          // Get all data rows (skip the header row)
          const rows = container.querySelectorAll('tbody tr');

          let result = true;
          problems.forEach((problem, index) => {
            const row = rows[index];
            if (!row) { result = false; return; }

            const greenCheck = row.querySelector('svg.text-green-600');
            const grayCircle = row.querySelector('svg.text-gray-300');

            if (acceptedIds.has(problem.id)) {
              if (!greenCheck) result = false;
            } else {
              if (grayCircle === null) result = false;
            }
          });

          unmount();
          return result;
        }
      ),
      { numRuns: 100 }
    );
  });
});
