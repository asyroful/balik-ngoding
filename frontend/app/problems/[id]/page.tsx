'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { getProblemById, getProblems, submitSolution } from '@/lib/api';
import { Problem } from '@/lib/types';
import { useSubmissionStore } from '@/store/submissionStore';
import ProblemDescription from '@/components/ProblemDetail/ProblemDescription';
import ExampleBlock from '@/components/ProblemDetail/ExampleBlock';
import CodeEditor from '@/components/Editor/CodeEditor';
import LanguageSelector, { getDefaultLanguage } from '@/components/Editor/LanguageSelector';
import SubmitButton from '@/components/Editor/SubmitButton';
import ResultPanel from '@/components/Result/ResultPanel';
import { ProblemDescriptionSkeleton, EditorSkeleton, Skeleton } from '@/components/Skeleton';

export default function ProblemDetailPage() {
  const { id } = useParams<{ id: string }>();
  const router = useRouter();

  const [problem, setProblem] = useState<Problem | null>(null);
  const [problems, setProblems] = useState<Problem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [activeTab, setActiveTab] = useState<'soal' | 'editor'>('soal');

  const { code, setCode, reset, language, setLanguage, isLoading: isSubmitting, setLoading, setResult, setError: setSubmitError } = useSubmissionStore();

  async function handleSubmit() {
    if (!problem || isSubmitting || code.trim() === '') return;
    setLoading(true);
    setSubmitError(null);
    setResult(null);
    try {
      const result = await submitSolution({ problemId: problem.id, code, language });
      setResult(result);
    } catch (err: unknown) {
      setSubmitError(err instanceof Error ? err.message : 'Terjadi kesalahan jaringan.');
      setResult(null);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    if (!id) return;
    setIsLoading(true);
    setError(null);
    reset();

    Promise.all([getProblemById(id), getProblems()])
      .then(([data, allProblems]) => {
        setProblem(data);
        setProblems(allProblems);
        setCode(data.starterCode);
        setLanguage(getDefaultLanguage(data.category));
      })
      .catch((err: Error) => {
        const msg = err.message ?? '';
        if (msg.includes('404') || msg.toLowerCase().includes('not found')) {
          router.push('/problems?error=Soal+tidak+ditemukan');
        } else {
          setError(msg || 'Gagal memuat soal. Silakan coba lagi.');
        }
      })
      .finally(() => setIsLoading(false));
  }, [id, router, setCode, reset]);

  if (isLoading) {
    return (
      <div className="flex overflow-hidden" style={{ height: 'calc(100vh - 53px)' }}>
        {/* Left skeleton */}
        <div className="w-[45%] overflow-hidden border-r border-gray-200 bg-white">
          {/* Top bar skeleton */}
          <div className="flex items-center justify-between border-b border-gray-100 px-6 py-3">
            <div className="flex gap-2">
              <Skeleton className="h-5 w-16 rounded-md" />
              <Skeleton className="h-5 w-16 rounded-md" />
            </div>
            <Skeleton className="h-3 w-20" />
          </div>
          <ProblemDescriptionSkeleton />
        </div>
        {/* Right skeleton */}
        <div className="flex w-[55%] flex-col overflow-hidden bg-[#1e1e1e]">
          <div className="flex items-center justify-between border-b border-white/10 px-5 py-2.5 bg-[#252526]">
            <Skeleton className="h-3 w-20 bg-white/10" />
            <div className="flex gap-1.5">
              <span className="h-2.5 w-2.5 rounded-full bg-[#ff5f57]/40" />
              <span className="h-2.5 w-2.5 rounded-full bg-[#febc2e]/40" />
              <span className="h-2.5 w-2.5 rounded-full bg-[#28c840]/40" />
            </div>
          </div>
          <EditorSkeleton />
          <div className="border-t border-gray-200 bg-white p-4">
            <Skeleton className="h-8 w-36 rounded-lg" />
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-[80vh] px-4">
        <div className="rounded-xl border border-red-200 bg-red-50 p-6 text-red-600 text-sm max-w-md text-center">
          {error}
        </div>
      </div>
    );
  }

  if (!problem) return null;

  return (
    <div className="flex flex-col overflow-hidden" style={{ height: 'calc(100vh - 53px)' }}>

      {/* ── MOBILE TAB NAVIGATION ── */}
      <div className="flex flex-col md:hidden border-b border-gray-200 bg-white">
        <div className="flex">
          <button
            onClick={() => setActiveTab('soal')}
            className={`flex-1 py-2.5 text-sm font-medium transition ${
              activeTab === 'soal'
                ? 'border-b-2 border-indigo-600 text-indigo-600 bg-indigo-50'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            Soal
          </button>
          <button
            onClick={() => setActiveTab('editor')}
            className={`flex-1 py-2.5 text-sm font-medium transition ${
              activeTab === 'editor'
                ? 'border-b-2 border-indigo-600 text-indigo-600 bg-indigo-50'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            Editor
          </button>
        </div>
      </div>

      {/* ── MAIN CONTENT ── */}
      <div className="flex flex-1 overflow-hidden md:flex-row">

      {/* ── LEFT: problem description — scrollable ── */}
      <div className={`${activeTab === 'soal' ? 'flex' : 'hidden'} flex-col w-full md:w-[45%] overflow-y-auto border-r border-gray-200 bg-white md:flex`}>
        {/* Sticky top bar */}
        <div className="sticky top-0 z-10 flex items-center justify-between border-b border-gray-100 bg-white/95 backdrop-blur-sm px-6 py-3">
          <div className="flex items-center gap-2">
            <span className={`rounded-md border px-2 py-0.5 text-xs font-medium capitalize
              ${problem.difficulty === 'easy' ? 'border-green-200 bg-green-50 text-green-700'
                : problem.difficulty === 'medium' ? 'border-yellow-200 bg-yellow-50 text-yellow-700'
                : 'border-red-200 bg-red-50 text-red-700'}`}>
              {problem.difficulty}
            </span>
            <span className="rounded-md border border-indigo-200 bg-indigo-50 px-2 py-0.5 text-xs font-medium capitalize text-indigo-700">
              {problem.category}
            </span>
          </div>
          <button
            onClick={() => router.push('/problems')}
            className="text-xs text-gray-400 hover:text-gray-700 transition"
          >
            ← Daftar Soal
          </button>
        </div>

        {/* Content */}
        <div className="space-y-6 px-6 py-5">
          <ProblemDescription problem={problem} />
          {problem.testCases && problem.testCases.length > 0 && (
            <ExampleBlock testCases={problem.testCases} />
          )}
        </div>
      </div>

      {/* ── RIGHT: editor + submit + result ── */}
      <div className={`${activeTab === 'editor' ? 'flex' : 'hidden'} flex-col w-full md:w-[55%] overflow-hidden bg-[#1e1e1e] md:flex`}>
        {/* Editor top bar */}
        <div className="flex items-center justify-between border-b border-white/10 px-5 py-2.5 bg-[#252526]">
          <LanguageSelector selected={language} onChange={setLanguage} problemCategory={problem.category} />
          <div className="flex gap-1.5">
            <span className="h-2.5 w-2.5 rounded-full bg-[#ff5f57]" />
            <span className="h-2.5 w-2.5 rounded-full bg-[#febc2e]" />
            <span className="h-2.5 w-2.5 rounded-full bg-[#28c840]" />
          </div>
        </div>

        {/* Monaco editor */}
        <div className="flex-1 overflow-hidden">
          <CodeEditor value={code} onChange={(v) => setCode(v)} language={language} onSubmit={handleSubmit} />
        </div>

        {/* Submit + Result — light panel at bottom */}
        <div
          className="overflow-y-auto border-t border-gray-200 bg-white"
          style={{ maxHeight: '38vh' }}
        >
          <div className="p-4 space-y-3">
            <SubmitButton problemId={problem.id} starterCode={problem.starterCode} language={language} />
            <ResultPanel
              starterCode={problem.starterCode}
              problems={problems}
              currentProblemId={problem.id}
            />
          </div>
        </div>
      </div>

      </div>
    </div>
  );
}
