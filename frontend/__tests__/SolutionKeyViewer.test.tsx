// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import { SolutionKeyViewer } from '@/components/Admin/SolutionKeyViewer';

// Mock next/dynamic for Monaco Editor
vi.mock('next/dynamic', () => ({
  default: (_importFn: unknown, _opts?: unknown) => {
    const Stub = () => <div data-testid="monaco-editor">Monaco Editor</div>;
    Stub.displayName = 'DynamicMonacoStub';
    return Stub;
  },
}));

// Mock the API
vi.mock('@/lib/api', () => ({
  getSolutionKey: vi.fn(),
}));

import { getSolutionKey } from '@/lib/api';

beforeEach(() => {
  vi.clearAllMocks();
});

describe('SolutionKeyViewer', () => {
  // ─── Task 19.1: Loading skeleton while fetching ───
  it('TestSolutionKeyViewerLoading: displays loading skeleton while fetching', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockImplementation(
      () => new Promise(() => {}) // Never resolves
    );

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic test"
      />
    );

    // Should show loading skeleton
    expect(screen.queryByTestId('monaco-editor')).not.toBeInTheDocument();
  });

  // ─── Task 19.2: Displays code in read-only editor ───
  it('TestSolutionKeyViewerDisplay: displays solution key code in read-only editor', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'function add(a, b) { return a + b; }',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic test"
      />
    );

    await waitFor(() => {
      expect(screen.getByTestId('monaco-editor')).toBeInTheDocument();
    });
  });

  // ─── Task 19.3: Displays metadata ───
  it('TestSolutionKeyViewerMetadata: displays metadata (title, language, updatedAt)', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'SELECT * FROM users;',
      language: 'sql',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-15T10:30:00Z',
    });

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="User Query Problem"
        authToken="Basic test"
      />
    );

    await waitFor(() => {
      expect(screen.getByText('User Query Problem')).toBeInTheDocument();
      expect(screen.getByText(/sql/i)).toBeInTheDocument();
    });
  });

  // ─── Task 19.4: Displays error message on fetch failure ───
  it('TestSolutionKeyViewerError: displays error message on fetch failure', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockRejectedValue(
      new Error('Failed to fetch solution key')
    );

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic test"
      />
    );

    await waitFor(() => {
      expect(screen.getByText(/gagal memuat solution key/i)).toBeInTheDocument();
    });
  });

  // ─── Task 19.5: Does not render without authToken ───
  it('TestSolutionKeyViewerNoToken: does not render without authToken', () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);

    const { container } = render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken={undefined}
      />
    );

    // Should render nothing
    expect(container.firstChild).toBeNull();
    expect(mockGetSolutionKey).not.toHaveBeenCalled();
  });

  // Additional: Verify 404 error handling
  it('displays specific message when solution key not found (404)', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockRejectedValue(
      new Error('404: Solution key not found')
    );

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic test"
      />
    );

    await waitFor(() => {
      expect(
        screen.getByText(/solution key tidak ditemukan/i)
      ).toBeInTheDocument();
    });
  });

  // Additional: Verify API is called with correct parameters
  it('calls getSolutionKey with correct problemId and authToken', async () => {
    const mockGetSolutionKey = vi.mocked(getSolutionKey);
    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '123',
      code: 'test code',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    render(
      <SolutionKeyViewer
        problemId="123"
        problemTitle="Test"
        authToken="Basic xyz"
      />
    );

    await waitFor(() => {
      expect(mockGetSolutionKey).toHaveBeenCalledWith('123', 'Basic xyz');
    });
  });
});
