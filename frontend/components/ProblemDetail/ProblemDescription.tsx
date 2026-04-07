import { Problem } from '@/lib/types';

interface ProblemDescriptionProps {
  problem: Problem;
}

export default function ProblemDescription({ problem }: ProblemDescriptionProps) {
  return (
    <div className="space-y-3">
      <h1 className="text-lg font-bold text-gray-900 leading-snug">{problem.title}</h1>
      <p className="text-sm text-gray-600 whitespace-pre-wrap leading-relaxed">
        {problem.description}
      </p>
    </div>
  );
}
