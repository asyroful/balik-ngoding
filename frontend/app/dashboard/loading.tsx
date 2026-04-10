export default function DashboardLoading() {
  return (
    <main className="mx-auto max-w-4xl px-4 py-10">
      <div className="mb-8 h-8 w-48 animate-pulse rounded bg-gray-700" />
      <div className="mb-10 grid grid-cols-2 gap-4 sm:grid-cols-4">
        {Array.from({ length: 4 }).map((_, i) => (
          <div key={i} className="h-24 animate-pulse rounded-lg bg-gray-700" />
        ))}
      </div>
      <div className="h-48 animate-pulse rounded-lg bg-gray-700" />
    </main>
  );
}
