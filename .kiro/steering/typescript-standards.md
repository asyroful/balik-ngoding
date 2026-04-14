# TypeScript Coding Standards — Balik Ngoding

## Naming Conventions

### Components
- Use PascalCase for React components
- Use descriptive names that indicate purpose
- Examples: `CodeEditor`, `ResultPanel`, `ProblemTable`

### Functions & Variables
- Use camelCase for functions and variables
- Use descriptive names that indicate action/purpose
- Examples: `handleSubmit`, `fetchProblems`, `isLoading`

### Constants
- Use UPPER_SNAKE_CASE for constants
- Use camelCase for configuration objects
- Examples: `MAX_CODE_LENGTH`, `apiConfig`

### Types & Interfaces
- Use PascalCase for type names
- Use descriptive names that indicate data structure
- Prefix interfaces with `I` if needed for clarity
- Examples: `Problem`, `Submission`, `SubmissionResult`

## Type Safety

### Type Annotations
- Always annotate function parameters
- Always annotate function return types
- Avoid using `any` type
- Use `unknown` when type is truly unknown

```typescript
// Good
function handleSubmit(code: string, problemId: string): Promise<SubmissionResult> {
  // Implementation
}

// Bad
function handleSubmit(code, problemId) {
  // Missing types
}
```

### Null/Undefined Handling
- Use optional chaining (`?.`) untuk safe property access
- Use nullish coalescing (`??`) untuk default values
- Use type guards untuk narrowing types
- Avoid non-null assertions (`!`) unless absolutely necessary

```typescript
// Good
const title = problem?.title ?? "Untitled";

// Bad
const title = problem!.title;
```

## Component Patterns

### Functional Components
- Use arrow functions untuk components
- Use `'use client'` directive only when needed
- Keep components focused dan single-responsibility
- Extract logic into custom hooks

```typescript
// Good
'use client';

import { useState } from 'react';

export function CodeEditor({ problemId }: { problemId: string }) {
  const [code, setCode] = useState('');
  
  return (
    <div>
      {/* Component JSX */}
    </div>
  );
}

// Bad
export default function CodeEditor(props) {
  // Missing type annotations
}
```

### Props Typing
- Define props interface explicitly
- Use destructuring untuk props
- Document props dengan JSDoc comments
- Use optional properties untuk optional props

```typescript
// Good
interface CodeEditorProps {
  /** The problem ID to edit code for */
  problemId: string;
  /** Initial code to display */
  initialCode?: string;
  /** Callback when code changes */
  onChange?: (code: string) => void;
}

export function CodeEditor({ problemId, initialCode = '', onChange }: CodeEditorProps) {
  // Implementation
}
```

## State Management

### Zustand Store
- Define store interface explicitly
- Use descriptive action names
- Keep store focused pada single domain
- Document store actions

```typescript
// Good
interface SubmissionStore {
  code: string;
  status: 'idle' | 'loading' | 'success' | 'error';
  error: string | null;
  setCode: (code: string) => void;
  submit: (problemId: string) => Promise<void>;
  reset: () => void;
}

export const useSubmissionStore = create<SubmissionStore>((set) => ({
  code: '',
  status: 'idle',
  error: null,
  setCode: (code) => set({ code }),
  submit: async (problemId) => {
    // Implementation
  },
  reset: () => set({ code: '', status: 'idle', error: null }),
}));
```

### React Hooks
- Use custom hooks untuk reusable logic
- Name hooks dengan `use` prefix
- Document hook parameters dan return values
- Handle cleanup dalam useEffect

```typescript
// Good
function useProgress() {
  const [progress, setProgress] = useState(0);
  
  useEffect(() => {
    // Setup
    return () => {
      // Cleanup
    };
  }, []);
  
  return progress;
}
```

## API Integration

### API Client
- Define request/response types explicitly
- Use async/await untuk API calls
- Handle errors properly
- Add request/response logging untuk debugging

```typescript
// Good
interface SubmitRequest {
  problemId: string;
  code: string;
  language: string;
  anonymousId?: string;
}

interface SubmitResponse {
  data: SubmissionResult;
}

export async function submitCode(req: SubmitRequest): Promise<SubmissionResult> {
  try {
    const response = await fetch(`${API_URL}/submit`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    });
    
    if (!response.ok) {
      throw new Error(`HTTP ${response.status}`);
    }
    
    const data: SubmitResponse = await response.json();
    return data.data;
  } catch (error) {
    console.error('Submit failed:', error);
    throw error;
  }
}
```

## Testing Patterns

### Unit Tests
- Test file naming: `{component}.test.tsx` atau `{function}.test.ts`
- Test function naming: `test('description', () => {})`
- Use descriptive test names
- Test both success dan failure cases

```typescript
// Good
describe('CodeEditor', () => {
  test('renders code input field', () => {
    render(<CodeEditor problemId="1" />);
    expect(screen.getByRole('textbox')).toBeInTheDocument();
  });
  
  test('calls onChange when code changes', async () => {
    const onChange = vi.fn();
    render(<CodeEditor problemId="1" onChange={onChange} />);
    
    const input = screen.getByRole('textbox');
    await userEvent.type(input, 'console.log("test")');
    
    expect(onChange).toHaveBeenCalled();
  });
});
```

### Property-Based Tests
- Use `fast-check` untuk property-based testing
- Test invariants dan properties
- Generate random test cases

```typescript
// Good
import fc from 'fast-check';

test('normalization is idempotent', () => {
  fc.assert(
    fc.property(fc.string(), (output) => {
      const normalized1 = normalizeOutput(output);
      const normalized2 = normalizeOutput(normalized1);
      return normalized1 === normalized2;
    })
  );
});
```

## Error Handling

### Error Types
- Define custom error types untuk different scenarios
- Use error boundaries untuk component errors
- Log errors dengan context
- Show user-friendly error messages

```typescript
// Good
class APIError extends Error {
  constructor(
    public statusCode: number,
    public message: string,
  ) {
    super(message);
  }
}

try {
  await submitCode(req);
} catch (error) {
  if (error instanceof APIError) {
    console.error(`API error ${error.statusCode}: ${error.message}`);
  } else {
    console.error('Unknown error:', error);
  }
}
```

## Performance Optimization

### Memoization
- Use `React.memo` untuk expensive components
- Use `useMemo` untuk expensive computations
- Use `useCallback` untuk stable function references
- Profile components dengan React DevTools

```typescript
// Good
const CodeEditor = React.memo(function CodeEditor({ code, onChange }: Props) {
  const handleChange = useCallback((newCode: string) => {
    onChange(newCode);
  }, [onChange]);
  
  return <Editor value={code} onChange={handleChange} />;
});
```

### Code Splitting
- Use dynamic imports untuk large components
- Use `next/dynamic` dengan `ssr: false` untuk client-only components
- Lazy load routes dengan Next.js App Router

```typescript
// Good
import dynamic from 'next/dynamic';

const CodeEditor = dynamic(() => import('./CodeEditor'), {
  ssr: false,
  loading: () => <Skeleton />,
});
```

## Accessibility

### ARIA Attributes
- Add `aria-label` untuk icon buttons
- Add `aria-describedby` untuk form fields
- Add `role` attributes untuk custom components
- Test dengan screen readers

```typescript
// Good
<button aria-label="Submit code">
  <SubmitIcon />
</button>

<input
  type="text"
  aria-label="Code input"
  aria-describedby="code-help"
/>
```

### Keyboard Navigation
- Ensure all interactive elements are keyboard accessible
- Use semantic HTML elements
- Test dengan keyboard navigation
- Provide focus indicators

## Documentation

### JSDoc Comments
- Document exported functions dan components
- Include parameter descriptions
- Include return value descriptions
- Include usage examples

```typescript
// Good
/**
 * Submits user code for evaluation
 * @param req - The submission request containing code and problem ID
 * @returns Promise resolving to submission result
 * @throws APIError if submission fails
 * @example
 * const result = await submitCode({ problemId: '1', code: 'console.log("hi")' });
 */
export async function submitCode(req: SubmitRequest): Promise<SubmissionResult> {
  // Implementation
}
```

### Inline Comments
- Write comments untuk "why", not "what"
- Use clear, concise language
- Update comments when code changes

```typescript
// Good - explains why
// Normalize output before comparison because users might have different whitespace
const normalized = normalizeOutput(output);

// Bad - explains what (obvious from code)
// Normalize the output
```

