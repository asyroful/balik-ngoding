'use client';

import { useState } from 'react';

interface ThinkingGuideProps {
  content: string | null | undefined;
}

export default function ThinkingGuide({ content }: ThinkingGuideProps) {
  const [isOpen, setIsOpen] = useState(false);
  if (!content) return null;

  return (
    <div className="rounded-xl border border-indigo-100 bg-indigo-50 overflow-hidden">
      <button
        onClick={() => setIsOpen((prev) => !prev)}
        className="w-full flex items-center justify-between px-4 py-3 text-left hover:bg-indigo-100 transition-colors"
      >
        <span className="flex items-center gap-2 text-sm font-semibold text-indigo-700">
          <span>💡</span>
          <span>Cara Berpikir</span>
        </span>
        <svg
          className={`w-4 h-4 text-indigo-500 transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`}
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          strokeWidth={2}
        >
          <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
        </svg>
      </button>
      {isOpen && (
        <div className="px-4 py-3 border-t border-indigo-100">
          <p className="text-sm text-indigo-900 whitespace-pre-wrap">{content}</p>
        </div>
      )}
    </div>
  );
}
