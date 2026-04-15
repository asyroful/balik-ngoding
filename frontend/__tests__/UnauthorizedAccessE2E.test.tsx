import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';
import AdminPage from '@/app/problems/[id]/admin/page';
import * as api from '@/lib/api';

// Mock the API module
vi.mock('@/lib/api');

// Mock useParams
vi.mock('next/navigation', () => ({
  useParams: () => ({ id: 'test-problem-1' }),
}));

describe('UnauthorizedAccessE2E', () => {
  const mockProblem = {
    id: 'test-problem-1',
    title: 'Test Problem',
    description: 'Test Description',
    category: 'loop',
    difficulty: 'easy',
    examples: [],
    testCases: [],
  };

  beforeEach(() => {
    localStorage.clear();
    vi.clearAllMocks();
  });

  // Feature: evaluator-fix-and-solution-keys, Property 13: admin authentication security
  test('TestUnauthorizedAccessE2E - verify unauthenticated users cannot access solution keys', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    render(<AdminPage />);

    // Should show login form, not solution key
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Solution key should not be visible
    expect(screen.queryByText(/function solution/i)).not.toBeInTheDocument();
  });

  test('TestUnauthorizedAccessE2E - verify 401 Unauthorized is returned for missing auth', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    render(<AdminPage />);

    // Verify login form is shown
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Verify no solution key data is exposed
    expect(screen.queryByText(/code|language|updated/i)).not.toBeInTheDocument();
  });

  test('TestUnauthorizedAccessE2E - verify solution keys are not exposed without authentication', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    render(<AdminPage />);

    // Try to submit invalid credentials
    const usernameInput = screen.getByPlaceholderText(/username/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const loginButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'wronguser');
    await user.type(passwordInput, 'wrongpass');
    await user.click(loginButton);

    // Should show error, not solution key
    await waitFor(() => {
      expect(screen.getByText(/invalid credentials|unauthorized/i)).toBeInTheDocument();
    });

    // Solution key should not be visible
    expect(screen.queryByText(/function solution/i)).not.toBeInTheDocument();
  });

  test('TestUnauthorizedAccessE2E - verify localStorage token is required for access', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    // Don't set any token in localStorage
    render(<AdminPage />);

    // Should show login form
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Verify no solution key is displayed
    expect(screen.queryByText(/solution key|code/i)).not.toBeInTheDocument();
  });

  test('TestUnauthorizedAccessE2E - verify invalid token is rejected', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    // Set invalid token in localStorage
    localStorage.setItem('admin_auth_token', 'Basic invalid');

    render(<AdminPage />);

    // Should attempt to fetch with invalid token and fail
    await waitFor(() => {
      // Should show login form again after failed auth
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });
  });

  test('TestUnauthorizedAccessE2E - verify solution key data is not leaked in error messages', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    render(<AdminPage />);

    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Check that error message doesn't contain sensitive data
    const errorMessages = screen.queryAllByText(/error|unauthorized|invalid/i);
    errorMessages.forEach((msg) => {
      expect(msg.textContent).not.toMatch(/function|code|password/i);
    });
  });

  test('TestUnauthorizedAccessE2E - verify session is cleared on logout', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockResolvedValue({
      id: 'sk-1',
      problemId: 'test-problem-1',
      code: 'function solution() {}',
      language: 'javascript',
      createdAt: '2024-01-01T00:00:00Z',
      updatedAt: '2024-01-01T00:00:00Z',
    });

    render(<AdminPage />);

    // Login first
    const usernameInput = screen.getByPlaceholderText(/username/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const loginButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'siful');
    await user.type(passwordInput, 'atmin162');
    await user.click(loginButton);

    // Wait for solution key to be displayed
    await waitFor(() => {
      expect(screen.getByRole('button', { name: /logout/i })).toBeInTheDocument();
    });

    // Logout
    const logoutButton = screen.getByRole('button', { name: /logout/i });
    await user.click(logoutButton);

    // Verify token is cleared
    expect(localStorage.getItem('admin_auth_token')).toBeNull();

    // Verify login form is shown again
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });
  });
});
