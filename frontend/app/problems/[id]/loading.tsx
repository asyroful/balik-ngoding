import { ProblemDescriptionSkeleton, EditorSkeleton, Skeleton } from '@/components/Skeleton';

export default function Loading() {
  return (
    <div className="flex overflow-hidden" style={{ height: 'calc(100vh - 53px)' }}>
      {/* Left skeleton */}
      <div className="w-[45%] overflow-hidden border-r border-gray-200 bg-white">
        <div className="flex items-center justify-between border-b border-gray-100 px-6 py-3">
          <div className="flex gap-2">
            <Skeleton className="h-5 w-16 rounded-md" />
            <Skeleton className="h-5 w-16 rounded-md" />
          </div>
          <Skeleton className="h-3 w-20" />
        </div>
        <ProblemDescriptionSkeleton />
      </div>

      {/* Right skeleton */}
      <div className="flex w-[55%] flex-col overflow-hidden bg-[#1e1e1e]">
        <div className="flex items-center justify-between border-b border-white/10 px-5 py-2.5 bg-[#252526]">
          <Skeleton className="h-3 w-20 bg-white/10" />
          <div className="flex gap-1.5">
            <span className="h-2.5 w-2.5 rounded-full bg-[#ff5f57]/40" />
            <span className="h-2.5 w-2.5 rounded-full bg-[#febc2e]/40" />
            <span className="h-2.5 w-2.5 rounded-full bg-[#28c840]/40" />
          </div>
        </div>
        <EditorSkeleton />
        <div className="border-t border-gray-200 bg-white p-4">
          <Skeleton className="h-8 w-36 rounded-lg" />
        </div>
      </div>
    </div>
  );
}
