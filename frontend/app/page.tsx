import NavLink from '@/components/NavLink';

const stats = [
  { value: '80', label: 'Soal Tersedia', sub: 'Terus bertambah' },
  { value: '4', label: 'Kategori', sub: 'Loop, String, Array, SQL' },
  { value: '3x', label: 'Hint Per Soal', sub: 'Makin spesifik' },
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
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
      </svg>
    ),
    title: 'Thinking Guide',
    desc: 'Panduan cara berpikir yang bisa dibuka opsional — membantu memahami pendekatan tanpa langsung spoiler solusi.',
  },
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 4h13M3 8h9m-9 4h6m4 0l4-4m0 0l4 4m-4-4v12" />
      </svg>
    ),
    title: 'Hint Bertingkat',
    desc: 'Tiga level hint dari umum ke spesifik. Buka satu per satu sesuai seberapa banyak bantuan yang kamu butuhkan.',
  },
  {
    icon: (
      <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
      </svg>
    ),
    title: 'Jalur Belajar Terstruktur',
    desc: 'Soal disusun dengan prerequisite chain — dari warm-up syntax dasar sampai soal yang lebih kompleks.',
  },
]

const categories = [
  { label: 'Loop', color: 'bg-blue-100 text-blue-700', desc: 'Iterasi & perulangan' },
  { label: 'String', color: 'bg-purple-100 text-purple-700', desc: 'Manipulasi teks' },
  { label: 'Array', color: 'bg-orange-100 text-orange-700', desc: 'Struktur data dasar' },
  { label: 'SQL', color: 'bg-teal-100 text-teal-700', desc: 'Query database' },
]

const languages = [
  {
    label: 'JavaScript',
    color: 'bg-yellow-100 text-yellow-700',
    available: true,
    icon: (
      <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
        <rect width="32" height="32" rx="4" fill="#F7DF1E"/>
        <path d="M9 25.3l2.3-1.4c.4.8.8 1.4 1.7 1.4.9 0 1.4-.3 1.4-1.7V14h2.8v9.7c0 2.8-1.6 4-4 4-2.1 0-3.4-1.1-4.2-2.4zM19.3 25l2.3-1.3c.6 1 1.3 1.7 2.7 1.7 1.1 0 1.8-.6 1.8-1.3 0-.9-.7-1.2-1.9-1.8l-.7-.3c-1.9-.8-3.2-1.8-3.2-3.9 0-1.9 1.5-3.4 3.8-3.4 1.6 0 2.8.6 3.6 2l-2.2 1.4c-.4-.8-1-1.1-1.7-1.1-.8 0-1.3.5-1.3 1.1 0 .8.5 1.1 1.6 1.6l.7.3c2.3 1 3.5 2 3.5 4.1 0 2.4-1.8 3.6-4.3 3.6-2.4 0-3.9-1.1-4.7-2.7z" fill="#000"/>
      </svg>
    ),
  },
  {
    label: 'SQL',
    color: 'bg-teal-100 text-teal-700',
    available: true,
    icon: (
      <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
        <rect width="32" height="32" rx="4" fill="#0f766e"/>
        <ellipse cx="16" cy="10" rx="8" ry="3.5" fill="#99f6e4"/>
        <path d="M8 10v4c0 1.9 3.6 3.5 8 3.5s8-1.6 8-3.5v-4" stroke="#99f6e4" strokeWidth="1.2"/>
        <path d="M8 14v4c0 1.9 3.6 3.5 8 3.5s8-1.6 8-3.5v-4" stroke="#99f6e4" strokeWidth="1.2"/>
        <path d="M8 18v4c0 1.9 3.6 3.5 8 3.5s8-1.6 8-3.5v-4" stroke="#99f6e4" strokeWidth="1.2"/>
      </svg>
    ),
  },
  {
    label: 'Python',
    color: 'bg-blue-100 text-blue-700',
    available: false,
    icon: (
      <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
        <rect width="32" height="32" rx="4" fill="#e2e8f0"/>
        <path d="M16 6c-4.4 0-4.1 1.9-4.1 1.9v2h4.2v.6H9.8S7 10.2 7 14.6c0 4.4 2.4 4.2 2.4 4.2h1.5v-2s-.1-2.4 2.4-2.4h4.1s2.3.04 2.3-2.2V8.1S20.1 6 16 6zm-2.3 1.3c.4 0 .8.3.8.8s-.3.8-.8.8-.8-.3-.8-.8.4-.8.8-.8z" fill="#3b82f6"/>
        <path d="M16 26c4.4 0 4.1-1.9 4.1-1.9v-2h-4.2v-.6h6.3s2.8.3 2.8-4.1c0-4.4-2.4-4.2-2.4-4.2h-1.5v2s.1 2.4-2.4 2.4h-4.1s-2.3-.04-2.3 2.2v3.9S11.9 26 16 26zm2.3-1.3c-.4 0-.8-.3-.8-.8s.3-.8.8-.8.8.3.8.8-.4.8-.8.8z" fill="#60a5fa"/>
      </svg>
    ),
  },
  {
    label: 'Java',
    color: 'bg-orange-100 text-orange-700',
    available: false,
    icon: (
      <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
        <rect width="32" height="32" rx="4" fill="#e2e8f0"/>
        <path d="M13 21.5s-.9.5.6.7c1.8.2 2.8.2 4.8-.2 0 0 .5.3 1.3.6-4.7 2-10.6-.1-6.7-1.1z" fill="#ea580c"/>
        <path d="M12.4 19s-1 .7.5.9c2 .2 3.5.2 6.2-.3 0 0 .4.4 1 .6-5.5 1.6-11.6.1-7.7-1.2z" fill="#ea580c"/>
        <path d="M17.3 14.5c1.1 1.3-.3 2.4-.3 2.4s2.8-1.4 1.5-3.2c-1.2-1.6-2.1-2.4 2.9-5.2 0 0-7.9 2-4.1 6z" fill="#ea580c"/>
        <path d="M22.5 23.2s.7.5-.7.9c-2.6.8-10.8 1-13.1 0-.8-.4.7-.9.7-.9s-.9-.6-2 .1c-1.4.9 7.8 1.9 13.3.7 1.3-.3 1.8-1.1 1.8-1.1l-2-.7z" fill="#ea580c"/>
        <path d="M13.6 16.8s-4.5 1.1-1.6 1.5c1.2.2 3.6.1 5.8-.1 1.8-.2 3.6-.5 3.6-.5s-.6.3-1 .5c-4.1 1.1-12 .6-9.7-.5 1.9-.9 2.9-.9 2.9-.9z" fill="#ea580c"/>
      </svg>
    ),
  },
  {
    label: 'PHP',
    color: 'bg-indigo-100 text-indigo-700',
    available: false,
    icon: (
      <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
        <rect width="32" height="32" rx="4" fill="#e2e8f0"/>
        <ellipse cx="16" cy="16" rx="12" ry="7" fill="#6366f1"/>
        <text x="16" y="20" textAnchor="middle" fontSize="8" fontWeight="bold" fill="white" fontFamily="monospace">php</text>
      </svg>
    ),
  }
  // {
  //   label: 'C',
  //   color: 'bg-gray-100 text-gray-600',
  //   available: false,
  //   icon: (
  //     <svg viewBox="0 0 32 32" className="w-10 h-10" fill="none">
  //       <rect width="32" height="32" rx="4" fill="#e2e8f0"/>
  //       <path d="M22 12.5A8 8 0 1 0 22 19.5" stroke="#4b5563" strokeWidth="2.5" strokeLinecap="round"/>
  //     </svg>
  //   ),
  // },
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
                Balik Ngoding<br />
                <span className="text-indigo-600">Belajar sampai beneran paham.</span>
              </h1>
              <p className="text-lg text-gray-500 leading-relaxed mb-8 max-w-lg">
                Bukan sekadar latihan soal ada panduan cara berpikir, hint bertahap, dan jalur belajar yang terstruktur.
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

            {/* Right: guided learning mockup */}
            <div className="hidden lg:block">
              <div className="relative">
                <div className="rounded-2xl border border-gray-200 bg-white shadow-2xl shadow-indigo-100/50 overflow-hidden">
                  {/* ThinkingGuide panel - open */}
                  <div className="border-b border-indigo-100 bg-indigo-50">
                    <div className="flex items-center justify-between px-4 py-3">
                      <span className="flex items-center gap-2 text-sm font-semibold text-indigo-700">
                        <span>💡</span>
                        <span>Cara Berpikir</span>
                      </span>
                      <span className="text-xs text-indigo-400">▲ Tutup</span>
                    </div>
                    <div className="px-4 pb-3 border-t border-indigo-100">
                      <p className="text-sm text-indigo-900 whitespace-pre-wrap">Mulai dengan memikirkan kondisi dasar:{'\n'}apa yang terjadi jika array kosong?{'\n'}Lalu iterasi satu per satu...</p>
                    </div>
                  </div>
                  {/* HintPanel */}
                  <div className="bg-amber-50">
                    <div className="flex items-center justify-between px-4 py-3">
                      <span className="flex items-center gap-2 text-sm font-semibold text-amber-700">
                        <span>🔍</span>
                        <span>Hint</span>
                      </span>
                      <span className="text-xs text-amber-500 font-medium">Hint 1 dari 3</span>
                    </div>
                    <div className="px-4 pb-3 border-t border-amber-100">
                      <p className="text-sm text-amber-900">✓ Coba gunakan for...of untuk iterasi</p>
                    </div>
                    <div className="px-4 pb-3">
                      <div className="w-full rounded-lg bg-amber-400 text-amber-900 text-sm font-semibold py-2 text-center">
                        Buka Hint Berikutnya →
                      </div>
                    </div>
                  </div>
                </div>
                {/* Result badge floating */}
                <div className="absolute -bottom-6 -right-6 flex items-center gap-2 rounded-xl bg-white border border-green-200 shadow-lg px-4 py-2.5">
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
                <div className="absolute -top-5 -left-6 flex items-center gap-1.5 rounded-xl bg-indigo-600 shadow-lg px-3 py-2">
                  <svg className="w-3.5 h-3.5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                  <span className="text-xs font-semibold text-white">⚡ Panduan lengkap</span>
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
            <div className="grid grid-cols-4 gap-8">
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

      {/* Cara Kerja */}
      <section className="mx-auto max-w-6xl px-6 py-16">
        <p className="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-2">Cara Kerja</p>
        <h2 className="text-2xl font-bold text-gray-900 mb-8">4 langkah menuju pemahaman nyata.</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
          {[
            { number: '①', title: 'Mulai dari Warm-up', desc: 'Soal dasar sebagai entry point untuk refresh syntax sebelum masuk ke soal yang lebih kompleks.', color: 'bg-amber-50 border-amber-200 text-amber-600' },
            { number: '②', title: 'Buka Hint Bertahap', desc: 'Hint 1 paling umum, hint 3 paling spesifik. Buka satu per satu sesuai kebutuhan.', color: 'bg-blue-50 border-blue-200 text-blue-600' },
            { number: '③', title: 'Baca Cara Berpikir', desc: 'Breakdown langkah pendekatan yang bisa dibuka opsional — tanpa langsung spoiler solusi.', color: 'bg-indigo-50 border-indigo-200 text-indigo-600' },
            { number: '④', title: 'Track Progress', desc: 'Badge selesai per soal tersimpan di browser. Lanjutkan kapan saja tanpa perlu login.', color: 'bg-green-50 border-green-200 text-green-600' },
          ].map((step) => (
            <div key={step.number} className={`rounded-2xl border p-5 ${step.color}`}>
              <div className="text-2xl font-bold mb-3">{step.number}</div>
              <h3 className="font-semibold text-gray-900 mb-1.5 text-sm">{step.title}</h3>
              <p className="text-sm text-gray-500 leading-relaxed">{step.desc}</p>
            </div>
          ))}
        </div>
      </section>

      {/* Learning Path Preview */}
      <section className="border-y border-gray-200 bg-white">
        <div className="mx-auto max-w-6xl px-6 py-16">
          <p className="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-2">Jalur Belajar</p>
          <h2 className="text-2xl font-bold text-gray-900 mb-8">Dari dasar ke kompleks, step by step.</h2>
          <div className="flex flex-col gap-4">
            {[
              { category: 'Loop', color: 'bg-blue-100 text-blue-700', problems: ['For Loop Dasar', 'FizzBuzz', 'Bilangan Prima'] },
              { category: 'String', color: 'bg-purple-100 text-purple-700', problems: ['String Dasar', 'Balik String', 'Palindrom'] },
              { category: 'Array', color: 'bg-orange-100 text-orange-700', problems: ['Array Dasar', 'Jumlah Array', 'Cari Duplikat'] },
            ].map((path) => (
              <div key={path.category} className="flex items-center gap-3 flex-wrap">
                <span className={`rounded-lg px-2.5 py-1 text-xs font-semibold w-16 text-center ${path.color}`}>{path.category}</span>
                {path.problems.map((problem, i) => (
                  <div key={problem} className="flex items-center gap-2">
                    <span className={`rounded-lg border px-3 py-1.5 text-xs font-medium ${i === 0 ? 'bg-amber-50 border-amber-200 text-amber-700' : 'bg-white border-gray-200 text-gray-700'}`}>
                      {i === 0 && <span className="mr-1 text-[10px] font-bold text-amber-600">Warm-up</span>}
                      {problem}
                    </span>
                    {i < path.problems.length - 1 && <span className="text-gray-300 text-sm">→</span>}
                  </div>
                ))}
              </div>
            ))}
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

      {/* Languages */}
      {/* Languages */}
      <section className="mx-auto max-w-6xl px-6 pb-16">
        <div className="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-3 mb-10">
          <div>
            <p className="text-xs font-semibold uppercase tracking-widest text-gray-400 mb-2">Bahasa Pemrograman</p>
            <h2 className="text-2xl font-bold text-gray-900">Pilih bahasa, mulai latihan.</h2>
          </div>
        </div>
        <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-4">
          {languages.map((l) => (
            <NavLink
              key={l.label}
              href={l.available ? (l.label === 'SQL' ? '/problems?category=sql' : '/problems') : '#'}
              className={`relative group flex flex-col items-center gap-3 rounded-2xl border p-5 text-center transition-all duration-200 ${
                l.available
                  ? 'bg-white border-gray-200 hover:border-indigo-300 hover:shadow-lg hover:shadow-indigo-100/60 hover:-translate-y-1 cursor-pointer'
                  : 'bg-gray-50/60 border-gray-100 pointer-events-none'
              }`}
            >
              {/* Icon */}
              <div className={`flex items-center justify-center w-12 h-12 rounded-2xl transition ${
                l.available ? 'shadow-sm group-hover:scale-110' : 'opacity-30 grayscale'
              }`}>
                {l.icon}
              </div>

              {/* Label */}
              <span className={`text-sm font-bold tracking-tight ${l.available ? 'text-gray-800' : 'text-gray-300'}`}>
                {l.label}
              </span>

              {/* Badge */}
              {l.available ? (
                <span className="inline-flex items-center gap-1 rounded-full bg-green-50 border border-green-200 text-green-600 text-[10px] font-semibold px-2.5 py-0.5">
                  <span className="w-1.5 h-1.5 rounded-full bg-green-500" />
                  Tersedia
                </span>
              ) : (
                <span className="rounded-full bg-gray-100 text-gray-400 text-[10px] font-medium px-2.5 py-0.5">
                  Coming Soon
                </span>
              )}
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
          <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
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
            <h3 className="text-2xl sm:text-3xl font-extrabold text-white mb-2">Siap belajar yang beneran nempel?</h3>
            <p className="text-indigo-200 mb-7 text-sm">Gratis, tanpa daftar, ada panduan lengkap.</p>
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
