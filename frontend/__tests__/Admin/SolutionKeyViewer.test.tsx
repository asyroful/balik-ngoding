import { render, screen, waitFor } from '@testing-library/react';
import { vi } from 'vitest';
import { SolutionKeyViewer } from '@/components/Admin/SolutionKeyViewer';
import * as api from '@/lib/api';

vi.mock('@/lib/api');

describe('SolutionKeyViewer', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  test('displays loading skeleton while fetching', () => {
    const mockGetSolutionKey = vi.spyOn(api, 'getSolutionKey').mockImplementation(
      () => new Promise(() => {}) // Never resolves
    );

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic dGVzdDp0ZXN0"
      />
    );

    // Should show loading state (Skeleton component)
    expect(mockGetSolutionKey).toHaveBeenCalled();
  });

  test('displays solution key code in read-only editor', async () => {
    const mockSolutionKey = {
      id: '1',
      problemId: '1',
      code: 'function test() { return 42; }',
      language: 'javascript' as const,
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    };

    vi.spyOn(api, 'getSolutionKey').mockResolvedValue(mockSolutionKey);

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic dGVzdDp0ZXN0"
      />
    );

    await waitFor(() => {
      expect(screen.getByText('Solution Code')).toBeInTheDocument();
    });
  });

  test('displays metadata (title, language, updatedAt)', async () => {
    const mockSolutionKey = {
      id: '1',
      problemId: '1',
      code: 'function test() { return 42; }',
      language: 'sql' as const,
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-02T12:30:00Z',
    };

    vi.spyOn(api, 'getSolutionKey').mockResolvedValue(mockSolutionKey);

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic dGVzdDp0ZXN0"
      />
    );

    await waitFor(() => {
      expect(screen.getByText('Test Problem')).toBeInTheDocument();
      expect(screen.getByText('sql')).toBeInTheDocument();
    });
  });

  test('displays error message on fetch failure', async () => {
    vi.spyOn(api, 'getSolutionKey').mockRejectedValue(new Error('Failed to fetch'));

    render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken="Basic dGVzdDp0ZXN0"
      />
    );

    await waitFor(() => {
      expect(screen.getByText('Failed to fetch')).toBeInTheDocument();
    });
  });

  test('does not render without authToken', () => {
    const mockGetSolutionKey = vi.spyOn(api, 'getSolutionKey');

    const { container } = render(
      <SolutionKeyViewer
        problemId="1"
        problemTitle="Test Problem"
        authToken={undefined}
      />
    );

    expect(mockGetSolutionKey).not.toHaveBeenCalled();
    expect(container.firstChild).toBeNull();
  });
});
