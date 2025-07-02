// components/ui/interactive-elements.jsx
"use client"

import { useRef, useEffect, useState } from "react"
import dynamic from 'next/dynamic'
import { Button } from "@/components/ui/button"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import Link from "next/link"
import Image from "next/image"
import { motion, AnimatePresence } from "framer-motion"

// Safe dynamic import for Lottie with robust error handling
const Lottie = dynamic(
  () => import('react-lottie')
    .then(mod => mod.default)
    .catch(() => {
      console.error("Lottie animation library failed to load")
      return () => null
    }),
  { ssr: false }
)

// Enhanced Placeholder component
export function Placeholder({ name = "content", className = "", children }) {
  return (
    <motion.div 
      className={`bg-gray-100 rounded-lg flex flex-col items-center justify-center p-4 ${className}`}
      whileHover={{ scale: 1.02 }}
    >
      <svg 
        xmlns="http://www.w3.org/2000/svg" 
        className="h-10 w-10 text-gray-400 mb-2" 
        fill="none" 
        viewBox="0 0 24 24" 
        stroke="currentColor"
      >
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
      </svg>
      <span className="text-gray-500 text-center">
        {name || children || 'Placeholder content'}
      </span>
    </motion.div>
  )
}

export function VideoModal({ isOpen, onClose }) {
  const videoRef = useRef(null)

  useEffect(() => {
    if (isOpen) {
... (truncated for brevity)