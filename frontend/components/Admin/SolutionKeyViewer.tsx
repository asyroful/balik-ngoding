'use client';

import { useEffect, useState } from 'react';
import dynamic from 'next/dynamic';
import { getSolutionKey } from '@/lib/api';
import { SolutionKey } from '@/lib/types';
import { Skeleton } from '@/components/Skeleton';

const CodeEditor = dynamic(() => import('@monaco-editor/react').then((mod) => mod.default), {
  ssr: false,
  loading: () => <Skeleton />,
});

interface SolutionKeyViewerProps {
  problemId: string;
  problemTitle: string;
  authToken?: string;
}

export function SolutionKeyViewer({ problemId, problemTitle, authToken }: SolutionKeyViewerProps) {
  const [solutionKey, setSolutionKey] = useState<SolutionKey | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!authToken) return;

    const fetchSolutionKey = async () => {
      setIsLoading(true);
      setError(null);

      try {
        const sk = await getSolutionKey(problemId, authToken);
        setSolutionKey(sk);
      } catch (err) {
        const message = err instanceof Error ? err.message : 'Failed to fetch solution key';
        setError(message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchSolutionKey();
  }, [authToken, problemId]);

  if (!authToken) {
    return null;
  }

  if (isLoading) {
    return <Skeleton />;
  }

  if (error) {
    return (
      <div className="p-4 bg-red-100 border border-red-400 text-red-700 rounded-lg" role="alert">
        {error}
      </div>
    );
  }

  if (!solutionKey) {
    return (
      <div className="p-4 bg-yellow-100 border border-yellow-400 text-yellow-700 rounded-lg">
        No solution key found for this problem.
      </div>
    );
  }

  const updatedAt = new Date(solutionKey.updatedAt).toLocaleString();

  return (
    <div className="space-y-4">
      <div className="bg-white p-4 rounded-lg shadow-sm border border-gray-200">
        <div className="grid grid-cols-2 gap-4">
          <div>
            <p className="text-sm text-gray-600">Problem</p>
            <p className="font-medium text-gray-900">{problemTitle}</p>
          </div>
          <div>
            <p className="text-sm text-gray-600">Language</p>
            <p className="font-medium text-gray-900 capitalize">{solutionKey.language}</p>
          </div>
          <div className="col-span-2">
            <p className="text-sm text-gray-600">Last Updated</p>
            <p className="font-medium text-gray-900">{updatedAt}</p>
          </div>
        </div>
      </div>

      <div className="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div className="p-4 border-b border-gray-200 bg-gray-50">
          <h3 className="font-medium text-gray-900">Solution Code</h3>
        </div>
        <div className="h-96">
          <CodeEditor
            value={solutionKey.code}
            language={solutionKey.language === 'sql' ? 'sql' : 'javascript'}
            theme="vs-light"
            options={{
              readOnly: true,
              minimap: { enabled: false },
              scrollBeyondLastLine: false,
            }}
          />
        </div>
      </div>
    </div>
  );
}
