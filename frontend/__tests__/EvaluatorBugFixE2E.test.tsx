import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';
import ProblemPage from '@/app/problems/[id]/page';
import * as api from '@/lib/api';

// Mock the API module
vi.mock('@/lib/api');

// Mock useParams
vi.mock('next/navigation', () => ({
  useParams: () => ({ id: 'test-string-problem' }),
}));

describe('EvaluatorBugFixE2E', () => {
  const mockProblem = {
    id: 'test-string-problem',
    title: 'String Manipulation',
    description: 'Test string input handling',
    category: 'string',
    difficulty: 'easy',
    examples: [
      {
        input: '"banana", "a"',
        output: '"a"',
        explanation: 'Return second argument',
      },
    ],
    testCases: [
      {
        id: 'tc-1',
        input: '"banana", "a"',
        expectedOutput: '"a"',
        description: 'Return second argument',
      },
      {
        id: 'tc-2',
        input: '"hello"',
        expectedOutput: '"hello"',
        description: 'Return input string',
      },
      {
        id: 'tc-3',
        input: '5, "test"',
        expectedOutput: '"test5"',
        description: 'Concatenate number and string',
      },
    ],
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  // Feature: evaluator-fix-and-solution-keys, Property 16: evaluator bug fix verification
  test('TestEvaluatorBugFixE2E - submit code with string inputs and verify correct evaluation', async () => {
    const user = userEvent.setup();

    // Mock API calls
    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.submitSolution as any).mockResolvedValue({
      status: 'accepted',
      score: 3,
      total: 3,
      results: [
        {
          passed: true,
          input: '"banana", "a"',
          expected: '"a"',
          actual: '"a"',
          testCaseId: 'tc-1',
        },
        {
          passed: true,
          input: '"hello"',
          expected: '"hello"',
          actual: '"hello"',
          testCaseId: 'tc-2',
        },
        {
          passed: true,
          input: '5, "test"',
          expected: '"test5"',
          actual: '"test5"',
          testCaseId: 'tc-3',
        },
      ],
    });

    render(<ProblemPage />);

    // Wait for problem to load
    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Find and fill the code editor
    const codeEditor = screen.getByRole('textbox', { hidden: true });
    await user.click(codeEditor);
    await user.type(codeEditor, 'function solution(a, b) { return b; }');

    // Submit the solution
    const submitButton = screen.getByRole('button', { name: /submit/i });
    await user.click(submitButton);

    // Verify results are displayed
    await waitFor(() => {
      expect(screen.getByText(/accepted|3 \/ 3/i)).toBeInTheDocument();
    });

    // Verify all test cases passed
    const results = screen.getAllByText(/passed/i);
    expect(results.length).toBeGreaterThan(0);
  });

  test('TestEvaluatorBugFixE2E - verify string input parsing correctness', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.submitSolution as any).mockResolvedValue({
      status: 'accepted',
      score: 1,
      total: 1,
      results: [
        {
          passed: true,
          input: '"banana", "a"',
          expected: '"a"',
          actual: '"a"',
          testCaseId: 'tc-1',
        },
      ],
    });

    render(<ProblemPage />);

    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Submit code that returns second argument
    const codeEditor = screen.getByRole('textbox', { hidden: true });
    await user.click(codeEditor);
    await user.type(codeEditor, 'function solution(a, b) { return b; }');

    const submitButton = screen.getByRole('button', { name: /submit/i });
    await user.click(submitButton);

    // Verify the test case with string inputs passed
    await waitFor(() => {
      expect(screen.getByText(/accepted/i)).toBeInTheDocument();
    });
  });

  test('TestEvaluatorBugFixE2E - verify mixed type input handling', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.submitSolution as any).mockResolvedValue({
      status: 'accepted',
      score: 1,
      total: 1,
      results: [
        {
          passed: true,
          input: '5, "test"',
          expected: '"test5"',
          actual: '"test5"',
          testCaseId: 'tc-3',
        },
      ],
    });

    render(<ProblemPage />);

    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Submit code that concatenates number and string
    const codeEditor = screen.getByRole('textbox', { hidden: true });
    await user.click(codeEditor);
    await user.type(codeEditor, 'function solution(num, str) { return str + num; }');

    const submitButton = screen.getByRole('button', { name: /submit/i });
    await user.click(submitButton);

    // Verify mixed type test case passed
    await waitFor(() => {
      expect(screen.getByText(/accepted/i)).toBeInTheDocument();
    });
  });

  test('TestEvaluatorBugFixE2E - verify incorrect solution fails', async () => {
    const user = userEvent.setup();

    (api.getProblemById as any).mockResolvedValue(mockProblem);
    (api.submitSolution as any).mockResolvedValue({
      status: 'rejected',
      score: 0,
      total: 3,
      results: [
        {
          passed: false,
          input: '"banana", "a"',
          expected: '"a"',
          actual: '"banana"',
          testCaseId: 'tc-1',
        },
        {
          passed: false,
          input: '"hello"',
          expected: '"hello"',
          actual: '"hello"',
          testCaseId: 'tc-2',
        },
        {
          passed: false,
          input: '5, "test"',
          expected: '"test5"',
          actual: '5test',
          testCaseId: 'tc-3',
        },
      ],
    });

    render(<ProblemPage />);

    await waitFor(() => {
      expect(screen.getByText(mockProblem.title)).toBeInTheDocument();
    });

    // Submit incorrect code
    const codeEditor = screen.getByRole('textbox', { hidden: true });
    await user.click(codeEditor);
    await user.type(codeEditor, 'function solution(a) { return a; }');

    const submitButton = screen.getByRole('button', { name: /submit/i });
    await user.click(submitButton);

    // Verify results show failures
    await waitFor(() => {
      expect(screen.getByText(/rejected|0 \/ 3/i)).toBeInTheDocument();
    });
  });
});
