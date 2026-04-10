'use client';

import Link from 'next/link';
import { useProgress } from '@/hooks/useProgress';

interface PrerequisiteWarningProps {
  prerequisiteId: string | null;
  prerequisiteTitle: string | null;
}

export default function PrerequisiteWarning({
  prerequisiteId,
  prerequisiteTitle,
}: PrerequisiteWarningProps) {
  const { isAccepted } = useProgress();

  if (prerequisiteId === null) return null;
  if (isAccepted(prerequisiteId)) return null;

  const displayTitle = prerequisiteTitle ?? 'soal prerequisite';

  return (
    <div className="rounded-xl border border-amber-200 bg-amber-50 px-4 py-3">
      <p className="text-sm text-amber-800">
        ⚠️ Disarankan selesaikan{' '}
        <Link
          href={`/problems/${prerequisiteId}`}
          className="font-semibold underline hover:text-amber-900"
        >
          {displayTitle}
        </Link>{' '}
        dulu
      </p>
    </div>
  );
}
