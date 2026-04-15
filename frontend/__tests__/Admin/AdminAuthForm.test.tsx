import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';
import { AdminAuthForm } from '@/components/Admin/AdminAuthForm';

describe('AdminAuthForm', () => {
  test('renders form with username and password fields', () => {
    const mockOnAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} />);

    expect(screen.getByLabelText('Username')).toBeInTheDocument();
    expect(screen.getByLabelText('Password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /login/i })).toBeInTheDocument();
  });

  test('calls onAuthenticate with token on valid submission', async () => {
    const mockOnAuthenticate = vi.fn();
    const user = userEvent.setup();

    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} />);

    const usernameInput = screen.getByLabelText('Username');
    const passwordInput = screen.getByLabelText('Password');
    const submitButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'siful');
    await user.type(passwordInput, 'atmin162');
    await user.click(submitButton);

    expect(mockOnAuthenticate).toHaveBeenCalled();
    expect(mockOnAuthenticate).toHaveBeenCalledWith(expect.stringContaining('Basic'));
  });

  test('displays error message on authentication failure', async () => {
    const mockOnAuthenticate = vi.fn();
    const user = userEvent.setup();

    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} error="Invalid credentials" />);

    expect(screen.getByText('Invalid credentials')).toBeInTheDocument();
  });

  test('disables form while loading', () => {
    const mockOnAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} isLoading={true} />);

    expect(screen.getByLabelText('Username')).toBeDisabled();
    expect(screen.getByLabelText('Password')).toBeDisabled();
    expect(screen.getByRole('button', { name: /logging in/i })).toBeDisabled();
  });

  test('disables submit button when fields are empty', () => {
    const mockOnAuthenticate = vi.fn();
    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} />);

    const submitButton = screen.getByRole('button', { name: /login/i });
    expect(submitButton).toBeDisabled();
  });

  test('enables submit button when both fields are filled', async () => {
    const mockOnAuthenticate = vi.fn();
    const user = userEvent.setup();

    render(<AdminAuthForm onAuthenticate={mockOnAuthenticate} />);

    const usernameInput = screen.getByLabelText('Username');
    const passwordInput = screen.getByLabelText('Password');
    const submitButton = screen.getByRole('button', { name: /login/i });

    await user.type(usernameInput, 'siful');
    await user.type(passwordInput, 'atmin162');

    expect(submitButton).not.toBeDisabled();
  });
});
