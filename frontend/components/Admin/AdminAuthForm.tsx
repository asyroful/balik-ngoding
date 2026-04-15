'use client';

import { useState } from 'react';
import { encodeAdminCredentials } from '@/lib/api';

interface AdminAuthFormProps {
  onAuthenticate: (token: string) => void;
  isLoading?: boolean;
  error?: string | null;
}

export function AdminAuthForm({ onAuthenticate, isLoading = false, error }: AdminAuthFormProps) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [localError, setLocalError] = useState<string | null>(error || null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLocalError(null);
    setIsSubmitting(true);

    try {
      // Encode credentials as Basic Auth
      const token = encodeAdminCredentials(username, password);

      // Call onAuthenticate callback
      onAuthenticate(token);
    } catch (err) {
      setLocalError('Failed to authenticate. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  };

  const isDisabled = isSubmitting || isLoading || !username || !password;

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6 text-center text-gray-800">Admin Login</h1>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">
              Username
            </label>
            <input
              id="username"
              type="text"
              placeholder="Enter username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              disabled={isSubmitting || isLoading}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
              aria-label="Username"
            />
          </div>

          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
              Password
            </label>
            <input
              id="password"
              type="password"
              placeholder="Enter password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              disabled={isSubmitting || isLoading}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
              aria-label="Password"
            />
          </div>

          {(localError || error) && (
            <div className="p-3 bg-red-100 border border-red-400 text-red-700 rounded-lg text-sm" role="alert">
              {localError || error}
            </div>
          )}

          <button
            type="submit"
            disabled={isDisabled}
            className="w-full px-4 py-2 bg-blue-600 text-white font-medium rounded-lg hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors"
            aria-busy={isSubmitting || isLoading}
          >
            {isSubmitting || isLoading ? 'Logging in...' : 'Login'}
          </button>
        </form>
      </div>
    </div>
  );
}
