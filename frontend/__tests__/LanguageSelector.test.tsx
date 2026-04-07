import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import LanguageSelector from '../components/Editor/LanguageSelector';

// Helper: render LanguageSelector with default props
function renderSelector(selected = 'javascript', onChange = vi.fn()) {
  return render(
    <LanguageSelector
      selected={selected}
      onChange={onChange}
      problemCategory="loop"
    />
  );
}

// ─── 1.1: Render menampilkan 6 bahasa ────────────────────────────────────────
describe('LanguageSelector — 6 bahasa ditampilkan (Req 1.1)', () => {
  it('should render exactly 6 buttons', () => {
    renderSelector();
    const buttons = screen.getAllByRole('button');
    expect(buttons).toHaveLength(6);
  });
});

// ─── 1.2: JS dan SQL aktif (tidak disabled, tidak ada "Coming Soon") ─────────
describe('LanguageSelector — JS dan SQL aktif (Req 1.2)', () => {
  it('JavaScript button should not be disabled', () => {
    renderSelector();
    const jsBtn = screen.getByRole('button', { name: /javascript/i });
    expect(jsBtn.hasAttribute('disabled')).toBe(false);
  });

  it('SQL button should not be disabled', () => {
    renderSelector();
    const sqlBtn = screen.getByRole('button', { name: /^sql$/i });
    expect(sqlBtn.hasAttribute('disabled')).toBe(false);
  });

  it('JavaScript button should not contain "Coming Soon" text', () => {
    renderSelector();
    const jsBtn = screen.getByRole('button', { name: /javascript/i });
    expect(jsBtn.textContent).not.toMatch(/coming soon/i);
  });

  it('SQL button should not contain "Coming Soon" text', () => {
    renderSelector();
    const sqlBtn = screen.getByRole('button', { name: /^sql$/i });
    expect(sqlBtn.textContent).not.toMatch(/coming soon/i);
  });
});

// ─── 1.3 & 1.4: Python/Java/PHP/C disabled dengan "Coming Soon" ──────────────
describe('LanguageSelector — Coming Soon languages disabled (Req 1.3, 1.4)', () => {
  // Accessible name includes "Coming Soon" span text, e.g. "Python Coming Soon"
  const comingSoonLanguages = [
    { label: 'Python', namePattern: /python/i },
    { label: 'Java',   namePattern: /^java coming soon$/i },
    { label: 'PHP',    namePattern: /php/i },
    { label: 'C',      namePattern: /^c coming soon$/i },
  ];

  comingSoonLanguages.forEach(({ label, namePattern }) => {
    it(`${label} button should be disabled`, () => {
      renderSelector();
      const btn = screen.getByRole('button', { name: namePattern });
      expect(btn.hasAttribute('disabled')).toBe(true);
    });

    it(`${label} button should contain "Coming Soon" text`, () => {
      renderSelector();
      const btn = screen.getByRole('button', { name: namePattern });
      expect(btn.textContent).toMatch(/coming soon/i);
    });
  });
});

// ─── 1.6: Klik JS/SQL memanggil onChange; klik Python tidak ──────────────────
describe('LanguageSelector — onClick behavior (Req 1.6)', () => {
  it('clicking JavaScript should call onChange with "javascript"', () => {
    const onChange = vi.fn();
    renderSelector('sql', onChange);
    fireEvent.click(screen.getByRole('button', { name: /javascript/i }));
    expect(onChange).toHaveBeenCalledOnce();
    expect(onChange).toHaveBeenCalledWith('javascript');
  });

  it('clicking SQL should call onChange with "sql"', () => {
    const onChange = vi.fn();
    renderSelector('javascript', onChange);
    fireEvent.click(screen.getByRole('button', { name: /^sql$/i }));
    expect(onChange).toHaveBeenCalledOnce();
    expect(onChange).toHaveBeenCalledWith('sql');
  });

  it('clicking Python should NOT call onChange', () => {
    const onChange = vi.fn();
    renderSelector('javascript', onChange);
    fireEvent.click(screen.getByRole('button', { name: /python/i }));
    expect(onChange).not.toHaveBeenCalled();
  });

  it('clicking Java should NOT call onChange', () => {
    const onChange = vi.fn();
    renderSelector('javascript', onChange);
    fireEvent.click(screen.getByRole('button', { name: /^java coming soon$/i }));
    expect(onChange).not.toHaveBeenCalled();
  });

  it('clicking PHP should NOT call onChange', () => {
    const onChange = vi.fn();
    renderSelector('javascript', onChange);
    fireEvent.click(screen.getByRole('button', { name: /php/i }));
    expect(onChange).not.toHaveBeenCalled();
  });

  it('clicking C should NOT call onChange', () => {
    const onChange = vi.fn();
    renderSelector('javascript', onChange);
    fireEvent.click(screen.getByRole('button', { name: /^c coming soon$/i }));
    expect(onChange).not.toHaveBeenCalled();
  });
});
