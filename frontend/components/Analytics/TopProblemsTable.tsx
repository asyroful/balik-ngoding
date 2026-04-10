import { TopProblem } from '@/lib/types';

interface TopProblemsTableProps {
  problems: TopProblem[];
}

export function TopProblemsTable({ problems }: TopProblemsTableProps) {
  if (problems.length === 0) {
    return <p className="text-gray-400 text-sm">Belum ada data submission.</p>;
  }

  return (
    <table className="w-full text-sm text-left">
      <thead>
        <tr className="border-b border-gray-700 text-gray-400">
          <th className="pb-2 font-medium">#</th>
          <th className="pb-2 font-medium">Soal</th>
          <th className="pb-2 font-medium text-right">Submissions</th>
        </tr>
      </thead>
      <tbody>
        {problems.map((p, i) => (
          <tr key={p.problemId} className="border-b border-gray-800">
            <td className="py-3 text-gray-500">{i + 1}</td>
            <td className="py-3 text-black">{p.title}</td>
            <td className="py-3 text-right text-black">{p.submissionCount}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
