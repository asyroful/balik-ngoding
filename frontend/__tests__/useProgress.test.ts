// @vitest-environment happy-dom
import { describe, it, expect, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { useProgress } from '../hooks/useProgress';

const STORAGE_KEY = 'balik-ngoding-progress';

beforeEach(() => {
  localStorage.clear();
});

// ─── markAccepted writes to localStorage ─────────────────────────────────────
describe('useProgress — markAccepted writes to localStorage', () => {
  it('should write to localStorage after markAccepted', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-1');
    });

    const raw = localStorage.getItem(STORAGE_KEY);
    expect(raw).not.toBeNull();
  });

  it('should store problemId with solved: true in localStorage JSON', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-abc');
    });

    const raw = localStorage.getItem(STORAGE_KEY);
    const parsed = JSON.parse(raw!);
    expect(parsed['problem-abc'].solved).toBe(true);
  });
});

// ─── Corrupt localStorage → progress is empty ────────────────────────────────
describe('useProgress — corrupt localStorage fallback', () => {
  it('should return empty progress when localStorage has corrupt JSON', () => {
    localStorage.setItem(STORAGE_KEY, '{invalid json');

    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('any-id')).toBe(false);
    expect(result.current.getHintsUnlocked('any-id')).toBe(0);
  });

  it('should return empty progress when localStorage has a plain string', () => {
    localStorage.setItem(STORAGE_KEY, 'not-json');

    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('any-id')).toBe(false);
  });

  it('should return empty progress when localStorage has an array (not an object)', () => {
    localStorage.setItem(STORAGE_KEY, '["a","b"]');

    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('any-id')).toBe(false);
  });
});

// ─── No localStorage entry → progress is empty ───────────────────────────────
describe('useProgress — no localStorage entry', () => {
  it('should return empty progress when key is absent', () => {
    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('problem-unknown')).toBe(false);
    expect(result.current.getHintsUnlocked('problem-unknown')).toBe(0);
  });
});

// ─── isAccepted ───────────────────────────────────────────────────────────────
describe('useProgress — isAccepted', () => {
  it('should return true for a problemId that was marked accepted', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-xyz');
    });

    expect(result.current.isAccepted('problem-xyz')).toBe(true);
  });

  it('should return false for a problemId that was never marked accepted', () => {
    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('problem-unknown')).toBe(false);
  });
});

// ─── markAccepted idempotent ──────────────────────────────────────────────────
describe('useProgress — markAccepted idempotent', () => {
  it('should preserve accepted status when markAccepted is called again', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-1');
    });

    act(() => {
      result.current.markAccepted('problem-1');
    });

    expect(result.current.isAccepted('problem-1')).toBe(true);
  });

  it('should read existing solved state from localStorage on init', () => {
    const existing = { 'problem-2': { solved: true, hintsUnlocked: 0 } };
    localStorage.setItem(STORAGE_KEY, JSON.stringify(existing));

    const { result } = renderHook(() => useProgress());

    expect(result.current.isAccepted('problem-2')).toBe(true);
  });
});

// ─── getHintsUnlocked ─────────────────────────────────────────────────────────
describe('useProgress — getHintsUnlocked', () => {
  it('should return 0 for a problem with no hints unlocked', () => {
    const { result } = renderHook(() => useProgress());

    expect(result.current.getHintsUnlocked('problem-1')).toBe(0);
  });

  it('should return the correct count after unlocking hints', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.unlockNextHint('problem-1', 3);
    });

    expect(result.current.getHintsUnlocked('problem-1')).toBe(1);

    act(() => {
      result.current.unlockNextHint('problem-1', 3);
    });

    expect(result.current.getHintsUnlocked('problem-1')).toBe(2);
  });
});

// ─── unlockNextHint ───────────────────────────────────────────────────────────
describe('useProgress — unlockNextHint', () => {
  it('should not exceed totalHints', () => {
    const { result } = renderHook(() => useProgress());

    act(() => { result.current.unlockNextHint('problem-1', 2); });
    act(() => { result.current.unlockNextHint('problem-1', 2); });
    act(() => { result.current.unlockNextHint('problem-1', 2); }); // capped at 2

    expect(result.current.getHintsUnlocked('problem-1')).toBe(2);
  });

  it('should be idempotent at upper bound', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.unlockNextHint('problem-1', 1);
    });
    act(() => {
      result.current.unlockNextHint('problem-1', 1);
    });

    expect(result.current.getHintsUnlocked('problem-1')).toBe(1);
  });

  it('should preserve solved state when unlocking hints', () => {
    const { result } = renderHook(() => useProgress());

    act(() => { result.current.markAccepted('problem-1'); });
    act(() => { result.current.unlockNextHint('problem-1', 3); });

    expect(result.current.isAccepted('problem-1')).toBe(true);
    expect(result.current.getHintsUnlocked('problem-1')).toBe(1);
  });

  it('should not affect other problems', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.unlockNextHint('problem-A', 3);
    });

    expect(result.current.getHintsUnlocked('problem-B')).toBe(0);
    expect(result.current.isAccepted('problem-B')).toBe(false);
  });
});

// ─── localStorage fallback (in-memory) ───────────────────────────────────────
describe('useProgress — localStorage unavailable fallback', () => {
  it('should work in-memory when localStorage throws on setItem', () => {
    const original = localStorage.setItem.bind(localStorage);
    localStorage.setItem = () => { throw new Error('QuotaExceededError'); };

    const { result } = renderHook(() => useProgress());

    expect(() => {
      act(() => {
        result.current.markAccepted('problem-1');
      });
    }).not.toThrow();

    expect(result.current.isAccepted('problem-1')).toBe(true);

    localStorage.setItem = original;
  });
});
