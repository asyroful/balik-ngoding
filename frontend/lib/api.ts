import { Problem, SubmissionResult, SubmitRequest, CategorySummary, AnalyticsStats, SolutionKey, CreateSolutionKeyRequest, UpdateSolutionKeyRequest } from './types';
import { getOrCreateAnonymousId } from './anonymousId';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

async function handleResponse<T>(res: Response): Promise<T> {
  if (!res.ok) {
    const body = await res.json().catch(() => ({}));
    const message = body?.message || `Request failed with status ${res.status}`;
    throw new Error(message);
  }
  const json = await res.json();
  return json.data as T;
}

export async function getProblems(category?: string): Promise<Problem[]> {
  const url = new URL(`${BASE_URL}/problems`);
  if (category) url.searchParams.set('category', category);
  const res = await fetch(url.toString());
  return handleResponse<Problem[]>(res);
}

export async function getProblemById(id: string): Promise<Problem> {
  const res = await fetch(`${BASE_URL}/problems/${id}`);
  return handleResponse<Problem>(res);
}

export async function submitSolution(req: SubmitRequest): Promise<SubmissionResult> {
  const anonymousId = getOrCreateAnonymousId();
  const res = await fetch(`${BASE_URL}/submit`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ ...req, anonymousId }),
  });
  return handleResponse<SubmissionResult>(res);
}

export async function getProblemsSummary(): Promise<CategorySummary[]> {
  const res = await fetch(`${BASE_URL}/problems/summary`);
  return handleResponse<CategorySummary[]>(res);
}

export async function getAnalyticsStats(): Promise<AnalyticsStats> {
  const res = await fetch(`${BASE_URL}/analytics/stats`, { cache: 'no-store' });
  return handleResponse<AnalyticsStats>(res);
}

export async function getSolutionKey(problemId: string, token: string): Promise<SolutionKey | null> {
  const res = await fetch(`${BASE_URL}/api/solution-keys/${problemId}`, {
    headers: {
      'Authorization': token,
    },
  });
  return handleResponse<SolutionKey | null>(res);
}

export async function createSolutionKey(req: CreateSolutionKeyRequest, token: string): Promise<SolutionKey> {
  const res = await fetch(`${BASE_URL}/api/solution-keys`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify(req),
  });
  return handleResponse<SolutionKey>(res);
}

export async function updateSolutionKey(id: string, req: UpdateSolutionKeyRequest, token: string): Promise<SolutionKey> {
  const res = await fetch(`${BASE_URL}/api/solution-keys/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify(req),
  });
  return handleResponse<SolutionKey>(res);
}

export async function deleteSolutionKey(id: string, token: string): Promise<void> {
  const res = await fetch(`${BASE_URL}/admin/solution-keys/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': token,
    },
  });
  if (!res.ok) {
    throw new Error(`Failed to delete solution key: ${res.status}`);
  }
}

export function encodeAdminCredentials(username: string, password: string): string {
  if (typeof window === 'undefined') {
    // Server-side: use Node.js Buffer
    return 'Basic ' + Buffer.from(`${username}:${password}`).toString('base64');
  }
  // Client-side: use btoa
  return 'Basic ' + btoa(`${username}:${password}`);
}
