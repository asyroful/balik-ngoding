import Link from 'next/link'

export default function NotFound() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-gray-50 px-4">
      <div className="max-w-md w-full text-center space-y-4">
        <p className="text-6xl font-bold text-gray-300">404</p>
        <h1 className="text-2xl font-semibold text-gray-800">
          Halaman tidak ditemukan
        </h1>
        <p className="text-gray-500">
          Halaman yang kamu cari tidak ada atau sudah dipindahkan.
        </p>
        <Link
          href="/problems"
          className="inline-block mt-4 px-6 py-3 bg-blue-600 text-white font-medium rounded-lg hover:bg-blue-700 transition-colors"
        >
          Kembali ke Daftar Soal
        </Link>
      </div>
    </main>
  )
}
