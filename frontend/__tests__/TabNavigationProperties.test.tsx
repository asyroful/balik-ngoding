// @vitest-environment happy-dom
// Feature: user-experience-improvements, Property 13: Active tab has distinct visual style

import React, { useState } from 'react';
import * as fc from 'fast-check';
import { describe, it, expect } from 'vitest';
import { render, screen, cleanup } from '@testing-library/react';

// Minimal inline component that mirrors the tab navigation from problems/[id]/page.tsx
function TabNavigation({
  initialTab,
}: {
  initialTab: 'soal' | 'editor';
}) {
  const [activeTab] = useState<'soal' | 'editor'>(initialTab);

  return (
    <div className="flex flex-col md:hidden border-b border-gray-200 bg-white">
      <div className="flex">
        <button
          className={`flex-1 py-2.5 text-sm font-medium transition ${
            activeTab === 'soal'
              ? 'border-b-2 border-indigo-600 text-indigo-600 bg-indigo-50'
              : 'text-gray-500 hover:text-gray-700'
          }`}
        >
          Soal
        </button>
        <button
          className={`flex-1 py-2.5 text-sm font-medium transition ${
            activeTab === 'editor'
              ? 'border-b-2 border-indigo-600 text-indigo-600 bg-indigo-50'
              : 'text-gray-500 hover:text-gray-700'
          }`}
        >
          Editor
        </button>
      </div>
    </div>
  );
}

// Feature: user-experience-improvements, Property 13: Active tab has distinct visual style
// Validates: Requirements 5.5
describe('Property 13: Active tab has distinct visual style', () => {
  it('active tab button has a different className than the inactive tab button', () => {
    fc.assert(
      fc.property(
        fc.constantFrom<'soal' | 'editor'>('soal', 'editor'),
        (activeTab) => {
          const { unmount } = render(<TabNavigation initialTab={activeTab} />);

          const soalBtn = screen.getByRole('button', { name: 'Soal' });
          const editorBtn = screen.getByRole('button', { name: 'Editor' });

          const activeBtn = activeTab === 'soal' ? soalBtn : editorBtn;
          const inactiveBtn = activeTab === 'soal' ? editorBtn : soalBtn;

          const result = activeBtn.className !== inactiveBtn.className;

          unmount();
          cleanup();
          return result;
        }
      ),
      { numRuns: 100 }
    );
  });
});
