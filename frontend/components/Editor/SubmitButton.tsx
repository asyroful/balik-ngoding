'use client';

import { useSubmissionStore } from '@/store/submissionStore';
import { submitSolution } from '@/lib/api';

interface SubmitButtonProps {
  problemId: string;
  starterCode: string;
  language: string;
}

export default function SubmitButton({ problemId, language }: SubmitButtonProps) {
  const { code, isLoading, setLoading, setResult, setError } = useSubmissionStore();

  const isEmpty = code.trim() === '';
  const isDisabled = isEmpty || isLoading;

  async function handleSubmit() {
    if (isLoading || isEmpty) return;
    setLoading(true);
    setError(null);
    setResult(null);
    try {
      const result = await submitSolution({ problemId, code, language });
      setResult(result);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Terjadi kesalahan jaringan.');
      setResult(null);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex items-center gap-3">
      <button
        onClick={handleSubmit}
        disabled={isDisabled}
        className={`flex items-center gap-2 rounded-lg px-5 py-2 text-sm font-semibold transition
          ${isDisabled
            ? 'cursor-not-allowed bg-gray-100 text-gray-400'
            : 'bg-indigo-700 text-white hover:bg-indigo-800 active:scale-[0.98]'
          }`}
      >
        {isLoading ? (
          <>
            <svg className="animate-spin h-3.5 w-3.5" fill="none" viewBox="0 0 24 24">
              <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
              <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
            </svg>
            Mengirim...
          </>
        ) : 'Submit Jawaban'}
      </button>
      {isEmpty && !isLoading && (
        <p className="text-xs text-red-500">Kode tidak boleh kosong</p>
      )}
    </div>
  );
}
