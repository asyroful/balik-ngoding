'use client';

import { useState } from 'react';

type ProgressMap = Record<string, 'accepted'>;

interface UseProgressReturn {
  progress: ProgressMap;
  markAccepted: (problemId: string) => void;
  isAccepted: (problemId: string) => boolean;
}

function readProgressCookie(): ProgressMap {
  try {
    const match = document.cookie
      .split('; ')
      .find((row) => row.startsWith('bn_progress='));
    if (!match) return {};
    const value = decodeURIComponent(match.split('=').slice(1).join('='));
    const parsed = JSON.parse(value);
    if (typeof parsed === 'object' && parsed !== null && !Array.isArray(parsed)) {
      return parsed as ProgressMap;
    }
    return {};
  } catch {
    return {};
  }
}

function writeProgressCookie(progress: ProgressMap): void {
  const value = encodeURIComponent(JSON.stringify(progress));
  document.cookie = `bn_progress=${value}; max-age=31536000; path=/; SameSite=Lax`;
}

export function useProgress(): UseProgressReturn {
  const [progress, setProgress] = useState<ProgressMap>(() => readProgressCookie());

  function markAccepted(problemId: string): void {
    if (progress[problemId] === 'accepted') return;
    const updated = { ...progress, [problemId]: 'accepted' as const };
    writeProgressCookie(updated);
    setProgress(updated);
  }

  function isAccepted(problemId: string): boolean {
    return progress[problemId] === 'accepted';
  }

  return { progress, markAccepted, isAccepted };
}
