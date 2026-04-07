// @vitest-environment happy-dom
// Feature: user-experience-improvements, Property 11: Accepted problems show visual indicator in table

import React from 'react';
import * as fc from 'fast-check';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, within } from '@testing-library/react';

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
});

// Validates: Requirements 3.4, 3.5, 3.6
describe('Property 11: Accepted problems show visual indicator in table', () => {
  it('each accepted row shows "Selesai" badge; non-accepted rows do not', () => {
    fc.assert(
      fc.property(
        fc.array(arbProblem, { minLength: 1, maxLength: 10 }),
        (problems) => {
          // Accept every other problem (deterministic subset)
          const acceptedIds = new Set(problems.filter((_, i) => i % 2 === 0).map((p) => p.id));

          mockUseProgress.mockReturnValue({
            progress: Object.fromEntries([...acceptedIds].map((id) => [id, 'accepted'])),
            markAccepted: vi.fn(),
            isAccepted: (id: string) => acceptedIds.has(id),
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

            const hasSelesai = within(row as HTMLElement).queryByText(/Selesai/i) !== null;

            if (acceptedIds.has(problem.id)) {
              if (!hasSelesai) result = false;
            } else {
              if (hasSelesai) result = false;
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
