export interface TestCase {
  id: string;
  input: string;
  expectedOutput: string;
  isHidden: boolean;
}

export interface Problem {
  id: string;
  title: string;
  description: string;
  category: 'loop' | 'string' | 'array' | 'sql';
  difficulty: 'easy' | 'medium' | 'hard';
  starterCode: string;
  isActive: boolean;
  testCases?: TestCase[];
  createdAt: string;
  thinkingGuide: string | null;
  hints: string[] | null;
  prerequisiteId: string | null;
  prerequisiteTitle: string | null;
}

export interface TestCaseResult {
  passed: boolean;
  input: string;
  expected: string;
  actual: string;
  error?: string;
}

export interface SubmissionResult {
  status: 'accepted' | 'wrong_answer' | 'error';
  score: number;
  total: number;
  results: TestCaseResult[];
  errorMessage?: string;
}

export interface SubmitRequest {
  problemId: string;
  code: string;
  language: string;
}

export interface CategorySummary {
  category: 'loop' | 'string' | 'array' | 'sql';
  total: number;
}
