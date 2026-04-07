'use client';

interface Language {
  value: string;
  label: string;
  active: boolean;
}

interface LanguageSelectorProps {
  selected: string;
  onChange: (lang: string) => void;
  problemCategory: string;
}

const LANGUAGES: Language[] = [
  { value: 'javascript', label: 'JavaScript', active: true },
  { value: 'sql',        label: 'SQL',        active: true },
  { value: 'python',     label: 'Python',     active: false },
  { value: 'java',       label: 'Java',       active: false },
  { value: 'php',        label: 'PHP',        active: false },
  { value: 'c',          label: 'C',          active: false },
];

export function getDefaultLanguage(category: string): string {
  return category === 'sql' ? 'sql' : 'javascript';
}

export default function LanguageSelector({ selected, onChange, problemCategory: _ }: LanguageSelectorProps) {
  return (
    <div className="flex flex-wrap gap-2 mb-3">
      {LANGUAGES.filter(lang => lang.active).map((lang) => {
        const isSelected = selected === lang.value;

        return (
          <button
            key={lang.value}
            onClick={() => onChange(lang.value)}
            className={`px-3 py-1.5 rounded text-sm font-medium border transition-colors ${
              isSelected
                ? 'bg-blue-600 text-white border-blue-500'
                : 'bg-gray-800 text-gray-300 border-gray-700 hover:bg-gray-700'
            }`}
          >
            {lang.label}
          </button>
        );
      })}
    </div>
  );
}
