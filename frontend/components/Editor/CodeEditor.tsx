'use client';

import dynamic from 'next/dynamic';
import { useState, Component, ReactNode } from 'react';

function FallbackTextarea({
  value,
  onChange,
}: {
  value: string;
  onChange: (value: string) => void;
}) {
  return (
    <textarea
      value={value}
      onChange={(e) => onChange(e.target.value)}
      style={{
        minHeight: '400px',
        fontFamily: 'monospace',
        width: '100%',
        backgroundColor: '#1e1e1e',
        color: '#d4d4d4',
        padding: '12px',
        border: '1px solid #3c3c3c',
        borderRadius: '4px',
        resize: 'vertical',
        fontSize: '14px',
      }}
      spellCheck={false}
    />
  );
}

interface ErrorBoundaryState { hasError: boolean; }

class MonacoErrorBoundary extends Component<
  { children: ReactNode; fallback: ReactNode },
  ErrorBoundaryState
> {
  constructor(props: { children: ReactNode; fallback: ReactNode }) {
    super(props);
    this.state = { hasError: false };
  }
  static getDerivedStateFromError(): ErrorBoundaryState {
    return { hasError: true };
  }
  render() {
    if (this.state.hasError) return this.props.fallback;
    return this.props.children;
  }
}

const MonacoEditor = dynamic(() => import('@monaco-editor/react'), {
  ssr: false,
  loading: () => (
    <div style={{ minHeight: '400px', backgroundColor: '#1e1e1e', display: 'flex', alignItems: 'center', justifyContent: 'center', color: '#888', fontSize: '14px' }}>
      Memuat editor...
    </div>
  ),
});

interface CodeEditorProps {
  value: string;
  onChange: (value: string) => void;
  language?: string;
}

export default function CodeEditor({ value, onChange, language }: CodeEditorProps) {
  const [monacoFailed, setMonacoFailed] = useState(false);

  if (monacoFailed) {
    return <FallbackTextarea value={value} onChange={onChange} />;
  }

  return (
    <MonacoErrorBoundary fallback={<FallbackTextarea value={value} onChange={onChange} />}>
      <div style={{ minHeight: '400px' }}>
        <MonacoEditor
          height="400px"
          defaultLanguage={language === 'sql' ? 'sql' : 'javascript'}
          value={value}
          onChange={(v) => onChange(v ?? '')}
          theme="vs-dark"
          options={{ lineNumbers: 'on', minimap: { enabled: false }, fontSize: 14, scrollBeyondLastLine: false, automaticLayout: true }}
          onMount={(_editor, monaco) => { if (!monaco) setMonacoFailed(true); }}
        />
      </div>
    </MonacoErrorBoundary>
  );
}
