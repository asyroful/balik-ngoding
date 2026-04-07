import { TestCase } from '@/lib/types';

interface ExampleBlockProps {
  testCases: TestCase[];
}

export default function ExampleBlock({ testCases }: ExampleBlockProps) {
  const visible = testCases.filter((tc) => !tc.isHidden);
  if (visible.length === 0) return null;

  return (
    <div className="space-y-3">
      <h2 className="text-xs font-semibold uppercase tracking-widest text-gray-400">Contoh</h2>
      {visible.map((tc, i) => (
        <div key={tc.id} className="rounded-xl border border-gray-100 bg-gray-50 p-4 space-y-3">
          <p className="text-xs font-semibold text-gray-400">Contoh {i + 1}</p>
          <div className="space-y-1">
            <span className="text-xs font-medium text-gray-500">Input</span>
            <pre className="rounded-lg bg-white border border-gray-200 px-3 py-2 text-xs font-mono text-gray-800 whitespace-pre-wrap">
              {tc.input}
            </pre>
          </div>
          <div className="space-y-1">
            <span className="text-xs font-medium text-gray-500">Output</span>
            <pre className="rounded-lg bg-white border border-gray-200 px-3 py-2 text-xs font-mono text-indigo-700 whitespace-pre-wrap">
              {tc.expectedOutput}
            </pre>
          </div>
        </div>
      ))}
    </div>
  );
}
