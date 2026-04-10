'use client';

import { useEffect } from 'react';
import { useSubmissionStore } from '@/store/submissionStore';
import { useProgress } from '@/hooks/useProgress';

interface ResultPanelProps {
  starterCode: string;
  problems?: { id: string }[];
  currentProblemId?: string;
}

export default function ResultPanel({ starterCode, problems, currentProblemId }: ResultPanelProps) {
  const { result, error, setCode, setResult, setError } = useSubmissionStore();
  const { markAccepted } = useProgress();

  useEffect(() => {
    if (result?.status === 'accepted' && currentProblemId) {
      markAccepted(currentProblemId);
    }
  }, [result, currentProblemId]);

  if (!result && !error) return null;

  function handleRetry() {
    setCode(starterCode);
    setResult(null);
    setError(null);
  }

  const nextProblemId = (() => {
    if (!problems || !currentProblemId) return null;
    const idx = problems.findIndex((p) => p.id === currentProblemId);
    if (idx === -1 || idx >= problems.length - 1) return null;
    return problems[idx + 1].id;
  })();

  return (
    <div className="rounded-xl border border-gray-200 overflow-hidden text-sm">
      {/* Network error */}
      {error && !result && (
        <div className="border-b border-red-100 bg-red-50 px-4 py-3">
          <p className="font-semibold text-red-600 text-xs">Gagal mengirim jawaban</p>
          <p className="text-red-500 text-xs mt-0.5">{error}</p>
        </div>
      )}

      {/* Result summary */}
      {result && (
        <>
          <div className={`px-4 py-3 border-b ${
            result.status === 'accepted'
              ? 'bg-green-50 border-green-100'
              : 'bg-yellow-50 border-yellow-100'
          }`}>
            <p className={`font-semibold text-sm ${
              result.status === 'accepted' ? 'text-green-700' : 'text-yellow-700'
            }`}>
              {result.status === 'accepted' ? '✓ Semua test case passed!'
                : result.status === 'wrong_answer' ? '✗ Jawaban salah'
                : '✗ Terjadi error'}
            </p>
            <p className="text-gray-500 text-xs mt-0.5">
              {result.score} dari {result.total} test case passed
            </p>
          </div>

          <div className="divide-y divide-gray-100">
            {result.results.map((tc, i) => (
              <div key={i} className="px-4 py-3">
                <div className="flex items-center gap-2 mb-2">
                  <span className={`text-xs font-semibold ${tc.passed ? 'text-green-600' : 'text-red-500'}`}>
                    {tc.passed ? '✓ Passed' : '✗ Failed'}
                  </span>
                  <span className="text-gray-400 text-xs">Test case {i + 1}</span>
                </div>
                <div className="grid grid-cols-3 gap-2 text-xs">
                  {(['Input', 'Expected', 'Actual'] as const).map((label, li) => {
                    const val = [tc.input, tc.expected, tc.actual][li];
                    return (
                      <div key={label}>
                        <span className="block text-gray-400 mb-1">{label}</span>
                        <code className={`block rounded-md border px-2 py-1 font-mono truncate
                          ${li === 2 && !tc.passed
                            ? 'border-red-200 bg-red-50 text-red-600'
                            : 'border-gray-200 bg-gray-50 text-gray-700'
                          }`}>
                          {val}
                        </code>
                      </div>
                    );
                  })}
                </div>
                {tc.error && <p className="mt-2 text-red-500 text-xs">{tc.error}</p>}
              </div>
            ))}
          </div>
        </>
      )}

      {/* Actions */}
      <div className="flex gap-2 px-4 py-3 bg-gray-50 border-t border-gray-100">
        <button
          onClick={handleRetry}
          className="rounded-lg border border-gray-200 bg-white px-4 py-1.5 text-xs font-medium text-gray-600 hover:bg-gray-50 transition"
        >
          Coba Lagi
        </button>
        {nextProblemId && result?.status === 'accepted' && (
          <a
            href={`/problems/${nextProblemId}`}
            className="rounded-lg bg-indigo-700 px-4 py-1.5 text-xs font-semibold text-white hover:bg-indigo-800 transition"
          >
            Soal Berikutnya →
          </a>
        )}
      </div>
    </div>
  );
}
