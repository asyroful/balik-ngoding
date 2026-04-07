import { create } from 'zustand';
import { SubmissionResult } from '../lib/types';

interface SubmissionStore {
  code: string;
  isLoading: boolean;
  result: SubmissionResult | null;
  error: string | null;
  language: string;
  setCode: (code: string) => void;
  setLoading: (loading: boolean) => void;
  setResult: (result: SubmissionResult | null) => void;
  setError: (error: string | null) => void;
  setLanguage: (language: string) => void;
  reset: () => void;
}

const initialState = {
  code: '',
  isLoading: false,
  result: null,
  error: null,
  language: 'javascript',
};

export const useSubmissionStore = create<SubmissionStore>((set) => ({
  ...initialState,
  setCode: (code) => set({ code }),
  setLoading: (isLoading) => set({ isLoading }),
  setResult: (result) => set({ result }),
  setError: (error) => set({ error }),
  setLanguage: (language) => set({ language }),
  reset: () => set(initialState),
}));
