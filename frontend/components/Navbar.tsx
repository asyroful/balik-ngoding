import Link from 'next/link';
import NavLink from './NavLink';

export default function Navbar() {
  return (
    <nav className="sticky top-0 z-50 border-b border-gray-200 bg-white/90 backdrop-blur-sm">
      <div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-3.5">
        <Link href="/" className="flex items-center gap-2.5">
          <span className="flex h-7 w-7 items-center justify-center rounded-md bg-indigo-700 text-xs font-bold text-white tracking-tight">
            BN
          </span>
          <span className="text-sm font-bold text-gray-900 tracking-tight">Balik Ngoding</span>
        </Link>
      </div>
    </nav>
  );
}
