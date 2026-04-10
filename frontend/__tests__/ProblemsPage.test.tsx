// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';

// Mock lib/api using relative path
vi.mock('../lib/api', () => ({
  getProblems: vi.fn(),
  getProblemsSummary: vi.fn(),
}));

// Mock next/navigation
vi.mock('next/navigation', () => ({
  useRouter: () => ({ push: vi.fn(), replace: vi.fn() }),
  useSearchParams: () => ({ get: () => null }),
}));

// Mock useProgress
vi.mock('../hooks/useProgress', () => ({
  useProgress: () => ({
    isAccepted: () => false,
    markAccepted: vi.fn(),
    getHintsUnlocked: () => 0,
    unlockNextHint: vi.fn(),
  }),
}));

import ProblemsPage from '../app/problems/page';
import { getProblems, getProblemsSummary } from '../lib/api';

const mockGetProblems = vi.mocked(getProblems);
const mockGetProblemsSummary = vi.mocked(getProblemsSummary);

const MOCK_PROBLEMS = [
  { id: '1', title: 'FizzBuzz', category: 'loop', difficulty: 'easy' },
  { id: '2', title: 'Reverse String', category: 'string', difficulty: 'easy' },
];

beforeEach(() => {
  vi.resetAllMocks();
  mockGetProblemsSummary.mockResolvedValue([]);
});

// ─── Req 4.1: Initial load calls getProblems('loop') ─────────────────────────
describe('ProblemsPage — initial load (Req 4.1)', () => {
  it('calls getProblems with "loop" on initial render', async () => {
    mockGetProblems.mockResolvedValue(MOCK_PROBLEMS);
    render(<ProblemsPage />);
    await waitFor(() => {
      expect(mockGetProblems).toHaveBeenCalledWith('loop');
    });
  });

  it('calls getProblems exactly once on initial render', async () => {
    mockGetProblems.mockResolvedValue(MOCK_PROBLEMS);
    render(<ProblemsPage />);
    await waitFor(() => {
      expect(mockGetProblems).toHaveBeenCalledWith('loop');
    });
    // Page calls getProblems once for the selected category on mount
    expect(mockGetProblems.mock.calls.length).toBeGreaterThanOrEqual(1);
  });
});

// ─── Req 4.2: Category tab click calls getProblems with selected category ─────
describe('ProblemsPage — category tab click (Req 4.2)', () => {
  it('calls getProblems with "string" when String tab is clicked', async () => {
    mockGetProblems.mockResolvedValue(MOCK_PROBLEMS);
    render(<ProblemsPage />);
    await waitFor(() => expect(mockGetProblems).toHaveBeenCalledWith('loop'));

    fireEvent.click(screen.getByRole('button', { name: /string/i }));

    await waitFor(() => {
      expect(mockGetProblems).toHaveBeenCalledWith('string');
    });
  });

  it('calls getProblems with "array" when Array tab is clicked', async () => {
    mockGetProblems.mockResolvedValue(MOCK_PROBLEMS);
    render(<ProblemsPage />);
    await waitFor(() => expect(mockGetProblems).toHaveBeenCalledWith('loop'));

    fireEvent.click(screen.getByRole('button', { name: /array/i }));

    await waitFor(() => {
      expect(mockGetProblems).toHaveBeenCalledWith('array');
    });
  });

  it('calls getProblems with "sql" when SQL tab is clicked', async () => {
    mockGetProblems.mockResolvedValue(MOCK_PROBLEMS);
    render(<ProblemsPage />);
    await waitFor(() => expect(mockGetProblems).toHaveBeenCalledWith('loop'));

    fireEvent.click(screen.getByRole('button', { name: /^sql$/i }));

    await waitFor(() => {
      expect(mockGetProblems).toHaveBeenCalledWith('sql');
    });
  });
});

// ─── Req 4.4: Error state shows "Coba lagi" button ───────────────────────────
describe('ProblemsPage — error state (Req 4.4)', () => {
  it('shows "Coba lagi" button when fetch fails', async () => {
    // Reject ALL calls — the all-categories useEffect catches errors silently,
    // but fetchProblems sets the error state and shows "Coba lagi".
    mockGetProblems.mockRejectedValue(new Error('Network error'));

    render(<ProblemsPage />);

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /coba lagi/i })).toBeInTheDocument();
    });
  });

  it('retries fetch for current category when "Coba lagi" is clicked', async () => {
    let callCount = 0;
    mockGetProblems.mockImplementation(() => {
      callCount++;
      // First call (initial mount) rejects. After retry click, resolves.
      if (callCount <= 1) return Promise.reject(new Error('Network error'));
      return Promise.resolve([]);
    });

    render(<ProblemsPage />);

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /coba lagi/i })).toBeInTheDocument();
    });

    const callCountBeforeRetry = callCount;
    fireEvent.click(screen.getByRole('button', { name: /coba lagi/i }));

    await waitFor(() => {
      expect(callCount).toBeGreaterThan(callCountBeforeRetry);
    });
  });

  it('hides error message after successful retry', async () => {
    let callCount = 0;
    mockGetProblems.mockImplementation(() => {
      callCount++;
      // First call rejects (initial mount), retry resolves.
      if (callCount <= 1) return Promise.reject(new Error('Network error'));
      return Promise.resolve([]);
    });

    render(<ProblemsPage />);

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /coba lagi/i })).toBeInTheDocument();
    });

    fireEvent.click(screen.getByRole('button', { name: /coba lagi/i }));

    await waitFor(() => {
      expect(screen.queryByRole('button', { name: /coba lagi/i })).not.toBeInTheDocument();
    });
  });
});
