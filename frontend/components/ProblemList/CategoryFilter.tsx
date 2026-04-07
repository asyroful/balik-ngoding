'use client';

interface CategoryFilterProps {
  selected: string;
  onChange: (category: string) => void;
}

const TABS = [
  { label: 'Loop', value: 'loop' },
  { label: 'String', value: 'string' },
  { label: 'Array', value: 'array' },
  { label: 'SQL', value: 'sql' },
];

export default function CategoryFilter({ selected, onChange }: CategoryFilterProps) {
  return (
    <div className="mb-5 flex gap-2 flex-wrap">
      {TABS.map((tab) => (
        <button
          key={tab.value}
          onClick={() => onChange(tab.value)}
          className={`rounded-lg px-4 py-1.5 text-sm font-medium transition
            ${selected === tab.value
              ? 'bg-indigo-700 text-white'
              : 'border border-gray-200 bg-white text-gray-600 hover:border-gray-300 hover:text-gray-900'
            }`}
        >
          {tab.label}
        </button>
      ))}
    </div>
  );
}
