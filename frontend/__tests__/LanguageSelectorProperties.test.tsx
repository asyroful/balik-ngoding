// @vitest-environment happy-dom
import React from 'react';
import * as fc from 'fast-check';
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import LanguageSelector, { getDefaultLanguage } from '../components/Editor/LanguageSelector';

// Helper: render LanguageSelector with arbitrary selected/category
function renderSelector(selected = 'javascript', category = 'loop') {
  return render(
    <LanguageSelector
      selected={selected}
      onChange={() => {}}
      problemCategory={category}
    />
  );
}

// Feature: user-experience-improvements, Property 1: Language list completeness
// Validates: Requirements 1.1
describe('Property 1: Language list completeness', () => {
  it('render always produces exactly 6 language buttons', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('javascript', 'sql', 'python', 'java', 'php', 'c'),
        fc.constantFrom('loop', 'string', 'array', 'sql'),
        (selected, category) => {
          const { unmount } = renderSelector(selected, category);
          const buttons = screen.getAllByRole('button');
          const count = buttons.length;
          unmount();
          return count === 6;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 2: Active languages are selectable
// Validates: Requirements 1.2
describe('Property 2: Active languages are selectable', () => {
  it('JavaScript and SQL are never disabled', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('javascript', 'sql', 'python', 'java', 'php', 'c'),
        fc.constantFrom('loop', 'string', 'array', 'sql'),
        (selected, category) => {
          const { unmount } = renderSelector(selected, category);
          const buttons = screen.getAllByRole('button');

          const jsBtn = buttons.find(
            (b) => b.textContent?.toLowerCase().includes('javascript') && !b.textContent?.toLowerCase().includes('coming soon')
          );
          const sqlBtn = buttons.find(
            (b) => b.textContent?.trim().toLowerCase() === 'sql'
          );

          const jsNotDisabled = jsBtn ? !jsBtn.hasAttribute('disabled') : false;
          const sqlNotDisabled = sqlBtn ? !sqlBtn.hasAttribute('disabled') : false;

          unmount();
          return jsNotDisabled && sqlNotDisabled;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 3: Coming soon languages are disabled and labeled
// Validates: Requirements 1.3, 1.4, 1.6
describe('Property 3: Coming soon languages are disabled and labeled', () => {
  // Map value -> display label for precise matching
  const comingSoonMap: Record<string, string> = {
    python: 'Python',
    java: 'Java',
    php: 'PHP',
    c: 'C',
  };

  it('Python/Java/PHP/C are always disabled and labeled "Coming Soon"', () => {
    fc.assert(
      fc.property(
        fc.constantFrom('python', 'java', 'php', 'c') as fc.Arbitrary<string>,
        fc.constantFrom('loop', 'string', 'array', 'sql'),
        (comingSoonLang, category) => {
          const { unmount, container } = renderSelector('javascript', category);
          const buttons = container.querySelectorAll('button');
          const label = comingSoonMap[comingSoonLang];

          // Find button whose first text node exactly matches the label
          // (e.g. "Java" not "JavaScript"), using the button's first child text
          let btn: Element | null = null;
          buttons.forEach((b) => {
            // The button's first text node should be the language label
            const firstText = Array.from(b.childNodes)
              .filter((n) => n.nodeType === Node.TEXT_NODE)
              .map((n) => n.textContent?.trim())
              .find((t) => t);
            if (firstText === label) {
              btn = b;
            }
          });

          const isDisabled = btn ? (btn as HTMLButtonElement).hasAttribute('disabled') : false;
          const hasComingSoon = btn
            ? /coming soon/i.test((btn as HTMLButtonElement).textContent ?? '')
            : false;

          unmount();
          return isDisabled && hasComingSoon;
        }
      ),
      { numRuns: 100 }
    );
  });
});

// Feature: user-experience-improvements, Property 4: Default language matches problem category
// Validates: Requirements 1.5
describe('Property 4: Default language matches problem category', () => {
  it('getDefaultLanguage returns "sql" only when category === "sql", otherwise "javascript"', () => {
    fc.assert(
      fc.property(
        fc.oneof(
          fc.constant('sql'),
          fc.string().filter((s) => s !== 'sql')
        ),
        (category) => {
          const lang = getDefaultLanguage(category);
          if (category === 'sql') {
            return lang === 'sql';
          }
          return lang === 'javascript';
        }
      ),
      { numRuns: 100 }
    );
  });
});
