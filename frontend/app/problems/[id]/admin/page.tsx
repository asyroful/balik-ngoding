'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import { getProblemById } from '@/lib/api';
import { Problem } from '@/lib/types';
import { AdminAuthForm } from '@/components/Admin/AdminAuthForm';
import { SolutionKeyViewer } from '@/components/Admin/SolutionKeyViewer';
import { Skeleton } from '@/components/Skeleton';

export default function AdminPage() {
  const params = useParams();
  const problemId = params.id as string;

  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [authToken, setAuthToken] = useState<string | null>(null);
  const [problem, setProblem] = useState<Problem | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [authError, setAuthError] = useState<string | null>(null);

  // Check localStorage for existing auth token on mount
  useEffect(() => {
    const storedToken = localStorage.getItem('admin_auth_token');
    if (storedToken) {
      setAuthToken(storedToken);
      setIsAuthenticated(true);
    }

    // Fetch problem details
    const fetchProblem = async () => {
      try {
        const p = await getProblemById(problemId);
        setProblem(p);
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to fetch problem';
        setError(message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchProblem();
  }, [problemId]);

  const handleAuthenticate = (token: string) => {
    try {
      // Validate token format (should be "Basic ...")
      if (!token.startsWith('Basic ')) {
        setAuthError('Invalid authentication token');
        return;
      }

      // Store token in localStorage
      localStorage.setItem('admin_auth_token', token);
      setAuthToken(token);
      setIsAuthenticated(true);
      setAuthError(null);
    } catch (err) {
      setAuthError('Failed to authenticate. Please try again.');
    }
  };

  if (isLoading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <Skeleton />
      </div>
    );
  }

  if (error) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="p-4 bg-red-100 border border-red-400 text-red-700 rounded-lg" role="alert">
          {error}
        </div>
      </div>
    );
  }

  if (!problem) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="p-4 bg-yellow-100 border border-yellow-400 text-yellow-700 rounded-lg">
          Problem not found.
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return <AdminAuthForm onAuthenticate={handleAuthenticate} error={authError} />;
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-gray-900 mb-2">{problem.title}</h1>
        <p className="text-gray-600">Admin Solution Key Viewer</p>
      </div>

      <SolutionKeyViewer problemId={problemId} problemTitle={problem.title} authToken={authToken || undefined} />
    </div>
  );
}
