import type { Metadata } from 'next'
import './globals.css'
import Navbar from '../components/Navbar'

export const metadata: Metadata = {
  title: 'Balik Ngoding',
  description: 'Platform latihan logika pemrograman dasar untuk fresh graduate',
  metadataBase: new URL('https://balikngoding.my.id'),
  openGraph: {
    title: 'Balik Ngoding',
    description: 'Latih logika koding tanpa bergantung pada AI.',
    url: '/',
    siteName: 'Balik Ngoding',
    images: [
      {
        url: '/og-image.png',
        width: 1200,
        height: 630,
        alt: 'Balik Ngoding - Platform Latihan Logika Koding',
      },
    ],
    locale: 'id_ID',
    type: 'website',
  },
  icons: {
    icon: '/favicon.ico',
    shortcut: '/favicon-16x16.png',
    apple: '/apple-touch-icon.png',
  },
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="id">
      <body className="min-h-screen bg-[#f8f7f4] text-gray-900 antialiased">
        <Navbar />
        {children}
      </body>
    </html>
  )
}