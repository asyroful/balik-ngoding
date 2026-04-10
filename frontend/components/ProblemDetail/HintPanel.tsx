'use client';

import { useProgress } from '@/hooks/useProgress';

interface HintPanelProps {
  hints: string[] | null | undefined;
  problemId: string;
}

export default function HintPanel({ hints, problemId }: HintPanelProps) {
  const { getHintsUnlocked, unlockNextHint } = useProgress();
  const hintsUnlocked = getHintsUnlocked(problemId);
  if (!hints || hints.length === 0) return null;

  return (
    <div className="rounded-xl border border-amber-100 bg-amber-50 overflow-hidden">
      <div className="flex items-center justify-between px-4 py-3">
        <span className="flex items-center gap-2 text-sm font-semibold text-amber-700">
          <span>🔍</span>
          <span>Hint</span>
        </span>
        <span className="text-xs text-amber-500 font-medium">
          Hint {hintsUnlocked} dari {hints.length}
        </span>
      </div>

      {hintsUnlocked > 0 && (
        <div className="px-4 pb-3 flex flex-col gap-2 border-t border-amber-100">
          {hints.slice(0, hintsUnlocked).map((hint, index) => (
            <div key={index} className="pt-3">
              <p className="text-xs font-semibold text-amber-600 mb-1">Hint {index + 1}</p>
              <p className="text-sm text-amber-900 whitespace-pre-wrap">{hint}</p>
            </div>
          ))}
        </div>
      )}

      {hintsUnlocked < hints.length && (
        <div className={`px-4 pb-3 ${hintsUnlocked > 0 ? '' : 'pt-0'}`}>
          <button
            onClick={() => unlockNextHint(problemId, hints.length)}
            className="mt-2 w-full rounded-lg bg-amber-400 hover:bg-amber-500 text-amber-900 text-sm font-semibold py-2 transition-colors"
          >
            Buka Hint Berikutnya
          </button>
        </div>
      )}
    </div>
  );
}
