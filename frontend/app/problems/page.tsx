'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getProblems } from '@/lib/api';
import { Problem } from '@/lib/types';
import CategoryFilter from '@/components/ProblemList/CategoryFilter';
import ProblemTable from '@/components/ProblemList/ProblemTable';
import { ProblemListSkeleton } from '@/components/Skeleton';

export default function ProblemsPage() {
  const router = useRouter();
  const [problems, setProblems] = useState<Problem[]>([]);
  const [selectedCategory, setSelectedCategory] = useState('loop');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProblems = async () => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await getProblems();
      setProblems(data);
    } catch {
      setError('Tidak dapat terhubung ke server. Coba lagi nanti.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => { fetchProblems(); }, []);

  const filtered = selectedCategory
    ? problems.filter((p) => p.category === selectedCategory)
    : problems;

  return (
    <main className="mx-auto max-w-4xl px-6 py-10">
      <div className="mb-7">
        <h1 className="text-2xl font-bold text-gray-900">Daftar Soal</h1>
        <p className="mt-1 text-sm text-gray-500">Pilih soal dan mulai latihan.</p>
      </div>

      {error && (
        <div className="mb-5 flex items-center justify-between rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-600">
          <span>{error}</span>
          <button onClick={fetchProblems} className="ml-4 font-medium underline hover:no-underline">
            Coba lagi
          </button>
        </div>
      )}

      <CategoryFilter selected={selectedCategory} onChange={setSelectedCategory} />

      {isLoading ? (
        <ProblemListSkeleton />
      ) : (
        <ProblemTable problems={filtered} onRowClick={(id) => router.push(`/problems/${id}`)} />
      )}
    </main>
  );
}
