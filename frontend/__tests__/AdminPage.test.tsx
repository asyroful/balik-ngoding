// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import AdminPage from '@/app/problems/[id]/admin/page';

// Mock next/navigation
vi.mock('next/navigation', () => ({
  useParams: vi.fn(() => ({ id: '1' })),
  useRouter: vi.fn(() => ({
    push: vi.fn(),
  })),
}));

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
  getProblemById: vi.fn(),
  getSolutionKey: vi.fn(),
}));

import { getProblemById, getSolutionKey } from '@/lib/api';

beforeEach(() => {
  vi.clearAllMocks();
  localStorage.clear();
});

afterEach(() => {
  localStorage.clear();
});

describe('AdminPage', () => {
  // ─── Task 20.1: Displays login form when not authenticated ───
  it('TestAdminPageNotAuthenticated: displays login form when not authenticated', async () => {
    const mockGetProblemById = vi.mocked(getProblemById);
    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Test Problem',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText(/admin login/i)).toBeInTheDocument();
      expect(screen.getByLabelText(/username/i)).toBeInTheDocument();
      expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    });
  });

  // ─── Task 20.2: Displays solution key viewer when authenticated ───
  it('TestAdminPageAuthenticated: displays solution key viewer when authenticated', async () => {
    const mockGetProblemById = vi.mocked(getProblemById);
    const mockGetSolutionKey = vi.mocked(getSolutionKey);

    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Test Problem',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'function test() { return true; }',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    // Set auth token in localStorage
    localStorage.setItem('admin_auth_token', 'Basic test');

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText(/solution key/i)).toBeInTheDocument();
      expect(screen.getByText(/test problem/i)).toBeInTheDocument();
    });
  });

  // ─── Task 20.3: Persists auth token to localStorage ───
  it('TestAdminPageTokenPersistence: persists auth token to localStorage', async () => {
    const user = userEvent.setup();
    const mockGetProblemById = vi.mocked(getProblemById);
    const mockGetSolutionKey = vi.mocked(getSolutionKey);

    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Test Problem',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'function test() { return true; }',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByLabelText(/username/i)).toBeInTheDocument();
    });

    const usernameInput = screen.getByLabelText(/username/i);
    const passwordInput = screen.getByLabelText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'testuser');
    await user.type(passwordInput, 'testpass');
    await user.click(submitButton);

    await waitFor(() => {
      const storedToken = localStorage.getItem('admin_auth_token');
      expect(storedToken).toBeTruthy();
      expect(storedToken).toContain('Basic');
    });
  });

  // ─── Task 20.4: Loads auth token from localStorage on mount ───
  it('TestAdminPageTokenRecovery: loads auth token from localStorage on mount', async () => {
    const mockGetProblemById = vi.mocked(getProblemById);
    const mockGetSolutionKey = vi.mocked(getSolutionKey);

    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Test Problem',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'function test() { return true; }',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    // Pre-set token in localStorage
    localStorage.setItem('admin_auth_token', 'Basic existing');

    render(<AdminPage />);

    await waitFor(() => {
      // Should show solution key viewer, not login form
      expect(screen.getByText(/solution key/i)).toBeInTheDocument();
      expect(screen.queryByText(/admin login/i)).not.toBeInTheDocument();
    });
  });

  // ─── Task 20.5: Displays error when problem doesn't exist ───
  it('TestAdminPageProblemNotFound: displays error when problem does not exist', async () => {
    const mockGetProblemById = vi.mocked(getProblemById);
    mockGetProblemById.mockRejectedValue(new Error('404: Problem not found'));

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText(/soal tidak ditemukan/i)).toBeInTheDocument();
    });
  });

  // Additional: Verify logout functionality
  it('logout button clears localStorage and shows login form', async () => {
    const user = userEvent.setup();
    const mockGetProblemById = vi.mocked(getProblemById);
    const mockGetSolutionKey = vi.mocked(getSolutionKey);

    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Test Problem',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    mockGetSolutionKey.mockResolvedValue({
      id: 'sk-1',
      problemId: '1',
      code: 'function test() { return true; }',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    localStorage.setItem('admin_auth_token', 'Basic test');

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText(/solution key/i)).toBeInTheDocument();
    });

    const logoutButton = screen.getByRole('button', { name: /logout/i });
    await user.click(logoutButton);

    await waitFor(() => {
      expect(localStorage.getItem('admin_auth_token')).toBeNull();
      expect(screen.getByText(/admin login/i)).toBeInTheDocument();
    });
  });

  // Additional: Verify problem title is displayed
  it('displays problem title in header', async () => {
    const mockGetProblemById = vi.mocked(getProblemById);
    mockGetProblemById.mockResolvedValue({
      id: '1',
      title: 'Add Two Numbers',
      description: 'Test description',
      category: 'loop',
      difficulty: 'easy',
      starterCode: 'function test() {}',
      isActive: true,
      createdAt: '2024-01-01T00:00:00Z',
      thinkingGuide: null,
      hints: null,
      prerequisiteId: null,
      prerequisiteTitle: null,
    });

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText('Add Two Numbers')).toBeInTheDocument();
    });
  });
});
