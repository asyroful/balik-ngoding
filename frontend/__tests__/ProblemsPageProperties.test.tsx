// @vitest-environment happy-dom
// Feature: user-experience-improvements, Property 12: Category change triggers server fetch

import React from 'react';
import * as fc from 'fast-check';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';

// Mock @/lib/api
vi.mock('../lib/api', () => ({
  getProblems: vi.fn(),
}));

// Mock next/navigation
vi.mock('next/navigation', () => ({
  useRouter: () => ({ push: vi.fn() }),
}));

import ProblemsPage from '../app/problems/page';
import { getProblems } from '../lib/api';

const mockGetProblems = vi.mocked(getProblems);

beforeEach(() => {
  vi.resetAllMocks();
  mockGetProblems.mockResolvedValue([]);
});

// Validates: Requirements 4.2, 4.5
describe('Property 12: Category change triggers server fetch', () => {
  it('clicking any category tab calls getProblems with that category', async () => {
    await fc.assert(
      fc.asyncProperty(
        fc.constantFrom('loop', 'string', 'array', 'sql'),
        async (category) => {
          vi.resetAllMocks();
          mockGetProblems.mockResolvedValue([]);

          const { unmount } = render(<ProblemsPage />);

          // Wait for initial load to complete
          await waitFor(() => {
            expect(mockGetProblems).toHaveBeenCalledWith('loop');
          });

          // Click the tab for the generated category
          const tabLabel = category === 'sql' ? /^sql$/i : new RegExp(`^${category}$`, 'i');
          fireEvent.click(screen.getByRole('button', { name: tabLabel }));

          // Verify getProblems was called with the selected category
          await waitFor(() => {
            expect(mockGetProblems).toHaveBeenCalledWith(category);
          });

          unmount();
        }
      ),
      { numRuns: 100 }
    );
  }, 60000); // 60s timeout for 100 async component render iterations
});
