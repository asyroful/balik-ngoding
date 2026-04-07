'use client';

import { useRouter } from 'next/navigation';
import { useTransition } from 'react';

interface NavLinkProps {
  href: string;
  className?: string;
  children: React.ReactNode;
}

// NavLink yang langsung navigasi tanpa nunggu halaman tujuan selesai render
export default function NavLink({ href, className, children }: NavLinkProps) {
  const router = useRouter();
  const [, startTransition] = useTransition();

  function handleClick(e: React.MouseEvent) {
    e.preventDefault();
    startTransition(() => {
      router.push(href);
    });
  }

  return (
    <a href={href} onClick={handleClick} className={className}>
      {children}
    </a>
  );
}
