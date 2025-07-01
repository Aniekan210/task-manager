import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Script from "next/script";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata = {
  title: "Schedulee.app | Modern Booking & Scheduling Platform",
  description: "Simplify your scheduling process with Schedulee.app - The intuitive booking platform for businesses and professionals.",
  metadataBase: new URL('https://schedulee.app'),
  alternates: {
    canonical: '/',
  },
  openGraph: {
    title: "Schedulee.app | Modern Booking & Scheduling Platform",
    description: "Simplify your scheduling process with Schedulee.app - The intuitive booking platform for businesses and professionals.",
    url: "https://schedulee.app",
    siteName: "Schedulee.app",
    images: [
      {
        url: "https://schedulee.app/logo.png",
        width: 600,
        height: 600,
        alt: "Schedulee.app Logo",
      },
    ],
    locale: "en_CA",
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
    title: "Schedulee.app | Modern Booking & Scheduling Platform",
    description: "Simplify your scheduling process with Schedulee.app - The intuitive booking platform for businesses and professionals.",
    images: {
      url: "https://schedulee.app/logo.png",
      alt: "Schedulee.app Logo",
      width: 1200,
      height: 1200,
    },
  },
  robots: {
    index: true,
... (truncated for brevity)