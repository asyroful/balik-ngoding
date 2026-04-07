import NavLink from '@/components/NavLink';

const stats = [
  { value: '100+', label: 'Soal Tersedia', sub: 'Terus bertambah' },
  { value: '4', label: 'Kategori', sub: 'Loop, String, Array, SQL' },
  { value: '100%', label: 'Gratis', sub: 'Tanpa daftar' },
]

const features = [
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
      </svg>
    ),
    title: 'Editor Langsung di Browser',
    desc: 'Tulis kode tanpa install apapun. Monaco Editor dengan syntax highlighting JavaScript.',
  },
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
      </svg>
    ),
    title: 'Feedback Per Test Case',
    desc: 'Lihat input, expected output, dan actual output untuk setiap test case secara detail.',
  },
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
      </svg>
    ),
    title: 'Berbagai Topik',
    desc: 'Dari loop sederhana sampai manipulasi array — latih semua aspek logika pemrograman.',
  },
]

const categories = [
  { label: 'Loop', color: 'bg-blue-100 text-blue-700', desc: 'Iterasi & perulangan' },
  { label: 'String', color: 'bg-purple-100 text-purple-700', desc: 'Manipulasi teks' },
  { label: 'Array', color: 'bg-orange-100 text-orange-700', desc: 'Struktur data dasar' },
  { label: 'SQL', color: 'bg-teal-100 text-teal-700', desc: 'Query database' },
]

export default function Home() {
  return (
    <div className="min-h-screen bg-[#f8f7f4]">

      {/* Hero */}
      <section className="relative overflow-hidden">
        {/* Background decoration */}
        <div className="absolute inset-0 pointer-events-none">
          <div className="absolute -top-24 -right-24 w-96 h-96 rounded-full bg-indigo-100/60 blur-3xl" />
          <div className="absolute top-32 -left-16 w-64 h-64 rounded-full bg-violet-100/40 blur-3xl" />
          <div className="absolute bottom-0 right-1/3 w-72 h-72 rounded-full bg-indigo-50/80 blur-3xl" />
        </div>

        <div className="relative mx-auto max-w-6xl px-6 pt-20 pb-20">
          <div className="grid lg:grid-cols-2 gap-12 items-center">
            {/* Left: copy */}
            <div>
              <span className="inline-flex items-center gap-1.5 rounded-full border border-indigo-200 bg-indigo-50 px-3 py-1 text-xs font-medium text-indigo-700 mb-6">
                <span className="w-1.5 h-1.5 rounded-full bg-indigo-500 animate-pulse" />
                Platform Latihan Coding · Gratis
              </span>
              <h1 className="text-5xl sm:text-6xl font-extrabold leading-[1.08] tracking-tight text-gray-900 mb-5">
                Asah logika,<br />
                <span className="text-indigo-600">balik ngoding.</span>
              </h1>
              <p className="text-lg text-gray-500 leading-relaxed mb-8 max-w-lg">
                Latih logika dasar dengan cara yang simpel.
              </p>
              <div className="flex items-center gap-3 flex-wrap">
                <NavLink
                  href="/problems"
                  className="inline-flex items-center gap-2 rounded-xl bg-indigo-600 px-7 py-3 text-sm font-semibold text-white shadow-md shadow-indigo-200 transition hover:bg-indigo-700 active:scale-[0.98]"
                >
                  Mulai Latihan
                  <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
                  </svg>
                </NavLink>
                <NavLink
                  href="/problems"
                  className="inline-flex items-center gap-2 rounded-xl border border-gray-300 bg-white px-7 py-3 text-sm font-semibold text-gray-700 transition hover:border-gray-400 hover:bg-gray-50"
                >
                  Lihat Soal
                </NavLink>
              </div>
            </div>

            {/* Right: code mockup */}
            <div className="hidden lg:block">
              <div className="relative">
                {/* Editor window */}
                <div className="rounded-2xl border border-gray-200 bg-gray-900 shadow-2xl shadow-indigo-100/50 overflow-hidden">
                  {/* Title bar */}
                  <div className="flex items-center gap-1.5 px-4 py-3 bg-gray-800 border-b border-gray-700">
                    <span className="w-3 h-3 rounded-full bg-red-400" />
                    <span className="w-3 h-3 rounded-full bg-yellow-400" />
                    <span className="w-3 h-3 rounded-full bg-green-400" />
                    <span className="ml-3 text-xs text-gray-400 font-mono">solution.js</span>
                  </div>
                  {/* Code */}
                  <div className="px-5 py-5 font-mono text-sm leading-relaxed">
                    <div className="flex gap-4">
                      <div className="text-gray-600 select-none text-right" style={{minWidth:'1.5rem'}}>
                        {[1,2,3,4,5,6,7,8].map(n => <div key={n}>{n}</div>)}
                      </div>
                      <div>
                        <div><span className="text-purple-400">function</span> <span className="text-yellow-300">sumArray</span><span className="text-gray-300">(arr) {'{'}</span></div>
                        <div className="pl-4"><span className="text-purple-400">let</span> <span className="text-blue-300">total</span> <span className="text-gray-300">= </span><span className="text-orange-300">0</span><span className="text-gray-300">;</span></div>
                        <div className="pl-4"><span className="text-purple-400">for</span> <span className="text-gray-300">(</span><span className="text-purple-400">const</span> <span className="text-blue-300">n</span> <span className="text-purple-400">of</span> <span className="text-blue-300">arr</span><span className="text-gray-300">) {'{'}</span></div>
                        <div className="pl-8"><span className="text-blue-300">total</span> <span className="text-gray-300">+= </span><span className="text-blue-300">n</span><span className="text-gray-300">;</span></div>
                        <div className="pl-4"><span className="text-gray-300">{'}'}</span></div>
                        <div className="pl-4"><span className="text-purple-400">return</span> <span className="text-blue-300">total</span><span className="text-gray-300">;</span></div>
                        <div><span className="text-gray-300">{'}'}</span></div>
                        <div className="mt-1 text-gray-600">{'// ✓ 3/3 test cases passed'}</div>
                      </div>
                    </div>
                  </div>
                </div>

                {/* Result badge floating */}
                <div className="absolute -bottom-4 -right-4 flex items-center gap-2 rounded-xl bg-white border border-green-200 shadow-lg px-4 py-2.5">
                  <span className="flex items-center justify-center w-6 h-6 rounded-full bg-green-100">
                    <svg className="w-3.5 h-3.5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2.5} d="M5 13l4 4L19 7" />
                    </svg>
                  </span>
                  <div>
                    <p className="text-xs font-semibold text-gray-800">Semua Test Lulus</p>
                    <p className="text-[10px] text-gray-400">3 / 3 passed</p>
                  </div>
                </div>

                {/* Floating tag */}
                <div className="absolute -top-4 -left-4 flex items-center gap-1.5 rounded-xl bg-indigo-600 shadow-lg px-3 py-2">
                  <svg className="w-3.5 h-3.5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                  <span className="text-xs font-semibold text-white">Langsung di browser</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Stats */}
      <section className="border-y border-gray-200 bg-white">
        <div className="mx-auto max-w-6xl px-6 py-10">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-6">
            <div className="grid grid-cols-3 gap-8">
              {stats.map((s) => (
                <div key={s.label} className="group">
                  <p className="text-3xl font-extrabold text-indigo-600">{s.value}</p>
                  <p className="mt-0.5 text-sm font-semibold text-gray-800">{s.label}</p>
                  <p className="text-xs text-gray-400">{s.sub}</p>
                </div>
              ))}
            </div>
            <div className="hidden sm:flex items-center gap-2 rounded-xl border border-indigo-100 bg-indigo-50 px-5 py-3">
              <svg className="w-4 h-4 text-indigo-500 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <p className="text-sm text-indigo-700">Mulai latihan <span className="font-semibold">langsung</span> tanpa daftar</p>
            </div>
          </div>
        </div>
      </section>

      {/* Categories */}
      <section className="mx-auto max-w-6xl px-6 py-16">
        <p className="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-2">Topik Latihan</p>
        <h2 className="text-2xl font-bold text-gray-900 mb-8">Pilih sesuai kebutuhanmu.</h2>
        <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
          {categories.map((c) => (
            <NavLink
              key={c.label}
              href="/problems"
              className="group flex flex-col gap-2 rounded-2xl border border-gray-200 bg-white p-5 transition hover:border-indigo-200 hover:shadow-md hover:-translate-y-0.5"
            >
              <span className={`self-start rounded-lg px-2.5 py-1 text-xs font-semibold ${c.color}`}>{c.label}</span>
              <span className="text-sm text-gray-500 group-hover:text-gray-700 transition">{c.desc}</span>
            </NavLink>
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="bg-white border-y border-gray-200">
        <div className="mx-auto max-w-6xl px-6 py-16">
          <p className="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-2">Kenapa Balik Ngoding?</p>
          <h2 className="text-2xl font-bold text-gray-900 mb-10 max-w-sm leading-snug">
            Semua yang kamu butuhkan untuk latihan.
          </h2>
          <div className="grid gap-5 sm:grid-cols-3">
            {features.map((f) => (
              <div
                key={f.title}
                className="group rounded-2xl border border-gray-200 bg-[#f8f7f4] p-6 transition hover:border-indigo-200 hover:bg-indigo-50/30"
              >
                <span className="inline-flex items-center justify-center w-9 h-9 rounded-xl bg-indigo-100 text-indigo-600 mb-4 group-hover:bg-indigo-200 transition">
                  {f.icon}
                </span>
                <h3 className="font-semibold text-gray-900 mb-1.5 text-sm">{f.title}</h3>
                <p className="text-sm text-gray-500 leading-relaxed">{f.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA bottom */}
      <section className="mx-auto max-w-6xl px-6 py-20">
        <div className="relative overflow-hidden rounded-3xl bg-indigo-600 px-8 py-12 sm:px-12 text-center">
          <div className="absolute -top-12 -right-12 w-48 h-48 rounded-full bg-white/10 blur-2xl pointer-events-none" />
          <div className="absolute -bottom-8 -left-8 w-36 h-36 rounded-full bg-white/10 blur-2xl pointer-events-none" />
          <div className="relative">
            <h3 className="text-2xl sm:text-3xl font-extrabold text-white mb-2">Siap mulai latihan?</h3>
            <p className="text-indigo-200 mb-7 text-sm">Gratis, tanpa daftar, langsung koding.</p>
            <NavLink
              href="/problems"
              className="inline-flex items-center gap-2 rounded-xl bg-white px-8 py-3 text-sm font-semibold text-indigo-700 shadow-lg transition hover:bg-indigo-50 active:scale-[0.98]"
            >
              Mulai Sekarang
              <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
              </svg>
            </NavLink>
          </div>
        </div>
      </section>

    </div>
  );
}
