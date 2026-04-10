import { getAnalyticsStats } from '@/lib/api';
import { StatsCard } from '@/components/Analytics/StatsCard';
import { TopProblemsTable } from '@/components/Analytics/TopProblemsTable';

export default async function DashboardPage() {
  let stats;
  try {
    stats = await getAnalyticsStats();
  } catch {
    return (
      <main className="mx-auto max-w-4xl px-4 py-10">
        <h1 className="mb-6 text-2xl font-bold text-black">Dashboard</h1>
        <p className="text-red-400">Gagal memuat statistik. Coba lagi nanti.</p>
      </main>
    );
  }

  return (
    <main className="mx-auto max-w-4xl px-4 py-10">
      <h1 className="mb-6 text-2xl font-bold text-black">Dashboard</h1>

      <div className="mb-10 grid grid-cols-2 gap-4 sm:grid-cols-4">
        <StatsCard label="Perangkat Unik" value={stats.uniqueDevices} />
        <StatsCard label="Total Submission" value={stats.totalSubmissions} />
        <StatsCard label="Diterima" value={stats.totalAccepted} />
        <StatsCard label="Perangkat Diterima" value={stats.devicesWithAccepted} />
      </div>

      <section>
        <h2 className="mb-4 text-lg font-semibold text-black">Soal Terpopuler</h2>
        <TopProblemsTable problems={stats.topProblems} />
      </section>
    </main>
  );
}
