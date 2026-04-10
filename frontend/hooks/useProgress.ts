'use client';

import { useState } from 'react';

export const STORAGE_KEY = 'balik-ngoding-progress';

export interface ProblemProgress {
  solved: boolean;
  hintsUnlocked: number; // 0..N, never decreases
}

export type ProgressStore = Record<string, ProblemProgress>;

interface UseProgressReturn {
  markAccepted: (problemId: string) => void;
  isAccepted: (problemId: string) => boolean;
  getHintsUnlocked: (problemId: string) => number;
  unlockNextHint: (problemId: string, totalHints: number) => void;
}

export function readFromStorage(): ProgressStore {
  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) return {};
    const parsed = JSON.parse(raw);
    if (typeof parsed === 'object' && parsed !== null && !Array.isArray(parsed)) {
      return parsed as ProgressStore;
    }
    return {};
  } catch {
    return {};
  }
}

export function writeToStorage(store: ProgressStore): void {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(store));
  } catch {
    // localStorage not available — in-memory state is the fallback
  }
}

export function applyMarkAccepted(store: ProgressStore, problemId: string): ProgressStore {
  const current = store[problemId];
  if (current?.solved) return store;
  return {
    ...store,
    [problemId]: {
      solved: true,
      hintsUnlocked: current?.hintsUnlocked ?? 0,
    },
  };
}

export function applyUnlockNextHint(
  store: ProgressStore,
  problemId: string,
  totalHints: number
): ProgressStore {
  const current = store[problemId];
  const currentCount = current?.hintsUnlocked ?? 0;
  if (currentCount >= totalHints) return store;
  return {
    ...store,
    [problemId]: {
      solved: current?.solved ?? false,
      hintsUnlocked: currentCount + 1,
    },
  };
}

export function useProgress(): UseProgressReturn {
  const [progress, setProgress] = useState<ProgressStore>(() => readFromStorage());

  function markAccepted(problemId: string): void {
    const updated = applyMarkAccepted(progress, problemId);
    if (updated === progress) return;
    writeToStorage(updated);
    setProgress(updated);
  }

  function isAccepted(problemId: string): boolean {
    return progress[problemId]?.solved === true;
  }

  function getHintsUnlocked(problemId: string): number {
    return progress[problemId]?.hintsUnlocked ?? 0;
  }

  function unlockNextHint(problemId: string, totalHints: number): void {
    const updated = applyUnlockNextHint(progress, problemId, totalHints);
    if (updated === progress) return;
    writeToStorage(updated);
    setProgress(updated);
  }

  return { markAccepted, isAccepted, getHintsUnlocked, unlockNextHint };
}
