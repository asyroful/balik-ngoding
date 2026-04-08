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

  const StatusIcon = ({ id }: { id: string }) =>
    isAccepted(id) ? (
      <svg className="w-4 h-4 shrink-0 text-green-600" viewBox="0 0 20 20" fill="currentColor">
        <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
      </svg>
    ) : (
      <svg className="w-4 h-4 shrink-0 text-gray-300" viewBox="0 0 20 20" fill="none" stroke="currentColor" strokeWidth="1.5">
        <circle cx="10" cy="10" r="8" />
      </svg>
    );

  return (
    <>
      {/* Mobile: card list */}
      <div className="flex flex-col gap-2 md:hidden">
        {problems.map((problem, index) => (
          <div
            key={problem.id}
            onClick={() => onRowClick(problem.id)}
            className="cursor-pointer rounded-xl border border-gray-200 bg-white px-4 py-3 transition hover:bg-indigo-50/50 active:bg-indigo-100"
          >
            <div className="flex justify-between items-center">
              <div>
                <div className="text-sm font-medium text-gray-800 truncate">
                  {index + 1}. {problem.title}
                </div>
                <div className="mt-2 flex items-center gap-2">
                  <span className={`inline-block rounded-md border px-2 py-0.5 text-xs font-medium capitalize ${CATEGORY_STYLE[problem.category] ?? 'bg-gray-50 text-gray-600 border-gray-200'}`}>
                    {problem.category}
                  </span>
                  <span className={`inline-block rounded-md border px-2 py-0.5 text-xs font-medium capitalize ${DIFFICULTY_STYLE[problem.difficulty] ?? 'bg-gray-50 text-gray-600 border-gray-200'}`}>
                    {problem.difficulty}
                  </span>
                </div>
              </div>
              <StatusIcon id={problem.id} />
            </div>
          </div>
        ))}
      </div>

      {/* Desktop: table */}
      <div className="hidden md:block overflow-hidden rounded-xl border border-gray-200 bg-white">
        <table className="min-w-full">
          <thead>
            <tr className="border-b border-gray-100 bg-gray-50">
              <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-12">#</th>
              <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400">Judul</th>
              <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-28">Kategori</th>
              <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-28">Level</th>
              <th className="px-5 py-3 text-left text-xs font-semibold uppercase tracking-wider text-gray-400 w-24"></th>
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
                  <StatusIcon id={problem.id} />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}
