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

describe('AdminPageE2E', () => {
  const mockProblem = {
    id: 'test-problem-1',
    title: 'Test Problem',
    description: 'Test Description',
    category: 'loop',
    difficulty: 'easy',
    examples: [],
    testCases: [],
  };

  const mockSolutionKey = {
    id: 'sk-1',
    problemId: 'test-problem-1',
    code: 'function solution(n) { return n * 2; }',
    language: 'javascript',
    createdAt: '2024-01-01T00:00:00Z',
    updatedAt: '2024-01-01T00:00:00Z',
  };

  beforeEach(() => {
    localStorage.clear();
    vi.clearAllMocks();
  });

  // Feature: evaluator-fix-and-solution-keys, Property 15: session persistence
  test('TestAdminPageE2E - login, view solution key, logout, verify re-authentication required', async () => {
    const user = userEvent.setup();

    // Mock API calls
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockResolvedValue(mockSolutionKey);

    // Step 1: Render page without authentication
    const { rerender } = render(<AdminPage />);

    // Should show login form
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Step 2: Submit valid credentials
    const usernameInput = screen.getByPlaceholderText(/username/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const loginButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'siful');
    await user.type(passwordInput, 'atmin162');
    await user.click(loginButton);

    // Step 3: Verify solution key is displayed
    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Verify auth token is stored in localStorage
    const storedToken = localStorage.getItem('admin_auth_token');
    expect(storedToken).toBeTruthy();

    // Step 4: Logout
    const logoutButton = screen.getByRole('button', { name: /logout/i });
    await user.click(logoutButton);

    // Step 5: Verify login form is shown again
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Verify auth token is removed from localStorage
    expect(localStorage.getItem('admin_auth_token')).toBeNull();

    // Step 6: Verify re-authentication is required
    // Clear the mock to ensure fresh API call
    (api.getProblemById as any).mockClear();
    (api.getProblemById as any).mockResolvedValue(mockProblem);

    // Rerender to simulate page refresh
    rerender(<AdminPage />);

    // Should show login form again
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });
  });

  test('TestAdminPageE2E - persist auth token across page reloads', async () => {
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockResolvedValue(mockSolutionKey);

    // Simulate stored auth token
    const token = 'Basic c2lmdWw6YXRtaW4xNjI='; // base64 encoded siful:atmin162
    localStorage.setItem('admin_auth_token', token);

    // Render page
    render(<AdminPage />);

    // Should directly show solution key viewer without login form
    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Verify logout button is present
    expect(screen.getByRole('button', { name: /logout/i })).toBeInTheDocument();
  });

  test('TestAdminPageE2E - handle problem not found error', async () => {
    (api.getProblemById as any).mockRejectedValue(new Error('404 Not Found'));

    render(<AdminPage />);

    // Should show error message
    await waitFor(() => {
      expect(screen.getByText(/soal tidak ditemukan/i)).toBeInTheDocument();
    });
  });

  test('TestAdminPageE2E - handle invalid credentials', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockRejectedValue(new Error('401 Unauthorized'));

    render(<AdminPage />);

    // Should show login form
    await waitFor(() => {
      expect(screen.getByText(/username/i)).toBeInTheDocument();
    });

    // Submit invalid credentials
    const usernameInput = screen.getByPlaceholderText(/username/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const loginButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'wronguser');
    await user.type(passwordInput, 'wrongpass');
    await user.click(loginButton);

    // Should show error message
    await waitFor(() => {
      expect(screen.getByText(/invalid credentials|unauthorized/i)).toBeInTheDocument();
    });

    // Login form should still be visible
    expect(screen.getByPlaceholderText(/username/i)).toBeInTheDocument();
  });

  test('TestAdminPageE2E - display solution key metadata', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.getSolutionKey as any).mockResolvedValue(mockSolutionKey);

    render(<AdminPage />);

    // Login
    const usernameInput = screen.getByPlaceholderText(/username/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const loginButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'siful');
    await user.type(passwordInput, 'atmin162');
    await user.click(loginButton);

    // Verify metadata is displayed
    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
      expect(screen.getByText(/javascript/i)).toBeInTheDocument();
    });
  });
});
