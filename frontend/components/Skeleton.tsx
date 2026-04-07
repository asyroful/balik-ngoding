// Reusable skeleton shimmer primitive
import React from 'react';
export function Skeleton({ className = '', style }: { className?: string; style?: React.CSSProperties }) {
  return (
    <div
      className={`animate-pulse rounded-md bg-gray-200 ${className}`}
      style={style}
    />
  );
}

// Skeleton for problem list table
export function ProblemListSkeleton() {
  return (
    <div className="overflow-hidden rounded-xl border border-gray-200 bg-white">
      {/* Table header */}
      <div className="flex gap-4 border-b border-gray-100 bg-gray-50 px-5 py-3">
        <Skeleton className="h-3 w-6" />
        <Skeleton className="h-3 w-48" />
        <Skeleton className="ml-auto h-3 w-16" />
        <Skeleton className="h-3 w-16" />
      </div>
      {/* Rows */}
      {Array.from({ length: 8 }).map((_, i) => (
        <div key={i} className="flex items-center gap-4 border-b border-gray-100 px-5 py-4 last:border-0">
          <Skeleton className="h-3 w-4" />
          <Skeleton className="h-3.5" style={{ width: `${140 + (i % 5) * 30}px` }} />
          <div className="ml-auto flex gap-3">
            <Skeleton className="h-5 w-14 rounded-md" />
            <Skeleton className="h-5 w-14 rounded-md" />
          </div>
        </div>
      ))}
    </div>
  );
}

// Skeleton for problem detail — left panel
export function ProblemDescriptionSkeleton() {
  return (
    <div className="space-y-4 px-6 py-5">
      {/* Title */}
      <Skeleton className="h-5 w-3/4" />
      {/* Body lines */}
      <div className="space-y-2 pt-1">
        <Skeleton className="h-3.5 w-full" />
        <Skeleton className="h-3.5 w-full" />
        <Skeleton className="h-3.5 w-5/6" />
        <Skeleton className="h-3.5 w-4/6" />
      </div>
      {/* Example block */}
      <div className="mt-4 space-y-2">
        <Skeleton className="h-3 w-16" />
        <div className="rounded-xl border border-gray-100 bg-gray-50 p-4 space-y-3">
          <Skeleton className="h-3 w-20" />
          <Skeleton className="h-8 w-full rounded-lg" />
          <Skeleton className="h-3 w-20" />
          <Skeleton className="h-8 w-full rounded-lg" />
        </div>
      </div>
    </div>
  );
}

// Skeleton for problem detail — right panel (editor area)
export function EditorSkeleton() {
  return (
    <div className="flex-1 bg-[#1e1e1e] flex flex-col">
      {/* Fake code lines */}
      <div className="flex-1 p-5 space-y-2.5">
        {[60, 40, 80, 30, 70, 50, 90, 45, 65, 35].map((w, i) => (
          <div key={i} className="flex items-center gap-3">
            <div className="w-6 h-3 rounded bg-white/5 flex-shrink-0" />
            <div
              className="h-3 rounded bg-white/10"
              style={{ width: `${w}%` }}
            />
          </div>
        ))}
      </div>
    </div>
  );
}
