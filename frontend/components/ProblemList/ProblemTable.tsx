'use client';

import { Problem } from '@/lib/types';
import { useProgress } from '@/hooks/useProgress';

interface ProblemTableProps {
  problems: Problem[];
  onRowClick: (id: string) => void;
}

const DIFFICULTY_STYLE: Record<string, string> = {
  easy: 'bg-green-50 text-green-700 border-green-200',
  medium: 'bg-yellow-50 text-yellow-700 border-yellow-200',
  hard: 'bg-red-50 text-red-700 border-red-200',
};

const CATEGORY_STYLE: Record<string, string> = {
  loop: 'bg-blue-50 text-blue-700 border-blue-200',
  string: 'bg-purple-50 text-purple-700 border-purple-200',
  array: 'bg-orange-50 text-orange-700 border-orange-200',
  sql: 'bg-teal-50 text-teal-700 border-teal-200',
};

export default function ProblemTable({ problems, onRowClick }: ProblemTableProps) {
  const { isAccepted } = useProgress();
  if (problems.length === 0) {
    return (
      <div className="rounded-xl border border-gray-200 bg-white py-16 text-center text-sm text-gray-400">
        Tidak ada soal yang tersedia.
      </div>
    );
  }

  return (
    <div className="overflow-hidden rounded-xl border border-gray-200 bg-white">
      <table className="min-w-full">
        <thead>
          <tr className="border-b border-gray-100 bg-gray-50">
            <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-12">#</th>
            <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400">Judul</th>
            <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-28">Kategori</th>
            <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-28">Level</th>
            <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-24">Status</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-gray-100">
          {problems.map((problem, index) => (
            <tr
              key={problem.id}
              onClick={() => onRowClick(problem.id)}
              className="cursor-pointer transition hover:bg-indigo-50/50 group"
            >
              <td className="px-5 py-4 text-sm text-gray-400">{index + 1}</td>
              <td className="px-5 py-4 text-sm font-medium text-gray-800 group-hover:text-indigo-700 transition">
                {problem.title}
              </td>
              <td className="px-5 py-4">
                <span className={`inline-block rounded-md border px-2 py-0.5 text-xs font-medium capitalize ${CATEGORY_STYLE[problem.category] ?? 'bg-gray-50 text-gray-600 border-gray-200'}`}>
                  {problem.category}
                </span>
              </td>
              <td className="px-5 py-4">
                <span className={`inline-block rounded-md border px-2 py-0.5 text-xs font-medium capitalize ${DIFFICULTY_STYLE[problem.difficulty] ?? 'bg-gray-50 text-gray-600 border-gray-200'}`}>
                  {problem.difficulty}
                </span>
              </td>
              <td className="px-5 py-4">
                {isAccepted(problem.id) && (
                  <span className="inline-flex items-center gap-1 rounded-md border border-green-200 bg-green-50 px-2 py-0.5 text-xs font-medium text-green-700">
                    ✓ Selesai
                  </span>
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
