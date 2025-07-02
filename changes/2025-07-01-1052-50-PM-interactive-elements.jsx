// components/interactive-elements.jsx
"use client"

import { useRef, useEffect, useState } from "react"
import dynamic from 'next/dynamic'
import { Button } from "@/components/ui/button"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import Link from "next/link"
import Image from "next/image"

// Safe dynamic import for Lottie with robust error handling
const Lottie = dynamic(
  () => import('react-lottie')
    .then(mod => mod.default)
    .catch(() => {
      console.error("Lottie animation library failed to load")
      return () => <div className="w-full h-full bg-gray-100 flex items-center justify-center">
        <span className="text-gray-500">Animation player not available</span>
      </div>
    }),
  { ssr: false }
)

// Generic placeholder component
const Placeholder = ({ name = "content", className = "" }) => (
  <div className={`bg-gray-100 rounded-lg flex items-center justify-center ${className}`}>
    <span className="text-gray-500">{name} placeholder</span>
  </div>
)

export function MobileMenu({ testimonials = [] }) {
  const [isMenuOpen, setIsMenuOpen] = useState(false)
  const [lastScrollY, setLastScrollY] = useState(0)
  const [navVisible, setNavVisible] = useState(true)
  const [logoError, setLogoError] = useState(false)

  // Handle scroll for navbar hide/show
  useEffect(() => {
    const handleScroll = () => {
      const currentScrollY = window.scrollY
      if (currentScrollY > lastScrollY && currentScrollY > 100) {
        setNavVisible(false)
      } else {
        setNavVisible(true)
      }
      setLastScrollY(currentScrollY)
    }

    window.addEventListener('scroll', handleScroll, { passive: true })
    return () => window.removeEventListener('scroll', handleScroll)
... (truncated for brevity)