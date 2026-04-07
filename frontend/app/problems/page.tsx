'use client';

import { useEffect, useState, Suspense } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { getProblems } from '@/lib/api';
import { Problem } from '@/lib/types';
import CategoryFilter from '@/components/ProblemList/CategoryFilter';
import ProblemTable from '@/components/ProblemList/ProblemTable';
import { ProblemListSkeleton } from '@/components/Skeleton';
import { useProgress } from '@/hooks/useProgress';

function ProblemsPageInner() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { progress } = useProgress();
  const [problems, setProblems] = useState<Problem[]>([]);
  const [allProblems, setAllProblems] = useState<Problem[]>([]);
  const [selectedCategory, setSelectedCategory] = useState(() => searchParams.get('category') ?? 'loop');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProblems = async (category: string) => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await getProblems(category);
      setProblems(data);
    } catch {
      setError('Tidak dapat terhubung ke server. Coba lagi nanti.');
    } finally {
      setIsLoading(false);
    }
  };

  // fetch all problems once for global progress
  useEffect(() => {
    Promise.all([
      getProblems('loop'),
      getProblems('string'),
      getProblems('array'),
      getProblems('sql'),
    ]).then((results) => setAllProblems(results.flat())).catch(() => {});
  }, []);

  useEffect(() => {
    fetchProblems(selectedCategory);
    window.scrollTo({ top: 0, behavior: 'smooth' });
    router.replace(`/problems?category=${selectedCategory}`, { scroll: false });
  }, [selectedCategory]);

  const completedCount = allProblems.filter((p) => progress[p.id] === 'accepted').length;
  const totalCount = allProblems.length;
  const progressPercent = totalCount > 0 ? Math.round((completedCount / totalCount) * 100) : 0;

  return (
    <main className="mx-auto max-w-4xl px-6 py-10">
      <div className="flex justify-between items-center">
        <div className="mb-7">
          <h1 className="text-2xl font-bold text-gray-900">Daftar Soal</h1>
          <p className="mt-1 text-sm text-gray-500">Pilih soal dan mulai latihan.</p>
        </div>

        <div className="mb-6 rounded-2xl border border-indigo-100 bg-gradient-to-br from-indigo-50 to-purple-50 px-5 py-4">
          <div className="mb-3 flex items-center justify-between">
            <div>
              <p className="text-xs font-medium uppercase tracking-wide text-indigo-400">Progress</p>
              <p className="mt-0.5 text-sm text-gray-700">
                <span className="text-lg font-bold text-indigo-600">{completedCount}</span>
                <span className="text-gray-400"> / {totalCount} soal</span>
              </p>
            </div>
            <div className="flex h-14 w-14 ml-5 items-center justify-center rounded-full bg-white shadow-sm ring-1 ring-indigo-100">
              <span className="text-base font-bold text-indigo-600">{progressPercent}%</span>
            </div>
          </div>
          <div className="h-2 w-full overflow-hidden rounded-full bg-indigo-100">
            <div
              className="h-full rounded-full bg-gradient-to-r from-indigo-500 to-purple-500 transition-all duration-500"
              style={{ width: `${progressPercent}%` }}
            />
          </div>
        </div>
      </div>

      {error && (
        <div className="mb-5 flex items-center justify-between rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-600">
          <span>{error}</span>
          <button onClick={() => fetchProblems(selectedCategory)} className="ml-4 font-medium underline hover:no-underline">
            Coba lagi
          </button>
        </div>
      )}

      <CategoryFilter selected={selectedCategory} onChange={setSelectedCategory} />

      {isLoading ? (
        <ProblemListSkeleton />
      ) : (
        <ProblemTable problems={problems} onRowClick={(id) => router.push(`/problems/${id}`)} />
      )}
    </main>
  );
}

export default function ProblemsPage() {
  return (
    <Suspense>
      <ProblemsPageInner />
    </Suspense>
  );
}
