// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render, screen, fireEvent, waitFor, cleanup } from '@testing-library/react';

// Mock next/navigation - use stable references to avoid useEffect re-runs
vi.mock('next/navigation', () => {
  const push = vi.fn();
  const router = { push };
  return {
    useParams: () => ({ id: 'test-id' }),
    useRouter: () => router,
  };
});

// Mock @/lib/api
vi.mock('../lib/api', () => ({
  getProblemById: vi.fn(),
  getProblems: vi.fn(),
  submitSolution: vi.fn(),
}));

// Mock next/dynamic to avoid Monaco loading
vi.mock('next/dynamic', () => ({
  default: (_fn: unknown, _opts: unknown) => {
    const MockComponent = () => <textarea data-testid="mock-editor" />;
    MockComponent.displayName = 'MockDynamic';
    return MockComponent;
  },
}));

// Mock @monaco-editor/react
vi.mock('@monaco-editor/react', () => ({
  default: () => <textarea data-testid="monaco-editor" />,
}));

// Mock @/hooks/useProgress
vi.mock('../hooks/useProgress', () => ({
  useProgress: () => ({
    progress: {},
    markAccepted: vi.fn(),
    isAccepted: vi.fn(() => false),
  }),
}));

// Mock the Zustand submission store to prevent async state updates outside act
const mockSetCode = vi.fn();
const mockSetLanguage = vi.fn();
const mockReset = vi.fn();
const mockSetLoading = vi.fn();
const mockSetResult = vi.fn();
const mockSetError = vi.fn();

vi.mock('../store/submissionStore', () => ({
  useSubmissionStore: () => ({
    code: 'function solution() {}',
    isLoading: false,
    result: null,
    error: null,
    language: 'javascript',
    setCode: mockSetCode,
    setLanguage: mockSetLanguage,
    reset: mockReset,
    setLoading: mockSetLoading,
    setResult: mockSetResult,
    setError: mockSetError,
  }),
}));

import ProblemDetailPage from '../app/problems/[id]/page';
import { getProblemById, getProblems, submitSolution } from '../lib/api';

const mockGetProblemById = vi.mocked(getProblemById);
const mockGetProblems = vi.mocked(getProblems);
const mockSubmitSolution = vi.mocked(submitSolution);

const MOCK_PROBLEM = {
  id: 'test-id',
  title: 'Test Problem',
  description: 'Ini adalah deskripsi soal untuk pengujian.',
  category: 'loop' as const,
  difficulty: 'easy' as const,
  starterCode: 'function solution() {}',
  isActive: true,
  testCases: [
    { id: 'tc-1', input: '5', expectedOutput: '10', isHidden: false },
  ],
  createdAt: '2024-01-01T00:00:00Z',
};

const MOCK_SUBMIT_RESULT = {
  status: 'accepted' as const,
  score: 100,
  total: 1,
  results: [{ passed: true, input: '5', expected: '10', actual: '10' }],
};

beforeEach(() => {
  vi.clearAllMocks();
  mockGetProblemById.mockResolvedValue(MOCK_PROBLEM);
  mockGetProblems.mockResolvedValue([]);
  mockSubmitSolution.mockResolvedValue(MOCK_SUBMIT_RESULT);
});

afterEach(() => {
  cleanup();
});

async function renderAndWait() {
  render(<ProblemDetailPage />);
  await waitFor(() => {
    expect(screen.getByText('Ini adalah deskripsi soal untuk pengujian.')).toBeInTheDocument();
  }, { timeout: 3000 });
}

// ─── Req 5.1, 5.2: Tab "Soal" shows problem description (default tab) ────────
describe('ProblemDetailMobile — Tab Soal (Req 5.1, 5.2)', () => {
  it('default tab is "soal" and description panel is visible', async () => {
    await renderAndWait();
    const descriptionText = screen.getByText('Ini adalah deskripsi soal untuk pengujian.');
    const leftPanel = descriptionText.closest('[class*="flex-col"]');
    expect(leftPanel).toBeTruthy();
    expect(leftPanel?.className).toContain('flex');
    // Check that the panel is not hidden (class starts with 'hidden' or has ' hidden ')
    expect(leftPanel?.className).not.toMatch(/(^|\s)hidden(\s|$)/);
  });

  it('tab "Soal" button is rendered in mobile tab navigation', async () => {
    await renderAndWait();
    const soalTab = screen.getByRole('button', { name: 'Soal' });
    expect(soalTab).toBeInTheDocument();
  });

  it('tab "Soal" is visually active by default (has active class)', async () => {
    await renderAndWait();
    const soalTab = screen.getByRole('button', { name: 'Soal' });
    expect(soalTab.className).toContain('border-indigo-600');
  });
});

// ─── Req 5.3, 5.4: Tab "Editor" shows editor and submit button ───────────────
describe('ProblemDetailMobile — Tab Editor (Req 5.3, 5.4)', () => {
  it('clicking Editor tab makes editor panel visible', async () => {
    await renderAndWait();
    fireEvent.click(screen.getByRole('button', { name: 'Editor' }));
    await waitFor(() => {
      expect(screen.getByRole('button', { name: /submit jawaban/i })).toBeInTheDocument();
    });
  });

  it('clicking Editor tab shows submit button', async () => {
    await renderAndWait();
    fireEvent.click(screen.getByRole('button', { name: 'Editor' }));
    await waitFor(() => {
      expect(screen.getByRole('button', { name: /submit jawaban/i })).toBeInTheDocument();
    });
  });

  it('clicking Editor tab makes right panel have flex class (not hidden)', async () => {
    await renderAndWait();
    fireEvent.click(screen.getByRole('button', { name: 'Editor' }));
    await waitFor(() => {
      const submitButton = screen.getByRole('button', { name: /submit jawaban/i });
      const rightPanel = submitButton.closest('[class*="flex-col"]');
      expect(rightPanel).toBeTruthy();
      expect(rightPanel?.className).toContain('flex');
      // Check that the panel is not hidden (class starts with 'hidden' or has ' hidden ')
      expect(rightPanel?.className).not.toMatch(/(^|\s)hidden(\s|$)/);
    });
  });
});

// ─── Req 5.6: Desktop layout shows both panels simultaneously ─────────────────
describe('ProblemDetailMobile — Desktop layout (Req 5.6)', () => {
  it('left panel has md:flex class for desktop visibility', async () => {
    await renderAndWait();
    const descriptionText = screen.getByText('Ini adalah deskripsi soal untuk pengujian.');
    const leftPanel = descriptionText.closest('[class*="md:flex"]');
    expect(leftPanel).toBeTruthy();
    expect(leftPanel?.className).toContain('md:flex');
  });

  it('right panel has md:flex class for desktop visibility', async () => {
    await renderAndWait();
    const submitButton = screen.getByRole('button', { name: /submit jawaban/i });
    const rightPanel = submitButton.closest('[class*="md:flex"]');
    expect(rightPanel).toBeTruthy();
    expect(rightPanel?.className).toContain('md:flex');
  });

  it('tab navigation has md:hidden class (only visible on mobile)', async () => {
    await renderAndWait();
    const soalTab = screen.getByRole('button', { name: 'Soal' });
    const tabNav = soalTab.closest('[class*="md:hidden"]');
    expect(tabNav).toBeTruthy();
    expect(tabNav?.className).toContain('md:hidden');
  });
});
