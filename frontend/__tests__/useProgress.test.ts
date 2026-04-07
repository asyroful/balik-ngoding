// @vitest-environment happy-dom
import { describe, it, expect, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { useProgress } from '../hooks/useProgress';

// Clear all cookies before each test
beforeEach(() => {
  document.cookie.split(';').forEach((cookie) => {
    const name = cookie.split('=')[0].trim();
    document.cookie = `${name}=; max-age=0; path=/`;
  });
});

// ─── Req 3.1, 3.2: markAccepted writes to document.cookie ────────────────────
describe('useProgress — markAccepted writes to cookie (Req 3.1, 3.2)', () => {
  it('should write bn_progress cookie after markAccepted', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-1');
    });

    expect(document.cookie).toContain('bn_progress=');
  });

  it('should store problemId with "accepted" value in cookie JSON', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-abc');
    });

    const match = document.cookie
      .split('; ')
      .find((row) => row.startsWith('bn_progress='));
    expect(match).toBeDefined();
    const value = decodeURIComponent(match!.split('=').slice(1).join('='));
    const parsed = JSON.parse(value);
    expect(parsed['problem-abc']).toBe('accepted');
  });
});

// ─── Req 3.7: Corrupt cookie → progress is {} ────────────────────────────────
describe('useProgress — corrupt cookie fallback (Req 3.7)', () => {
  it('should return empty progress when cookie is corrupt JSON', () => {
    document.cookie = `bn_progress=${encodeURIComponent('{invalid json')}; path=/`;

    const { result } = renderHook(() => useProgress());

    expect(result.current.progress).toEqual({});
  });

  it('should return empty progress when cookie is a plain string', () => {
    document.cookie = `bn_progress=${encodeURIComponent('not-json')}; path=/`;

    const { result } = renderHook(() => useProgress());

    expect(result.current.progress).toEqual({});
  });

  it('should return empty progress when cookie is an array (not an object)', () => {
    document.cookie = `bn_progress=${encodeURIComponent('["a","b"]')}; path=/`;

    const { result } = renderHook(() => useProgress());

    expect(result.current.progress).toEqual({});
  });
});

// ─── Req 3.7: No cookie → progress is {} ─────────────────────────────────────
describe('useProgress — no cookie fallback (Req 3.7)', () => {
  it('should return empty progress when bn_progress cookie is absent', () => {
    const { result } = renderHook(() => useProgress());

    expect(result.current.progress).toEqual({});
  });
});

// ─── Req 3.3: isAccepted returns true for accepted problemId ─────────────────
describe('useProgress — isAccepted (Req 3.3)', () => {
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

// ─── Req 3.8: markAccepted does not overwrite existing 'accepted' status ──────
describe('useProgress — markAccepted idempotent (Req 3.8)', () => {
  it('should preserve accepted status when markAccepted is called again', () => {
    const { result } = renderHook(() => useProgress());

    act(() => {
      result.current.markAccepted('problem-1');
    });

    // Call again — should not throw and status should remain 'accepted'
    act(() => {
      result.current.markAccepted('problem-1');
    });

    expect(result.current.isAccepted('problem-1')).toBe(true);
    expect(result.current.progress['problem-1']).toBe('accepted');
  });

  it('should not overwrite accepted status when cookie already has it', () => {
    // Pre-populate cookie with accepted status
    const existing = { 'problem-2': 'accepted' };
    document.cookie = `bn_progress=${encodeURIComponent(JSON.stringify(existing))}; path=/`;

    const { result } = renderHook(() => useProgress());

    // Verify it was read correctly
    expect(result.current.isAccepted('problem-2')).toBe(true);

    // Calling markAccepted again should be a no-op
    act(() => {
      result.current.markAccepted('problem-2');
    });

    expect(result.current.isAccepted('problem-2')).toBe(true);
  });
});
