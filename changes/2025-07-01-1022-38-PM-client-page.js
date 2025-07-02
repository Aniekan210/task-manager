"use client"
import { useRef, useEffect, useState } from "react"
import dynamic from 'next/dynamic'
import Link from "next/link"
import Image from "next/image"
import { Button } from "@/components/ui/button"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"

// Safe dynamic import for Lottie with error handling
const Lottie = dynamic(
    () => import('react-lottie')
        .then(mod => mod.default)
        .catch(() => null), // Return null if import fails
    { ssr: false }
)

export default function ClientPage({ testimonials = [] }) {
    const [isMenuOpen, setIsMenuOpen] = useState(false)
    const [lastScrollY, setLastScrollY] = useState(0)
    const [navVisible, setNavVisible] = useState(true)
    const [animationError, setAnimationError] = useState(false)
    const [videoError, setVideoError] = useState(false)
    const [currentTestimonials, setCurrentTestimonials] = useState(testimonials.slice(0, 3))
    const [currentIndex, setCurrentIndex] = useState(0)
    const videoRef = useRef(null)

    // Safe animation options with error handling
    const getAnimationOptions = (animationName) => {
        try {
            const animationData = require(`@/public/animations/${animationName}.json`)
            return {
                loop: true,
                autoplay: true,
                animationData,
                rendererSettings: {
                    preserveAspectRatio: 'xMidYMid slice'
                }
            }
        } catch (e) {
            console.error(`Animation not found: ${animationName}`)
            setAnimationError(true)
            return null
        }
    }

    // Safe animation options - will be null if animations don't exist
    const calendarOptions = getAnimationOptions('calendar')
    const customizationOptions = getAnimationOptions('customization')
    const noLoginOptions = getAnimationOptions('no-login')
    const setupOptions = getAnimationOptions('setup')
... (truncated for brevity)