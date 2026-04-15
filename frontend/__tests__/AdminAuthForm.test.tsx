// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { AdminAuthForm } from '@/components/Admin/AdminAuthForm';

beforeEach(() => {
  vi.clearAllMocks();
});

describe('AdminAuthForm', () => {
  // ─── Task 18.1: Form renders with username and password fields ───
  it('TestAdminAuthFormRenders: renders form with username and password fields', () => {
    const onAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={onAuthenticate} />);

    expect(screen.getByLabelText(/username/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /login/i })).toBeInTheDocument();
  });

  // ─── Task 18.2: Valid submission calls onAuthenticate ───
  it('TestAdminAuthFormValidSubmission: form submission with valid credentials calls onAuthenticate', async () => {
    const user = userEvent.setup();
    const onAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={onAuthenticate} />);

    const usernameInput = screen.getByLabelText(/username/i);
    const passwordInput = screen.getByLabelText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'testuser');
    await user.type(passwordInput, 'testpass');
    await user.click(submitButton);

    expect(onAuthenticate).toHaveBeenCalledWith(expect.stringContaining('Basic'));
  });

  // ─── Task 18.3: Invalid credentials show error ───
  it('TestAdminAuthFormInvalidCredentials: empty credentials show error message', async () => {
    const user = userEvent.setup();
    const onAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={onAuthenticate} />);

    const submitButton = screen.getByRole('button', { name: /login/i });
    await user.click(submitButton);

    expect(screen.getByText(/username dan password harus diisi/i)).toBeInTheDocument();
    expect(onAuthenticate).not.toHaveBeenCalled();
  });

  // ─── Task 18.4: Loading state disables form ───
  it('TestAdminAuthFormLoadingState: form is disabled while loading', () => {
    const onAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={onAuthenticate} isLoading={true} />);

    const usernameInput = screen.getByLabelText(/username/i) as HTMLInputElement;
    const passwordInput = screen.getByLabelText(/password/i) as HTMLInputElement;
    const submitButton = screen.getByRole('button', { name: /memproses/i }) as HTMLButtonElement;

    expect(usernameInput.disabled).toBe(true);
    expect(passwordInput.disabled).toBe(true);
    expect(submitButton.disabled).toBe(true);
  });

  // ─── Task 18.5: Error message is displayed ───
  it('TestAdminAuthFormErrorDisplay: error message is displayed on failure', () => {
    const onAuthenticate = vi.fn();
    const errorMessage = 'Invalid credentials';
    render(
      <AdminAuthForm
        onAuthenticate={onAuthenticate}
        error={errorMessage}
      />
    );

    expect(screen.getByText(errorMessage)).toBeInTheDocument();
  });

  // Additional: Verify button text changes during loading
  it('button text changes to "Memproses..." when loading', () => {
    const onAuthenticate = vi.fn();
    const { rerender } = render(
      <AdminAuthForm onAuthenticate={onAuthenticate} isLoading={false} />
    );

    expect(screen.getByRole('button', { name: /login/i })).toBeInTheDocument();

    rerender(<AdminAuthForm onAuthenticate={onAuthenticate} isLoading={true} />);

    expect(screen.getByRole('button', { name: /memproses/i })).toBeInTheDocument();
  });

  // Additional: Verify form submission with only username shows error
  it('form submission with only username shows error', async () => {
    const user = userEvent.setup();
    const onAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={onAuthenticate} />);

    const usernameInput = screen.getByLabelText(/username/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'testuser');
    await user.click(submitButton);

    expect(screen.getByText(/username dan password harus diisi/i)).toBeInTheDocument();
    expect(onAuthenticate).not.toHaveBeenCalled();
  });
});
