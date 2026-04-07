import { ProblemListSkeleton } from '@/components/Skeleton';

export default function Loading() {
  return (
    <main className="mx-auto max-w-4xl px-6 py-10">
      <div className="mb-7">
        <div className="h-7 w-32 animate-pulse rounded-md bg-gray-200" />
        <div className="mt-2 h-4 w-48 animate-pulse rounded-md bg-gray-200" />
      </div>
      <div className="mb-5 flex gap-2">
        {Array.from({ length: 5 }).map((_, i) => (
          <div key={i} className="h-8 w-20 animate-pulse rounded-lg bg-gray-200" />
        ))}
      </div>
      <ProblemListSkeleton />
    </main>
  );
}
