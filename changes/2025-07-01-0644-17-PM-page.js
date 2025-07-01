// app/page.js
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"

export default function Home() {
  return (
    <div className="min-h-screen bg-white">
      {/* Navigation */}
      <nav className="container mx-auto px-6 py-4 flex justify-between items-center">
        <Link href="#">
          <img src="/logo.avif" alt="schedulee.app Logo" className="h-10" />
        </Link>
        <div className="flex items-center gap-4">
          <Link href="#features" className="text-gray-600 hover:text-blue-500 transition-colors">
            Features
          </Link>
          <Link href="#how-it-works" className="text-gray-600 hover:text-blue-500 transition-colors">
            How It Works
          </Link>
          <Link href="#testimonials" className="text-gray-600 hover:text-blue-500 transition-colors">
            Testimonials
          </Link>
          <Link href="/dashboard/bookings">
            <Button className="bg-blue-500 hover:bg-blue-600">Get Started</Button>
          </Link>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="container mx-auto px-6 py-20 flex flex-col md:flex-row items-center gap-12">
        <div className="md:w-1/2">
          <h1 className="text-5xl font-bold text-gray-900 mb-6">
            Simplify Your Scheduling with <span className="text-blue-500">schedulee.app</span>
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            Let your clients book appointments effortlessly while you maintain full control over your availability.
          </p>
          <div className="flex flex-col sm:flex-row gap-4">
            <Link href="/dashboard/bookings">
              <Button className="bg-blue-500 hover:bg-blue-600 px-8 py-6 text-lg">
                Start Your 14-Day Free Trial
              </Button>
            </Link>
            <Button variant="outline" className="px-8 py-6 text-lg" onClick={() => document.getElementById('demo-video').scrollIntoView()}>
              Watch Demo
            </Button>
          </div>
          <p className="mt-4 text-gray-500">Just $8.99 CAD/month after trial</p>
        </div>
... (truncated for brevity)