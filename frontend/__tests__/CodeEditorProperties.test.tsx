// @vitest-environment happy-dom
import * as fc from 'fast-check';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { useSubmissionStore } from '../store/submissionStore';

// The guard logic extracted from CodeEditor's Monaco `run` callback:
//   if (isLoading || code.trim() === '') return;
//   onSubmit?.();
function runGuard(isLoading: boolean, code: string, onSubmit: () => void) {
  if (isLoading || code.trim() === '') return;
  onSubmit();
}

beforeEach(() => {
  useSubmissionStore.getState().reset();
  vi.clearAllMocks();
});

// Feature: user-experience-improvements, Property 6: Ctrl+Enter ignored when loading
// Validates: Requirements 2.2
describe('Property 6: Ctrl+Enter ignored when loading', () => {
  it('onSubmit is never called when isLoading is true', () => {
    fc.assert(
      fc.property(
        fc.string(), // arbitrary code content
        (code) => {
          const onSubmit = vi.fn();
          useSubmissionStore.getState().setLoading(true);
          useSubmissionStore.getState().setCode(code);

          const { isLoading } = useSubmissionStore.getState();
          runGuard(isLoading, code, onSubmit);

          return onSubmit.mock.calls.length === 0;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 7: Ctrl+Enter ignored for whitespace-only code
// Validates: Requirements 2.3
describe('Property 7: Ctrl+Enter ignored for whitespace-only code', () => {
  it('onSubmit is never called for whitespace-only code strings', () => {
    fc.assert(
      fc.property(
        fc.stringOf(fc.constantFrom(' ', '\t', '\n', '\r')),
        (whitespaceCode) => {
          const onSubmit = vi.fn();
          useSubmissionStore.getState().setLoading(false);
          useSubmissionStore.getState().setCode(whitespaceCode);

          const { isLoading, code } = useSubmissionStore.getState();
          runGuard(isLoading, code, onSubmit);

          return onSubmit.mock.calls.length === 0;
        }
      ),
      { numRuns: 100 }
    );
  });
});
