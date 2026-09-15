'use client';

import { useEffect } from 'react';
import Script from 'next/script';

export default function PWASetup() {
  useEffect(() => {
    // iOS home screen support
    if (window.navigator.standalone) {
      document.body.classList.add('ios-standalone');
    }
  }, []);

  return (
    <>
      <Script id="pwa-init" strategy="afterInteractive">
        {`
          if ('serviceWorker' in navigator) {
            window.addEventListener('load', function() {
              navigator.serviceWorker.register('/sw.js');
            });
          }
        `}
      </Script>
    </>
  );
}
